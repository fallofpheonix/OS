package v1

import (
	"context"
	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
)

// Snapshot represents a serialized/hashed state snapshot at a specific replay sequence.
type Snapshot interface {
	StateHash() string
	ReplayOffset() uint64
	Payload() []byte
}

// Reconstructor defines the behavior for rebuilding system state from events.
type Reconstructor interface {
	Reconstruct(ctx context.Context, events []eventsv1.EventEnvelope) (Snapshot, error)
}

// ReplayEngine defines the core interface for running deterministic event replay.
type ReplayEngine interface {
	Replay(ctx context.Context, events []eventsv1.EventEnvelope) error
	GetCurrentState() (Snapshot, error)
	SetLogicalTime(time uint64)
}
