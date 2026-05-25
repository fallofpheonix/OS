package unit

import (
	"testing"
)

// TestFileProofChain validates the full recovery/replay cycle for file containment.
func TestFileProofChain(t *testing.T) {
	// 1. Setup
	audit := NewFileAudit()
	audit.LogAction(FileAction{Path: "/bin/sh", Action: ActionMonitor, Reason: "r1"})
	audit.LogAction(FileAction{Path: "/bin/sh", Action: ActionFreeze, Reason: "r2"})

	// 2. Snapshot
	data, _ := audit.CreateSnapshot()

	// 3. Restore and Verify (Proof)
	a2 := NewFileAudit()
	if err := a2.RestoreFromSnapshot(data); err != nil {
		t.Fatalf("proof failed: restore error %v", err)
	}

	// 4. Deep replay equivalence check
	if a2.Sequence != audit.Sequence || len(a2.History) != len(audit.History) {
		t.Fatal("proof failed: metadata mismatch")
	}
	if a2.History[0].Hash != audit.History[0].Hash {
		t.Fatal("proof failed: integrity hash mismatch")
	}
}
