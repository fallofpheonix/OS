package clock

import (
	"fmt"
)

// ClockSynchronizer defines the interface for synchronizing logical clocks.
type ClockSynchronizer interface {
	// SyncTo advances the target clock to the value of the source clock,
	// or returns an error if the target clock would regress.
	SyncTo(target LogicalClock, source LogicalClock) error
	// AdvanceTargetTo advances the target clock to a specific value,
	// or returns an error if the target clock would regress.
	AdvanceTargetTo(target LogicalClock, value uint64) error
}

// NewClockSynchronizer creates a new instance of ClockSynchronizer.
func NewClockSynchronizer() ClockSynchronizer {
	return &simpleClockSynchronizer{}
}

type simpleClockSynchronizer struct{}

// SyncTo advances the target clock to the value of the source clock.
func (cs *simpleClockSynchronizer) SyncTo(target LogicalClock, source LogicalClock) error {
	if target == nil {
		return fmt.Errorf("target clock cannot be nil")
	}
	if source == nil {
		return fmt.Errorf("source clock cannot be nil")
	}
	return target.AdvanceTo(source.Now())
}

// AdvanceTargetTo advances the target clock to a specific value.
func (cs *simpleClockSynchronizer) AdvanceTargetTo(target LogicalClock, value uint64) error {
	if target == nil {
		return fmt.Errorf("target clock cannot be nil")
	}
	return target.AdvanceTo(value)
}
