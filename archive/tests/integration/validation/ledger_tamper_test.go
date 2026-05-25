package validation

import (
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/ledger/src"
)

// TestLedgerTamper verifies that tampering validation hashes or state fields is detected.
func TestLedgerTamper(t *testing.T) {
	l := ledger.NewLedger(nil)

	err := l.AddEntryV2("evt-1", "cause-1", []byte("payload"), "hash-1", "SAFE", "WATCH", "v1.0")
	if err != nil {
		t.Fatalf("failed to add ledger entry: %v", err)
	}

	// Mutate validation hash directly
	for key, entry := range l.Entries {
		entry.ValidationHash = []byte("badvalidationhash")
		l.Entries[key] = entry
	}

	if err := l.Verify(); err == nil {
		t.Error("expected verify to fail after tampering validation hash, but it succeeded")
	} else {
		t.Logf("tamper verify failed correctly: %v", err)
	}
}
