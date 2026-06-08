package protocol

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

// CalculateMerkleRoot computes the Balanced Binary Merkle Tree root with Carry-Up.
// NOTE: This follows the frozen PROTOCOL-014 pseudocode.
func CalculateMerkleRoot(events []contracts.Event) (contracts.Hash, error) {
	if len(events) == 0 {
		return sha256.Sum256([]byte("")), nil
	}

	// 1. Leaf Layer: SHA256(0x00 | leaf_digest)
	nodes := make([]contracts.Hash, len(events))
	for i, e := range events {
		digest, err := DigestEvent(e)
		if err != nil {
			return contracts.Hash{}, fmt.Errorf("merkle: failed to digest event %d: %w", i, err)
		}

		h := sha256.New()
		h.Write([]byte{0x00})
		h.Write(digest[:])
		copy(nodes[i][:], h.Sum(nil))
	}

	// 2. Iterative Reduction
	for len(nodes) > 1 {
		var nextLevel []contracts.Hash
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				// Pair: SHA256(0x01 | Left | Right)
				h := sha256.New()
				h.Write([]byte{0x01})
				h.Write(nodes[i][:])
				h.Write(nodes[i+1][:])
				var res contracts.Hash
				copy(res[:], h.Sum(nil))
				nextLevel = append(nextLevel, res)
			} else {
				// Carry-Up Rule: Odd node is promoted without hashing
				nextLevel = append(nextLevel, nodes[i])
			}
		}
		nodes = nextLevel
	}

	return nodes[0], nil
}

// GenerateMerkleProof constructs an inclusion proof for a leaf at the given index.
// NOTE: Follows the PROTOCOL-017 Carry-Up proof generation pseudocode.
func GenerateMerkleProof(events []contracts.Event, targetIndex int) ([]contracts.Hash, error) {
	if targetIndex < 0 || targetIndex >= len(events) {
		return nil, fmt.Errorf("merkle: index %d out of bounds (total %d)", targetIndex, len(events))
	}

	// Initial leaf digests
	nodes := make([]contracts.Hash, len(events))
	for i, e := range events {
		d, err := DigestEvent(e)
		if err != nil {
			return nil, err
		}
		h := sha256.New()
		h.Write([]byte{0x00})
		h.Write(d[:])
		copy(nodes[i][:], h.Sum(nil))
	}

	var proof []contracts.Hash
	currIdx := targetIndex
	for len(nodes) > 1 {
		isRightChild := (currIdx%2 == 1)
		isLastNode := (currIdx == len(nodes)-1)

		if isRightChild {
			proof = append(proof, nodes[currIdx-1])
		} else if !isLastNode {
			proof = append(proof, nodes[currIdx+1])
		}
		// Carry-up level: if !isRightChild && isLastNode, no hash is added to proof.

		// Build next level
		var nextLevel []contracts.Hash
		for i := 0; i < len(nodes); i += 2 {
			if i+1 < len(nodes) {
				h := sha256.New()
				h.Write([]byte{0x01})
				h.Write(nodes[i][:])
				h.Write(nodes[i+1][:])
				var res contracts.Hash
				copy(res[:], h.Sum(nil))
				nextLevel = append(nextLevel, res)
			} else {
				nextLevel = append(nextLevel, nodes[i])
			}
		}
		nodes = nextLevel
		currIdx /= 2
	}

	return proof, nil
}

// VerifyMerkleProof validates an inclusion proof using level-driven carry-up logic.
// NOTE: Follows the corrected PROTOCOL-017 level-driven verification pseudocode.
func VerifyMerkleProof(leaf contracts.Hash, index int, totalLeaves int, proof []contracts.Hash, root contracts.Hash) bool {
	if totalLeaves == 0 {
		return false
	}

	h := sha256.New()
	h.Write([]byte{0x00})
	h.Write(leaf[:])
	currentHash := h.Sum(nil)

	currIdx := index
	currTotal := totalLeaves
	proofPtr := 0

	for currTotal > 1 {
		isRightChild := (currIdx%2 == 1)
		isLastNode := (currIdx == currTotal-1)

		// If it's an odd node at the end, it was carried up (no sibling hash used)
		if !(!isRightChild && isLastNode) {
			if proofPtr >= len(proof) {
				return false
			}
			sibling := proof[proofPtr]
			proofPtr++

			h := sha256.New()
			h.Write([]byte{0x01})
			if isRightChild {
				h.Write(sibling[:])
				h.Write(currentHash)
			} else {
				h.Write(currentHash)
				h.Write(sibling[:])
			}
			currentHash = h.Sum(nil)
		}

		currIdx /= 2
		currTotal = (currTotal + 1) / 2
	}

	return bytes.Equal(currentHash, root[:]) && proofPtr == len(proof)
}
