package validation

import (
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/ledger/src"
)

// TestHashFork verifies that rollback and re-entry creates a new timeline with a distinct hash chain.
func TestHashFork(t *testing.T) {
	l := ledger.NewLedger(nil)

	l.AddEntry("evt-1", "cause-1", []byte("checkpoint-state"))
	chk, _ := l.Checkpoint()

	// Normal timeline progression
	l.AddEntry("evt-2", "cause-1", []byte("event-2"))
	head1, _ := l.Checkpoint()

	// Rollback and fork timeline with different action
	if err := l.RollbackTo(chk); err != nil {
		t.Fatalf("rollback failed: %v", err)
	}

	l.AddEntry("evt-3", "cause-1", []byte("event-3"))
	head2, _ := l.Checkpoint()

	if fmt.Sprintf("%x", head1) == fmt.Sprintf("%x", head2) {
		t.Error("expected rolled back and forked timelines to have different head hashes, but they matched")
	}
}
