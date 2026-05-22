package logical_clock

import (
	"sync/atomic"
)

// LogicalClock provides a monotonic sequence of ticks to replace wall-clock time
// in the deterministic execution path.
type LogicalClock struct {
	tick uint64
}

// NewLogicalClock creates a new clock instance initialized at 0.
func NewLogicalClock() *LogicalClock {
	return &LogicalClock{tick: 0}
}

// Tick increments the logical clock and returns the new value.
func (c *LogicalClock) Tick() uint64 {
	return atomic.AddUint64(&c.tick, 1)
}

// Current returns the current logical tick without incrementing it.
func (c *LogicalClock) Current() uint64 {
	return atomic.LoadUint64(&c.tick)
}
