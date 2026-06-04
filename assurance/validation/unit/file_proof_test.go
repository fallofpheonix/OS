/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/file"
	"testing"
)

// TestFileProofChain validates the full recovery/replay cycle for file containment.
func TestFileProofChain(t *testing.T) {
	// 1. Setup
	audit := file.NewFileAudit()
	audit.LogAction(file.FileAction{Path: "/bin/sh", Action: file.ActionMonitor, Reason: "r1"})
	audit.LogAction(file.FileAction{Path: "/bin/sh", Action: file.ActionFreeze, Reason: "r2"})

	// 2. Snapshot
	data, _ := audit.CreateSnapshot()

	// 3. Restore and Verify (Proof)
	a2 := file.NewFileAudit()
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
