package validation

import (
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/ledger/src"
)

// TestTruthMutation verifies that a modification of payload bytes in the ledger is detected.
func TestTruthMutation(t *testing.T) {
	l := ledger.NewLedger(nil)
	
	err := l.AddEntryV2("evt-1", "cause-1", []byte("valid payload"), "hash-1", "SAFE", "WATCH", "v1.0")
	if err != nil {
		t.Fatalf("failed to add ledger entry: %v", err)
	}

	// Verify ledger is initially clean
	if err := l.Verify(); err != nil {
		t.Fatalf("ledger initial verify failed: %v", err)
	}

	// Find the entry in map and mutate payload bytes
	for key, entry := range l.Entries {
		entry.Payload = []byte("mutated payload")
		l.Entries[key] = entry
	}

	// Verify must fail now due to hash mismatch
	if err := l.Verify(); err == nil {
		t.Error("expected verify to fail after mutating payload, but it succeeded")
	} else {
		t.Logf("verify failed correctly: %v", err)
	}
}
