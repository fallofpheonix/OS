/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: CYCLE 9j — NETWORK CONTAINMENT AUDIT (Layer 4.5)
 *
 * The NetworkAudit stores the history of all network containment actions
 * with deterministic hashing. Each action is hashed and sequenced for
 * tamper-evident audit trail.
 *
 * WORKFLOW:
 *   NetworkAudit.LogAction(action)
 *     → Assign sequence number
 *     → Compute deterministic hash (Src, Dst, Port, Action, Reason, Evidence)
 *     → Append to History
 *   → NetworkAudit.CreateSnapshot() → serialize for persistence
 *   → NetworkAudit.RestoreFromSnapshot() → deserialize and validate
 *
 * THREAD SAFETY: Uses sync.RWMutex for concurrent access.
 * ========================================================================= */
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
