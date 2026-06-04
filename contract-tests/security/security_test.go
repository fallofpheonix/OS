package security

import (
	"context"
	"testing"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// stubContainment implements securityv1.Containment
type stubContainment struct {
	target string
	level  securityv1.ContainmentLevel
	reason string
}

func (c *stubContainment) Target() string                      { return c.target }
func (c *stubContainment) Level() securityv1.ContainmentLevel  { return c.level }
func (c *stubContainment) Reason() string                      { return c.reason }

// stubActuator implements securityv1.Actuator
type stubActuator struct {
	currentLevel securityv1.ContainmentLevel
}

func (a *stubActuator) Actuate(ctx context.Context, action securityv1.Containment) error {
	a.currentLevel = action.Level()
	return nil
}

func (a *stubActuator) Kill(ctx context.Context, pid int) error {
	return nil
}

func (a *stubActuator) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	return a.currentLevel, nil
}

// TestFITSec001 verifies securityv1.Actuator, Containment, and ContainmentLevel interfaces.
func TestFITSec001(t *testing.T) {
	var act interface{} = &stubActuator{
		currentLevel: securityv1.LevelNone,
	}

	if _, ok := act.(securityv1.Actuator); !ok {
		t.Fatal("stubActuator does not satisfy securityv1.Actuator contract interface")
	}

	var action interface{} = &stubContainment{
		target: "proc_1234",
		level:  securityv1.LevelIsolate,
		reason: "unauthorized network connection",
	}

	if _, ok := action.(securityv1.Containment); !ok {
		t.Fatal("stubContainment does not satisfy securityv1.Containment contract interface")
	}
}
