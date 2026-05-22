package logical_clock

import (
	"testing"
)

func TestLogicalClock_Tick(t *testing.T) {
	clock := NewLogicalClock()
	
	val1 := clock.Tick()
	if val1 != 1 {
		t.Errorf("Expected 1, got %d", val1)
	}
	
	val2 := clock.Tick()
	if val2 != 2 {
		t.Errorf("Expected 2, got %d", val2)
	}
}

func TestLogicalClock_Current(t *testing.T) {
	clock := NewLogicalClock()
	
	if clock.Current() != 0 {
		t.Errorf("Expected 0, got %d", clock.Current())
	}
	
	clock.Tick()
	if clock.Current() != 1 {
		t.Errorf("Expected 1, got %d", clock.Current())
	}
}
