package forensics

import (
	"os"
	"testing"
)

func TestForensicsAgent(t *testing.T) {
	tmpDir := "test_snapshots"
	defer os.RemoveAll(tmpDir)

	agent, err := NewForensicsAgent(tmpDir)
	if err != nil {
		t.Fatalf("failed to create agent: %v", err)
	}

	snap, err := agent.TriggerSnapshot(456)
	if err != nil {
		t.Fatalf("failed to trigger snapshot: %v", err)
	}

	if snap.PID != 456 {
		t.Errorf("expected PID 456, got %d", snap.PID)
	}

	if snap.Hash == "" {
		t.Error("snapshot hash is empty")
	}

	if _, err := os.Stat(snap.Path); os.IsNotExist(err) {
		t.Errorf("snapshot file does not exist: %s", snap.Path)
	}

	verifyHash, err := agent.VerifyHash(snap.Path)
	if err != nil {
		t.Fatalf("failed to verify hash: %v", err)
	}

	if verifyHash != snap.Hash {
		t.Errorf("hash mismatch: %s != %s", verifyHash, snap.Hash)
	}
}
