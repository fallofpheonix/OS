package containment

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// ProcessAudit stores the history of containment actions.
type ProcessAudit struct {
	mu       sync.RWMutex
	History  []ProcessAction
	Sequence int
}

func NewProcessAudit() *ProcessAudit {
	return &ProcessAudit{
		History:  []ProcessAction{},
		Sequence: 0,
	}
}

// computeHash calculates the hash of a process action based on fields.
func computeHash(a ProcessAction) string {
	b, _ := json.Marshal(struct {
		PID        int
		Action     ProcessActionType
		Reason     string
		EvidenceID string
		DecisionID string
		Sequence   int
	}{
		PID:        a.PID,
		Action:     a.Action,
		Reason:     a.Reason,
		EvidenceID: a.EvidenceID,
		DecisionID: a.DecisionID,
		Sequence:   a.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// LogAction binds a safe primitive action to the audit trail.
func (a *ProcessAudit) LogAction(action ProcessAction) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Sequence++
	action.Sequence = a.Sequence
	action.Timestamp = time.Now()
	action.Hash = computeHash(action)
	a.History = append(a.History, action)
}
