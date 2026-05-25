package rollback

import (
	"time"
)

type Component string

const (
	ComponentIsolation  Component = "ISOLATION"
	ComponentProcess    Component = "PROCESS"
	ComponentNetwork    Component = "NETWORK"
	ComponentFile       Component = "FILE"
)

// RollbackRecord binds a cross-component rollback event to snapshot evidence.
type RollbackRecord struct {
	Component     Component
	PreviousState string
	CurrentState  string
	SnapshotID    string
	RecoveryID    string
	EvidenceID    string
	DecisionID    string
	Sequence      int
	Hash          string
	Timestamp     time.Time
}
