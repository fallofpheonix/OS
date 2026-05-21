package types

import "time"

// TelemetryEvent conforms to RFC-001 Common Header and Category-specific payloads
type TelemetryEvent struct {
	Timestamp   time.Time `json:"timestamp"`
	EventID     string    `json:"event_id"`
	Category    string    `json:"category"` // process | syscall | filesystem | network | container | memory
	EventType   string    `json:"event_type"`
	HostID      string    `json:"host_id"`
	PID         uint32    `json:"pid"`
	PPID        uint32    `json:"ppid"`
	UID         uint32    `json:"uid"`
	GID         uint32    `json:"gid"`
	Comm        string    `json:"comm"`
	ExePath     string    `json:"exe_path"`
	ContainerID string    `json:"container_id,omitempty"`

	// Category Payloads (only one is populated per event type)
	Process    *ProcessPayload    `json:"process,omitempty"`
	Syscall    *SyscallPayload    `json:"syscall,omitempty"`
	Filesystem *FilesystemPayload `json:"filesystem,omitempty"`
	Network    *NetworkPayload    `json:"network,omitempty"`
	Container  *ContainerPayload  `json:"container,omitempty"`
	Memory     *MemoryPayload     `json:"memory,omitempty"`
}

type ProcessPayload struct {
	Args     []string `json:"args"`
	EnvVars  []string `json:"env_vars"`
	ExitCode int32    `json:"exit_code,omitempty"`
}

type SyscallPayload struct {
	SyscallNr uint64   `json:"syscall_nr"`
	Args      []uint64 `json:"args"`
	RetVal    int64    `json:"retval"`
}

type FilesystemPayload struct {
	FilePath         string `json:"file_path"`
	Flags            uint32 `json:"flags"`
	Mode             uint32 `json:"mode"`
	BytesRequested   uint64 `json:"bytes_requested"`
	BytesTransferred int64  `json:"bytes_transferred"`
}

type NetworkPayload struct {
	SAddr       string `json:"saddr"`
	DAddr       string `json:"daddr"`
	SPort       uint16 `json:"sport"`
	DPort       uint16 `json:"dport"`
	Protocol    string `json:"protocol"` // TCP | UDP | RAW
	StateChange string `json:"state_change"`
}

type ContainerPayload struct {
	NamespacePID uint32 `json:"namespace_pid"`
	NamespaceNet uint32 `json:"namespace_net"`
	NamespaceMnt uint32 `json:"namespace_mnt"`
	CgroupPath   string `json:"cgroup_path"`
}

type MemoryPayload struct {
	Address    uint64 `json:"address"`
	Length     uint64 `json:"length"`
	Protection uint32 `json:"protection"`
	Flags      uint32 `json:"flags"`
}

// IncidentGraph represents the current attack DAG or incident graph structure (L4)
type IncidentGraph struct {
	Nodes map[string]*ProcessNode `json:"nodes"`
	Edges map[string][]string     `json:"edges"`
}

type ProcessNode struct {
	PID         uint32    `json:"pid"`
	Comm        string    `json:"comm"`
	ExePath     string    `json:"exe_path"`
	Centrality  float64   `json:"centrality"`
	ThreatScore float64   `json:"threat_score"`
	LastSeen    time.Time `json:"last_seen"`
}

// SecurityState represents physical threat states (L6)
type SecurityState struct {
	Timestamp        time.Time `json:"timestamp"`
	Entropy          float64   `json:"entropy"`
	KLDivergence     float64   `json:"kl_divergence"`
	ThreatTemperature float64   `json:"threat_temperature"`
	SDI              float64   `json:"sdi"`
	IsAnomaly        bool      `json:"is_anomaly"`
}

// Strategy represents strategic game recommendations (L6.5)
type Strategy struct {
	ContainmentLevel int       `json:"containment_level"` // Use ContainmentLevel constants
	TargetPIDs       []uint32  `json:"target_pids"`
	StrategyType     string    `json:"strategy_type"` // Nash | Stackelberg | Bayesian
	Timestamp        time.Time `json:"timestamp"`
}

const (
	LevelObserve    = 0 // SAFE / WATCH
	LevelLimit      = 1 // SUSPICIOUS
	LevelThrottled  = 2 // SUSPICIOUS
	LevelFreeze     = 3 // CRITICAL
	LevelIsolate    = 4 // CRITICAL
	LevelKill       = 5 // COMPROMISED
)

// PIDMetrics represents control feedback metrics (L5)
type PIDMetrics struct {
	Setpoint  float64 `json:"setpoint"`
	Measured  float64 `json:"measured"`
	Output    float64 `json:"output"`
	LastError float64 `json:"last_error"`
	Integral  float64 `json:"integral"`
}

// SnapshotResult represents forensic snapshot metadata (DFIR)
type SnapshotResult struct {
	SnapshotID string    `json:"snapshot_id"`
	PID        uint32    `json:"pid"`
	Hash       string    `json:"hash"`
	Path       string    `json:"path"`
	Timestamp  time.Time `json:"timestamp"`
}
