package control

import (
	"testing"
)

func TestControllerTransitions(t *testing.T) {
	c := NewController()

	if c.CurrentState != StateSafe {
		t.Errorf("Expected initial state SAFE, got %s", c.CurrentState)
	}

	state, action := c.UpdateState(0.4)
	if state != StateWatch || action != "OBSERVE" {
		t.Errorf("Transition to WATCH failed: %s/%s", state, action)
	}

	state, action = c.UpdateState(0.85)
	if state != StateCritical || action != "FREEZE" {
		t.Errorf("Transition to CRITICAL failed: %s/%s", state, action)
	}

	state, action = c.UpdateState(1.0)
	if state != StateCompromised || action != "ISOLATE" {
		t.Errorf("Transition to COMPROMISED failed: %s/%s", state, action)
	}
}
