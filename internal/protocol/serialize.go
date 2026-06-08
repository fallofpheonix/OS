package protocol

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/fallofpheonix/phoenix/internal/contracts"
	"io"
)

const (
	// PROTOCOL-016: Strict Resource Limits
	MaxBlockSize      = 1024 * 1024 // 1 MB
	MaxPayloadSize    = 65535       // 64 KB
	MaxEventsPerBlock = 1024
	MaxQCSignatures   = 1024 // SPEC-CONSISTENCY: Reverted to 1024
	MaxSignatureSize  = 512
)

// MarshalEvent writes the canonical binary representation of an event.
func MarshalEvent(w io.Writer, e contracts.Event) error {
	if len(e.Payload) > MaxPayloadSize {
		return fmt.Errorf("protocol: event payload exceeds limit (%d > %d)", len(e.Payload), MaxPayloadSize)
	}

	if err := binary.Write(w, binary.BigEndian, e.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint16(e.Type)); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint16(len(e.Payload))); err != nil {
		return err
	}
	if _, err := w.Write(e.Payload); err != nil {
		return err
	}
	return nil
}

// UnmarshalEvent reads an event from the canonical binary representation.
func UnmarshalEvent(r io.Reader) (contracts.Event, error) {
	var e contracts.Event
	if err := binary.Read(r, binary.BigEndian, &e.Version); err != nil {
		return e, err
	}
	var eType uint16
	if err := binary.Read(r, binary.BigEndian, &eType); err != nil {
		return e, err
	}
	e.Type = contracts.EventType(eType)

	var payloadLen uint16
	if err := binary.Read(r, binary.BigEndian, &payloadLen); err != nil {
		return e, err
	}

	// Defensive check before allocation
	if int(payloadLen) > MaxPayloadSize {
		return e, fmt.Errorf("protocol: unmarshal event payload too large: %d", payloadLen)
	}

	e.Payload = make([]byte, payloadLen)
	if _, err := io.ReadFull(r, e.Payload); err != nil {
		return e, err
	}
	return e, nil
}

// MarshalBlockHeader writes the canonical binary representation of a block header.
func MarshalBlockHeader(w io.Writer, b contracts.FinalizedBlock) error {
	if err := binary.Write(w, binary.BigEndian, b.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, b.Height); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, b.Epoch); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, b.Round); err != nil {
		return err
	}
	if _, err := w.Write(b.Proposer[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.PrevBlockHash[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.MerkleRoot[:]); err != nil {
		return err
	}
	if _, err := w.Write(b.StateRoot[:]); err != nil {
		return err
	}
	return nil
}

// UnmarshalBlockHeader reads a block header from the canonical binary representation.
func UnmarshalBlockHeader(r io.Reader) (contracts.FinalizedBlock, error) {
	var b contracts.FinalizedBlock
	if err := binary.Read(r, binary.BigEndian, &b.Version); err != nil {
		return b, err
	}
	if err := binary.Read(r, binary.BigEndian, &b.Height); err != nil {
		return b, err
	}
	if err := binary.Read(r, binary.BigEndian, &b.Epoch); err != nil {
		return b, err
	}
	if err := binary.Read(r, binary.BigEndian, &b.Round); err != nil {
		return b, err
	}
	if _, err := io.ReadFull(r, b.Proposer[:]); err != nil {
		return b, err
	}
	if _, err := io.ReadFull(r, b.PrevBlockHash[:]); err != nil {
		return b, err
	}
	if _, err := io.ReadFull(r, b.MerkleRoot[:]); err != nil {
		return b, err
	}
	if _, err := io.ReadFull(r, b.StateRoot[:]); err != nil {
		return b, err
	}
	return b, nil
}

// MarshalQC writes the canonical binary representation of a Quorum Certificate.
func MarshalQC(w io.Writer, qc contracts.QuorumCertificate) error {
	if err := binary.Write(w, binary.BigEndian, qc.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, qc.Epoch); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, qc.Round); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, qc.Height); err != nil {
		return err
	}
	if _, err := w.Write(qc.BlockID[:]); err != nil {
		return err
	}
	if _, err := w.Write(qc.StateRoot[:]); err != nil {
		return err
	}
	if _, err := w.Write(qc.ValidatorSetHash[:]); err != nil {
		return err
	}

	if err := binary.Write(w, binary.BigEndian, uint16(len(qc.Signatures))); err != nil {
		return err
	}
	for i, entry := range qc.Signatures {
		if _, err := w.Write(entry.ValidatorID[:]); err != nil {
			return err
		}

		// PROTOCOL-018: Signatures must be length-prefixed
		if err := binary.Write(w, binary.BigEndian, uint16(len(entry.Signature))); err != nil {
			return err
		}
		if _, err := w.Write(entry.Signature); err != nil {
			return err
		}

		// PROTOCOL-013: Signatures must be sorted.
		if i > 0 {
			cmp := bytes.Compare(qc.Signatures[i-1].ValidatorID[:], entry.ValidatorID[:])
			if cmp > 0 {
				return fmt.Errorf("protocol: out-of-order signature at index %d", i)
			}
			if cmp == 0 {
				return fmt.Errorf("protocol: duplicate signature for validator %x", entry.ValidatorID)
			}
		}
	}
	return nil
}

