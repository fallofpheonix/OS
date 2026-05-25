package kernel

import (
	"sync"
	"testing"
)

func TestRingMonitor_StartStopMonitoring(t *testing.T) {
	rm := NewRingMonitor()

	// Test starting with valid capacity
	err := rm.StartMonitoring(10)
	if err != nil {
		t.Fatalf("StartMonitoring failed: %v", err)
	}

	// Test starting again (should fail)
	err = rm.StartMonitoring(10)
	if err == nil {
		t.Error("Expected error when starting an already running monitor, got nil")
	} else if err.Error() != "ring monitor is already running" {
		t.Errorf("Expected 'ring monitor is already running' error, got: %v", err)
	}

	// Test stopping
	err = rm.StopMonitoring()
	if err != nil {
		t.Fatalf("StopMonitoring failed: %v", err)
	}

	// Test stopping again (should fail)
	err = rm.StopMonitoring()
	if err == nil {
		t.Error("Expected error when stopping an already stopped monitor, got nil")
	} else if err.Error() != "ring monitor is not running" {
		t.Errorf("Expected 'ring monitor is not running' error, got: %v", err)
	}

	// Test starting with zero capacity (should fail)
	rm = NewRingMonitor() // Reset to a new monitor
	err = rm.StartMonitoring(0)
	if err == nil {
		t.Error("Expected error when starting with zero capacity, got nil")
	} else if err.Error() != "capacity cannot be zero" {
		t.Errorf("Expected 'capacity cannot be zero' error, got: %v", err)
	}
}

func TestRingMonitor_ReportState(t *testing.T) {
	rm := NewRingMonitor()

	// Test ReportState when not running (should fail)
	_, err := rm.ReportState()
	if err == nil {
		t.Error("Expected error when reporting state of not running monitor, got nil")
	} else if err.Error() != "ring monitor is not running" {
		t.Errorf("Expected 'ring monitor is not running' error, got: %v", err)
	}

	err = rm.StartMonitoring(100)
	if err != nil {
		t.Fatalf("StartMonitoring failed: %v", err)
	}

	state, err := rm.ReportState()
	if err != nil {
		t.Fatalf("ReportState failed: %v", err)
	}
	if state.Capacity != 100 {
		t.Errorf("Expected capacity 100, got %d", state.Capacity)
	}
	if state.CurrentSize != 0 {
		t.Errorf("Expected current size 0, got %d", state.CurrentSize)
	}
	if state.OverflowCount != 0 {
		t.Errorf("Expected overflow count 0, got %d", state.OverflowCount)
	}
	if state.DropCount != 0 {
		t.Errorf("Expected drop count 0, got %d", state.DropCount)
	}
}

func TestRingMonitor_SimulateEvent(t *testing.T) {
	rm := NewRingMonitor()
	capacity := uint64(10)
	err := rm.StartMonitoring(capacity)
	if err != nil {
		t.Fatalf("StartMonitoring failed: %v", err)
	}

	// Test SimulateEvent when not running (should fail)
	err = rm.StopMonitoring()
	if err != nil {
		t.Fatalf("StopMonitoring failed: %v", err)
	}
	err = rm.SimulateEvent(1)
	if err == nil {
		t.Error("Expected error when simulating event on not running monitor, got nil")
	} else if err.Error() != "ring monitor is not running" {
		t.Errorf("Expected 'ring monitor is not running' error, got: %v", err)
	}
	err = rm.StartMonitoring(capacity) // Restart for next tests
	if err != nil {
		t.Fatalf("StartMonitoring failed: %v", err)
	}

	// Simulate events without overflow
	err = rm.SimulateEvent(5)
	if err != nil {
		t.Fatalf("SimulateEvent failed: %v", err)
	}
	state, _ := rm.ReportState()
	if state.CurrentSize != 5 {
		t.Errorf("Expected current size 5, got %d", state.CurrentSize)
	}
	if state.OverflowCount != 0 {
		t.Errorf("Expected overflow count 0, got %d", state.OverflowCount)
	}
	if state.DropCount != 0 {
		t.Errorf("Expected drop count 0, got %d", state.DropCount)
	}

	// Simulate events causing overflow
	err = rm.SimulateEvent(8) // CurrentSize is 5, capacity 10. Adding 8 makes it 13. Overflow 3.
	if err != nil {
		t.Fatalf("SimulateEvent failed: %v", err)
	}
	state, _ = rm.ReportState()
	if state.CurrentSize != capacity { // Should be capped at capacity
		t.Errorf("Expected current size %d, got %d", capacity, state.CurrentSize)
	}
	if state.OverflowCount != 3 {
		t.Errorf("Expected overflow count 3, got %d", state.OverflowCount)
	}
	if state.DropCount != 3 {
		t.Errorf("Expected drop count 3, got %d", state.DropCount)
	}

	// Simulate more events causing overflow
	err = rm.SimulateEvent(5) // CurrentSize is 10, capacity 10. Adding 5 makes it 15. Overflow 5. Total overflow 8.
	if err != nil {
		t.Fatalf("SimulateEvent failed: %v", err)
	}
	state, _ = rm.ReportState()
	if state.CurrentSize != capacity {
		t.Errorf("Expected current size %d, got %d", capacity, state.CurrentSize)
	}
	if state.OverflowCount != 8 {
		t.Errorf("Expected overflow count 8, got %d", state.OverflowCount)
	}
	if state.DropCount != 8 {
		t.Errorf("Expected drop count 8, got %d", state.DropCount)
	}
}

func TestRingMonitor_Concurrency(t *testing.T) {
	rm := NewRingMonitor()
	capacity := uint64(100)
	err := rm.StartMonitoring(capacity)
	if err != nil {
		t.Fatalf("StartMonitoring failed: %v", err)
	}
	defer func() {
		_ = rm.StopMonitoring()
	}()

	numGoroutines := 10
	eventsPerGoroutine := 50

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < eventsPerGoroutine; j++ {
				_ = rm.SimulateEvent(1) // Simulate one event at a time
			}
		}()
	}
	wg.Wait()

	state, err := rm.ReportState()
	if err != nil {
		t.Fatalf("ReportState failed: %v", err)
	}

	totalEvents := uint64(numGoroutines * eventsPerGoroutine)
	expectedOverflows := uint64(0)
	if totalEvents > capacity {
		expectedOverflows = totalEvents - capacity
	}

	if state.OverflowCount != expectedOverflows {
		t.Errorf("Expected %d overflows, got %d", expectedOverflows, state.OverflowCount)
	}
	if state.DropCount != expectedOverflows {
		t.Errorf("Expected %d drops, got %d", expectedOverflows, state.DropCount)
	}
	if state.CurrentSize != capacity {
		t.Errorf("Expected current size %d, got %d", capacity, state.CurrentSize)
	}

	// Ensure that after many events, the system is stable
	err = rm.SimulateEvent(1) // One more event should cause one more overflow/drop
	if err != nil {
		t.Fatalf("SimulateEvent failed after concurrency: %v", err)
	}
	state, _ = rm.ReportState()
	if state.OverflowCount != expectedOverflows+1 {
		t.Errorf("Expected %d overflows, got %d", expectedOverflows+1, state.OverflowCount)
	}
}
