/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package clock

import (
	"fmt"
	"sync/atomic"
)

// MonotonicGuard defines an interface for ensuring monotonicity of a sequence.
type MonotonicGuard interface {
	// CheckAndSet atomically updates the current value to `newValue` if `newValue` is
	// strictly greater than the current value. If `newValue` is less than or equal
	// to the current value, it returns an error indicating regression.
	CheckAndSet(newValue uint64) error
	// CurrentValue returns the currently held monotonic value.
	CurrentValue() uint64
}

// NewMonotonicGuard creates a new instance of MonotonicGuard with an initial value.
func NewMonotonicGuard(initialValue uint64) MonotonicGuard {
	mg := &simpleMonotonicGuard{}
	mg.currentValue.Store(initialValue)
	return mg
}

type simpleMonotonicGuard struct {
	currentValue atomic.Uint64
}

// CheckAndSet atomically updates the current value.
func (mg *simpleMonotonicGuard) CheckAndSet(newValue uint64) error {
	for {
		current := mg.currentValue.Load()
		if newValue <= current {
			return fmt.Errorf("monotonicity violation: new value %d is not strictly greater than current value %d", newValue, current)
		}
		if mg.currentValue.CompareAndSwap(current, newValue) {
			return nil
		}
	}
}

// CurrentValue returns the currently held monotonic value.
func (mg *simpleMonotonicGuard) CurrentValue() uint64 {
	return mg.currentValue.Load()
}
