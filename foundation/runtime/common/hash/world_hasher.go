// Package hash provides deterministic hashing utilities for the Phoenix system state.
// Core Domain Logic: Implements canonical world state hashing via a Merkle-inspired hash chain,
// ensuring system-wide consistency across distributed nodes by enforcing deterministic ordering.
package hash

import (
	"crypto/sha256"
	"fmt"
	"sort"
	"strings"
)

// SubsystemManifest defines the deterministic ordering of subsystems in the Merkle tree.
// Internal State: Encapsulates the set of subsystem identifiers participating in the global state.
// API Scope: Public within the Nucleus for state consistency enforcement.
// Concurrency: Modification of Subsystems slice is not thread-safe; intended as immutable after initialization.
type SubsystemManifest struct {
	Subsystems []string
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// GetHash returns the canonical hash of the manifest itself.
// I/O: None.
// Side Effects: None.
// Complexity: O(N log N) where N is the number of subsystems (due to internal sorting).
func (m *SubsystemManifest) GetHash() string {
	sorted := make([]string, len(m.Subsystems))
	copy(sorted, m.Subsystems)
	sort.Strings(sorted)
	
	h := sha256.New()
	h.Write([]byte(strings.Join(sorted, ",")))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// WorldStateHasher computes the Merkle root of the system state.
// Internal State: References a SubsystemManifest to ensure consistent leaf ordering.
// API Scope: Public for world state integrity verification.
// Concurrency: Thread-safe for concurrent hash computations as it performs no internal state mutation.
type WorldStateHasher struct {
	Manifest *SubsystemManifest
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// ComputeHash computes the Merkle root of the provided subsystem hashes.
// I/O: None.
// Side Effects: None.
// Complexity: O(N log N) to sort keys + O(N) to process entries, where N is the number of subsystems in the manifest.
func (w *WorldStateHasher) ComputeHash(subsystemHashes map[string]string) string {
	// 1. Sort subsystems according to manifest (determinism contract)
	sortedKeys := make([]string, len(w.Manifest.Subsystems))
	copy(sortedKeys, w.Manifest.Subsystems)
	sort.Strings(sortedKeys)

	// 2. Build Merkle tree (or concatenated hash chain for simplicity as per WorldState definition)
	// We'll use a sequential hash of hashes in manifest order for this implementation.
	h := sha256.New()
	
	// Add manifest hash first as leaf
	h.Write([]byte(w.Manifest.GetHash()))
	
	for _, key := range sortedKeys {
		hash, ok := subsystemHashes[key]
		if !ok {
			// If missing, use zero-value leaf hash (or handle as error)
			h.Write([]byte("0000000000000000000000000000000000000000000000000000000000000000"))
		} else {
			h.Write([]byte(hash))
		}
	}
	
	return fmt.Sprintf("%x", h.Sum(nil))
}
