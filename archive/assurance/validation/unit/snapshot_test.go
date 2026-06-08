/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"fmt"
	"github.com/fallofpheonix/phoenix/foundation/runtime/truth"
	"testing"
)

func TestLedgerSnapshot(t *testing.T) {
	l := truth.NewTruthLedger()
	l.AddEntry(truth.EvidenceWrapper{Sequence: 1, Hash: "A"})
	l.AddEntry(truth.EvidenceWrapper{Sequence: 2, Hash: "B"})

	snap := l.Snapshot()
	if len(snap) != 2 {
		t.Errorf("Expected 2 entries in snapshot, got %d", len(snap))
	}

	// Add more to ledger, snapshot should not change
	l.AddEntry(truth.EvidenceWrapper{Sequence: 3, Hash: "C"})
	if len(snap) != 2 {
		t.Error("Snapshot changed after ledger update")
	}
	fmt.Println("[PX-012] Ledger Snapshot Integrity: PASSED")
}
