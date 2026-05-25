package network

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// NetworkAudit stores the history of network containment actions.
type NetworkAudit struct {
	mu           sync.RWMutex
	History      []NetworkAction
	Sequence     int
	ReplayCursor int
}

func NewNetworkAudit() *NetworkAudit {
	return &NetworkAudit{
		History:      []NetworkAction{},
		Sequence:     0,
		ReplayCursor: 0,
	}
}

// computeHash calculates the deterministic hash for a network action.
func computeHash(a NetworkAction) string {
	b, _ := json.Marshal(struct {
		Src        string
		Dst        string
		Port       int
		Action     NetworkActionType
		Reason     string
		EvidenceID string
		DecisionID string
		Sequence   int
	}{
		Src:        a.Src,
		Dst:        a.Dst,
		Port:       a.Port,
		Action:     a.Action,
		Reason:     a.Reason,
		EvidenceID: a.EvidenceID,
		DecisionID: a.DecisionID,
		Sequence:   a.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// LogAction binds a network action to the audit trail.
func (a *NetworkAudit) LogAction(action NetworkAction) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Sequence++
	action.Sequence = a.Sequence
	action.Timestamp = time.Now()
	action.Hash = computeHash(action)
	a.History = append(a.History, action)
}
