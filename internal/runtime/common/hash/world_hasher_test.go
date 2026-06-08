package hash_test

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/hash"
)

func TestDeterministicReconstructionHash(t *testing.T) {
	manifest := &hash.SubsystemManifest{Subsystems: []string{"Authority", "PhoenixMind"}}
	hasher := &hash.WorldStateHasher{Manifest: manifest}

	// 100 iterations of identical ledger input
	const iterations = 100
	hashes := make([]string, iterations)

	for i := 0; i < iterations; i++ {
		// Simulate subsystem state hashes
		subsystemHashes := map[string]string{
			"Authority":   "a1b2c3d4e5f6...",
			"PhoenixMind": "1a2b3c4d5e6f...",
		}
		hashes[i] = hasher.ComputeHash(subsystemHashes)
	}

	// Verify all hashes are identical
	for i := 1; i < iterations; i++ {
		if hashes[i] != hashes[0] {
			t.Fatalf("Hash divergence at iteration %d: expected %s, got %s", i, hashes[0], hashes[i])
		}
	}
}

func TestAdversarialHashDivergence(t *testing.T) {
	manifest := &hash.SubsystemManifest{Subsystems: []string{"Authority", "PhoenixMind"}}
	hasher := &hash.WorldStateHasher{Manifest: manifest}

	subsystemHashes := map[string]string{
		"Authority":   "a1b2c3d4e5f6...",
		"PhoenixMind": "1a2b3c4d5e6f...",
	}

	hash1 := hasher.ComputeHash(subsystemHashes)

	// Adversarial: mutate one subsystem hash
	subsystemHashesMutated := map[string]string{
		"Authority":   "b1b2c3d4e5f6...", // Mutated
		"PhoenixMind": "1a2b3c4d5e6f...",
	}

	hash2 := hasher.ComputeHash(subsystemHashesMutated)

	if hash1 == hash2 {
		t.Fatal("Adversarial mutation failed: hash did not diverge")
	}
}
