package security

import (
	"context"
	"testing"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	warden "github.com/fallofpheonix/phoenix/assurance/security"
	"github.com/fallofpheonix/phoenix/assurance/security/actuators"
)

// FIT-GUARD-001 verifies Warden and actuators implement securityv1.Actuator
func TestFITGuard001(t *testing.T) {
	w := warden.NewWarden(nil)
	var _ securityv1.Actuator = w

	pa := actuators.NewProcessActuator()
	var _ securityv1.Actuator = pa

	ea := actuators.NewEBPFActuator(nil)
	var _ securityv1.Actuator = ea
}

type mockContainment struct {
	target string
	level  securityv1.ContainmentLevel
	reason string
}

func (m mockContainment) Target() string                      { return m.target }
func (m mockContainment) Level() securityv1.ContainmentLevel  { return m.level }
func (m mockContainment) Reason() string                      { return m.reason }

// FIT-GUARD-002 verifies that containment levels mapped and processed correctly
func TestFITGuard002(t *testing.T) {
	w := warden.NewWarden(nil)
	w.ShadowMode = false

	// Register an actuator to observe level changes
	pa := actuators.NewProcessActuator()
	w.RegisterActuator(pa)

	ctx := context.Background()

	// Initial level: LevelNone (StateSafe)
	lvl, err := w.GetCurrentLevel()
	if err != nil {
		t.Fatalf("failed to get current level: %v", err)
	}
	if lvl != securityv1.LevelNone {
		t.Errorf("expected LevelNone, got %v", lvl)
	}

	// Elevate to LevelMonitor
	action := mockContainment{target: "PID:1234", level: securityv1.LevelMonitor, reason: "high memory usage"}
	if err := w.Actuate(ctx, action); err != nil {
		t.Fatalf("failed to actuate: %v", err)
	}

	lvl, _ = w.GetCurrentLevel()
	if lvl != securityv1.LevelMonitor {
		t.Errorf("expected LevelMonitor, got %v", lvl)
	}

	paLvl, _ := pa.GetCurrentLevel()
	if paLvl != securityv1.LevelMonitor {
		t.Errorf("expected ProcessActuator to be LevelMonitor, got %v", paLvl)
	}
}

// FIT-GUARD-003 verifies FSM transition constraints (no state skipping)
func TestFITGuard003(t *testing.T) {
	w := warden.NewWarden(nil)
	w.ShadowMode = false

	ctx := context.Background()

	// SAFE -> LevelSandbox (StateSuspicious) is illegal (skips LevelMonitor / WATCH)
	actionSkip := mockContainment{target: "PID:1234", level: securityv1.LevelSandbox, reason: "skip state"}
	if err := w.Actuate(ctx, actionSkip); err == nil {
		t.Errorf("expected error for illegal state skipping, but got none")
	}

	// SAFE -> LevelMonitor (StateWatch) is legal
	actionMonitor := mockContainment{target: "PID:1234", level: securityv1.LevelMonitor, reason: "legal step"}
	if err := w.Actuate(ctx, actionMonitor); err != nil {
		t.Fatalf("failed legal state transition SAFE -> LevelMonitor: %v", err)
	}
}
