/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment"
	"testing"
)

func TestIsolationLifecycle(t *testing.T) {
	e := containment.NewIsolationEngine(containment.StateObserve)

	// Valid sequence: Observe -> Watch -> Throttle -> Isolate -> Recover -> Observe
	transitions := []struct {
		next containment.IsolationState
	}{
		{containment.StateWatch},
		{containment.StateThrottle},
		{containment.StateIsolate},
		{containment.StateRecover},
		{containment.StateObserve},
	}

	for _, tr := range transitions {
		if err := e.Transition(tr.next, "ev-1", "dec-1"); err != nil {
			t.Errorf("transition to %s failed: %v", tr.next, err)
		}
	}
}

func TestIllegalContainmentTransition(t *testing.T) {
	e := containment.NewIsolationEngine(containment.StateObserve)

	// Illegal jump: Observe -> Isolate
	if err := e.Transition(containment.StateIsolate, "ev-1", "dec-1"); err == nil {
		t.Error("expected error for illegal containment transition")
	}
}
