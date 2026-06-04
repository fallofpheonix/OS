package v1

import (
	"context"
)

// ContainmentLevel represents a level of containment on the containment ladder.
type ContainmentLevel int

const (
	LevelNone ContainmentLevel = iota
	LevelMonitor
	LevelSandbox
	LevelIsolate
	LevelQuench
)

// Containment defines the interface for containment policies/decisions.
type Containment interface {
	Target() string
	Level() ContainmentLevel
	Reason() string
}

// Escalation defines the interface for raising containment severity.
type Escalation interface {
	CurrentLevel() ContainmentLevel
	TargetLevel() ContainmentLevel
	TriggerReason() string
}

// Actuator defines the canonical interface for enforcing security policies (e.g. wardens).
type Actuator interface {
	Actuate(ctx context.Context, action Containment) error
	Kill(ctx context.Context, pid int) error
	GetCurrentLevel() (ContainmentLevel, error)
}
