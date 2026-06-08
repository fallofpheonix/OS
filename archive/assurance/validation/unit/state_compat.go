/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/state"
)

type SystemState = state.SystemState
type Snapshot = state.Snapshot
type StateRegistry = state.StateRegistry
type StateRecord = state.StateRecord

const (
	StateSafe     = state.StateSafe
	StateWatch    = state.StateWatch
	StateAlert    = state.StateAlert
	StateContain  = state.StateContain
	StateRecovery = state.StateRecovery
)

var GlobalMetrics = &state.GlobalMetrics

func NewStateRegistry(initial SystemState) *StateRegistry {
	return state.NewStateRegistry(initial)
}
