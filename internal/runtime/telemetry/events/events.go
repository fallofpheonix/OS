/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: TYPE DEFINITIONS — NORMALIZED EVENT FORMAT
//
// This file defines the canonical Event structure used throughout the
// telemetry pipeline. All events entering the Bus must conform to this format.
//
// WORKFLOW: Raw kernel events → Normalizer → Event struct → Bus distribution
//
// FIELDS:
//   Timestamp: When the event occurred (wall clock)
//   EventID: Unique identifier (assigned by normalizer or eBPF enforcer)
//   Category: Event classification (process, syscall, filesystem, network)
//   EventType: Specific event type (execve, write, open, etc.)
//   HostID: Which node generated this event
//   PID/PPID/UID/GID: Process identity
//   Comm/ExePath: Process name and executable path
//   Payload: Category-specific data (args, file_path, etc.)
// =========================================================================
package events

import "time"

// Event represents a normalized telemetry event.
type Event struct {
	Timestamp time.Time      `json:"timestamp"`
	EventID   string         `json:"event_id"`
	Category  string         `json:"category"`
	EventType string         `json:"event_type"`
	HostID    string         `json:"host_id"`
	PID       uint32         `json:"pid"`
	PPID      uint32         `json:"ppid"`
	UID       uint32         `json:"uid"`
	GID       uint32         `json:"gid"`
	Comm      string         `json:"comm"`
	ExePath   string         `json:"exe_path"`
	Payload   map[string]any `json:"payload"`
}
