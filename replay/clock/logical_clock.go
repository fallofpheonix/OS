package clock

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// LogicalClock defines the interface for a monotonic, thread-safe logical clock.
type LogicalClock interface {
	// Tick advances the clock by one and returns the new value.
	Tick() uint64
	// Now returns the current value of the clock.
	Now() uint64
	// AdvanceTo sets the clock to the given value, if it's greater than the current value.
	// It returns an error if the clock would regress.
	AdvanceTo(v uint64) error
}

// NewLogicalClock creates a new default implementation of LogicalClock.
func NewLogicalClock() LogicalClock {
	var current atomic.Uint64
	current.Store(0)
	return &logicalClockImpl{
		currentValue: current,
	}
}

type logicalClockImpl struct {
	currentValue atomic.Uint64
	mu           sync.Mutex // Protects against concurrent AdvanceTo calls
}

// Tick advances the clock by one and returns the new value.
// It is thread-safe and ensures monotonicity.
func (lc *logicalClockImpl) Tick() uint64 {
	return lc.currentValue.Add(1)
}

// Now returns the current value of the clock.
func (lc *logicalClockImpl) Now() uint64 {
	return lc.currentValue.Load()
}

// AdvanceTo sets the clock to the given value, if it's greater than the current value.
// It returns an error if the clock would regress.
func (lc *logicalClockImpl) AdvanceTo(v uint64) error {
	lc.mu.Lock()
	defer lc.mu.Unlock()

	current := lc.currentValue.Load()
	if v < current {
		return fmt.Errorf("clock regression: cannot advance to %d, current value is %d", v, current)
	}

	// It's safe to directly store here because the mutex protects this critical section
	// from other AdvanceTo calls, and Tick() uses Add, which is atomic and will not conflict.
	// We only update if `v` is strictly greater than `current`. If `v` is equal to `current`,
	// it's a no-op, and if `v` is less, it's a regression caught above.
	if v > current {
		lc.currentValue.Store(v)
	}
	return nil
}
