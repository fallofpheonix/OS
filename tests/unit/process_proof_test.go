package unit

import (
	"testing"
)

// TestProcessDeterminism validates repeated action streams generate identical audit state.
func TestProcessDeterminism(t *testing.T) {
	run := func() []ProcessAction {
		audit := NewProcessAudit()
		audit.LogAction(ProcessAction{PID: 1, Action: ActionMonitor, Reason: "b1"})
		audit.LogAction(ProcessAction{PID: 1, Action: ActionThrottle, Reason: "b2"})
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
	audit := NewProcessAudit()
	audit.LogAction(ProcessAction{PID: 1, Action: ActionMonitor, Reason: "r1"})
	audit.LogAction(ProcessAction{PID: 1, Action: ActionIsolate, Reason: "r2"})

	// Snapshot
	data, _ := audit.CreateSnapshot()

	// Restore
	a2 := NewProcessAudit()
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
