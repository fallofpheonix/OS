package network

import (
	"time"
)

type NetworkActionType string

const (
	ActionMonitor   NetworkActionType = "MONITOR_CONNECTION"
	ActionThrottle  NetworkActionType = "THROTTLE_BANDWIDTH"
	ActionPause     NetworkActionType = "PAUSE_EGRESS"
	ActionQuarantine NetworkActionType = "QUARANTINE_NAMESPACE"
	ActionRestore    NetworkActionType = "RESTORE_NETWORK"
)

// NetworkAction defines the safe primitive set for network quarantine.
type NetworkAction struct {
	Src        string
	Dst        string
	Port       int
	Action     NetworkActionType
	Reason     string
	EvidenceID string
	DecisionID string
	Timestamp  time.Time
	Sequence   int
	Hash       string
}
