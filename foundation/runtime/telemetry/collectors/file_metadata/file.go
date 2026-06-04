/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: TYPE DEFINITIONS — FILESYSTEM EVENT FORMAT
 *
 * This file defines the canonical FilesystemEvent and FilesystemPayload
 * structures used by the eBPF filesystem monitoring probes.
 *
 * WORKFLOW:
 *   eBPF probe captures filesystem event → FilesystemEvent struct
 *     → Normalizer.Normalize() → bus.TelemetryEvent
 *       → Bus distributes to Monitor, Graph, TCS
 *
 * FIELDS:
 *   FilePath: Full path to the accessed file
 *   Flags: Open flags (O_RDONLY, O_WRONLY, etc.)
 *   Mode: File permissions (0644, 0755, etc.)
 *   BytesRequested/BytesTransferred: I/O sizes
 *   EntropyScore: Shannon entropy of file contents (for ransomware detection)
 *
 * SECURITY: High entropy in file contents (> 7.5) may indicate encryption.
 * Combined with write operations, this is a strong ransomware indicator.
 * ========================================================================= */
package file_metadata

import "time"

// FilesystemPayload contains category-specific fields for filesystem events
type FilesystemPayload struct {
	FilePath         string  `json:"file_path"`
	Flags            uint32  `json:"flags"`
	Mode             uint32  `json:"mode"`
	BytesRequested   uint64  `json:"bytes_requested"`
	BytesTransferred int64   `json:"bytes_transferred"`
	EntropyScore     float64 `json:"entropy_score,omitempty"`
}

// FilesystemEvent represents a system filesystem event complying with the schema
type FilesystemEvent struct {
	Timestamp   time.Time         `json:"timestamp"`
	EventID     string            `json:"event_id"`
	Category    string            `json:"category"`   // "filesystem"
	EventType   string            `json:"event_type"` // e.g. "open", "write", "unlink"
	HostID      string            `json:"host_id"`
	PID         uint32            `json:"pid"`
	PPID        uint32            `json:"ppid"`
	UID         uint32            `json:"uid"`
	GID         uint32            `json:"gid"`
	Comm        string            `json:"comm"`
	ExePath     string            `json:"exe_path"`
	ContainerID string            `json:"container_id,omitempty"`
	Payload     FilesystemPayload `json:"payload"`
}
