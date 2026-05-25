package unit

import (
	"testing"
)

// S9.1 Snapshot-Rollback Recovery Chain
func TestSnapshotRollbackRecovery(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	r.Transition(StateAlert, "spike", "ev-2", "dec-2")

	// 1. Snapshot
	data, _ := r.CreateSnapshot()

	// 2. Rollback
	r.Rollback("rollback", "ev-3", "dec-3")
	if r.CurrentState != StateWatch {
		t.Errorf("expected Watch, got %s", r.CurrentState)
	}

	// 3. Restore snapshot
	r.RestoreFromSnapshot(data)

	// 4. Verify
	if r.CurrentState != StateAlert {
		t.Errorf("expected Alert, got %s", r.CurrentState)
	}
	if len(r.History) != 3 {
		t.Errorf("expected history 3, got %d", len(r.History))
	}
}
