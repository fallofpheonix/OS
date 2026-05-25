package unit

import (
	"testing"
)

// TestFileDeterminism validates repeated action streams generate identical audit state.
func TestFileDeterminism(t *testing.T) {
	run := func() []FileAction {
		audit := NewFileAudit()
		audit.LogAction(FileAction{Path: "/etc/fstab", Action: ActionMonitor, Reason: "b1"})
		audit.LogAction(FileAction{Path: "/etc/fstab", Action: ActionFreeze, Reason: "b2"})
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
