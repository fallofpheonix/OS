/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: TYPE DEFINITIONS — FILE REPLAY STRUCTURE
 *
 * FileReplay defines the structure for replaying file containment actions.
 * It captures the state before and after each action for deterministic replay.
 *
 * WORKFLOW:
 *   1. Record file action with previous/current state
 *   2. Store replay cursor for sequential playback
 *   3. Include snapshot and recovery IDs for audit trail
 *   4. Hash the replay record for integrity verification
 *
 * PURPOSE: Enables deterministic replay of file containment operations.
 * If the system needs to replay a sequence of file actions, this structure
 * provides all the information needed for faithful reconstruction.
 * ========================================================================= */
package file

import (
	"time"
)

// FileReplay defines the structure for replaying file containment actions.
type FileReplay struct {
	Path          string
	Action        FileAction
	PreviousState string
	CurrentState  string
	ReplayCursor  int
	SnapshotID    string
	RecoveryID    string
	Sequence      int
	Hash          string
	EvidenceID    string
	DecisionID    string
	Timestamp     time.Time
}
