package file

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sync"
	"time"
)
// FileAudit stores the history of file containment actions.
type FileAudit struct {
	mu           sync.RWMutex
	History      []FileAction
	Sequence     int
	ReplayCursor int
}

func NewFileAudit() *FileAudit {
	return &FileAudit{
		History:      []FileAction{},
		Sequence:     0,
		ReplayCursor: 0,
	}
}

// computeHash calculates the deterministic hash for a file action.
func computeHash(a FileAction) string {
	b, _ := json.Marshal(struct {
		Path       string
		Action     FileActionType
		Reason     string
		EvidenceID string
		DecisionID string
		Sequence   int
	}{
		Path:       a.Path,
		Action:     a.Action,
		Reason:     a.Reason,
		EvidenceID: a.EvidenceID,
		DecisionID: a.DecisionID,
		Sequence:   a.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// LogAction binds a file action to the audit trail.
func (a *FileAudit) LogAction(action FileAction) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Sequence++
	action.Sequence = a.Sequence
	action.Timestamp = time.Now()
	action.Hash = computeHash(action)
	a.History = append(a.History, action)
}
