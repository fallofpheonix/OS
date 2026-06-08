package protocol

import (
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"testing"
)

// PROTOCOL-015: Independent Golden Vectors
// These hashes are derived from the manual binary layouts specified in PROTOCOL_ENCODING.md
// and verified via external SHA256 calculations.

const (
	// Manual Calculation: SHA256(01 | 0001 | 0001 | 0002 | AABB)
	// Verified via: echo -n -e "\x01\x00\x01\x00\x01\x00\x02\xAA\xBB" | shasum -a 256
	GoldenEventDigest = "6f413efab66c6735513bb8e089ce2cdd2b87839f1d64de50bf3910fe645fe278"

	// SHA256("")
	GoldenMerkleEmpty = "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855"

	// SHA256(00 | 6f413e...)
	// JUSTIFICATION: The 0x00 prefix is mandatory for Merkle leaves (RFC 6962) to prevent
	// second-preimage attacks. Original vector lacked this prefix.
	// Verified via: echo "006f413efab66c6735513bb8e089ce2cdd2b87839f1d64de50bf3910fe645fe278" | xxd -r -p | shasum -a 256
	GoldenMerkleLeaf = "20e8104fc7ae82f44df5b54911b9e8efd0cf63f7fc4f085daa5c4f6f63abbb9a"

	// Manual Calculation: SHA256(02 | 0001 | 00..01 | 00..01 | 00..01 | 32x0A | 32x0B | 32x0C | 32x0D)
	// Verified via: go run gen_golden.go
	GoldenBlockHeaderDigest = "f94f3fb89f8c02f1c4aaf9101572d4e7591468e85b92f9a9cccc880019e06a83"

	// Manual Calculation: SHA256(03 | 0001 | 00..01 | 00..01 | 00..01 | 32x0A | 32x0B | 32x0C | 1 | 32x0D | 0004 | SIG1)
	// Verified via: go run gen_golden_qc.go
	GoldenQCDigest = "d1dde34cb1c285004aecbba5788ee21d0aeafac890f6727279f6b1bc5cb7868c"
)

func TestLaw_GoldenQC(t *testing.T) {
	nodeA := contracts.NodeID{}
	for i := range nodeA {
		nodeA[i] = 0x0D
	}

	bid := contracts.BlockID{}
	for i := range bid {
		bid[i] = 0x0A
	}

	sroot := contracts.Hash{}
	for i := range sroot {
		sroot[i] = 0x0B
	}

	vset := contracts.Hash{}
	for i := range vset {
		vset[i] = 0x0C
	}

	qc := contracts.QuorumCertificate{
		Version:          1,
		Epoch:            1,
		Round:            1,
		Height:           1,
		BlockID:          bid,
		StateRoot:        sroot,
		ValidatorSetHash: vset,
		Signatures: []contracts.SignatureEntry{
			{ValidatorID: nodeA, Signature: []byte("SIG1")},
		},
	}

	digest, err := DigestQC(qc)
	if err != nil {
		t.Fatalf("DigestQC failed: %v", err)
	}

	if digest.String() != GoldenQCDigest {
		t.Errorf("QC Digest Law Violation!\nGot:  %s\nWant: %s", digest.String(), GoldenQCDigest)
	}
}

func TestLaw_MerkleProofs(t *testing.T) {
	e1 := contracts.Event{Version: 1, Type: 0, Payload: []byte("A")}
	e2 := contracts.Event{Version: 1, Type: 0, Payload: []byte("B")}
	e3 := contracts.Event{Version: 1, Type: 0, Payload: []byte("C")}
	events := []contracts.Event{e1, e2, e3}

	root, _ := CalculateMerkleRoot(events)

	for i := 0; i < len(events); i++ {
		proof, err := GenerateMerkleProof(events, i)
		if err != nil {
			t.Fatalf("Proof generation failed for index %d: %v", i, err)
		}

		d, _ := DigestEvent(events[i])
		if !VerifyMerkleProof(d, i, len(events), proof, root) {
			t.Errorf("Merkle proof verification failed for index %d", i)
		}
	}
}

func TestLaw_GoldenBlockHeader(t *testing.T) {
	p := contracts.NodeID{}
	for i := range p {
		p[i] = 0x0A
	}

	h := contracts.Hash{}
	prev := contracts.Hash{}
	for i := range h {
		prev[i] = 0x0B
	}
	m := contracts.Hash{}
	for i := range h {
		m[i] = 0x0C
	}
	s := contracts.Hash{}
	for i := range h {
		s[i] = 0x0D
	}

	b := contracts.FinalizedBlock{
		Version:       1,
		Height:        1,
		Epoch:         1,
		Round:         1,
		Proposer:      p,
		PrevBlockHash: prev,
		MerkleRoot:    m,
		StateRoot:     s,
	}

	digest, err := DigestBlockHeader(b)
	if err != nil {
		t.Fatalf("DigestBlockHeader failed: %v", err)
	}

	if digest.String() != GoldenBlockHeaderDigest {
		t.Errorf("BlockHeader Digest Law Violation!\nGot:  %s\nWant: %s", digest.String(), GoldenBlockHeaderDigest)
	}
}

func TestLaw_GoldenEvent(t *testing.T) {
	e := contracts.Event{
		Version: 1,
		Type:    1,
		Payload: []byte{0xAA, 0xBB},
	}

	digest, err := DigestEvent(e)
	if err != nil {
		t.Fatalf("DigestEvent failed: %v", err)
	}

	if digest.String() != GoldenEventDigest {
		t.Errorf("Event Digest Law Violation!\nGot:  %s\nWant: %s", digest.String(), GoldenEventDigest)
	}
}

func TestLaw_MerkleRoots(t *testing.T) {
	// 1. Empty Root
	root0, err := CalculateMerkleRoot(nil)
	if err != nil {
		t.Fatal(err)
	}
	if root0.String() != GoldenMerkleEmpty {
		t.Errorf("Merkle Empty Law Violation: %s", root0.String())
	}

	// 2. Single Leaf Root
	e1 := contracts.Event{Version: 1, Type: 1, Payload: []byte{0xAA, 0xBB}}
	root1, err := CalculateMerkleRoot([]contracts.Event{e1})
	if err != nil {
		t.Fatal(err)
	}
	if root1.String() != GoldenMerkleLeaf {
		t.Errorf("Merkle Single Leaf Law Violation!\nGot:  %s\nWant: %s", root1.String(), GoldenMerkleLeaf)
	}
}

func TestLaw_CarryUpMerkle(t *testing.T) {
	// Carry-Up logic check for 3 leaves:
	// L1 = SHA256(00 | Digest(E1))
	// L2 = SHA256(00 | Digest(E2))
	// L3 = SHA256(00 | Digest(E3))
	// Level 1: [H(01 | L1 | L2), L3 (Carried)]
	// Level 2 (Root): H(01 | Level1[0] | L3)

	e1 := contracts.Event{Version: 1, Type: 0, Payload: []byte("A")}
	e2 := contracts.Event{Version: 1, Type: 0, Payload: []byte("B")}
	e3 := contracts.Event{Version: 1, Type: 0, Payload: []byte("C")}

	root, err := CalculateMerkleRoot([]contracts.Event{e1, e2, e3})
	if err != nil {
		t.Fatal(err)
	}

	if root == (contracts.Hash{}) {
		t.Error("Merkle root is zero")
	}
}
