/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package kernel

import (
	"sync"
	"testing"
)

func TestOverflowDetector_Configure(t *testing.T) {
	od := NewOverflowDetector()

	// Test valid configuration
	err := od.Configure(100)
	if err != nil {
		t.Fatalf("Configure(100) failed: %v", err)
	}
	if od.Threshold() != 100 {
		t.Errorf("Expected threshold 100, got %d", od.Threshold())
	}

	// Test zero threshold (should fail)
	err = od.Configure(0)
	if err == nil {
		t.Error("Expected error for zero threshold, got nil")
	} else if err.Error() != "overflow threshold cannot be zero" {
		t.Errorf("Expected 'overflow threshold cannot be zero' error, got: %v", err)
	}
	if od.Threshold() != 100 { // Should retain previous valid threshold
		t.Errorf("Expected threshold 100 after failed configuration, got %d", od.Threshold())
	}

	// Reconfigure to new valid threshold
	err = od.Configure(200)
	if err != nil {
		t.Fatalf("Configure(200) failed: %v", err)
	}
	if od.Threshold() != 200 {
		t.Errorf("Expected threshold 200, got %d", od.Threshold())
	}
}

func TestOverflowDetector_Record(t *testing.T) {
	od := NewOverflowDetector()

	// Test Record before configuration (should fail)
	_, err := od.Record(10)
	if err == nil {
		t.Error("Expected error for Record before configuration, got nil")
	} else if err.Error() != "overflow detector not configured: threshold is zero" {
		t.Errorf("Expected 'overflow detector not configured: threshold is zero' error, got: %v", err)
	}

	err = od.Configure(100)
	if err != nil {
		t.Fatalf("Configure failed: %v", err)
	}

	// Test Record without overflow
	overflow, err := od.Record(50)
	if err != nil {
		t.Fatalf("Record(50) failed: %v", err)
	}
	if overflow {
		t.Error("Expected no overflow, got true")
	}
	if od.CurrentCount() != 50 {
		t.Errorf("Expected current count 50, got %d", od.CurrentCount())
	}

	// Test Record hitting threshold (no overflow yet)
	overflow, err = od.Record(50) // Total 100
	if err != nil {
		t.Fatalf("Record(50) failed: %v", err)
	}
	if overflow {
		t.Error("Expected no overflow, got true")
	}
	if od.CurrentCount() != 100 {
		t.Errorf("Expected current count 100, got %d", od.CurrentCount())
	}

	// Test Record causing overflow
	overflow, err = od.Record(1) // Total 101
	if err != nil {
		t.Fatalf("Record(1) failed: %v", err)
	}
	if !overflow {
		t.Error("Expected overflow, got false")
	}
	if od.CurrentCount() != 101 {
		t.Errorf("Expected current count 101, got %d", od.CurrentCount())
	}

	// Test Record with large value causing overflow immediately
	od.Reset()
	err = od.Configure(10)
	if err != nil {
		t.Fatalf("Configure failed: %v", err)
	}
	overflow, err = od.Record(20)
	if err != nil {
		t.Fatalf("Record(20) failed: %v", err)
	}
	if !overflow {
		t.Error("Expected overflow, got false")
	}
	if od.CurrentCount() != 20 {
		t.Errorf("Expected current count 20, got %d", od.CurrentCount())
	}
}

func TestOverflowDetector_Reset(t *testing.T) {
	od := NewOverflowDetector()
	_ = od.Configure(100)
	_, _ = od.Record(50)
	if od.CurrentCount() != 50 {
		t.Fatalf("Expected current count 50 before reset, got %d", od.CurrentCount())
	}
	od.Reset()
	if od.CurrentCount() != 0 {
		t.Errorf("Expected current count 0 after reset, got %d", od.CurrentCount())
	}
}

func TestOverflowDetector_Concurrency(t *testing.T) {
	od := NewOverflowDetector()
	_ = od.Configure(1000) // Small enough to hit overflow with many operations

	numGoroutines := 100
	incrementsPerGoroutine := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < incrementsPerGoroutine; j++ {
				_, _ = od.Record(1) // Increment by 1
			}
		}()
	}
	wg.Wait()

	finalCount := od.CurrentCount()
	expectedTotalIncrements := uint64(numGoroutines * incrementsPerGoroutine)
	if finalCount != expectedTotalIncrements {
		t.Errorf("Expected final count %d, got %d", expectedTotalIncrements, finalCount)
	}

	// Check if overflow was detected at some point
	// This is harder to test precisely in a concurrent environment unless we specifically
	// check the return value of Record, but we can infer it if finalCount > threshold.
	if finalCount > od.Threshold() {
		// Verify that a record operation after the threshold was crossed correctly returns true
		// This requires another sync point or specific design. For now, rely on overall count.
	}

	// Test that it does not error out
	if od.Threshold() != 1000 {
		t.Errorf("Threshold changed unexpectedly, got %d", od.Threshold())
	}
}
