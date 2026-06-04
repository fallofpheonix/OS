/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package emergency

import (
	"testing"

	"github.com/fallofpheonix/phoenix/assurance/security/engine"
)

func TestKillSwitch_Engage(t *testing.T) {
	w := engine.NewWarden()
	ks := NewKillSwitch(w)

	if ks.IsEngaged() {
		t.Error("KillSwitch engaged prematurely")
	}

	ks.Engage()

	if !ks.IsEngaged() {
		t.Error("KillSwitch failed to engage")
	}

	// Verify Warden is locked
	if err := w.Transition(engine.StateWatch); err == nil {
		t.Error("Warden FSM was not locked by KillSwitch")
	}
}
