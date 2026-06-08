/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package identity

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// HashVerifier defines the interface for generating and verifying cryptographic hashes.
type HashVerifier interface {
	// GenerateHash takes arbitrary data and returns its SHA-256 hash as a string.
	GenerateHash(data []byte) (string, error)
	// VerifyHash compares a given hash with the hash of provided data.
	VerifyHash(data []byte, expectedHash string) (bool, error)
}

// NewHashVerifier creates a new instance of HashVerifier.
func NewHashVerifier() HashVerifier {
	return &sha256HashVerifier{}
}

type sha256HashVerifier struct{}

// GenerateHash generates a SHA-256 hash of the input data.
func (h *sha256HashVerifier) GenerateHash(data []byte) (string, error) {
	if data == nil {
		return "", fmt.Errorf("data cannot be nil")
	}
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:]), nil
}

// VerifyHash verifies if the generated hash of data matches the expected hash.
func (h *sha256HashVerifier) VerifyHash(data []byte, expectedHash string) (bool, error) {
	generatedHash, err := h.GenerateHash(data)
	if err != nil {
		return false, err
	}
	return generatedHash == expectedHash, nil
}
