package consensus

import (
	"bytes"
	"crypto/ed25519"
	"encoding/hex"
	"fmt"
	"sort"

	"github.com/fallofpheonix/phoenix/internal/contracts"
	"github.com/fallofpheonix/phoenix/internal/protocol"
)

// Re-export common types for convenience
type Hash = contracts.Hash
type NodeID = contracts.NodeID
type QuorumCertificate = contracts.QuorumCertificate
type SignedEnvelope = contracts.SignedEnvelope
type SignatureEntry = contracts.SignatureEntry

const (
	MsgEvent       = uint8(0)
	MsgVote        = uint8(1)
	MsgCertificate = uint8(2)
)

// Sign signs the envelope using the provided private key.
func SignEnvelope(e *SignedEnvelope, priv ed25519.PrivateKey) error {
	e.Validator = contracts.NodeID(priv.Public().(ed25519.PublicKey))
	digest, err := protocol.DigestEnvelope(*e)
	if err != nil {
		return err
	}
	e.Signature = ed25519.Sign(priv, digest[:])
	return nil
}

// VerifyEnvelope checks the envelope signature.
func VerifyEnvelope(e *SignedEnvelope) bool {
	if len(e.Signature) != ed25519.SignatureSize {
		return false
	}
	digest, err := protocol.DigestEnvelope(*e)
	if err != nil {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(e.Validator[:]), digest[:], e.Signature)
}

// CheckQuorum verifies if a set of signatures meets the 2f+1 requirement.
// CONSENSUS-014/015: Quorum verification with uniqueness check.
func CheckQuorum(validators [][]byte, signatures []SignatureEntry, digest Hash) (bool, error) {
	if len(validators) == 0 {
		return true, nil // Bootstrap mode
	}

	validUniqueCount := 0
	f := (len(validators) - 1) / 3
	threshold := 2*f + 1

	seen := make(map[string]bool)

	for _, entry := range signatures {
		pubKeyHex := hex.EncodeToString(entry.ValidatorID[:])
		if seen[pubKeyHex] {
			continue // Deduplicate multiple signatures from same identity
		}

		// Verify membership
		authorized := false
		for _, v := range validators {
			if bytes.Equal(v, entry.ValidatorID[:]) {
				authorized = true
				break
			}
		}

		if authorized {
			if ed25519.Verify(ed25519.PublicKey(entry.ValidatorID[:]), digest[:], entry.Signature) {
				validUniqueCount++
				seen[pubKeyHex] = true
			}
		}
	}

	return validUniqueCount >= threshold, nil
}

// ElectLeader selects a deterministic leader for a given round using round-robin.
// CONSENSUS-134: Deterministic Leader Selection.
func ElectLeader(validators [][]byte, round uint64) ([]byte, error) {
	if len(validators) == 0 {
		return nil, fmt.Errorf("no validators available for election")
	}

	// 1. Sort validators to ensure identical ordering on all nodes
	sorted := make([][]byte, len(validators))
	copy(sorted, validators)
	sort.Slice(sorted, func(i, j int) bool {
		return bytes.Compare(sorted[i], sorted[j]) < 0
	})

	// 2. Select leader using round-robin
	index := round % uint64(len(sorted))
	return sorted[index], nil
}
