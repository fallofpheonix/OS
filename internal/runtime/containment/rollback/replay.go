/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package rollback

import (
	"time"
)

// RollbackReplay defines the structure for replaying cross-component rollback actions.
type RollbackReplay struct {
	Component     Component
	Record        RollbackRecord
	ReplayCursor  int
	SnapshotID    string
	RecoveryID    string
	PreviousState string
	CurrentState  string
	Sequence      int
	Hash          string
	EvidenceID    string
	DecisionID    string
	Timestamp     time.Time
}
