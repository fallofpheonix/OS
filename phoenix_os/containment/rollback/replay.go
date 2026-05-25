package rollback

import (
	"time"
)

// RollbackReplay defines the structure for replaying cross-component rollback actions.
type RollbackReplay struct {
	Component     Component
	Record        RollbackRecord
	ReplayCursor  int
	SnapshotID    string
	RecoveryID    string
	PreviousState string
	CurrentState  string
	Sequence      int
	Hash          string
	EvidenceID    string
	DecisionID    string
	Timestamp     time.Time
}
