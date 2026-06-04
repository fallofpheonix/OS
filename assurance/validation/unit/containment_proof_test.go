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

	"github.com/fallofpheonix/phoenix/foundation/runtime/containment"
)

// C1.5c Corruption Rejection
func TestCorruptSnapshot(t *testing.T) {
	e := containment.NewIsolationEngine(containment.StateObserve)
	e.Transition(containment.StateWatch, "ev-1", "dec-1")
	data, _ := e.CreateSnapshot()

	var s containment.Snapshot
	json.Unmarshal(data, &s)
	s.Hash = "invalid"
	badData, _ := json.Marshal(s)

	if err := e.RestoreFromSnapshot(badData); err == nil {
		t.Error("expected error for corrupted snapshot hash")
	}
}

// C1.8 Executable Proof Chain
func TestContainmentProof(t *testing.T) {
	// 1. Lifecycle
	e := containment.NewIsolationEngine(containment.StateObserve)
	e.Transition(containment.StateWatch, "ev-1", "dec-1")
	e.Transition(containment.StateThrottle, "ev-2", "dec-2")

	// 2. Snapshot
	data, _ := e.CreateSnapshot()

	// 3. Deterministic restore & verify (Field-level equality)
	e2 := containment.NewIsolationEngine(containment.StateObserve)
	if err := e2.RestoreFromSnapshot(data); err != nil {
		t.Fatalf("restore failed: %v", err)
	}

	if e2.CurrentState != containment.StateThrottle {
		t.Errorf("state mismatch: expected Throttle, got %s", e2.CurrentState)
	}

	for i := range e.History {
		if e.History[i].Current != e2.History[i].Current || e.History[i].EvidenceID != e2.History[i].EvidenceID {
			t.Errorf("deep history equality failed at step %d", i)
		}
	}
}
