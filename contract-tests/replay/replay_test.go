package replay

import (
	"context"
	"testing"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	replayv1 "github.com/fallofpheonix/phoenix/foundation/contracts/replay/v1"
)

// stubSnapshot implements replayv1.Snapshot
type stubSnapshot struct {
	stateHash    string
	replayOffset uint64
	payload      []byte
}

func (s *stubSnapshot) StateHash() string      { return s.stateHash }
func (s *stubSnapshot) ReplayOffset() uint64   { return s.replayOffset }
func (s *stubSnapshot) Payload() []byte        { return s.payload }

// stubReplayer implements replayv1.ReplayEngine
type stubReplayer struct {
	snapshot replayv1.Snapshot
}

func (r *stubReplayer) Replay(ctx context.Context, events []eventsv1.EventEnvelope) error {
	return nil
}

func (r *stubReplayer) GetCurrentState() (replayv1.Snapshot, error) {
	return r.snapshot, nil
}

func (r *stubReplayer) SetLogicalTime(time uint64) {}

// TestFITReplay001 verifies replayv1.ReplayEngine and Snapshot interfaces.
func TestFITReplay001(t *testing.T) {
	var snap interface{} = &stubSnapshot{
		stateHash: "sha256_snapshot_hash",
	}

	if _, ok := snap.(replayv1.Snapshot); !ok {
		t.Fatal("stubSnapshot does not satisfy replayv1.Snapshot contract interface")
	}

	var engine interface{} = &stubReplayer{
		snapshot: snap.(replayv1.Snapshot),
	}

	if _, ok := engine.(replayv1.ReplayEngine); !ok {
		t.Fatal("stubReplayer does not satisfy replayv1.ReplayEngine contract interface")
	}
}
