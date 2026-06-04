/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package security

import (
	"github.com/fallofpheonix/phoenix/assurance/security/control"
	"testing"
)

func TestControllerTransitions(t *testing.T) {
	c := control.NewController()

	if c.CurrentState != control.StateSafe {
		t.Errorf("Expected initial state SAFE, got %s", c.CurrentState)
	}

	state, action := c.UpdateState(0.4)
	if state != control.StateWatch || action != "OBSERVE" {
		t.Errorf("Transition to WATCH failed: %s/%s", state, action)
	}

	state, action = c.UpdateState(0.85)
	if state != control.StateCritical || action != "FREEZE" {
		t.Errorf("Transition to CRITICAL failed: %s/%s", state, action)
	}

	state, action = c.UpdateState(1.0)
	if state != control.StateCompromised || action != "ISOLATE" {
		t.Errorf("Transition to COMPROMISED failed: %s/%s", state, action)
	}
}
