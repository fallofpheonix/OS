package ledger

import (
	"bytes"
	"path/filepath"
	"testing"
)

func TestEndToEndPersistenceAndRecovery(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "e2e_ledger.log")

	// Phase 1 - Write
	p1, err := NewPersistor(path)
	if err != nil {
		t.Fatalf("Phase 1: Failed to create persistor: %v", err)
	}

	header := LedgerFileHeader{Version: "1.0", GenesisID: "GENESIS", Timestamp: 12345, Algorithm: "SHA256", FixedPointDivisor: 1000000}
	if err := p1.WriteHeader(header); err != nil {
		t.Fatalf("Phase 1: Failed to write header: %v", err)
	}

	liveLedger := NewLedger(nil).WithPersistor(p1)

	for i := 0; i < 10; i++ {
		// Use distinct EventIDs to verify data later
		eventID := "EVT-E2E"
		if i == 0 {
			eventID = "EVT-START"
		} else if i == 9 {
			eventID = "EVT-END"
		}

		err := liveLedger.AddEntryV2(eventID, "CAUSE", uint64(i), []byte("payload"), "", []byte("PRE"), []byte("POST"), "1.0")
		if err != nil {
			t.Fatalf("Phase 1: Failed to add entry %d: %v", i, err)
		}
	}

	if len(liveLedger.Heads) != 1 {
		t.Fatalf("Phase 1: Expected exactly 1 head, got %d", len(liveLedger.Heads))
	}
	recordedHeadHash := liveLedger.Heads[0]

	// Simulate process death by discarding all references
	p1 = nil
	liveLedger = nil

	// Phase 2 - Recover
	replayer, err := NewReplayer(path)
	if err != nil {
		t.Fatalf("Phase 2: Failed to create replayer: %v", err)
	}

	recoveredLedger, err := replayer.Replay()
	if err != nil {
		t.Fatalf("Phase 2: Failed to replay ledger: %v", err)
	}

	// Assert Head Hash matches
	if len(recoveredLedger.Heads) != 1 {
		t.Fatalf("Phase 2: Expected exactly 1 head on recovered ledger, got %d", len(recoveredLedger.Heads))
	}

	if !bytes.Equal(recoveredLedger.Heads[0], recordedHeadHash) {
		t.Fatalf("Phase 2: Head Hash mismatch: expected %x, got %x", recordedHeadHash, recoveredLedger.Heads[0])
	}

	// Assert entry count
	if len(recoveredLedger.Entries) != 10 {
		t.Fatalf("Phase 2: Expected 10 entries in recovered ledger, got %d", len(recoveredLedger.Entries))
	}

	// In the map we can iterate to find the specific events, but since the replayer creates a fresh map,
	// we just need to verify the specific hashes or just search the values.
	foundStart := false
	foundEnd := false
	for _, entry := range recoveredLedger.Entries {
		if entry.EventID == "EVT-START" && entry.LogicalTick == 0 {
			foundStart = true
		}
		if entry.EventID == "EVT-END" && entry.LogicalTick == 9 {
			foundEnd = true
		}
	}

	if !foundStart {
		t.Error("Phase 2: Could not find entry 0 (EVT-START)")
	}
	if !foundEnd {
		t.Error("Phase 2: Could not find entry 9 (EVT-END)")
	}
}
