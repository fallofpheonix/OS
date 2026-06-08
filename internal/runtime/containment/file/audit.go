/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: CYCLE 9i — FILE CONTAINMENT AUDIT (Layer 4.5)
 *
 * The FileAudit stores the history of all file containment actions with
 * deterministic hashing. Each action is hashed and sequenced for
 * tamper-evident audit trail.
 *
 * WORKFLOW:
 *   FileAudit.LogAction(action)
 *     → Assign sequence number
 *     → Compute deterministic hash (Path, Action, Reason, Evidence, Decision)
 *     → Append to History
 *   → FileAudit.CreateSnapshot() → serialize for persistence
 *   → FileAudit.RestoreFromSnapshot() → deserialize and validate
 *
 * THREAD SAFETY: Uses sync.RWMutex for concurrent access.
 * ========================================================================= */
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
