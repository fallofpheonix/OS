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

// TestFileDeterminism validates repeated action streams generate identical audit state.
func TestFileDeterminism(t *testing.T) {
	run := func() []file.FileAction {
		audit := file.NewFileAudit()
		audit.LogAction(file.FileAction{Path: "/etc/fstab", Action: file.ActionMonitor, Reason: "b1"})
		audit.LogAction(file.FileAction{Path: "/etc/fstab", Action: file.ActionFreeze, Reason: "b2"})
		return audit.History
	}

	h1 := run()
	for i := 0; i < 10; i++ {
		h2 := run()
		for j := range h1 {
			if h1[j].Hash != h2[j].Hash || h1[j].Sequence != h2[j].Sequence {
				t.Fatalf("determinism failure: mismatch at run %d, step %d", i, j)
			}
		}
	}
}
