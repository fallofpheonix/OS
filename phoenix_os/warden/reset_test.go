package warden

import (
	"testing"

	"phoenix/bus"
)

func TestWardenManualBudgetReset(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	w.recoveryBudget = 1 // only 1 de-escalation allowed

	// Escalate to Suspicious
	w.Actuate(StateSuspicious, ClassLog, 1.0, 1, 1000, 10)

	// De-escalate to Normal (1st de-escalation)
	w.Actuate(StateNormal, ClassObserve, 1.0, 2, 1001, 41) // succeeds

	// Escalate to Suspicious again
	w.Actuate(StateSuspicious, ClassLog, 1.0, 3, 1002, 52)

	// Try to de-escalate to Normal again (2nd de-escalation - should block due to budget)
	transition := w.Actuate(StateNormal, ClassObserve, 1.0, 4, 1003, 83)
	if transition {
		t.Fatal("Expected de-escalation to fail since recovery budget of 1 is exhausted")
	}

	// Call manual budget reset
	w.ResetBudget()

	// Retrying de-escalation should now succeed
	transition = w.Actuate(StateNormal, ClassObserve, 1.0, 5, 1004, 83)
	if !transition {
		t.Error("Expected de-escalation to succeed after recovery budget reset")
	}
	if w.State != StateNormal {
		t.Errorf("Expected state to transition to Normal, got %s", w.State)
	}
}
