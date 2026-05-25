package contracts

import (
	"fmt"
	"time"
)

// Version defines the semantic versioning for PhoenixOS components.
type Version struct {
	Major      int
	Minor      int
	Patch      int
	PreRelease string
	Build      string
	ReleasedAt time.Time
}

func (v Version) String() string {
	if v.PreRelease != "" {
		return fmt.Sprintf("%d.%d.%d-%s", v.Major, v.Minor, v.Patch, v.PreRelease)
	}
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}

func CompatibilityCheck(v Version) bool {
	// Simple rule for test: Major 0 is compatible, Major 1 is not compatible with 0
	return v.Major == 0
}

const (
	// CurrentAPILevel defines the compatibility layer.
	CurrentAPILevel = 1
	// DefaultContractHash is the SHA-256 root of the contract definition.
	DefaultContractHash = "phx_root_0x1234"
)

// ActuationClass defines the severity and type of response action.
type ActuationClass int

const (
	ClassObserve ActuationClass = iota
	ClassLog
	ClassThrottle
	ClassLocalIsolate
	ClassClusterIsolate
	ClassKernelEmergency
)

// SystemState defines the state for the system.
type SystemState string

const (
	StateSafe     SystemState = "SAFE"
	StateWatch    SystemState = "WATCH"
	StateAlert    SystemState = "ALERT"
	StateContain  SystemState = "CONTAIN"
	StateRecovery SystemState = "RECOVERY"
)

// PolicyContext provides the necessary state for a policy evaluation.
type PolicyContext struct {
	ID        string
	Timestamp time.Time
	Evidence  []interface{}
	Metadata  map[string]string
}

// CompatibilityMatrix tracks breaking changes between versions.
type CompatibilityMatrix struct {
	SupportedLevels []int
	MinVersion      Version
}

// DeprecationEntry tracks deprecated interfaces.
type DeprecationEntry struct {
	Name        string
	Deprecated  Version
	Removal     Version
	Alternative string
}
