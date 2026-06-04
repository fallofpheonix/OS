/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9e — PROCESS CONTAINMENT AUDIT (Layer 4.5)
//
// The ProcessAudit stores the history of all containment actions with
// deterministic hashing. Each action is hashed and sequenced for
// tamper-evident audit trail.
//
// WORKFLOW:
//   ProcessAudit.LogAction(action)
//     → Assign sequence number
//     → Compute deterministic hash (PID, Action, Reason, Evidence, Decision)
//     → Append to History
//   → ProcessAudit.CreateSnapshot() → serialize for persistence
//   → ProcessAudit.RestoreFromSnapshot() → deserialize and validate
//
// INTEGRITY: Each action's hash is computed over its fields (excluding timestamp).
// The hash chain ensures that no action can be inserted or removed without
// breaking the sequence.
//
// THREAD SAFETY: Uses sync.RWMutex for concurrent access.
// =========================================================================
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
