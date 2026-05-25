package clock

import (
	"fmt"
	"sync/atomic"
	"testing"
)

// MockLogicalClock for testing ClockSynchronizer
type mockLogicalClock struct {
	value atomic.Uint64
}

func (m *mockLogicalClock) Tick() uint64 {
	return m.value.Add(1)
}

func (m *mockLogicalClock) Now() uint64 {
	return m.value.Load()
}

func (m *mockLogicalClock) AdvanceTo(v uint64) error {
	current := m.value.Load()
	if v < current {
		return fmt.Errorf("mock clock regression: cannot advance to %d, current value is %d", v, current)
	}
	m.value.Store(v)
	return nil
}

func newMockLogicalClock(initialValue uint64) *mockLogicalClock {
	m := &mockLogicalClock{}
	m.value.Store(initialValue)
	return m
}

func TestClockSynchronizer_SyncTo(t *testing.T) {
	cs := NewClockSynchronizer()

	target := newMockLogicalClock(10)
	source := newMockLogicalClock(50)

	// Test successful sync: target advances to source's value
	err := cs.SyncTo(target, source)
	if err != nil {
		t.Fatalf("SyncTo failed: %v", err)
	}
	if target.Now() != 50 {
		t.Errorf("Expected target clock to be 50, got %d", target.Now())
	}

	// Test sync where source is lower than target (should cause target to report regression)
	sourceLower := newMockLogicalClock(30) // Create a new source that is lower than current target (50)
	err = cs.SyncTo(target, sourceLower) // target (50) trying to advance to sourceLower (30)
	if err == nil {
		t.Error("Expected error for clock regression, got nil")
	} else if err.Error() != "mock clock regression: cannot advance to 30, current value is 50" {
		t.Errorf("Expected specific clock regression error, got: %v", err)
	}
	if target.Now() != 50 { // Should not have changed
		t.Errorf("Expected target clock to remain 50, got %d", target.Now())
	}

	// Test nil target clock
	err = cs.SyncTo(nil, source)
	if err == nil {
		t.Error("Expected error for nil target clock, got nil")
	} else if err.Error() != "target clock cannot be nil" {
		t.Errorf("Expected 'target clock cannot be nil' error, got: %v", err)
	}

	// Test nil source clock
	err = cs.SyncTo(target, nil)
	if err == nil {
		t.Error("Expected error for nil source clock, got nil")
	} else if err.Error() != "source clock cannot be nil" {
		t.Errorf("Expected 'source clock cannot be nil' error, got: %v", err)
	}
}

func TestClockSynchronizer_AdvanceTargetTo(t *testing.T) {
	cs := NewClockSynchronizer()
	target := newMockLogicalClock(20)

	// Test successful advance
	err := cs.AdvanceTargetTo(target, 60)
	if err != nil {
		t.Fatalf("AdvanceTargetTo failed: %v", err)
	}
	if target.Now() != 60 {
		t.Errorf("Expected target clock to be 60, got %d", target.Now())
	}

	// Test advance to lower value (should fail)
	err = cs.AdvanceTargetTo(target, 40)
	if err == nil {
		t.Error("Expected error for clock regression, got nil")
	} else if err.Error() != "mock clock regression: cannot advance to 40, current value is 60" {
		t.Errorf("Expected clock regression error, got: %v", err)
	}
	if target.Now() != 60 { // Should not have changed
		t.Errorf("Expected target clock to remain 60, got %d", target.Now())
	}

	// Test nil target clock
	err = cs.AdvanceTargetTo(nil, 100)
	if err == nil {
		t.Error("Expected error for nil target clock, got nil")
	} else if err.Error() != "target clock cannot be nil" {
		t.Errorf("Expected 'target clock cannot be nil' error, got: %v", err)
	}
}
