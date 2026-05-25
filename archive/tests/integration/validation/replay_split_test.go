package validation

import (
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/ledger/src"
)

// TestReplaySplit verifies that logging divergent action chains yields distinct, divergent head hashes.
func TestReplaySplit(t *testing.T) {
	l1 := ledger.NewLedger(nil)
	l2 := ledger.NewLedger(nil)

	// Same root
	l1.AddEntry("evt-1", "cause-1", []byte("event-1"))
	l2.AddEntry("evt-1", "cause-1", []byte("event-1"))

	// Split pathways
	l1.AddEntry("evt-2a", "cause-1", []byte("path-A"))
	l2.AddEntry("evt-2b", "cause-1", []byte("path-B"))

	h1, _ := l1.Checkpoint()
	h2, _ := l2.Checkpoint()

	if fmt.Sprintf("%x", h1) == fmt.Sprintf("%x", h2) {
		t.Error("expected divergent execution paths to yield different checkpoint hashes, but they were identical")
	}
}
