package warden

import (
	"testing"

	"phoenix/bus"
)

func TestWardenHysteresisDwellLimits(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	// Elevate state immediately
	transition1 := w.Actuate(StateSuspicious, ClassLog, 1.0, 1, 1000, 10)
	if !transition1 {
		t.Fatal("Initial transition should succeed")
	}

	// Try to de-escalate back to normal immediately (tick 11 < 10 + 30 dwell ticks)
	transition2 := w.Actuate(StateNormal, ClassObserve, 1.0, 2, 1001, 11)
	if transition2 {
		t.Error("Expected de-escalation to fail due to dwell limit hysteresis")
	}
	if w.State != StateSuspicious {
		t.Errorf("Expected state to remain Suspicious, got %s", w.State)
	}

	// De-escalate after dwell ticks pass (tick 41 >= 10 + 30 dwell ticks)
	transition3 := w.Actuate(StateNormal, ClassObserve, 1.0, 3, 1002, 41)
	if !transition3 {
		t.Error("Expected de-escalation to succeed after dwell ticks passed")
	}
	if w.State != StateNormal {
		t.Errorf("Expected state to transition to Normal, got %s", w.State)
	}
}

func TestWardenTransitionCooldownLock(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	// 1. Initial transition (tick 10)
	w.Actuate(StateSuspicious, ClassLog, 1.0, 1, 1000, 10)

	// 2. Try to transition to Contained during stabilization cooldown (tick 15 < 10 + 10 cooldown ticks)
	// with a non-critical actuation class (ClassLocalIsolate = 3 < ClassClusterIsolate/KernelEmergency)
	// Wait, ClassLocalIsolate is defined in warden.go:
	// ClassThrottle = 2, ClassLocalIsolate = 3, ClassClusterIsolate = 4, ClassKernelEmergency = 5
	// In warden.go, we defined bypass for ClassLocalIsolate or higher (class >= ClassLocalIsolate).
	// Let's test cooldown lockout with a ClassLog transition (ClassLog = 1).
	transition1 := w.Actuate(StateNormal, ClassLog, 1.0, 2, 1001, 15)
	if transition1 {
		t.Error("Expected transition during cooldown window to be blocked")
	}

	// 3. Try to transition during cooldown but with high-severity escalation (ClassLocalIsolate)
	transition2 := w.Actuate(StateContained, ClassLocalIsolate, 1.0, 3, 1002, 15)
	if !transition2 {
		t.Error("Expected high-severity escalation to bypass cooldown lock")
	}
}

func TestWardenRecoveryBudget(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	w.recoveryBudget = 2 // override budget to 2 deescalations for test

	// Cycle 1: Escalate & De-escalate
	w.Actuate(StateSuspicious, ClassLog, 1.0, 1, 1000, 10) // tick 10
	w.Actuate(StateNormal, ClassObserve, 1.0, 2, 1001, 41) // tick 41 (de-escalation 1)

	// Cycle 2: Escalate & De-escalate
	w.Actuate(StateSuspicious, ClassLog, 1.0, 3, 1002, 52) // tick 52
	w.Actuate(StateNormal, ClassObserve, 1.0, 4, 1003, 83) // tick 83 (de-escalation 2)

	// Cycle 3: Escalate & De-escalate (Should fail on de-escalate)
	w.Actuate(StateSuspicious, ClassLog, 1.0, 5, 1004, 94)  // tick 94
	transition := w.Actuate(StateNormal, ClassObserve, 1.0, 6, 1005, 125) // tick 125 (de-escalation 3)

	if transition {
		t.Error("Expected de-escalation to fail because recovery budget was exhausted")
	}
	if w.State != StateSuspicious {
		t.Errorf("Expected state to remain Suspicious, got %s", w.State)
	}
}
