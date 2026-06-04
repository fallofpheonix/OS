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

func TestFileAudit(t *testing.T) {
	audit := file.NewFileAudit()

	action := file.FileAction{
		Path:       "/etc/config",
		Action:     file.ActionThrottle,
		Reason:     "high I/O",
		EvidenceID: "ev-1",
		DecisionID: "dec-1",
	}

	audit.LogAction(action)

	if len(audit.History) != 1 {
		t.Errorf("expected 1 history entry, got %d", len(audit.History))
	}

	if audit.History[0].Sequence != 1 {
		t.Errorf("expected sequence 1, got %d", audit.History[0].Sequence)
	}
}
