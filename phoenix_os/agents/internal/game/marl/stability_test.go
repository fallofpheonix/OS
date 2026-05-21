package marl

import (
	"testing"
	"time"
)

func TestStability(t *testing.T) {
	sc := NewStabilityController(100*time.Millisecond, 1.0)

	// First action should pass
	if !sc.CanAct(0.5) {
		t.Error("Expected action to be allowed")
	}
	sc.RecordAction(0.5)

	// Immediate second action should fail due to cooldown
	if sc.CanAct(0.1) {
		t.Error("Expected action to be throttled due to cooldown")
	}

	// Wait for cooldown
	time.Sleep(150 * time.Millisecond)

	// Should pass if under containment limit
	if !sc.CanAct(0.4) {
		t.Error("Expected action to be allowed after cooldown")
	}
	sc.RecordAction(0.4)

	// Should fail if exceeding limit
	if sc.CanAct(0.2) {
		t.Error("Expected action to be throttled due to containment limit")
	}
}
