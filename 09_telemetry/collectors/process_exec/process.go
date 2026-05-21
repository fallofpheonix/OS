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
