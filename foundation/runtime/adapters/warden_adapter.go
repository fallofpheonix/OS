package adapters

import (
	"context"
	"fmt"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	"github.com/fallofpheonix/phoenix/assurance/security/engine"
)

// WardenAdapter implements securityv1.Actuator wrapping *engine.Warden.
type WardenAdapter struct {
	Warden *engine.Warden
}

// NewWardenAdapter creates a new WardenAdapter.
func NewWardenAdapter(w *engine.Warden) *WardenAdapter {
	return &WardenAdapter{
		Warden: w,
	}
}

// Actuate maps the Containment action to Warden SystemState and calls Transition.
func (a *WardenAdapter) Actuate(ctx context.Context, action securityv1.Containment) error {
	var targetState engine.SystemState
	switch action.Level() {
	case securityv1.LevelNone:
		targetState = engine.StateSafe
	case securityv1.LevelMonitor:
		targetState = engine.StateWatch
	case securityv1.LevelSandbox:
		targetState = engine.StateSuspicious
	case securityv1.LevelIsolate:
		targetState = engine.StateCritical
	case securityv1.LevelQuench:
		targetState = engine.StateCompromised
	default:
		return fmt.Errorf("unknown containment level: %v", action.Level())
	}

	return a.Warden.Transition(targetState)
}

// Kill simulates process termination or initiates emergency FSM lockout.
func (a *WardenAdapter) Kill(ctx context.Context, pid int) error {
	a.Warden.Lock()
	return nil
}

// GetCurrentLevel queries the current SystemState and maps it back to ContainmentLevel.
func (a *WardenAdapter) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	state := a.Warden.GetState()
	switch state {
	case engine.StateSafe:
		return securityv1.LevelNone, nil
	case engine.StateWatch:
		return securityv1.LevelMonitor, nil
	case engine.StateSuspicious:
		return securityv1.LevelSandbox, nil
	case engine.StateCritical:
		return securityv1.LevelIsolate, nil
	case engine.StateCompromised:
		return securityv1.LevelQuench, nil
	default:
		return securityv1.LevelNone, fmt.Errorf("unknown warden state: %s", state)
	}
}
