/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: TYPE DEFINITIONS — ROLLBACK COMPONENT TYPES
 *
 * This file defines the Component types and RollbackRecord structure
 * for cross-component rollback tracking.
 *
 * COMPONENTS:
 *   ISOLATION → Containment engine state
 *   PROCESS   → Process audit trail
 *   NETWORK   → Network audit trail
 *   FILE      → File audit trail
 *
 * WORKFLOW:
 *   RollbackRecord created with component, states, snapshot/recovery IDs
 *   → RollbackAudit.LogRollback(record) → audit trail
 *   → RollbackOrchestrator.RestoreGlobal() → cross-component restore
 *
 * PURPOSE: Tracks which component is being rolled back and provides
 * the evidence chain for forensic analysis.
 * ========================================================================= */
package rollback

import (
	"time"
)

type Component string

const (
	ComponentIsolation Component = "ISOLATION"
	ComponentProcess   Component = "PROCESS"
	ComponentNetwork   Component = "NETWORK"
	ComponentFile      Component = "FILE"
)

// RollbackRecord binds a cross-component rollback event to snapshot evidence.
type RollbackRecord struct {
	Component     Component
	PreviousState string
	CurrentState  string
	SnapshotID    string
	RecoveryID    string
	EvidenceID    string
	DecisionID    string
	Sequence      int
	Hash          string
	Timestamp     time.Time
}
