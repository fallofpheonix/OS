package ledger

import (
	"path/filepath"
	"testing"
)

func TestReplay_HeadHashInvariant(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "vertical_slice.jsonl")

	// Phase 1: Normal Operation
	p1, _ := NewPersistor(path)

	// Genesis
	e0 := NewEvent(0, EventGenesis, []byte("{}"), "", "SYSTEM")
	p1.Append(e0)

	// Some activity
	e1 := NewEvent(1, EventEnforce, []byte(`{"action":"QUARANTINE"}`), e0.Hash, "WARDEN")
	p1.Append(e1)

	expectedHash := e1.Hash
	p1.Close()

	// Phase 2: System Crash / Restart
	// We simulate a fresh start by creating a new Persistor and Replayer
	p2, _ := NewPersistor(path)
	defer p2.Close()

	replayer := NewReplayer(p2)

	// RECONSTRUCT
	actualHash, err := replayer.Reconstruct()
	if err != nil {
		t.Fatalf("Reconstruction failed: %v", err)
	}

	// VERIFY INVARIANT
	if actualHash != expectedHash {
		t.Errorf("Head Hash Invariant Broken!\nExpected: %s\nActual:   %s", expectedHash, actualHash)
	} else {
		t.Logf("Head Hash Invariant Verified: %s", actualHash)
	}
}

func TestReplay_EmptyLedger(t *testing.T) {
	tmpDir := t.TempDir()
	path := filepath.Join(tmpDir, "empty.jsonl")

	p, _ := NewPersistor(path)
	defer p.Close()

	replayer := NewReplayer(p)
	_, err := replayer.Reconstruct()
	if err == nil {
		t.Error("Expected error for empty ledger, got nil")
	}
}
