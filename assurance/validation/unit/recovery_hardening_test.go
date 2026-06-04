/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"encoding/json"
	"github.com/fallofpheonix/phoenix/foundation/runtime/state"
	"testing"
)

// S9.2 Corrupt Snapshot Restore
func TestCorruptSnapshotRestore(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	data, _ := r.CreateSnapshot()

	var s Snapshot
	json.Unmarshal(data, &s)
	s.Hash = "corrupt" // Corrupt hash
	corruptData, _ := json.Marshal(s)

	if err := r.RestoreFromSnapshot(corruptData); err == nil {
		t.Error("expected error for corrupted snapshot hash")
	}
}

// S9.3 Version Mismatch
func TestVersionMismatch(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	data, _ := r.CreateSnapshot()

	var s Snapshot
	json.Unmarshal(data, &s)
	s.Version = "2.0.0"             // Wrong version
	s.Hash = state.CalculateHash(s) // Re-sign with wrong version
	badVerData, _ := json.Marshal(s)

	if err := r.RestoreFromSnapshot(badVerData); err == nil {
		t.Error("expected error for version mismatch")
	}
}

// S9.5 Repeatability
func TestRecoveryRepeatability(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")
	data, _ := r.CreateSnapshot()

	for i := 0; i < 5; i++ {
		r.RestoreFromSnapshot(data)
		r.Rollback("rollback", "ev-2", "dec-2")
		r.RestoreFromSnapshot(data)
		if r.CurrentState != StateWatch {
			t.Errorf("repeat %d: expected StateWatch, got %s", i, r.CurrentState)
		}
	}
}
