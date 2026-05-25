package ordering

import (
	"testing"
)

func TestLogicalClock(t *testing.T) {
	clock := &MonotonicClock{}
	tick1 := clock.Tick()
	if tick1 != 1 {
		t.Errorf("Expected 1, got %d", tick1)
	}
	tick2 := clock.Tick()
	if tick2 != 2 {
		t.Errorf("Expected 2, got %d", tick2)
	}
}
