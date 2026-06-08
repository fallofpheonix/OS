package contracts

// Vote represents an individual validator's agreement.
type Vote struct {
	Version   uint16  `json:"version"`
	Epoch     uint64  `json:"epoch"`
	Round     uint32  `json:"round"`
	Height    uint64  `json:"height"`
	BlockID   BlockID `json:"block_id"`
	StateRoot Hash    `json:"state_root"`
}
