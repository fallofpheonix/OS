/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9c — PROCESS CONTAINMENT PRIMITIVES (Layer 4)
//
// This file defines the primitive actions available for process containment.
// These are the BUILDING BLOCKS that the SandboxWarden uses to implement
// process isolation.
//
// PRIMITIVES:
//   MONITOR  → Observe process without interference
//   THROTTLE → Rate-limit process CPU/IO
//   PAUSE    → Suspend process execution (cgroup freeze)
//   ISOLATE  → Full containment (freeze + namespace sever)
//   RECOVER  → Restore process to normal operation
//
// Each primitive maps to a specific kernel operation:
//   MONITOR  → /proc/<pid>/status monitoring
//   THROTTLE → cgroup cpu.max limit
//   PAUSE    → cgroup freezer.state = F
//   ISOLATE  → cgroup freeze + setns to blackhole namespace
//   RECOVER  → cgroup thaw + restore network namespace
// =========================================================================
package containment

import (
	"time"
)

type ProcessActionType string

const (
	ActionMonitor  ProcessActionType = "MONITOR"
	ActionThrottle ProcessActionType = "THROTTLE"
	ActionPause    ProcessActionType = "PAUSE"
	ActionIsolate  ProcessActionType = "ISOLATE"
	ActionRecover  ProcessActionType = "RECOVER"
)

// ProcessAction defines the safe primitive set for process containment.
type ProcessAction struct {
	PID        int
	Action     ProcessActionType
	Reason     string
	EvidenceID string
	DecisionID string
	Timestamp  time.Time
	Sequence   int
	Hash       string
}
