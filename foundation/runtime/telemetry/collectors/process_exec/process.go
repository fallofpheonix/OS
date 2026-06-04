/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: TYPE DEFINITIONS — PROCESS EVENT FORMAT
 *
 * This file defines the canonical ProcessEvent and ProcessPayload
 * structures used by the eBPF process monitoring probes.
 *
 * WORKFLOW:
 *   eBPF probe captures process event → ProcessEvent struct
 *     → Normalizer.Normalize() → bus.TelemetryEvent
 *       → Bus distributes to Monitor, Graph, TCS
 *
 * FIELDS:
 *   Args: Command-line arguments
 *   EnvVars: Environment variables (may contain secrets!)
 *   ExitCode: Process exit code (nil if still running)
 *
 * SECURITY: EnvVars may contain API keys, passwords, or tokens.
 * The normalizer should redact sensitive environment variables
 * before publishing to the Bus. Currently, no redaction is performed.
 * ========================================================================= */
package process_exec

import "time"

// ProcessPayload contains category-specific fields for process events
type ProcessPayload struct {
	Args     []string `json:"args"`
	EnvVars  []string `json:"env_vars"`
	ExitCode *int32   `json:"exit_code,omitempty"`
}

// ProcessEvent represents a system process event complying with the schema
type ProcessEvent struct {
	Timestamp   time.Time      `json:"timestamp"`
	EventID     string         `json:"event_id"`
	Category    string         `json:"category"`   // "process"
	EventType   string         `json:"event_type"` // e.g. "fork", "execve", "exit"
	HostID      string         `json:"host_id"`
	PID         uint32         `json:"pid"`
	PPID        uint32         `json:"ppid"`
	UID         uint32         `json:"uid"`
	GID         uint32         `json:"gid"`
	Comm        string         `json:"comm"`
	ExePath     string         `json:"exe_path"`
	ContainerID string         `json:"container_id,omitempty"`
	Payload     ProcessPayload `json:"payload"`
}
