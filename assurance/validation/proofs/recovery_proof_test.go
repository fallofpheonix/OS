package proofs

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/assurance/validation/replay"
	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	"github.com/fallofpheonix/phoenix/foundation/events"
	"github.com/fallofpheonix/phoenix/foundation/runtime/constitution"
	"github.com/fallofpheonix/phoenix/foundation/runtime/recovery"
)

type eventEnvelopeWrapper struct {
	ev event.Event
}

func (w *eventEnvelopeWrapper) GetEventID() string         { return w.ev.EventID }
func (w *eventEnvelopeWrapper) GetEventVersion() string    { return "1.0.0" }
func (w *eventEnvelopeWrapper) GetEventType() string       { return "PhoenixEvent" }
func (w *eventEnvelopeWrapper) GetTimestamp() time.Time    { return time.Now() }
func (w *eventEnvelopeWrapper) GetMonotonicTime() uint64   { return w.ev.LogicalTime }
func (w *eventEnvelopeWrapper) GetSourceRepo() string      { return w.ev.AuthorityID }
func (w *eventEnvelopeWrapper) GetSourceComponent() string { return w.ev.IdentityID }
func (w *eventEnvelopeWrapper) GetParentEvent() string     { return w.ev.ParentID }
func (w *eventEnvelopeWrapper) GetCorrelationID() string   { return "" }
func (w *eventEnvelopeWrapper) GetCausalChain() []string   { return nil }
func (w *eventEnvelopeWrapper) GetReplaySequence() uint64  { return w.ev.LogicalTime }
func (w *eventEnvelopeWrapper) GetEvidenceHash() string {
	if len(w.ev.Evidence) > 0 {
		return w.ev.Evidence[0]
	}
	return ""
}
func (w *eventEnvelopeWrapper) GetTrustScore() float64    { return 1.0 }
func (w *eventEnvelopeWrapper) GetSignature() string      { return w.ev.Signature }
func (w *eventEnvelopeWrapper) GetPayloadHash() string    { return "" }
func (w *eventEnvelopeWrapper) GetSchemaVersion() string  { return "" }
func (w *eventEnvelopeWrapper) GetCreatedAt() time.Time   { return time.Now() }
func (w *eventEnvelopeWrapper) GetUpdatedAt() time.Time   { return time.Now() }
func (w *eventEnvelopeWrapper) GetValidationHash() string { return "" }
func (w *eventEnvelopeWrapper) GetPayload() []byte        { return []byte(w.ev.Payload) }

// STATUS: EXPERIMENTAL
// Proof 2: Recovery (Destroy, Recover, Verify)
func TestRecoveryProof(t *testing.T) {
	events := []event.Event{
		{LogicalTime: 1, EventID: "E1", Payload: json.RawMessage(`{"status":"online"}`)},
	}

	// 1. Establish ground truth state hash
	engine := replay.NewEngine()
	ctx := context.Background()
	err := engine.Replay(ctx, []eventsv1.EventEnvelope{&eventEnvelopeWrapper{ev: events[0]}})
	if err != nil {
		t.Fatalf("Failed ground truth replay: %v", err)
	}
	snap, err := engine.GetCurrentState()
	if err != nil {
		t.Fatalf("Failed ground truth state: %v", err)
	}
	expectedHash := snap.StateHash()

	// Define the checkpoint representing the desired state
	checkpoint := event.Checkpoint{
		StateHash:    expectedHash,
		ReplayOffset: 0,
	}

	// 2. Destroy: Create a new node substrate from scratch (simulating resurrection)
	recoveryEngine := recovery.NewEngine(constitution.NewEngine(), replay.NewEngine())

	// 3. Recover: Restore state from Ledger + Checkpoint
	err = recoveryEngine.Recover(checkpoint, events, []event.ArtifactManifest{})
	if err != nil {
		t.Fatalf("Recovery failed to reconstruct authoritative state: %v", err)
	}

	// 4. Verify: The recovered state hash must perfectly match the ground truth
	snap, err = recoveryEngine.Replay.GetCurrentState()
	if err != nil {
		t.Fatalf("Failed to get current state: %v", err)
	}
	actualHash := snap.StateHash()
	if actualHash != expectedHash {
		t.Errorf("RECOVERY_PROOF_FAILED: hash mismatch (expected %s, got %s)", expectedHash, actualHash)
	}
}