// UnmarshalQC reads a QC from the canonical binary representation.
func UnmarshalQC(r io.Reader) (contracts.QuorumCertificate, error) {
	var qc contracts.QuorumCertificate
	if err := binary.Read(r, binary.BigEndian, &qc.Version); err != nil {
		return qc, err
	}
	if err := binary.Read(r, binary.BigEndian, &qc.Epoch); err != nil {
		return qc, err
	}
	if err := binary.Read(r, binary.BigEndian, &qc.Round); err != nil {
		return qc, err
	}
	if err := binary.Read(r, binary.BigEndian, &qc.Height); err != nil {
		return qc, err
	}
	if _, err := io.ReadFull(r, qc.BlockID[:]); err != nil {
		return qc, err
	}
	if _, err := io.ReadFull(r, qc.StateRoot[:]); err != nil {
		return qc, err
	}
	if _, err := io.ReadFull(r, qc.ValidatorSetHash[:]); err != nil {
		return qc, err
	}

	var sigCount uint16
	if err := binary.Read(r, binary.BigEndian, &sigCount); err != nil {
		return qc, err
	}

	// Defensive check before allocation
	if int(sigCount) > MaxQCSignatures {
		return qc, fmt.Errorf("protocol: unmarshal QC sig count too high: %d", sigCount)
	}

	qc.Signatures = make([]contracts.SignatureEntry, sigCount)
	for i := 0; i < int(sigCount); i++ {
		if _, err := io.ReadFull(r, qc.Signatures[i].ValidatorID[:]); err != nil {
			return qc, err
		}

		var sigLen uint16
		if err := binary.Read(r, binary.BigEndian, &sigLen); err != nil {
			return qc, err
		}

		if int(sigLen) > MaxSignatureSize {
			return qc, fmt.Errorf("protocol: unmarshal QC sig too large at index %d: %d", i, sigLen)
		}

		qc.Signatures[i].Signature = make([]byte, sigLen)
		if _, err := io.ReadFull(r, qc.Signatures[i].Signature); err != nil {
			return qc, err
		}
	}
	return qc, nil
}

// MarshalEnvelope writes the canonical binary representation of a SignedEnvelope.
func MarshalEnvelope(w io.Writer, e contracts.SignedEnvelope) error {
	if err := binary.Write(w, binary.BigEndian, e.Version); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, e.Type); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, e.Epoch); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, e.Sequence); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, e.Timestamp); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint32(len(e.Payload))); err != nil {
		return err
	}
	if _, err := w.Write(e.Payload); err != nil {
		return err
	}
	if _, err := w.Write(e.Validator[:]); err != nil {
		return err
	}
	if err := binary.Write(w, binary.BigEndian, uint16(len(e.Signature))); err != nil {
		return err
	}
	if _, err := w.Write(e.Signature); err != nil {
		return err
	}
	return nil
}

