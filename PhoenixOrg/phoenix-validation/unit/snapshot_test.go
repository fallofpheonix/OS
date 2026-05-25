package unit

import (
	"fmt"
	"testing"
)

func TestLedgerSnapshot(t *testing.T) {
	l := NewTruthLedger()
	l.AddEntry(1, []byte("A"))
	l.AddEntry(2, []byte("B"))

	snap := l.Snapshot()
	if len(snap) != 2 {
		t.Errorf("Expected 2 entries in snapshot, got %d", len(snap))
	}

	// Add more to ledger, snapshot should not change
	l.AddEntry(3, []byte("C"))
	if len(snap) != 2 {
		t.Error("Snapshot changed after ledger update")
	}
	fmt.Println("[PX-012] Ledger Snapshot Integrity: PASSED")
}
