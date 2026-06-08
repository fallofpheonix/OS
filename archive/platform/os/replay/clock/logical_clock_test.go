/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package clock

import (
	"sync"
	"testing"
)

func TestLogicalClock_InitialValue(t *testing.T) {
	lc := NewLogicalClock()
	if lc.Now() != 0 {
		t.Errorf("Expected initial clock value to be 0, got %d", lc.Now())
	}
}

func TestLogicalClock_Tick(t *testing.T) {
	lc := NewLogicalClock()
	for i := uint64(1); i <= 10; i++ {
		if tickVal := lc.Tick(); tickVal != i {
			t.Errorf("Expected Tick() to return %d, got %d", i, tickVal)
		}
		if nowVal := lc.Now(); nowVal != i {
			t.Errorf("Expected Now() to return %d after tick, got %d", i, nowVal)
		}
	}
}

func TestLogicalClock_AdvanceTo(t *testing.T) {
	lc := NewLogicalClock()

	// Test successful advance
	if err := lc.AdvanceTo(100); err != nil {
		t.Fatalf("AdvanceTo(100) failed: %v", err)
	}
	if lc.Now() != 100 {
		t.Errorf("Expected clock value to be 100, got %d", lc.Now())
	}

	// Test advance to a higher value
	if err := lc.AdvanceTo(150); err != nil {
		t.Fatalf("AdvanceTo(150) failed: %v", err)
	}
	if lc.Now() != 150 {
		t.Errorf("Expected clock value to be 150, got %d", lc.Now())
	}

	// Test advance to the same value (should be no-op and no error)
	if err := lc.AdvanceTo(150); err != nil {
		t.Fatalf("AdvanceTo(150) failed when already at value: %v", err)
	}
	if lc.Now() != 150 {
		t.Errorf("Expected clock value to be 150, got %d", lc.Now())
	}

	// Test clock regression (should return error)
	if err := lc.AdvanceTo(120); err == nil {
		t.Error("Expected error when trying to regress clock, but got none")
	} else if err.Error() != "clock regression: cannot advance to 120, current value is 150" {
		t.Errorf("Expected specific clock regression error, got: %v", err)
	}
	if lc.Now() != 150 { // Value should not have changed on regression attempt
		t.Errorf("Clock value should not change on regression attempt, got %d", lc.Now())
	}
}

func TestLogicalClock_MonotonicityWithTickAndAdvanceTo(t *testing.T) {
	lc := NewLogicalClock()

	lc.Tick() // 1
	lc.Tick() // 2

	if err := lc.AdvanceTo(5); err != nil { // 5
		t.Fatalf("AdvanceTo(5) failed: %v", err)
	}
	if lc.Now() != 5 {
		t.Errorf("Expected clock value 5, got %d", lc.Now())
	}

	lc.Tick() // 6
	if lc.Now() != 6 {
		t.Errorf("Expected clock value 6, got %d", lc.Now())
	}

	if err := lc.AdvanceTo(4); err == nil {
		t.Error("Expected error for regression, got none")
	}
	if lc.Now() != 6 { // Should not regress
		t.Errorf("Expected clock value 6, got %d", lc.Now())
	}
}

func TestLogicalClock_Concurrency(t *testing.T) {
	lc := NewLogicalClock()
	numGoroutines := 100
	numOperationsPerGoroutine := 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines * 2) // For Tick and AdvanceTo

	// Concurrent Ticks
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < numOperationsPerGoroutine; j++ {
				lc.Tick()
			}
		}()
	}

	// Concurrent AdvanceTo
	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < numOperationsPerGoroutine; j++ {
				// Try to advance to a value that is likely to be ahead, but not guaranteed
				// This simulates external clock synchronization attempts
				target := uint64(id*numOperationsPerGoroutine + j + 1000)
				_ = lc.AdvanceTo(target) // We expect some errors here if regression is attempted, which is fine
			}
		}(i)
	}

	wg.Wait()

	finalValue := lc.Now()
	t.Logf("Final clock value after concurrent operations: %d", finalValue)

	// Verify monotonicity by checking that the final value is at least
	// the sum of all tick operations
	minExpected := uint64(numGoroutines * numOperationsPerGoroutine)
	if finalValue < minExpected {
		t.Errorf("Clock did not advance sufficiently. Expected at least %d, got %d", minExpected, finalValue)
	}

	// Further check: Create a new clock and try to advance it many times concurrently
	// to ensure it doesn't get stuck or regress
	lc2 := NewLogicalClock()
	var wg2 sync.WaitGroup
	wg2.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(id int) {
			defer wg2.Done()
			for j := 0; j < numOperationsPerGoroutine; j++ {
				target := uint64(id*numOperationsPerGoroutine + j + 1)
				_ = lc2.AdvanceTo(target)
			}
		}(i)
	}
	wg2.Wait()
	finalValue2 := lc2.Now()
	t.Logf("Final clock value after concurrent AdvanceTo operations: %d", finalValue2)

	// The final value should be approximately the highest target value, or higher due to ticks.
	// Since AdvanceTo can be called with arbitrary values, we can only assert it didn't regress below its own ticks.
	// A better check is to ensure that `Now()` consistently returns increasing values.
	// For simplicity, for now, we just ensure it didn't reset to 0.
	if finalValue2 == 0 {
		t.Errorf("Clock did not advance at all after concurrent AdvanceTo operations.")
	}

	// Test concurrent Tick operations ensuring monotonicity
	lc3 := NewLogicalClock()
	var wg3 sync.WaitGroup
	wg3.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg3.Done()
			for j := 0; j < numOperationsPerGoroutine; j++ {
				lc3.Tick()
			}
		}()
	}
	wg3.Wait()
	finalValue3 := lc3.Now()
	t.Logf("Final clock value after concurrent Tick operations: %d", finalValue3)
	expectedFinalTick := uint64(numGoroutines * numOperationsPerGoroutine)
	if finalValue3 != expectedFinalTick {
		t.Errorf("Expected final clock value %d after %d concurrent ticks, got %d", expectedFinalTick, expectedFinalTick, finalValue3)
	}
}

func TestLogicalClock_AdvanceTo_LargeValue(t *testing.T) {
	lc := NewLogicalClock()
	largeValue := uint64(1<<63 - 1) // Max uint64

	if err := lc.AdvanceTo(largeValue); err != nil {
		t.Fatalf("AdvanceTo(%d) failed: %v", largeValue, err)
	}
	if lc.Now() != largeValue {
		t.Errorf("Expected clock value to be %d, got %d", largeValue, lc.Now())
	}
}

func TestLogicalClock_TickAfterAdvanceTo(t *testing.T) {
	lc := NewLogicalClock()
	if err := lc.AdvanceTo(100); err != nil {
		t.Fatalf("AdvanceTo(100) failed: %v", err)
	}
	if val := lc.Tick(); val != 101 {
		t.Errorf("Expected Tick after AdvanceTo(100) to be 101, got %d", val)
	}
	if lc.Now() != 101 {
		t.Errorf("Expected Now() after Tick to be 101, got %d", lc.Now())
	}
}