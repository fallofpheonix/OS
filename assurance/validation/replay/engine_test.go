package replay

import (
	"context"
	"testing"
	"time"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
)

type stubEnvelope struct {
	id      string
	seq     uint64
	payload []byte
}

func (e *stubEnvelope) GetEventID() string         { return e.id }
func (e *stubEnvelope) GetEventVersion() string    { return "1.0.0" }
func (e *stubEnvelope) GetEventType() string       { return "test" }
func (e *stubEnvelope) GetTimestamp() time.Time    { return time.Now() }
func (e *stubEnvelope) GetMonotonicTime() uint64   { return e.seq }
func (e *stubEnvelope) GetSourceRepo() string      { return "" }
func (e *stubEnvelope) GetSourceComponent() string { return "" }
func (e *stubEnvelope) GetParentEvent() string     { return "" }
func (e *stubEnvelope) GetCorrelationID() string   { return "" }
func (e *stubEnvelope) GetCausalChain() []string   { return nil }
func (e *stubEnvelope) GetReplaySequence() uint64  { return e.seq }
func (e *stubEnvelope) GetEvidenceHash() string    { return "" }
func (e *stubEnvelope) GetTrustScore() float64     { return 1.0 }
func (e *stubEnvelope) GetSignature() string       { return "" }
func (e *stubEnvelope) GetPayloadHash() string     { return "" }
func (e *stubEnvelope) GetSchemaVersion() string   { return "" }
func (e *stubEnvelope) GetCreatedAt() time.Time    { return time.Now() }
func (e *stubEnvelope) GetUpdatedAt() time.Time    { return time.Now() }
func (e *stubEnvelope) GetValidationHash() string  { return "" }
func (e *stubEnvelope) GetPayload() []byte         { return e.payload }

func TestDeterministicReconstruction(t *testing.T) {
	engine := NewEngine()

	// Create a stream of events
	events := []eventsv1.EventEnvelope{
		&stubEnvelope{
			id:      "E1",
			seq:     1,
			payload: []byte(`{"key1":"val1"}`),
		},
		&stubEnvelope{
			id:      "E2",
			seq:     2,
			payload: []byte(`{"key2":"val2"}`),
		},
	}

	// 1. Reconstruct state
	ctx := context.Background()
	err := engine.Replay(ctx, events)
	if err != nil {
		t.Fatalf("Reconstruction failed: %v", err)
	}

	// 2. Verify state values
	if val, ok := engine.State.Values["key2"]; !ok || val != "val2" {
		t.Errorf("Expected key2 to be val2, got %q", val)
	}

	// 3. Verify hashing is deterministic (run again)
	engine2 := NewEngine()
	engine2.Replay(ctx, events)
	snap2, _ := engine2.GetCurrentState()
	finalHash2Val := snap2.StateHash()

	snap1, _ := engine.GetCurrentState()
	finalHash1Val := snap1.StateHash()

	if finalHash1Val != finalHash2Val {
		t.Errorf("Non-deterministic hashing: %s vs %s", finalHash1Val, finalHash2Val)
	}

	// 4. Test Deterministic Drift
	engine3 := NewEngine()
	events[1].(*stubEnvelope).payload = []byte(`{"key2":"val3"}`) // Modify event payload
	engine3.Replay(ctx, events)
	snap3, _ := engine3.GetCurrentState()
	finalHash3Val := snap3.StateHash()

	if finalHash1Val == finalHash3Val {
		t.Error("Failed to detect modification in event stream")
	}
}

func TestFATAL_DETERMINISTIC_DRIFT(t *testing.T) {
	verifier := NewAuthorityVerifier()
	err := verifier.VerifyAuthority("hash-a", "hash-b")
	if err == nil {
		t.Error("Authority verifier failed to detect drift")
	}
}
