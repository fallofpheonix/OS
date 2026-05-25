package identity

import (
	"testing"
)

func TestReplayIdentity_StartSession(t *testing.T) {
	ri := NewReplayIdentity()
	replayID := "test-replay-001"

	session, err := ri.StartSession(replayID)
	if err != nil {
		t.Fatalf("StartSession failed: %v", err)
	}

	if session.ID != replayID {
		t.Errorf("Expected session ID %s, got %s", replayID, session.ID)
	}
	if session.StartTime.IsZero() {
		t.Error("Expected StartTime to be set, but it's zero")
	}
	if session.InitialHash != "initial_placeholder_hash" {
		t.Errorf("Expected initial hash to be 'initial_placeholder_hash', got %s", session.InitialHash)
	}

	// Test with empty replay ID
	_, err = ri.StartSession("")
	if err == nil {
		t.Error("Expected error for empty replay ID, got nil")
	} else if err.Error() != "replay ID cannot be empty" {
		t.Errorf("Expected 'replay ID cannot be empty' error, got: %v", err)
	}
}

func TestReplayIdentity_EndSession(t *testing.T) {
	ri := NewReplayIdentity()
	replayID := "test-replay-002"
	session, _ := ri.StartSession(replayID)

	// Test verified case
	finalHashVerified := "initial_placeholder_hash" // Same as InitialHash
	report, err := ri.EndSession(session, finalHashVerified)
	if err != nil {
		t.Fatalf("EndSession for verified case failed: %v", err)
	}
	if report.ReplayID != replayID {
		t.Errorf("Expected report ReplayID %s, got %s", replayID, report.ReplayID)
	}
	if report.HashBefore != session.InitialHash {
		t.Errorf("Expected report HashBefore %s, got %s", session.InitialHash, report.HashBefore)
	}
	if report.HashAfter != finalHashVerified {
		t.Errorf("Expected report HashAfter %s, got %s", finalHashVerified, report.HashAfter)
	}
	if report.Divergence {
		t.Error("Expected no divergence for verified case")
	}
	if report.Status != "VERIFIED" {
		t.Errorf("Expected status VERIFIED, got %s", report.Status)
	}

	// Test diverged case
	finalHashDiverged := "diverged_hash_value"
	report, err = ri.EndSession(session, finalHashDiverged)
	if err != nil {
		t.Fatalf("EndSession for diverged case failed: %v", err)
	}
	if !report.Divergence {
		t.Error("Expected divergence for diverged case")
	}
	if report.Status != "DIVERGED" {
		t.Errorf("Expected status DIVERGED, got %s", report.Status)
	}

	// Test with nil session
	_, err = ri.EndSession(nil, "some_hash")
	if err == nil {
		t.Error("Expected error for nil session, got nil")
	} else if err.Error() != "replay session cannot be nil" {
		t.Errorf("Expected 'replay session cannot be nil' error, got: %v", err)
	}

	// Test with empty final hash
	_, err = ri.EndSession(session, "")
	if err == nil {
		t.Error("Expected error for empty final hash, got nil")
	} else if err.Error() != "final state hash cannot be empty" {
		t.Errorf("Expected 'final state hash cannot be empty' error, got: %v", err)
	}
}

func TestReplayIdentity_CheckDivergence(t *testing.T) {
	ri := NewReplayIdentity()
	if ri.CheckDivergence("hash1", "hash1") {
		t.Error("Expected no divergence for identical hashes")
	}
	if !ri.CheckDivergence("hash1", "hash2") {
		t.Error("Expected divergence for different hashes")
	}
}

func TestHashVerifier_GenerateAndVerifyHash(t *testing.T) {
	hv := NewHashVerifier()
	data1 := []byte("hello world")
	data2 := []byte("hello world")
	data3 := []byte("another string")

	hash1, err := hv.GenerateHash(data1)
	if err != nil {
		t.Fatalf("GenerateHash failed for data1: %v", err)
	}
	hash2, err := hv.GenerateHash(data2)
	if err != nil {
		t.Fatalf("GenerateHash failed for data2: %v", err)
	}
	hash3, err := hv.GenerateHash(data3)
	if err != nil {
		t.Fatalf("GenerateHash failed for data3: %v", err)
	}

	if hash1 != hash2 {
		t.Errorf("Expected identical hashes for identical data, got %s and %s", hash1, hash2)
	}
	if hash1 == hash3 {
		t.Error("Expected different hashes for different data")
	}

	// Verify hash
	ok, err := hv.VerifyHash(data1, hash1)
	if err != nil {
		t.Fatalf("VerifyHash failed for valid hash: %v", err)
	}
	if !ok {
		t.Error("Expected valid hash to verify successfully")
	}

	ok, err = hv.VerifyHash(data1, hash3) // data1 with hash3
	if err != nil {
		t.Fatalf("VerifyHash failed for invalid hash comparison: %v", err)
	}
	if ok {
		t.Error("Expected invalid hash to not verify successfully")
	}

	// Test nil data
	_, err = hv.GenerateHash(nil)
	if err == nil {
		t.Error("Expected error for nil data in GenerateHash, got nil")
	} else if err.Error() != "data cannot be nil" {
		t.Errorf("Expected 'data cannot be nil' error, got: %v", err)
	}
}

func TestDivergenceDetector_Detect(t *testing.T) {
	dd := NewDivergenceDetector()
	if dd.Detect("hashA", "hashA") {
		t.Error("Expected no divergence for identical hashes")
	}
	if !dd.Detect("hashA", "hashB") {
		t.Error("Expected divergence for different hashes")
	}
}

func TestDivergenceDetector_Analyze(t *testing.T) {
	dd := NewDivergenceDetector()

	// No divergence
	hashes1 := []string{"a", "b", "c"}
	hashes2 := []string{"a", "b", "c"}
	diverged, msg := dd.Analyze(hashes1, hashes2)
	if diverged {
		t.Errorf("Expected no divergence, got true. Message: %s", msg)
	}
	if msg != "No divergence detected." {
		t.Errorf("Expected 'No divergence detected.' message, got: %s", msg)
	}

	// Divergence at an early step
	hashes3 := []string{"a", "X", "c"}
	hashes4 := []string{"a", "Y", "c"}
	diverged, msg = dd.Analyze(hashes3, hashes4)
	if !diverged {
		t.Error("Expected divergence, got false")
	}
	expectedMsg := "Divergence detected at step 1: initial hash X, final hash Y"
	if msg != expectedMsg {
		t.Errorf("Expected message '%s', got: '%s'", expectedMsg, msg)
	}

	// Different lengths
	hashes5 := []string{"a", "b"}
	hashes6 := []string{"a", "b", "c"}
	diverged, msg = dd.Analyze(hashes5, hashes6)
	if !diverged {
		t.Error("Expected divergence due to length difference, got false")
	}
	expectedMsg = "Divergence detected: hash chain lengths differ. Initial: 2, Final: 3"
	if msg != expectedMsg {
		t.Errorf("Expected message '%s', got: '%s'", expectedMsg, msg)
	}

	// Empty hash chains
	diverged, msg = dd.Analyze([]string{}, []string{})
	if diverged {
		t.Errorf("Expected no divergence for empty chains, got true. Message: %s", msg)
	}
}
