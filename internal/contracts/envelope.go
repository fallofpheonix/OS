package contracts

// SignedEnvelope is the network container for pre-consensus messages.
type SignedEnvelope struct {
	Version   uint16 `json:"version"`
	Type      uint8  `json:"type"`
	Epoch     uint64 `json:"epoch"`
	Sequence  uint64 `json:"sequence"`
	Timestamp int64  `json:"timestamp"`
	Payload   []byte `json:"payload"`
	Validator NodeID `json:"validator"`
	Signature []byte `json:"signature"`
}
