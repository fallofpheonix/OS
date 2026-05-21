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
