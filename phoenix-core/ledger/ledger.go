package ledger

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
)

// LedgerEntry represents a node in the Evidence Merkle DAG.
type LedgerEntry struct {
	LogicalTick   uint64
	EventID       string // ActionID renamed to EventID
	CauseID       string
	Source        string
	PolicyVersion string
	StateBefore   []byte
	StateAfter    []byte
	ParentIDs     [][]byte
	Payload       []byte
	PrevHash      []byte
	ReplayHash    []byte
	Hash          []byte
}

// Ledger is the append-only Evidence Merkle DAG.
type Ledger struct {
	Entries map[string]LedgerEntry // keyed by hex(Hash) for DAG traversal
	Heads   [][]byte               // hashes of current leaf nodes
	Counter uint64                 // logical tick source
}

func NewLedger() *Ledger {
	return &Ledger{
		Entries: make(map[string]LedgerEntry),
		Heads:   nil,
	}
}

// computeHash calculates the hash of an entry
func (l *Ledger) computeHash(entry LedgerEntry) []byte {
	h := sha256.New()
	binary.Write(h, binary.BigEndian, entry.LogicalTick)
	h.Write([]byte(entry.EventID))
	h.Write([]byte(entry.CauseID))
	h.Write([]byte(entry.Source))
	h.Write([]byte(entry.PolicyVersion))
	h.Write([]byte(entry.StateBefore))
	h.Write([]byte(entry.StateAfter))
	
	// Hash parents in deterministic order
	for _, p := range entry.ParentIDs {
		h.Write(p)
	}
	
	h.Write(entry.Payload)
	h.Write(entry.PrevHash)
	h.Write(entry.ReplayHash)

	return h.Sum(nil)
}

func (l *Ledger) Verify() error {
	for hashHex, entry := range l.Entries {
		computed := l.computeHash(entry)
		
		if fmt.Sprintf("%x", computed) != hashHex {
			return fmt.Errorf("invalid hash at tick %d", entry.LogicalTick)
		}
		
		// Verify parents exist
		for _, p := range entry.ParentIDs {
			if _, ok := l.Entries[fmt.Sprintf("%x", p)]; !ok {
				return fmt.Errorf("missing parent hash %x at tick %d", p, entry.LogicalTick)
			}
		}
	}
	return nil
}

// AddEntry adds a new event to the ledger, automatically linking it to current heads.
func (l *Ledger) AddEntry(eventID, causeID string, payload []byte) error {
	// ... (Implementation needed)
	return nil
}
