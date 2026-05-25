package warden

import (
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestWardenStateTransitions(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	// 1. Initial State: SAFE
	if w.State != StateSafe {
		t.Errorf("Expected initial state SAFE, got %s", w.State)
	}

	// 2. Trigger: Fork Burst -> WATCH
	w.EvaluateTrigger(TriggerForkBurst, 1001, 10)
	if w.State != StateWatch {
		t.Errorf("Expected state WATCH after Fork Burst, got %s", w.State)
	}

	// 3. Trigger: Network Beacon -> ALERT
	w.EvaluateTrigger(TriggerNetworkBeacon, 1001, 20)
	if w.State != StateAlert {
		t.Errorf("Expected state ALERT after Network Beacon, got %s", w.State)
	}

	// 4. Trigger: Reverse Shell -> CONTAIN (High Severity, bypasses intermediate)
	w.EvaluateTrigger(TriggerReverseShell, 1001, 30)
	if w.State != StateContain {
		t.Errorf("Expected state CONTAIN after Reverse Shell, got %s", w.State)
	}

	// 5. Trigger: De-escalation (Hysteresis Check)
	// Try to de-escalate to Recovery before DwellTicks (30) have passed since tick 30
	w.EvaluateTrigger(TriggerHumanOverride, 1001, 40)
	if w.State != StateContain {
		t.Errorf("Expected Hysteresis block to keep state at CONTAIN at tick 40, got %s", w.State)
	}

	// Try again after DwellTicks (30 + 30 = 60)
	w.EvaluateTrigger(TriggerHumanOverride, 1001, 70)
	if w.State != StateRecovery {
		t.Errorf("Expected state RECOVERY after Human Override at tick 70, got %s", w.State)
	}
}

func TestWardenDeterminism(t *testing.T) {
	b := bus.NewBus()

	// Create two wardens
	w1 := NewWarden(b)
	w2 := NewWarden(b)

	// Inputs
	triggers := []TriggerType{TriggerForkBurst, TriggerMassWrite, TriggerReverseShell}
	ticks := []uint64{10, 20, 30}

	// Process on both
	for i := range triggers {
		w1.EvaluateTrigger(triggers[i], 2000, ticks[i])
		w2.EvaluateTrigger(triggers[i], 2000, ticks[i])
	}

	// Verify states are identical
	if w1.State != w2.State {
		t.Errorf("Divergence detected: Warden1 state %s != Warden2 state %s", w1.State, w2.State)
	}
}
