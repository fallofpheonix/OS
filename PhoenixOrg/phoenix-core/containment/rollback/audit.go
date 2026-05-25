package rollback

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)

// RollbackAudit manages the history of cross-layer rollbacks.
type RollbackAudit struct {
	mu           sync.RWMutex
	History      []RollbackRecord
	ReplayCursor int
}

func NewRollbackAudit() *RollbackAudit {
	return &RollbackAudit{
		History:      []RollbackRecord{},
		ReplayCursor: 0,
	}
}

// computeHash calculates the deterministic hash for a rollback record.
func computeHash(r RollbackRecord) string {
	b, _ := json.Marshal(struct {
		Component     Component
		SnapshotID    string
		RecoveryID    string
		EvidenceID    string
		DecisionID    string
		Sequence      int
	}{
		Component:  r.Component,
		SnapshotID: r.SnapshotID,
		RecoveryID: r.RecoveryID,
		EvidenceID: r.EvidenceID,
		DecisionID: r.DecisionID,
		Sequence:   r.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// LogRollback commits a new rollback event to the audit log.
func (a *RollbackAudit) LogRollback(record RollbackRecord) {
	a.mu.Lock()
	defer a.mu.Unlock()

	record.Timestamp = time.Now()
	record.Hash = computeHash(record)
	a.History = append(a.History, record)
}
