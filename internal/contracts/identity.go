package contracts

import (
	"encoding/hex"
	"fmt"
)

// Hash represents a canonical 32-byte SHA-256 digest.
type Hash [32]byte

func (h Hash) String() string {
	return hex.EncodeToString(h[:])
}

// ParseHash converts a hex string to a Hash.
func ParseHash(s string) (Hash, error) {
	var h Hash
	b, err := hex.DecodeString(s)
	if err != nil {
		return h, err
	}
	if len(b) != 32 {
		return h, fmt.Errorf("invalid hash length: %d", len(b))
	}
	copy(h[:], b)
	return h, nil
}

// NodeID represents a 32-byte Ed25519 Public Key.
// It is the primary identity used for both gossip and consensus.
type NodeID [32]byte

// BlockID represents the unique hash of a block header.
type BlockID Hash
