// Package ledger provides the append-only forensic record for PhoenixOS.
// Domain Logic: Implements the "Evidence Merkle DAG", an immutable ledger that records events and their state transitions with cryptographic integrity and causal chaining.
// Responsibility: Ensures tamper-proof logging and causal traceability of system events.
package ledger

import (
	"bytes"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"io"
	"sort"
	"sync"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// LedgerEntry is the canonical forensic record in the Evidence Merkle DAG.
// Concurrency: Read-only instances are thread-safe.
// State Management: Encapsulates metadata, payload, and cryptographic hashes for a single event.
type LedgerEntry struct {
	EventID     string   `json:"event_id"`
	CauseID     string   `json:"cause_id"`
	Sequence    uint64   `json:"sequence"` // WHY: Ledger-assigned insertion index for tie-breaking.
	ParentIDs   [][]byte `json:"parent_ids"`
	LogicalTick uint64   `json:"logical_tick"`
	Source      string   `json:"source"`
	Payload     []byte   `json:"payload"`
	Hash        []byte   `json:"hash"`
	TraceHash   string   `json:"trace_hash"`
	// Ledger V2 Fields
	StateBefore    []byte `json:"state_before"`
	StateAfter     []byte `json:"state_after"`
	PolicyVersion  string `json:"policy_version"`
	ValidationHash []byte `json:"validation_hash"`
}

// ResourceAllocator defines an interface for deterministic resource bounding.
type ResourceAllocator interface {
	Allocate(bytes uint64) error
	Deallocate(bytes uint64)
}

// Ledger is the append-only Evidence Merkle DAG.
// Concurrency: Thread-safe via sync.RWMutex.
// State Management: Maintains a map of entries, tracks leaf nodes (heads), and uses an allocator for memory management.
type Ledger struct {
	mu        sync.RWMutex
	Entries   map[string]LedgerEntry
	Heads     [][]byte
	Counter   uint64
	Sequence  uint64 // Tracks next monotonic sequence number.
	allocator ResourceAllocator
	persistor *Persistor
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewLedger initializes a new ledger instance.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewLedger(alloc ResourceAllocator) *Ledger {
	return &Ledger{
		Entries:   make(map[string]LedgerEntry),
		Heads:     nil,
		allocator: alloc,
		Sequence:  0,
	}
}

// WithPersistor attaches a durabilty layer to the ledger.
func (l *Ledger) WithPersistor(p *Persistor) *Ledger {
	l.persistor = p
	return l
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// AddEntryV2 appends a new entry to the Merkle DAG with V2 state transition metadata.
// I/O: Indirect via ResourceAllocator (memory bounding) and Persistor (disk).
// Side Effects: Modifies internal map (Entries) and slice (Heads), appends to disk.
// Complexity: O(H + D) where H is the number of heads and D is the payload size.
func (l *Ledger) AddEntryV2(eventID, causeID string, tick uint64, payload []byte, traceHash string, stateBefore, stateAfter []byte, policyVersion string) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	// Assign sequence under lock
	l.Sequence++
	seq := l.Sequence

	// Deterministic Resource Bounding
	if l.allocator != nil {
		if err := l.allocator.Allocate(uint64(len(payload) + 256)); err != nil {
			return err
		}
	}

	entry := LedgerEntry{
		LogicalTick:   tick,
		Sequence:      seq,
		EventID:       eventID,
		CauseID:       causeID,
		Payload:       payload,
		ParentIDs:     l.Heads,
		TraceHash:     traceHash,
		StateBefore:   stateBefore,
		StateAfter:    stateAfter,
		PolicyVersion: policyVersion,
	}

	entry.Hash = l.computeHash(entry)
	entry.ValidationHash = l.computeValidationHash(entry)

	if l.persistor != nil {
		if err := l.persistor.Append(entry); err != nil {
			return err
		}
	}

	if l.Entries == nil {
		l.Entries = make(map[string]LedgerEntry)
	}

	l.Entries[fmt.Sprintf("%x", entry.Hash)] = entry
	l.Heads = [][]byte{entry.Hash}
	return nil
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// AddEntry appends a standard entry to the ledger (v1 compatibility).
// I/O: Indirect via ResourceAllocator.
// Side Effects: Modifies internal state.
// Complexity: O(H + D).
func (l *Ledger) AddEntry(eventID, causeID string, tick uint64, payload []byte) error {
	return l.AddEntryV2(eventID, causeID, tick, payload, "", nil, nil, "0.0.1")
}

const HashSchemaVersion = 2

// writeLP writes a uint64 length prefix followed by the data to the writer.
func writeLP(w io.Writer, data []byte) {
	binary.Write(w, binary.BigEndian, uint64(len(data)))
	w.Write(data)
}

// writeStringLP writes a uint64 length prefix followed by the string bytes to the writer.
func writeStringLP(w io.Writer, s string) {
	writeLP(w, []byte(s))
}

func (l *Ledger) computeHash(entry LedgerEntry) []byte {
	h := sha256.New()
	// Versioned Hash Schema
	binary.Write(h, binary.BigEndian, uint32(HashSchemaVersion))

	if err := binary.Write(h, binary.BigEndian, entry.LogicalTick); err != nil {
		// Panic is acceptable here as hashing failure indicates a system-level invariant violation
		panic(fmt.Sprintf("ledger: failed to write logical tick to hash: %v", err))
	}

	if err := binary.Write(h, binary.BigEndian, entry.Sequence); err != nil {
		panic(fmt.Sprintf("ledger: failed to write sequence to hash: %v", err))
	}

	writeStringLP(h, entry.EventID)
	writeStringLP(h, entry.CauseID)

	// Write ParentIDs length then each parent
	binary.Write(h, binary.BigEndian, uint64(len(entry.ParentIDs)))
	for _, p := range entry.ParentIDs {
		writeLP(h, p)
	}

	writeLP(h, entry.Payload)
	writeStringLP(h, entry.TraceHash)

	// V2 Fields
	writeLP(h, entry.StateBefore)
	writeLP(h, entry.StateAfter)
	writeStringLP(h, entry.PolicyVersion)

	return h.Sum(nil)
}

func (l *Ledger) computeValidationHash(entry LedgerEntry) []byte {
	h := sha256.New()

	writeLP(h, entry.StateBefore)
	writeLP(h, entry.StateAfter)
	writeStringLP(h, entry.PolicyVersion)
	writeLP(h, entry.Hash)

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
			if !bytes.Equal(valHash, entry.ValidationHash) {
				return fmt.Errorf("invalid validation hash at tick %d", entry.LogicalTick)
			}
		}

		for _, p := range entry.ParentIDs {
			parentHex := fmt.Sprintf("%x", p)
			parent, ok := l.Entries[parentHex]
			if !ok {
				return fmt.Errorf("missing parent hash %x at tick %d", p, entry.LogicalTick)
			}

			if len(entry.StateBefore) > 0 && len(parent.StateAfter) > 0 {
				if !bytes.Equal(entry.StateBefore, parent.StateAfter) {
					return fmt.Errorf("state transition gap: before state %x does not match parent after state %x at tick %d",
						entry.StateBefore, parent.StateAfter, entry.LogicalTick)
				}
			}
		}
	}
	return nil
}

// LABEL: [READ_ONLY] [UNCONSTRAINED] [STABLE]
// Checkpoint returns the current head hash of the ledger.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func (l *Ledger) Checkpoint() ([]byte, error) {
	l.mu.RLock()
	defer l.mu.RUnlock()

	if len(l.Heads) == 0 {
		return nil, fmt.Errorf("ledger has no heads")
	}
	return l.Heads[0], nil
}

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// Prune removes historical entries beyond a specified retention depth.
// I/O: Indirect via ResourceAllocator (deallocation).
// Side Effects: Modifies Entries map.
// Complexity: O(E) where E is the number of entries.
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

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// RollbackTo restores the ledger to a specific checkpoint.
// I/O: Indirect via ResourceAllocator.
// Side Effects: Modifies Entries map, Heads, and Counter.
// Complexity: O(E) where E is the number of entries.
func (l *Ledger) RollbackTo(checkpoint []byte) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	checkpointHex := fmt.Sprintf("%x", checkpoint)
	targetEntry, ok := l.Entries[checkpointHex]
	if !ok {
		return fmt.Errorf("checkpoint not found: %s", checkpointHex)
	}

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
	}

	l.Heads = [][]byte{checkpoint}
	l.Counter = targetEntry.LogicalTick + 1
	return nil
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [EXPERIMENTAL]
// GenerateCertificate produces a cryptographic stub for event authorization.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func (l *Ledger) GenerateCertificate(eventID string, weight phxmath.FixedPoint) ([]byte, error) {
	return []byte(fmt.Sprintf("CERT:%s:%v", eventID, weight)), nil
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [EXPERIMENTAL]
// VerifyCertificate validates an authorization certificate.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func (l *Ledger) VerifyCertificate(eventID string, weight phxmath.FixedPoint, cert []byte) bool {
	expected := fmt.Sprintf("CERT:%s:%v", eventID, weight)
	return string(cert) == expected
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// SortedEntries returns all ledger entries ordered by their logical tick.
// I/O: None.
// Side Effects: None.
// Complexity: O(E log E).
func (l *Ledger) SortedEntries() []LedgerEntry {
	l.mu.RLock()
	defer l.mu.RUnlock()

	var entries []LedgerEntry
	for _, e := range l.Entries {
		entries = append(entries, e)
	}
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].LogicalTick != entries[j].LogicalTick {
			return entries[i].LogicalTick < entries[j].LogicalTick
		}
		return entries[i].Sequence < entries[j].Sequence
	})
	return entries
}
