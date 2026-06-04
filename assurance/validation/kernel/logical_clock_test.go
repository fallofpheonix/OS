/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package kernel

import (
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/logical_clock"
)

func TestLogicalClockMonotonicity(t *testing.T) {
	c := logical_clock.NewClock()
	last := uint64(0)
	for i := 0; i < 1000; i++ {
		curr := c.Tick()
		if curr <= last {
			t.Errorf("Clock not monotonic: %d followed %d", curr, last)
		}
		last = curr
	}
	fmt.Println("[PX-001] Logical Clock Monotonicity: PASSED")
}

func TestLogicalClockJump(t *testing.T) {
	c := logical_clock.NewClock()
	c.AdvanceTo(5000)
	if c.Now() != 5000 {
		t.Errorf("Clock AdvanceTo failed: expected 5000, got %d", c.Now())
	}

	c.AdvanceTo(1000)
	if c.Now() != 5000 {
		t.Errorf("Clock backward AdvanceTo should be ignored: expected 5000, got %d", c.Now())
	}
	fmt.Println("[PX-001] Logical Clock Jump: PASSED")
}

func TestReplayClockSync(t *testing.T) {
	// Simulated replay sync test
	runtimeClock := logical_clock.NewClock()
	replayClock := logical_clock.NewClock()

	// Replay event at tick 100
	eventTick := uint64(100)
	replayClock.AdvanceTo(eventTick)
	runtimeClock.AdvanceTo(replayClock.Now())

	if runtimeClock.Now() != 100 {
		t.Errorf("Runtime clock sync failed: expected 100, got %d", runtimeClock.Now())
	}
	fmt.Println("[PX-001] Replay Clock Sync: PASSED")
}
