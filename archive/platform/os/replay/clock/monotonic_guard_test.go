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

func TestMonotonicGuard_InitialValue(t *testing.T) {
	mg := NewMonotonicGuard(100)
	if mg.CurrentValue() != 100 {
		t.Errorf("Expected initial value 100, got %d", mg.CurrentValue())
	}
}

func TestMonotonicGuard_CheckAndSet_Success(t *testing.T) {
	mg := NewMonotonicGuard(10)

	err := mg.CheckAndSet(20)
	if err != nil {
		t.Fatalf("CheckAndSet(20) failed: %v", err)
	}
	if mg.CurrentValue() != 20 {
		t.Errorf("Expected value 20, got %d", mg.CurrentValue())
	}

	err = mg.CheckAndSet(30)
	if err != nil {
		t.Fatalf("CheckAndSet(30) failed: %v", err)
	}
	if mg.CurrentValue() != 30 {
		t.Errorf("Expected value 30, got %d", mg.CurrentValue())
	}
}

func TestMonotonicGuard_CheckAndSet_Regression(t *testing.T) {
	mg := NewMonotonicGuard(50)

	// Test smaller value
	err := mg.CheckAndSet(40)
	if err == nil {
		t.Error("Expected error for smaller value, got nil")
	} else if err.Error() != "monotonicity violation: new value 40 is not strictly greater than current value 50" {
		t.Errorf("Expected specific error for regression, got: %v", err)
	}
	if mg.CurrentValue() != 50 { // Value should not change on regression
		t.Errorf("Expected value 50, got %d", mg.CurrentValue())
	}

	// Test same value
	err = mg.CheckAndSet(50)
	if err == nil {
		t.Error("Expected error for same value, got nil")
	} else if err.Error() != "monotonicity violation: new value 50 is not strictly greater than current value 50" {
		t.Errorf("Expected specific error for same value, got: %v", err)
	}
	if mg.CurrentValue() != 50 { // Value should not change
		t.Errorf("Expected value 50, got %d", mg.CurrentValue())
	}
}

func TestMonotonicGuard_Concurrency(t *testing.T) {
	mg := NewMonotonicGuard(0)
	numGoroutines := 100
	incrementsPerGoroutine := 1000

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	// Each goroutine tries to set increasing values
	for i := 0; i < numGoroutines; i++ {
		go func(gid int) {
			defer wg.Done()
			startVal := uint64(gid * incrementsPerGoroutine)
			for j := 0; j < incrementsPerGoroutine; j++ {
				newValue := startVal + uint64(j)
				// We don't expect all these to succeed, as some will be less than current global max.
				// The goal is to ensure no panics and final value is consistent.
				_ = mg.CheckAndSet(newValue)
			}
		}(i)
	}
	wg.Wait()

	// The final value should be the maximum value successfully set
	finalValue := mg.CurrentValue()
	t.Logf("Final monotonic value: %d", finalValue)

	// The final value must be greater than 0 (initial)
	if finalValue == 0 {
		t.Error("Expected final value to be greater than 0")
	}

	// Any value set must be strictly greater than previous
	// This is verified by the CheckAndSet logic itself.
	// A good check here is to try to set a value just below finalValue
	// and ensure it fails.
	if finalValue > 0 {
		err := mg.CheckAndSet(finalValue)
		if err == nil {
			t.Errorf("Expected error when setting to current finalValue %d, got nil", finalValue)
		}
		err = mg.CheckAndSet(finalValue - 1)
		if err == nil {
			t.Errorf("Expected error when setting to finalValue-1 %d, got nil", finalValue-1)
		}
		err = mg.CheckAndSet(finalValue + 1)
		if err != nil {
			t.Errorf("Expected success when setting to finalValue+1 %d, got %v", finalValue+1, err)
		}
	}
}
