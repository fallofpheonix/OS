package logical_clock

import (
	"sync/atomic"
)

// Clock provides a deterministic monotonic counter for logical time.
type Clock struct {
	tick uint64
}

// NewClock initializes a clock starting at 0.
func NewClock() *Clock {
	return &Clock{tick: 0}
}

// Tick increments the logical time and returns the new value.
func (c *Clock) Tick() uint64 {
	return atomic.AddUint64(&c.tick, 1)
}

// Now returns the current logical time without incrementing it.
func (c *Clock) Now() uint64 {
	return atomic.LoadUint64(&c.tick)
}

// AdvanceTo ensures the clock is at least at the specified tick.
// Useful for syncing clocks from received events.
func (c *Clock) AdvanceTo(target uint64) {
	for {
		current := atomic.LoadUint64(&c.tick)
		if current >= target {
			break
		}
		if atomic.CompareAndSwapUint64(&c.tick, current, target) {
			break
		}
	}
}
