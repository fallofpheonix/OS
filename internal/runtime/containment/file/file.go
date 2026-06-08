/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9g — FILE CONTAINMENT PRIMITIVES (Layer 4)
//
// This file defines the primitive actions available for file containment.
// These are the BUILDING BLOCKS that the SandboxWarden uses to implement
// file system isolation.
//
// PRIMITIVES:
//   MONITOR_FILE    → Observe file access without interference
//   THROTTLE_WRITE  → Rate-limit file write operations
//   FREEZE_PATH     → Block all access to a specific path
//   ISOLATE_WORKSPACE → Restrict process to a sandboxed directory
//   RESTORE_FILE_ACCESS → Remove file restrictions
//
// Each primitive maps to a specific kernel operation:
//   MONITOR_FILE    → fanotify/inotify watch
//   THROTTLE_WRITE  → cgroup I/O limit
//   FREEZE_PATH     → LSM hooks (security_file_open)
//   ISOLATE_WORKSPACE → chroot/overlayfs
//   RESTORE_FILE_ACCESS → remove watches and limits
// =========================================================================
package file

import (
	"time"
)

type FileActionType string

const (
	ActionMonitor  FileActionType = "MONITOR_FILE"
	ActionThrottle FileActionType = "THROTTLE_WRITE"
	ActionFreeze   FileActionType = "FREEZE_PATH"
	ActionIsolate  FileActionType = "ISOLATE_WORKSPACE"
	ActionRestore  FileActionType = "RESTORE_FILE_ACCESS"
)

// FileAction defines the safe primitive set for file containment.
type FileAction struct {
	Path       string
	Action     FileActionType
	Reason     string
	EvidenceID string
	DecisionID string
	Timestamp  time.Time
	Sequence   int
	Hash       string
}
