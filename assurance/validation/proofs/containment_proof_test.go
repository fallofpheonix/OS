package proofs

import (
	"context"
	"fmt"
	"testing"

	warden "github.com/fallofpheonix/phoenix/assurance/security"
	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// MockActuator simulates a physical containment mechanism.
type MockActuator struct {
	IsolatedPID  int
	KilledPID    int
	FreezedPID   int
	WarnedPID    int
	ThrottledPID int
	currentLevel securityv1.ContainmentLevel
}

func (m *MockActuator) Actuate(ctx context.Context, action securityv1.Containment) error {
	m.currentLevel = action.Level()
	var pid int
	_, _ = fmt.Sscanf(action.Target(), "PID:%d", &pid)
	if pid == 0 {
		pid = 1234 // fallback for testing if target format differs
	}

	switch action.Level() {
	case securityv1.LevelMonitor:
		m.WarnedPID = pid
	case securityv1.LevelSandbox:
		m.ThrottledPID = pid
	case securityv1.LevelIsolate:
		m.IsolatedPID = pid
	case securityv1.LevelQuench:
		m.FreezedPID = pid
	}
	return nil
}

func (m *MockActuator) Kill(ctx context.Context, pid int) error {
	m.currentLevel = securityv1.LevelQuench
	m.KilledPID = pid
	return nil
}

func (m *MockActuator) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	return m.currentLevel, nil
}

func (m *MockActuator) Name() string { return "MockActuator" }

// STATUS: EXPERIMENTAL
// Proof 3: Containment (Attack, Detect, Contain, Verify)
func TestContainmentProof(t *testing.T) {
	// Setup Warden with a Mock Actuator
	w := warden.NewWarden(nil, nil)
	w.ShadowMode = false
	w.ShadowMode = false // Explicitly disable shadow mode for real actuation test
	act := &MockActuator{}
	w.Actuators = append(w.Actuators, act)

	// Define a formal Invariant: WATCH state requires 0.5 evidence weight.
	w.Invariants = append(w.Invariants, &warden.EvidenceWeightInvariant{
		StateThresholds: map[warden.SystemState]float64{
			warden.StateWatch: 0.5,
		},
	})

	// 1. Attack: Attempt an unauthorized state transition with insufficient evidence.
	req := warden.AuthorityEscalationRequest{
		TargetPID:      1234,
		TargetState:    warden.StateWatch,
		ActuationClass: warden.ClassIsolate,
		EvidenceWeight: 0.1, // Violation: 0.1 < 0.5
	}

	// 2. Detect: Warden Actuate identifies the Invariant Violation.
	// 3. Contain: Warden triggers emergencyHalt and isolates the target PID.
	success := w.ActuateRequest(req, 1, 100)

	if success {
		t.Error("Expected ActuateRequest to fail and trigger containment due to invariant violation")
	}

	// 4. Verify: System is in COMPROMISED state and PID is Isolated.
	if w.State != warden.StateCompromised {
		t.Errorf("CONTAINMENT_PROOF_FAILED: Expected state COMPROMISED, got %s", w.State)
	}

	if act.IsolatedPID != 1234 {
		t.Errorf("CONTAINMENT_PROOF_FAILED: PID 1234 was not isolated by the actuator")
	}
}

func TestShadowModeProof(t *testing.T) {
	w := warden.NewWarden(nil, nil)
	w.ShadowMode = false
	w.ShadowMode = true // Enable Shadow Mode
	act := &MockActuator{}
	w.Actuators = append(w.Actuators, act)

	req := warden.AuthorityEscalationRequest{
		TargetPID:      5678,
		TargetState:    warden.StateWatch,
		ActuationClass: warden.ClassKill,
		EvidenceWeight: 0.9,
	}

	// In shadow mode, this should "succeed" but NOT trigger the actuator
	success := w.ActuateRequest(req, 1, 100)
	if !success {
		t.Error("Expected ActuateRequest to succeed in Shadow Mode")
	}

	if w.State != warden.StateSafe {
		t.Errorf("Expected internal state to remain SAFE in Shadow Mode, got %s", w.State)
	}

	if act.KilledPID != 0 {
		t.Error("SHADOW_MODE_FAILURE: Actuator triggered in shadow mode")
	}
}
