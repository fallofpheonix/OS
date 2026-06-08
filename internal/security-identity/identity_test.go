package identity

import (
	"os"
	"testing"
)

func TestIdentity_EncryptedPersistence(t *testing.T) {
	path := "test_identity.enc"
	passphrase := "correct-horse-battery-staple"
	defer os.Remove(path)

	id, err := GenerateNewIdentity()
	if err != nil {
		t.Fatalf("Generation failed: %v", err)
	}

	// 1. Save
	if err := id.SaveEncrypted(path, passphrase); err != nil {
		t.Fatalf("Save failed: %v", err)
	}

	// 2. Load correctly
	loaded, err := LoadEncrypted(path, passphrase)
	if err != nil {
		t.Fatalf("Load failed: %v", err)
	}
	if loaded.ExportPublicKey() != id.ExportPublicKey() {
		t.Error("Public key mismatch after reload")
	}

	// 3. Load with wrong passphrase
	_, err = LoadEncrypted(path, "wrong-password")
	if err == nil {
		t.Error("Expected error with wrong passphrase, got nil")
	}
}
