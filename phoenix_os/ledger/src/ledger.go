package ledger

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"sort"
)

// LedgerEntry is the canonical forensic record in the Evidence Merkle DAG.
type LedgerEntry struct {
	EventID     string   `json:"event_id"`
	CauseID     string   `json:"cause_id"`
	ParentIDs   [][]byte `json:"parent_ids"`
	LogicalTick uint64   `json:"logical_tick"`
	Source      string   `json:"source"`
	Payload     []byte   `json:"payload"`
	Hash        []byte   `json:"hash"`
}

type ResourceAllocator interface {
	Allocate(bytes uint64) error
	Deallocate(bytes uint64)
}

// Ledger is the append-only Evidence Merkle DAG.
type Ledger struct {
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

func (l *Ledger) AddEntry(eventID, causeID string, payload []byte) error {
	entry := LedgerEntry{
		LogicalTick: l.Counter,
		EventID:     eventID,
		CauseID:     causeID,
		Payload:     payload,
		ParentIDs:   l.Heads,
	}
	l.Counter++

	entry.Hash = l.computeHash(entry)
	
	// Deterministic Resource Bounding
	if l.allocator != nil {
		if err := l.allocator.Allocate(uint64(len(payload) + 128)); err != nil {
			return err
		}
	}

	l.Entries[fmt.Sprintf("%x", entry.Hash)] = entry
	l.Heads = [][]byte{entry.Hash}
	return nil
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
	return h.Sum(nil)
}

func (l *Ledger) Verify() error {
	for hashHex, entry := range l.Entries {
		computed := l.computeHash(entry)
		if fmt.Sprintf("%x", computed) != hashHex {
			return fmt.Errorf("invalid hash at tick %d", entry.LogicalTick)
		}
		for _, p := range entry.ParentIDs {
			if _, ok := l.Entries[fmt.Sprintf("%x", p)]; !ok {
				return fmt.Errorf("missing parent hash %x at tick %d", p, entry.LogicalTick)
			}
		}
	}
	return nil
}

// Checkpoint returns the current head hash for rollback/branching.
func (l *Ledger) Checkpoint() ([]byte, error) {
	if len(l.Heads) == 0 {
		return nil, fmt.Errorf("ledger has no heads")
	}
	// Return the primary head (simplified for MVP)
	return l.Heads[0], nil
}

// Prune removes entries older than the specified retention depth.
func (l *Ledger) Prune(depth uint64) (int, error) {
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

// SortedEntries returns entries sorted by LogicalTick (for testing/display)
func (l *Ledger) SortedEntries() []LedgerEntry {
	var entries []LedgerEntry
	for _, e := range l.Entries {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].LogicalTick < entries[j].LogicalTick
	})
	return entries
}
