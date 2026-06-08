package protocol

import (
	"crypto/sha256"
	"fmt"
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

const (
	PrefixEvent       byte = 0x01
	PrefixBlockHeader byte = 0x02
	PrefixQC          byte = 0x03
	PrefixEnvelope    byte = 0x04
)

// DigestEvent produces the canonical binary digest for an event.
func DigestEvent(e contracts.Event) (contracts.Hash, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{PrefixEvent}); err != nil {
		return contracts.Hash{}, err
	}
	if err := MarshalEvent(h, e); err != nil {
		return contracts.Hash{}, err
	}
	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res, nil
}

// DigestBlockHeader produces the canonical digest for a block header.
func DigestBlockHeader(b contracts.FinalizedBlock) (contracts.Hash, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{PrefixBlockHeader}); err != nil {
		return contracts.Hash{}, err
	}
	if err := MarshalBlockHeader(h, b); err != nil {
		return contracts.Hash{}, err
	}
	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res, nil
}

// DigestQC produces the canonical digest for a Quorum Certificate.
func DigestQC(qc contracts.QuorumCertificate) (contracts.Hash, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{PrefixQC}); err != nil {
		return contracts.Hash{}, err
	}

	// MarshalQC enforces the sorted-signatures invariant and duplicate-rejection.
	if err := MarshalQC(h, qc); err != nil {
		return contracts.Hash{}, fmt.Errorf("protocol: qc digest failed: %w", err)
	}

	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res, nil
}

// DigestEnvelope produces the canonical binary digest for an envelope.
// NOTE: The Signature field is explicitly EXCLUDED to allow for signing and verification.
func DigestEnvelope(e contracts.SignedEnvelope) (contracts.Hash, error) {
	h := sha256.New()
	if _, err := h.Write([]byte{PrefixEnvelope}); err != nil {
		return contracts.Hash{}, err
	}

	// Create a shallow copy and strip signature for hashing
	e.Signature = nil
	if err := MarshalEnvelope(h, e); err != nil {
		return contracts.Hash{}, err
	}
	var res contracts.Hash
	copy(res[:], h.Sum(nil))
	return res, nil
}
