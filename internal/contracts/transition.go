package contracts

// AppliedEvent is the enriched semantic payload consumed by the State layer.
type AppliedEvent struct {
	Height uint64
	Epoch  uint64
	Event  Event
}
