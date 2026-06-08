/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: CYCLE 10b — ROLLBACK AUDIT (Layer 5)
 *
 * The RollbackAudit manages the history of all cross-component rollback
 * events. Each rollback is hashed and sequenced for tamper-evident audit.
 *
 * WORKFLOW:
 *   RollbackAudit.LogRollback(record)
 *     → Compute deterministic hash (Component, SnapshotID, RecoveryID, Evidence)
 *     → Append to History
 *   → RollbackAudit.CreateSnapshot() → serialize for persistence
 *   → RollbackAudit.RestoreFromSnapshot() → deserialize and validate
 *
 * PURPOSE: Provides an audit trail for all rollback operations.
 * If a rollback is attempted but fails, the audit record provides
 * the forensic evidence needed for investigation.
 *
 * THREAD SAFETY: Uses sync.RWMutex for concurrent access.
 * ========================================================================= */
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
		Component  Component
		SnapshotID string
		RecoveryID string
		EvidenceID string
		DecisionID string
		Sequence   int
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
