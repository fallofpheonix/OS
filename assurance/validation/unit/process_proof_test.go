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

// TestProcessDeterminism validates repeated action streams generate identical audit state.
func TestProcessDeterminism(t *testing.T) {
	run := func() []containment.ProcessAction {
		audit := containment.NewProcessAudit()
		audit.LogAction(containment.ProcessAction{PID: 1, Action: containment.ActionMonitor, Reason: "b1"})
		audit.LogAction(containment.ProcessAction{PID: 1, Action: containment.ActionThrottle, Reason: "b2"})
		return audit.History
	}

	h1 := run()
	for i := 0; i < 10; i++ {
		h2 := run()
		for j := range h1 {
			if h1[j].Hash != h2[j].Hash {
				t.Fatalf("determinism failure: hash mismatch at run %d, step %d", i, j)
			}
		}
	}
}

// TestProcessProofChain implements the C10 equivalent for Process Containment.
func TestProcessProofChain(t *testing.T) {
	audit := containment.NewProcessAudit()
	audit.LogAction(containment.ProcessAction{PID: 1, Action: containment.ActionMonitor, Reason: "r1"})
	audit.LogAction(containment.ProcessAction{PID: 1, Action: containment.ActionIsolate, Reason: "r2"})

	// Snapshot
	data, _ := audit.CreateSnapshot()

	// Restore
	a2 := containment.NewProcessAudit()
	if err := a2.RestoreFromSnapshot(data); err != nil {
		t.Fatalf("proof failed: %v", err)
	}

	// Verify deep equality
	if a2.Sequence != audit.Sequence || len(a2.History) != len(audit.History) {
		t.Fatal("proof failed: metadata mismatch")
	}
	if a2.History[0].Hash != audit.History[0].Hash {
		t.Fatal("proof failed: integrity hash mismatch")
	}
}
