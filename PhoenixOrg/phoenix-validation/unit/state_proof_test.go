package unit

import "testing"

// S10: Executable Proof Chain
func TestStateExecutableProof(t *testing.T) {
	// 1. Initial State Transition
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "b1", "e1", "d1")
	r.Transition(StateAlert, "b2", "e2", "d2")

	// 2. Snapshot
	data, _ := r.CreateSnapshot()

	// 3. Rollback & Replay chain
	r.Rollback("r1", "e3", "d3")
	if r.CurrentState != StateWatch {
		t.Fatalf("rollback failed: expected Watch, got %s", r.CurrentState)
	}

	// 4. Restore & Verify Determinism
	r.RestoreFromSnapshot(data)
	if r.CurrentState != StateAlert {
		t.Fatalf("restore failed: expected Alert, got %s", r.CurrentState)
	}

	// 5. Final Replay Validation
	if len(r.History) != 3 {
		t.Errorf("history mismatch: expected 3, got %d", len(r.History))
	}
}
