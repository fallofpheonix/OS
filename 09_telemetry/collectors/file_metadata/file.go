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