// UnmarshalEnvelope reads a SignedEnvelope from the canonical binary representation.
func UnmarshalEnvelope(r io.Reader) (contracts.SignedEnvelope, error) {
	var e contracts.SignedEnvelope
	if err := binary.Read(r, binary.BigEndian, &e.Version); err != nil {
		return e, err
	}
	if err := binary.Read(r, binary.BigEndian, &e.Type); err != nil {
		return e, err
	}
	if err := binary.Read(r, binary.BigEndian, &e.Epoch); err != nil {
		return e, err
	}
	if err := binary.Read(r, binary.BigEndian, &e.Sequence); err != nil {
		return e, err
	}
	if err := binary.Read(r, binary.BigEndian, &e.Timestamp); err != nil {
		return e, err
	}

	var payloadLen uint32
	if err := binary.Read(r, binary.BigEndian, &payloadLen); err != nil {
		return e, err
	}
	if int(payloadLen) > MaxBlockSize { // Using MaxBlockSize as bound for envelope payloads
		return e, fmt.Errorf("protocol: unmarshal envelope payload too large: %d", payloadLen)
	}

	e.Payload = make([]byte, payloadLen)
	if _, err := io.ReadFull(r, e.Payload); err != nil {
		return e, err
	}

	if _, err := io.ReadFull(r, e.Validator[:]); err != nil {
		return e, err
	}

	var sigLen uint16
	if err := binary.Read(r, binary.BigEndian, &sigLen); err != nil {
		return e, err
	}
	if int(sigLen) > MaxSignatureSize {
		return e, fmt.Errorf("protocol: unmarshal envelope signature too large: %d", sigLen)
	}

	e.Signature = make([]byte, sigLen)
	if _, err := io.ReadFull(r, e.Signature); err != nil {
		return e, err
	}
	return e, nil
}

// ValidateBlock performs the authoritative validation pipeline for a block.
func ValidateBlock(b contracts.FinalizedBlock) error {
	// 1. Static Size Limits
	var buf bytes.Buffer
	if err := MarshalBlockHeader(&buf, b); err != nil {
		return fmt.Errorf("protocol: block header marshal failed: %w", err)
	}
	totalSize := buf.Len()

	for _, e := range b.Events {
		var eBuf bytes.Buffer
		if err := MarshalEvent(&eBuf, e); err != nil {
			return fmt.Errorf("protocol: event marshal failed: %w", err)
		}
		totalSize += eBuf.Len()
	}

	var qcBuf bytes.Buffer
	if err := MarshalQC(&qcBuf, b.QC); err != nil {
		return fmt.Errorf("protocol: QC marshal failed: %w", err)
	}
	totalSize += qcBuf.Len()

	if totalSize > MaxBlockSize {
		return fmt.Errorf("protocol: total block size %d exceeds limit %d (ERR_LIMIT)", totalSize, MaxBlockSize)
	}

	// 2. Event Count Limit
	if len(b.Events) > MaxEventsPerBlock {
		return fmt.Errorf("protocol: block has %d events, exceeds limit %d (ERR_LIMIT)", len(b.Events), MaxEventsPerBlock)
	}

	// 3. Merkle Validation
	root, err := CalculateMerkleRoot(b.Events)
	if err != nil {
		return fmt.Errorf("protocol: merkle calculation failed: %w", err)
	}
	if root != b.MerkleRoot {
		return fmt.Errorf("protocol: merkle root mismatch. expected %s, got %s (ERR_MERKLE)", b.MerkleRoot, root)
	}

	return nil
}
