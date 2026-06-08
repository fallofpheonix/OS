/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package state

import "time"

// StateRecord captures a single transition with audit metadata.
type StateRecord struct {
	ID         int
	Previous   SystemState
	Current    SystemState
	Timestamp  time.Time
	Reason     string
	EvidenceID string
}
