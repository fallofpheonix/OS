package ledger

import "testing"

func TestLedgerIntegrity(t *testing.T) {
	l := NewLedger()
	
	// Genesis state
	if l.LastHash != "GENESIS" {
		t.Errorf("Expected genesis hash, got %s", l.LastHash)
	}

	// First commit
	h1 := l.Commit(Evidence{TraceHash: "0x1", SDI: 0.5})
	
	// Second commit
	l.Commit(Evidence{TraceHash: "0x2", SDI: 0.6})
	if l.Entries[1].PrevHash != h1 {
		t.Errorf("Hash chain broken: PrevHash of second entry != hash of first")
	}

	// Verification
	if !l.Verify() {
		t.Errorf("Ledger verification failed")
	}

	// Tamper simulation
	l.Entries[0].SDI = 0.99
	if l.Verify() {
		t.Errorf("Tampering was not detected!")
	}
	
	// Restore and check
	l.Entries[0].SDI = 0.5
	if !l.Verify() {
		t.Errorf("Verification failed after restoration")
	}

	// Linkage tampering
	l.Entries[1].PrevHash = "HACKED"
	if l.Verify() {
		t.Errorf("Linkage tampering was not detected!")
	}
}
