package protocol

import (
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"testing"
)

func TestDigest_IndependentGoldenVectors(t *testing.T) {
	// 1. Event Golden Vector (Manual Calculation Proof)
	// Input: Prefix=01, Ver=0001, Type=0001, Len=0002, Payload=AABB
	e := contracts.Event{
		Version: 1,
		Type:    1,
		Payload: []byte{0xAA, 0xBB},
	}
	digest, err := DigestEvent(e)
	if err != nil {
		t.Fatalf("DigestEvent failed: %v", err)
	}

	gotE := digest.String()
	// Verified via: echo -n -e "\x01\x00\x01\x00\x01\x00\x02\xAA\xBB" | shasum -a 256
	wantE := "6f413efab66c6735513bb8e089ce2cdd2b87839f1d64de50bf3910fe645fe278"

	if gotE != wantE {
		t.Errorf("Event Golden Vector mismatch!\nGot:  %s\nWant: %s", gotE, wantE)
	}

	// 2. Merkle Golden Vector (Empty)
	root0, _ := CalculateMerkleRoot(nil)
	gotM0 := root0.String()
	wantM0 := "e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855" // SHA256("")
	if gotM0 != wantM0 {
		t.Errorf("Empty Merkle Root mismatch: %s", gotM0)
	}
}

func TestProtocol_RejectionRules(t *testing.T) {
	// PROTOCOL-011: Reject Overflow
	largePayload := make([]byte, 65536)
	eLarge := contracts.Event{Version: 1, Type: 1, Payload: largePayload}
	_, err := DigestEvent(eLarge)
	if err == nil {
		t.Error("Expected error for >64KB payload, got nil")
	}

	// PROTOCOL-012: Reject Duplicates in QC
	nodeA := contracts.NodeID{0xA}
	qcDup := contracts.QuorumCertificate{
		Version: 1,
		Signatures: []contracts.SignatureEntry{
			{ValidatorID: nodeA, Signature: []byte("SIG1")},
			{ValidatorID: nodeA, Signature: []byte("SIG2")},
		},
	}
	_, err = DigestQC(qcDup)
	if err == nil {
		t.Error("Expected error for duplicate validator in QC, got nil")
	}

	// PROTOCOL-013: Reject Unsorted Signatures in QC
	nodeB := contracts.NodeID{0xB}
	qcUnsorted := contracts.QuorumCertificate{
		Version: 1,
		Signatures: []contracts.SignatureEntry{
			{ValidatorID: nodeB, Signature: []byte("SIG_B")},
			{ValidatorID: nodeA, Signature: []byte("SIG_A")},
		},
	}
	_, err = DigestQC(qcUnsorted)
	if err == nil {
		t.Error("Expected error for unsorted signatures in QC, got nil")
	}
}

func TestProtocol_PayloadBoundaries(t *testing.T) {
	cases := []int{0, 1, 65535}
	for _, size := range cases {
		e := contracts.Event{Version: 1, Type: 1, Payload: make([]byte, size)}
		_, err := DigestEvent(e)
		if err != nil {
			t.Errorf("Failed for valid payload size %d: %v", size, err)
		}
	}
}

func TestMerkle_CarryUp_Invariant(t *testing.T) {
	// Verify the odd-node carry-up behavior
	e1 := contracts.Event{Version: 1, Type: 0, Payload: []byte("A")}
	e2 := contracts.Event{Version: 1, Type: 0, Payload: []byte("B")}
	e3 := contracts.Event{Version: 1, Type: 0, Payload: []byte("C")}

	root3, _ := CalculateMerkleRoot([]contracts.Event{e1, e2, e3})
	if root3 == (contracts.Hash{}) {
		t.Error("Merkle root should not be zero")
	}
}
