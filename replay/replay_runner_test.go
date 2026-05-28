package replay

import (
	"testing"
	"github.com/fallofpheonix/PheonixTruth/src"
	"github.com/fallofpheonix/replay/identity"
)

type MockProvider struct {
	Entries []ledger.LedgerEntry
}

func (m *MockProvider) LoadTrace(traceID string) ([]ledger.LedgerEntry, error) {
	return m.Entries, nil
}

func TestReplayRunner_Success(t *testing.T) {
	l := ledger.NewLedger(nil)
	
	// Create a perfect trace
	entry1 := ledger.LedgerEntry{
		LogicalTick: 0,
		EventID:     "evt-1",
		Payload:     []byte("data-1"),
	}
	entry1.Hash = l.ComputeHash(entry1)

	entry2 := ledger.LedgerEntry{
		LogicalTick: 1,
		EventID:     "evt-2",
		Payload:     []byte("data-2"),
		ParentIDs:   [][]byte{entry1.Hash},
	}
	entry2.Hash = l.ComputeHash(entry2)

	provider := &MockProvider{
		Entries: []ledger.LedgerEntry{entry1, entry2},
	}

	runner := NewReplayRunner(provider, l)
	report, err := runner.ExecuteReplay("test-trace")

	if err != nil {
		t.Fatalf("ExecuteReplay failed: %v", err)
	}

	if report.Divergence {
		t.Errorf("Expected no divergence, got %s: %s", report.DivType, report.Message)
	}

	if report.Status != "VERIFIED" {
		t.Errorf("Expected status VERIFIED, got %s", report.Status)
	}
}

func TestReplayRunner_HashMismatch(t *testing.T) {
	l := ledger.NewLedger(nil)
	
	entry := ledger.LedgerEntry{
		LogicalTick: 0,
		EventID:     "evt-1",
		Payload:     []byte("data-1"),
	}
	entry.Hash = l.ComputeHash(entry)
	
	// Tamper with payload but keep old hash
	tamperedEntry := entry
	tamperedEntry.Payload = []byte("tampered-data")

	provider := &MockProvider{
		Entries: []ledger.LedgerEntry{tamperedEntry},
	}

	runner := NewReplayRunner(provider, l)
	report, err := runner.ExecuteReplay("tamper-trace")

	if err != nil {
		t.Fatalf("ExecuteReplay failed: %v", err)
	}

	if !report.Divergence {
		t.Fatal("Expected divergence detection for tampered trace")
	}

	if report.DivType != identity.DivHash {
		t.Errorf("Expected divergence type HASH_MISMATCH, got %s", report.DivType)
	}
}

func TestReplayRunner_SequenceGap(t *testing.T) {
	l := ledger.NewLedger(nil)
	
	entry1 := ledger.LedgerEntry{LogicalTick: 0, EventID: "e1"}
	entry1.Hash = l.ComputeHash(entry1)

	entry3 := ledger.LedgerEntry{LogicalTick: 2, EventID: "e3"} // Gap: tick 1 missing
	entry3.Hash = l.ComputeHash(entry3)

	provider := &MockProvider{
		Entries: []ledger.LedgerEntry{entry1, entry3},
	}

	runner := NewReplayRunner(provider, l)
	report, err := runner.ExecuteReplay("gap-trace")

	if err != nil {
		t.Fatalf("ExecuteReplay failed: %v", err)
	}

	if !report.Divergence {
		t.Fatal("Expected divergence detection for sequence gap")
	}

	if report.DivType != identity.DivSequence {
		t.Errorf("Expected divergence type SEQUENCE_GAP, got %s", report.DivType)
	}
}
