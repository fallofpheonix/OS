package containment

import (
	"time"
)

type ProcessActionType string

const (
	ActionMonitor   ProcessActionType = "MONITOR"
	ActionThrottle  ProcessActionType = "THROTTLE"
	ActionPause     ProcessActionType = "PAUSE"
	ActionIsolate   ProcessActionType = "ISOLATE"
	ActionRecover   ProcessActionType = "RECOVER"
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
