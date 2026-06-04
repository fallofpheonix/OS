package adapters

import (
	"context"
	"encoding/json"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	replayv1 "github.com/fallofpheonix/phoenix/foundation/contracts/replay/v1"
	"github.com/fallofpheonix/phoenix/assurance/validation/replay"
)

// ReplayAdapter implements replayv1.ReplayEngine and wraps the legacy replay.Engine.
type ReplayAdapter struct {
	Engine *replay.Engine
}

// NewReplayAdapter creates a new ReplayAdapter wrapping the legacy replay.Engine.
func NewReplayAdapter(engine *replay.Engine) *ReplayAdapter {
	return &ReplayAdapter{
		Engine: engine,
	}
}

// SnapshotAdapter implements replayv1.Snapshot.
type SnapshotAdapter struct {
	stateHash    string
	replayOffset uint64
	payload      []byte
}

func (s *SnapshotAdapter) StateHash() string    { return s.stateHash }
func (s *SnapshotAdapter) ReplayOffset() uint64 { return s.replayOffset }
func (s *SnapshotAdapter) Payload() []byte      { return s.payload }

// Replay converts EventEnvelopes and calls the wrapped legacy Replay method.
func (a *ReplayAdapter) Replay(ctx context.Context, envelopes []eventsv1.EventEnvelope) error {
	return a.Engine.Replay(ctx, envelopes)
}

// GetCurrentState returns the current snapshot of the legacy Replay Engine.
func (a *ReplayAdapter) GetCurrentState() (replayv1.Snapshot, error) {
	stateHash := a.Engine.CalculateStateHash()
	// Marshal state values as payload
	payload, _ := json.Marshal(a.Engine.State.Values)

	return &SnapshotAdapter{
		stateHash:    stateHash,
		replayOffset: a.Engine.State.LogicalTime,
		payload:      payload,
	}, nil
}

// SetLogicalTime sets the logical time offset of the legacy state.
func (a *ReplayAdapter) SetLogicalTime(time uint64) {
	a.Engine.State.LogicalTime = time
}
