package ledger

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"

	"phoenix/common/resource"
)

type LedgerEntry struct {
	LogicalTick uint64
	EventID     string
	CauseID     string
	Payload     []byte
	ParentIDs   [][]byte
	Hash        []byte
}

type Ledger struct {
	Entries   map[string]LedgerEntry
	Heads     [][]byte
	Counter   uint64
	Allocator *resource.BoundedAllocator
}

func NewLedger(alloc *resource.BoundedAllocator) *Ledger {
	return &Ledger{
		Entries:   make(map[string]LedgerEntry),
		Heads:     nil,
		Allocator: alloc,
	}
}

func (l *Ledger) AddEntry(eventID, causeID string, payload []byte) error {
	if l.Allocator != nil {
		if err := l.Allocator.Allocate(uint64(len(payload))); err != nil {
			return err
		}
	}

	entry := LedgerEntry{
		LogicalTick: l.Counter,
		EventID:     eventID,
		CauseID:     causeID,
		Payload:     payload,
		ParentIDs:   l.Heads,
	}
	l.Counter++

	entry.Hash = l.computeHash(entry)
	l.Entries[fmt.Sprintf("%x", entry.Hash)] = entry
	l.Heads = [][]byte{entry.Hash}
	
	return nil
}

// Checkpoint returns a deterministic snapshot of the current Ledger head and counter.
func (l *Ledger) Checkpoint() ([]byte, error) {
	h := sha256.New()
	binary.Write(h, binary.BigEndian, l.Counter)
	for _, head := range l.Heads {
		h.Write(head)
	}
	return h.Sum(nil), nil
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

// Prune removes old ledger entries older than the specified depth, ensuring determinism.
func (l *Ledger) Prune(depth uint64) (int, error) {
	if l.Counter <= depth {
		return 0, nil
	}
	
	cutoff := l.Counter - depth
	removed := 0
	
	for hash, entry := range l.Entries {
		if entry.LogicalTick < cutoff {
			delete(l.Entries, hash)
			removed++
		}
	}
	
	return removed, nil
}

func (l *Ledger) Verify() error {
	for hashHex, entry := range l.Entries {
		computed := l.computeHash(entry)
		if fmt.Sprintf("%x", computed) != hashHex {
			return fmt.Errorf("invalid hash at tick %d", entry.LogicalTick)
		}
	}
	return nil
}
