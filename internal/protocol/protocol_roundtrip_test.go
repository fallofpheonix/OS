package protocol

import (
	"bytes"
	"encoding/binary"
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"testing"
)

func TestProtocol_RoundTrip(t *testing.T) {
	// PROTOCOL-019: Encode(Decode(B)) == B and Decode(Encode(O)) == O

	t.Run("Event_BytePerfect", func(t *testing.T) {
		e := contracts.Event{
			Version: 1,
			Type:    contracts.EventMove,
			Payload: []byte("BYTE_PERFECT_TEST"),
		}

		// 1. O -> B
		var buf1 bytes.Buffer
		if err := MarshalEvent(&buf1, e); err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		b1 := buf1.Bytes()

		// 2. B -> O'
		ePrime, err := UnmarshalEvent(bytes.NewReader(b1))
		if err != nil {
			t.Fatalf("Unmarshal failed: %v", err)
		}

		// 3. O' -> B'
		var buf2 bytes.Buffer
		if err := MarshalEvent(&buf2, ePrime); err != nil {
			t.Fatalf("Marshal failed: %v", err)
		}
		b2 := buf2.Bytes()

		if !bytes.Equal(b1, b2) {
			t.Errorf("Byte-perfect round-trip failed for Event!\nB1: %x\nB2: %x", b1, b2)
		}
	})

	t.Run("BlockHeader_BytePerfect", func(t *testing.T) {
		b := contracts.FinalizedBlock{
			Version:       1,
			Height:        100,
			Epoch:         5,
			Round:         10,
			Proposer:      contracts.NodeID{0xAA},
			PrevBlockHash: contracts.Hash{0xBB},
			MerkleRoot:    contracts.Hash{0xCC},
			StateRoot:     contracts.Hash{0xDD},
		}

		var buf1 bytes.Buffer
		MarshalBlockHeader(&buf1, b)
		b1 := buf1.Bytes()

		bPrime, _ := UnmarshalBlockHeader(bytes.NewReader(b1))

		var buf2 bytes.Buffer
		MarshalBlockHeader(&buf2, bPrime)
		b2 := buf2.Bytes()

		if !bytes.Equal(b1, b2) {
			t.Errorf("Byte-perfect round-trip failed for BlockHeader")
		}
	})

	t.Run("QC_BytePerfect", func(t *testing.T) {
		qc := contracts.QuorumCertificate{
			Version: 1,
			Epoch:   1,
			Signatures: []contracts.SignatureEntry{
				{ValidatorID: contracts.NodeID{0x01}, Signature: []byte("SIG1")},
				{ValidatorID: contracts.NodeID{0x02}, Signature: []byte("SIG2")},
			},
		}

		var buf1 bytes.Buffer
		MarshalQC(&buf1, qc)
		b1 := buf1.Bytes()

		qcPrime, _ := UnmarshalQC(bytes.NewReader(b1))

		var buf2 bytes.Buffer
		MarshalQC(&buf2, qcPrime)
		b2 := buf2.Bytes()

		if !bytes.Equal(b1, b2) {
			t.Errorf("Byte-perfect round-trip failed for QC")
		}
	})
}

func TestProtocol_MemoryProtection(t *testing.T) {
	t.Run("OversizedPayload", func(t *testing.T) {
		var buf bytes.Buffer
		binary.Write(&buf, binary.BigEndian, uint16(1))     // Ver
		binary.Write(&buf, binary.BigEndian, uint16(1))     // Type
		binary.Write(&buf, binary.BigEndian, uint16(65535)) // Claimed Len
		buf.Write(make([]byte, 10))

		_, err := UnmarshalEvent(bytes.NewReader(buf.Bytes()))
		if err == nil {
			t.Error("Expected error due to incomplete large payload, got nil")
		}
	})

	t.Run("ExcessiveQCSignatures", func(t *testing.T) {
		var buf bytes.Buffer
		buf.Write(make([]byte, 2+8+4+8+32+32+32))          // QC Header
		binary.Write(&buf, binary.BigEndian, uint16(5000)) // sigCount > MaxQCSignatures (1024)

		_, err := UnmarshalQC(bytes.NewReader(buf.Bytes()))
		if err == nil || err.Error() != "protocol: unmarshal QC sig count too high: 5000" {
			t.Errorf("Expected sig count error, got: %v", err)
		}
	})
}

func TestProtocol_BlockEnforcement(t *testing.T) {
	t.Run("MaxBlockSizeRejection", func(t *testing.T) {
		// Create a block that exceeds 1MB
		largeEvent := contracts.Event{
			Version: 1,
			Type:    contracts.EventMove,
			Payload: make([]byte, 65000),
		}

		var events []contracts.Event
		for i := 0; i < 20; i++ { // 20 * 65000 > 1MB
			events = append(events, largeEvent)
		}

		root, _ := CalculateMerkleRoot(events)
		b := contracts.FinalizedBlock{
			Version:    1,
			Events:     events,
			MerkleRoot: root,
		}

		err := ValidateBlock(b)
		if err == nil {
			t.Error("Expected ValidateBlock to reject oversized block, got nil")
		}
	})
}
