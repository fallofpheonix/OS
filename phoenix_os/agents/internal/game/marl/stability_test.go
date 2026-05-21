package marl

import (
	"testing"
	"time"
)

func TestStability(t *testing.T) {
<<<<<<< HEAD
	start := time.Now()
	// 100ms cooldown, 1.0 max containment, decay 0.0 per sec (for testing strict limits)
	sc := NewStabilityController(100*time.Millisecond, 1.0, 0.0, start)
=======
	// 100ms cooldown, 1.0 max containment, decay 0.0 per sec (for testing strict limits)
	sc := NewStabilityController(100*time.Millisecond, 1.0, 0.0)

	start := time.Now()
>>>>>>> origin/main

	// First action should pass
	if !sc.TryRecordAction(0.5, start) {
		t.Error("Expected action to be allowed")
	}

	// Immediate second action should fail due to cooldown
	if sc.TryRecordAction(0.1, start.Add(10*time.Millisecond)) {
		t.Error("Expected action to be throttled due to cooldown")
	}

	// Advance time past cooldown
	afterCooldown := start.Add(150 * time.Millisecond)

	// Should pass if under containment limit
	if !sc.TryRecordAction(0.4, afterCooldown) {
		t.Error("Expected action to be allowed after cooldown")
	}

	// Should fail if exceeding limit
	if sc.TryRecordAction(0.2, afterCooldown.Add(150*time.Millisecond)) {
		t.Error("Expected action to be throttled due to containment limit")
	}
}

func TestStabilityDecay(t *testing.T) {
<<<<<<< HEAD
	start := time.Now()
	// Decay 1.0 per second
	sc := NewStabilityController(100*time.Millisecond, 1.0, 1.0, start)

=======
	// Decay 1.0 per second
	sc := NewStabilityController(100*time.Millisecond, 1.0, 1.0)

	start := time.Now()
>>>>>>> origin/main
	
	// Max out containment
	sc.TryRecordAction(1.0, start)

<<<<<<< HEAD
	// Should fail (debt decays by 0.15 to 0.85, adding 0.2 makes it 1.05 > 1.0)
	if sc.TryRecordAction(0.2, start.Add(150*time.Millisecond)) {
=======
	// Should fail
	if sc.TryRecordAction(0.1, start.Add(150*time.Millisecond)) {
>>>>>>> origin/main
		t.Error("Expected action to be throttled due to containment limit")
	}

	// Advance time by 0.5s, debt should decay by 0.5
	afterDecay := start.Add(500 * time.Millisecond)
	
	// Should pass now that debt has decayed
	if !sc.TryRecordAction(0.4, afterDecay) {
		t.Error("Expected action to be allowed after debt decay")
	}
}

