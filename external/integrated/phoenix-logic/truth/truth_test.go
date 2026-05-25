package truth

import (
	"testing"
)

func TestTruthLedgerIntegrity(t *testing.T) {
	l := NewTruthLedger()

	payloads := []string{
		"event 1",
		"event 2",
		"event 3",
	}

	for _, p := range payloads {
		if _, err := l.Append(p); err != nil {
			t.Fatalf("Failed to append payload: %v", err)
		}
	}

	if ok, err := l.Verify(); !ok || err != nil {
		t.Errorf("Ledger verification failed: %v", err)
	}

	// Tamper with payload
	l.Entries[1].Payload = "tampered event"
	if ok, _ := l.Verify(); ok {
		t.Error("Expected verification failure after tampering, but it passed")
	}
}

func TestTruthLedgerHashChainTamper(t *testing.T) {
	l := NewTruthLedger()

	l.Append("event 1")
	l.Append("event 2")

	if ok, err := l.Verify(); !ok || err != nil {
		t.Fatalf("Initial verification failed: %v", err)
	}

	// Tamper with hash chain
	l.Entries[0].Hash = "corrupted_hash"
	if ok, _ := l.Verify(); ok {
		t.Error("Expected verification failure after tampering with hash chain, but it passed")
	}
}
