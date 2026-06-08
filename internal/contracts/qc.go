package contracts

// QuorumCertificate (QC) is the proof of consensus finality.
type QuorumCertificate struct {
	Version          uint16           `json:"version"`
	Epoch            uint64           `json:"epoch"`
	Round            uint32           `json:"round"`
	Height           uint64           `json:"height"`
	BlockID          BlockID          `json:"block_id"`
	StateRoot        Hash             `json:"state_root"`
	ValidatorSetHash Hash             `json:"validator_set_hash"`
	Signatures       []SignatureEntry `json:"signatures"`
}

// SignatureEntry represents an individual validator's signature.
// Signatures must be sorted by ValidatorID raw bytes for determinism.
type SignatureEntry struct {
	ValidatorID NodeID `json:"validator_id"`
	Signature   []byte `json:"signature"`
}
