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
	"testing"
)

// S8.1 Deterministic snapshot repeatability
func TestSnapshotRepeatability(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")

	lastHash := ""
	for i := 0; i < 10; i++ {
		data, _ := r.CreateSnapshot()
		var s Snapshot
		json.Unmarshal(data, &s)

		if lastHash != "" && s.Hash != lastHash {
			t.Errorf("run %d: hash mismatch", i)
		}
		lastHash = s.Hash
	}
}

// S8.2 Replay equivalence
func TestReplayEquivalence(t *testing.T) {
	r := NewStateRegistry(StateSafe)
	r.Transition(StateWatch, "burst", "ev-1", "dec-1")

	data, _ := r.CreateSnapshot()

	r2 := NewStateRegistry(StateSafe)
	r2.RestoreFromSnapshot(data)

	if r2.CurrentState != r.CurrentState {
		t.Error("state mismatch after restore")
	}
	if len(r2.History) != len(r.History) {
		t.Error("history length mismatch")
	}
}
