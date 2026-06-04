/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — LAMPORT LOGICAL CLOCK
//
// The Clock provides a deterministic monotonic counter for logical time.
// It's used by the Normalizer to assign sequence IDs to events and by
// the Replay Engine to verify temporal ordering.
//
// WORKFLOW:
//   Normalizer.Normalize(raw) → clock.Tick() → monotonic SeqID
//   ReplayEngine.Apply(event) → verify event.MonotonicTime >= lastMonotonic
//   Clock.AdvanceTo(target) → sync clock to received event time
//
// ALGORITHM: Lamport clock — monotonically increasing counter.
// Each Tick() increments by 1. AdvanceTo() jumps to the maximum of
// current and target (ensuring monotonicity).
//
// THREAD SAFETY: Uses sync/atomic for lock-free concurrent access.
// Multiple goroutines can call Tick() simultaneously without races.
//
// LIMITATION: Lamport clocks capture causality within a single node.
// For multi-node operation, vector clocks would be needed.
// =========================================================================
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
