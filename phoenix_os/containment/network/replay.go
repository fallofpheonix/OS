package network

import (
	"time"
)

// NetworkReplay defines the structure for replaying network containment actions.
type NetworkReplay struct {
	ConnectionID   string
	Action         NetworkAction
	PreviousState  string
	CurrentState   string
	ReplayCursor   int
	SnapshotID     string
	RecoveryID     string
	Sequence       int
	Hash           string
	EvidenceID     string
	DecisionID     string
	Timestamp      time.Time
}
