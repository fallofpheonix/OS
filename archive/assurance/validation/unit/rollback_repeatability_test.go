/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import "testing"

// S5.8 Rollback determinism
func TestStateRegistryRollbackDeterminism(t *testing.T) {
	for i := 0; i < 10; i++ {
		r := NewStateRegistry(StateSafe)
		r.Transition(StateWatch, "burst", "ev-1", "dec-1")
		r.Rollback("rollback", "ev-2", "dec-2")

		if r.CurrentState != StateSafe {
			t.Errorf("run %d: expected StateSafe, got %s", i, r.CurrentState)
		}
	}
}

// S5.7 Replay rollback validation
func TestRollbackReplayRepeatability(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	r.Transition(StateAlert, "spike", "ev-2", "dec-2")
	if err := r.Rollback("rollback", "ev-3", "dec-3"); err != nil {
		t.Fatalf("rollback returned error: %v", err)
	}

	if r.CurrentState != StateWatch {
		t.Errorf("history: %v, current: %s", r.History, r.CurrentState)
	}
}
