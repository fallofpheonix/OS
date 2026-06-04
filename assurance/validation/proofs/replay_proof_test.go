package proofs

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	"github.com/fallofpheonix/phoenix/foundation/events"
	"github.com/fallofpheonix/phoenix/assurance/validation/replay"
)

type replayEventEnvelopeWrapper struct {
	ev event.Event
}

func (w *replayEventEnvelopeWrapper) GetEventID() string            { return w.ev.EventID }
func (w *replayEventEnvelopeWrapper) GetEventVersion() string       { return "1.0.0" }
func (w *replayEventEnvelopeWrapper) GetEventType() string          { return "PhoenixEvent" }
func (w *replayEventEnvelopeWrapper) GetTimestamp() time.Time       { return time.Now() }
func (w *replayEventEnvelopeWrapper) GetMonotonicTime() uint64      { return w.ev.LogicalTime }
func (w *replayEventEnvelopeWrapper) GetSourceRepo() string         { return w.ev.AuthorityID }
func (w *replayEventEnvelopeWrapper) GetSourceComponent() string    { return w.ev.IdentityID }
func (w *replayEventEnvelopeWrapper) GetParentEvent() string        { return w.ev.ParentID }
func (w *replayEventEnvelopeWrapper) GetCorrelationID() string      { return "" }
func (w *replayEventEnvelopeWrapper) GetCausalChain() []string      { return nil }
func (w *replayEventEnvelopeWrapper) GetReplaySequence() uint64     { return w.ev.LogicalTime }
func (w *replayEventEnvelopeWrapper) GetEvidenceHash() string       { if len(w.ev.Evidence) > 0 { return w.ev.Evidence[0] }; return "" }
func (w *replayEventEnvelopeWrapper) GetTrustScore() float64        { return 1.0 }
func (w *replayEventEnvelopeWrapper) GetSignature() string         { return w.ev.Signature }
func (w *replayEventEnvelopeWrapper) GetPayloadHash() string       { return "" }
func (w *replayEventEnvelopeWrapper) GetSchemaVersion() string      { return "" }
func (w *replayEventEnvelopeWrapper) GetCreatedAt() time.Time       { return time.Now() }
func (w *replayEventEnvelopeWrapper) GetUpdatedAt() time.Time       { return time.Now() }
func (w *replayEventEnvelopeWrapper) GetValidationHash() string    { return "" }
func (w *replayEventEnvelopeWrapper) GetPayload() []byte            { return []byte(w.ev.Payload) }

// STATUS: EXPERIMENTAL
// Proof 1: Replay (10 runs, 3 verifiers, 0 divergence)
func TestReplayProof(t *testing.T) {
	events := []event.Event{
		{LogicalTime: 1, EventID: "E1", Payload: json.RawMessage(`{"key":"val1"}`)},
		{LogicalTime: 2, EventID: "E2", Payload: json.RawMessage(`{"key":"val2"}`)},
	}

	envelopes := []eventsv1.EventEnvelope{
		&replayEventEnvelopeWrapper{ev: events[0]},
		&replayEventEnvelopeWrapper{ev: events[1]},
	}

	hashes := make(map[string]bool)
	ctx := context.Background()

	for run := 0; run < 10; run++ {
		// Simulate 3 different verifier instances per run
		for v := 0; v < 3; v++ {
			engine := replay.NewEngine()
			err := engine.Replay(ctx, envelopes)
			if err != nil {
				t.Fatalf("Run %d Verifier %d failed: %v", run, v, err)
			}
			snap, err := engine.GetCurrentState()
			if err != nil {
				t.Fatalf("Failed to get snapshot: %v", err)
			}
			hash := snap.StateHash()
			hashes[hash] = true
		}
	}

	// Determinism Proof: All runs must produce exactly one identical state hash.
	if len(hashes) != 1 {
		t.Errorf("DIVERGENCE DETECTED: expected 1 unique hash across all runs, got %d", len(hashes))
		for h := range hashes {
			t.Logf("Observed Hash: %s", h)
		}
	}
}
