package file

import (
	"time"
)

// FileReplay defines the structure for replaying file containment actions.
type FileReplay struct {
	Path           string
	Action         FileAction
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
