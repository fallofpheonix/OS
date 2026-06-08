/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import "testing"

func TestRollback(t *testing.T) {
	r := NewStateRegistry(StateSafe)

	r.Transition(StateWatch, "burst", "ev-1", "dec-1")

	if err := r.Rollback("rollback", "ev-2", "dec-2"); err != nil {
		t.Fatalf("rollback failed: %v", err)
	}

	if r.CurrentState != StateSafe {
		t.Errorf("expected Safe, got %s", r.CurrentState)
	}

	if len(r.History) != 3 {
		t.Errorf("expected history length 3, got %d", len(r.History))
	}
}

func TestIllegalRollback(t *testing.T) {
	r := NewStateRegistry(StateSafe)

	// Case 1: No history to rollback
	if err := r.Rollback("illegal", "ev-0", "dec-0"); err == nil {
		t.Error("expected error for rollback from initial state, but got nil")
	}

	// Case 2: Illegal path: ALERT -> SAFE
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	r.Transition(StateAlert, "spike", "ev-2", "dec-2")
	// Current state is ALERT. Trying to rollback to SAFE is illegal (should go to WATCH).
	// We force the state via manual mutation to simulate an illegal path attempt if necessary,
	// or just realize that the logic prevents this transition.
	// Actually, based on rollback rules, ALERT only rolls back to WATCH.
	// Since ALERT is the current state, `Rollback` tries to go to `lastRecord.Previous`.
	// If `lastRecord` was WATCH->ALERT, `Previous` is WATCH. That is a legal rollback.
	// To test an illegal rollback path, I need to manipulate the history.

	// Let's create an illegal transition in history to test validation.
	r.History = append(r.History, StateRecord{
		Previous: StateSafe,  // Pretending we came from SAFE
		Current:  StateAlert, // To ALERT
	})
	r.CurrentState = StateAlert

	if err := r.Rollback("illegal-path", "ev-3", "dec-3"); err == nil {
		t.Error("expected error for illegal rollback path ALERT->SAFE, but got nil")
	}
}
