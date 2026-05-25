package file

import (
	"time"
)

type FileActionType string

const (
	ActionMonitor   FileActionType = "MONITOR_FILE"
	ActionThrottle  FileActionType = "THROTTLE_WRITE"
	ActionFreeze    FileActionType = "FREEZE_PATH"
	ActionIsolate   FileActionType = "ISOLATE_WORKSPACE"
	ActionRestore   FileActionType = "RESTORE_FILE_ACCESS"
)

// FileAction defines the safe primitive set for file containment.
type FileAction struct {
	Path       string
	Action     FileActionType
	Reason     string
	EvidenceID string
	DecisionID string
	Timestamp  time.Time
	Sequence   int
	Hash       string
}
