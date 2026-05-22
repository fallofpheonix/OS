package ledger

import (
	"fmt"
)

// ReplayAPI defines the interface for forensic retrieval
type ReplayAPI interface {
	GetEntryByEventID(eventID string) (*LedgerEntry, error)
	ReconstructCausalChain(eventID string) ([]LedgerEntry, error)
}

// GetEntryByEventID searches the ledger for a specific action
func (l *Ledger) GetEntryByEventID(eventID string) (*LedgerEntry, error) {
	for _, entry := range l.Entries {
		if entry.EventID == eventID {
			return &entry, nil
		}
	}
	return nil, fmt.Errorf("event ID %s not found", eventID)
}

// ReconstructCausalChain walks backward through the ledger using ParentIDs (DAG)
func (l *Ledger) ReconstructCausalChain(eventID string) ([]LedgerEntry, error) {
	chain := []LedgerEntry{}
	currentEntry, err := l.GetEntryByEventID(eventID)
	if err != nil {
		return nil, err
	}
	chain = append(chain, *currentEntry)

	curr := *currentEntry
	for len(curr.ParentIDs) > 0 {
		parentHash := curr.ParentIDs[0] // Taking the first parent
		parentEntry, ok := l.Entries[fmt.Sprintf("%x", parentHash)]
		if !ok {
			return chain, fmt.Errorf("missing parent hash in DAG")
		}
		chain = append(chain, parentEntry)
		curr = parentEntry
	}
	return chain, nil
}
