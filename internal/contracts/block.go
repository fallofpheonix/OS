package contracts

// FinalizedBlock is the authoritative unit of the Ledger.
type FinalizedBlock struct {
	Version       uint16            `json:"version"`
	Height        uint64            `json:"height"`
	Epoch         uint64            `json:"epoch"`
	Round         uint32            `json:"round"`
	Proposer      NodeID            `json:"proposer"`
	PrevBlockHash Hash              `json:"prev_block_hash"`
	MerkleRoot    Hash              `json:"merkle_root"` // Root of Events Merkle Tree
	Events        []Event           `json:"events"`
	StateRoot     Hash              `json:"state_root"`
	QC            QuorumCertificate `json:"qc"`
}
