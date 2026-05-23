package ledger

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
	"sync"
)

// LedgerEntry is the canonical forensic record in the Evidence Merkle DAG.
type LedgerEntry struct {
	EventID        string   `json:"event_id"`
	CauseID        string   `json:"cause_id"`
	ParentIDs      [][]byte `json:"parent_ids"`
	LogicalTick    uint64   `json:"logical_tick"`
	Source         string   `json:"source"`
	Payload        []byte   `json:"payload"`
	Hash           []byte   `json:"hash"`
	// Ledger V2 Fields
	StateBefore    string   `json:"state_before"`
	StateAfter     string   `json:"state_after"`
	PolicyVersion  string   `json:"policy_version"`
	ValidationHash []byte   `json:"validation_hash"`
}

type ResourceAllocator interface {
	Allocate(bytes uint64) error
	Deallocate(bytes uint64)
}

// Ledger is the append-only Evidence Merkle DAG.
type Ledger struct {
	mu        sync.RWMutex
	Entries   map[string]LedgerEntry
	Heads     [][]byte
	Counter   uint64
	allocator ResourceAllocator
}

func NewLedger(alloc ResourceAllocator) *Ledger {
	return &Ledger{
		Entries:   make(map[string]LedgerEntry),
		Heads:     nil,
		allocator: alloc,
	}
}

func (l *Ledger) AddEntryV2(eventID, causeID string, payload []byte, stateBefore, stateAfter, policyVersion string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Deterministic Resource Bounding (check before modifying state)
	if l.allocator != nil {
		if err := l.allocator.Allocate(uint64(len(payload) + 256)); err != nil {
			return err
		}
	}

	entry := LedgerEntry{
		LogicalTick:   l.Counter,
		EventID:       eventID,
		CauseID:       causeID,
		Payload:       payload,
		ParentIDs:     l.Heads,
		StateBefore:   stateBefore,
		StateAfter:    stateAfter,
		PolicyVersion: policyVersion,
	}
	l.Counter++

	entry.Hash = l.computeHash(entry)
	entry.ValidationHash = l.computeValidationHash(entry)

	l.Entries[fmt.Sprintf("%x", entry.Hash)] = entry
	l.Heads = [][]byte{entry.Hash}
	return nil
}

func (l *Ledger) AddEntry(eventID, causeID string, payload []byte) error {
	return l.AddEntryV2(eventID, causeID, payload, "", "", "")
}

func (l *Ledger) computeHash(entry LedgerEntry) []byte {
	h := sha256.New()
	binary.Write(h, binary.BigEndian, entry.LogicalTick)
	h.Write([]byte(entry.EventID))
	h.Write([]byte(entry.CauseID))
	for _, p := range entry.ParentIDs {
		h.Write(p)
	}
	h.Write(entry.Payload)
	// V2 Fields
	h.Write([]byte(entry.StateBefore))
	h.Write([]byte(entry.StateAfter))
	h.Write([]byte(entry.PolicyVersion))
	return h.Sum(nil)
}

func (l *Ledger) computeValidationHash(entry LedgerEntry) []byte {
	h := sha256.New()
	h.Write([]byte(entry.StateBefore + "->" + entry.StateAfter))
	h.Write([]byte(entry.PolicyVersion))
	h.Write(entry.Hash)
	return h.Sum(nil)
}

func (l *Ledger) Verify() error {
	l.mu.RLock()
	defer l.mu.RUnlock()

	for hashHex, entry := range l.Entries {
		computed := l.computeHash(entry)
		if fmt.Sprintf("%x", computed) != hashHex {
			return fmt.Errorf("invalid hash at tick %d", entry.LogicalTick)
		}

		if len(entry.ValidationHash) > 0 {
			valHash := l.computeValidationHash(entry)
			if fmt.Sprintf("%x", valHash) != fmt.Sprintf("%x", entry.ValidationHash) {
				return fmt.Errorf("invalid validation hash at tick %d", entry.LogicalTick)
			}
		}

		for _, p := range entry.ParentIDs {
			parentHex := fmt.Sprintf("%x", p)
			parent, ok := l.Entries[parentHex]
			if !ok {
				return fmt.Errorf("missing parent hash %x at tick %d", p, entry.LogicalTick)
			}

			// Verify state transition consistency if V2 fields are present
			if entry.StateBefore != "" && parent.StateAfter != "" {
				if entry.StateBefore != parent.StateAfter {
					return fmt.Errorf("state transition gap: before state %s does not match parent after state %s at tick %d",
						entry.StateBefore, parent.StateAfter, entry.LogicalTick)
				}
			}
		}
	}
	return nil
}

// Checkpoint returns the current head hash for rollback/branching.
func (l *Ledger) Checkpoint() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if len(l.Heads) == 0 {
		return nil, fmt.Errorf("ledger has no heads")
	}
	// Return the primary head (simplified for MVP)
	return l.Heads[0], nil
}

// Prune removes entries older than the specified retention depth.
func (l *Ledger) Prune(depth uint64) (int, error) {
	l.mu.Lock()
	defer l.mu.Unlock()

	if l.Counter <= depth {
		return 0, nil
	}

	cutoff := l.Counter - depth
	removedCount := 0

	var toDelete []string
	for hash, entry := range l.Entries {
		if entry.LogicalTick < cutoff {
			toDelete = append(toDelete, hash)
		}
	}

	for _, hash := range toDelete {
		entry := l.Entries[hash]
		if l.allocator != nil {
			l.allocator.Deallocate(uint64(len(entry.Payload) + 128))
		}
		delete(l.Entries, hash)
		removedCount++
	}

	return removedCount, nil
}

// RollbackTo restores the ledger to a specific checkpoint hash.
// This addresses [PRO-009].
func (l *Ledger) RollbackTo(checkpoint []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	checkpointHex := fmt.Sprintf("%x", checkpoint)
	targetEntry, ok := l.Entries[checkpointHex]
	if !ok {
		return fmt.Errorf("checkpoint not found: %s", checkpointHex)
	}

	// Prune everything after the target logical tick
	removedCount := 0
	var toDelete []string
	for hash, entry := range l.Entries {
		if entry.LogicalTick > targetEntry.LogicalTick {
			toDelete = append(toDelete, hash)
		}
	}

	for _, hash := range toDelete {
		entry := l.Entries[hash]
		if l.allocator != nil {
			l.allocator.Deallocate(uint64(len(entry.Payload) + 256))
		}
		delete(l.Entries, hash)
		removedCount++
	}

	l.Heads = [][]byte{checkpoint}
	l.Counter = targetEntry.LogicalTick + 1
	return nil
}

// SortedEntries returns entries sorted by LogicalTick (for testing/display)
func (l *Ledger) SortedEntries() []LedgerEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var entries []LedgerEntry
	for _, e := range l.Entries {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].LogicalTick < entries[j].LogicalTick
	})
	return entries
}
