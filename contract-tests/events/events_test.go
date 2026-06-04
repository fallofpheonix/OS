package events

import (
	"encoding/json"
	"testing"
	"time"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
)

// stubEvent implements eventsv1.Event
type stubEvent struct {
	ID          string          `json:"event_id"`
	ParentID    string          `json:"parent_id"`
	AuthorityID string          `json:"authority_id"`
	IdentityID  string          `json:"identity_id"`
	LogicalTime uint64          `json:"logical_time"`
	Evidence    []string        `json:"evidence"`
	Signature   string          `json:"signature"`
	Payload     json.RawMessage `json:"payload"`
}

func (e *stubEvent) GetEventID() string            { return e.ID }
func (e *stubEvent) GetParentID() string           { return e.ParentID }
func (e *stubEvent) GetAuthorityID() string        { return e.AuthorityID }
func (e *stubEvent) GetIdentityID() string         { return e.IdentityID }
func (e *stubEvent) GetLogicalTime() uint64        { return e.LogicalTime }
func (e *stubEvent) GetEvidence() []string         { return e.Evidence }
func (e *stubEvent) GetSignature() string          { return e.Signature }
func (e *stubEvent) GetPayload() json.RawMessage   { return e.Payload }

// TestFITEvent001 verifies the single canonical Event contract definition.
func TestFITEvent001(t *testing.T) {
	var ev interface{} = &stubEvent{
		ID: "evt_123",
	}

	if _, ok := ev.(eventsv1.Event); !ok {
		t.Fatal("stubEvent does not satisfy eventsv1.Event contract interface")
	}
}

// stubEnvelope implements eventsv1.EventEnvelope
type stubEnvelope struct {
	EventIDVal         string          `json:"event_id"`
	EventVersionVal    string          `json:"event_version"`
	EventTypeVal       string          `json:"event_type"`
	TimestampVal       time.Time       `json:"timestamp"`
	MonotonicTimeVal   uint64          `json:"monotonic_time"`
	SourceRepoVal      string          `json:"source_repo"`
	SourceComponentVal string          `json:"source_component"`
	ParentEventVal     string          `json:"parent_event"`
	CorrelationIDVal   string          `json:"correlation_id"`
	CausalChainVal     []string        `json:"causal_chain"`
	ReplaySequenceVal  uint64          `json:"replay_sequence"`
	EvidenceHashVal    string          `json:"evidence_hash"`
	TrustScoreVal      float64         `json:"trust_score"`
	SignatureVal       string          `json:"signature"`
	PayloadHashVal     string          `json:"payload_hash"`
	SchemaVersionVal   string          `json:"schema_version"`
	CreatedAtVal       time.Time       `json:"created_at"`
	UpdatedAtVal       time.Time       `json:"updated_at"`
	ValidationHashVal  string          `json:"validation_hash"`
}

func (e *stubEnvelope) GetEventID() string          { return e.EventIDVal }
func (e *stubEnvelope) GetEventVersion() string     { return e.EventVersionVal }
func (e *stubEnvelope) GetEventType() string        { return e.EventTypeVal }
func (e *stubEnvelope) GetTimestamp() time.Time     { return e.TimestampVal }
func (e *stubEnvelope) GetMonotonicTime() uint64    { return e.MonotonicTimeVal }
func (e *stubEnvelope) GetSourceRepo() string       { return e.SourceRepoVal }
func (e *stubEnvelope) GetSourceComponent() string  { return e.SourceComponentVal }
func (e *stubEnvelope) GetParentEvent() string      { return e.ParentEventVal }
func (e *stubEnvelope) GetCorrelationID() string    { return e.CorrelationIDVal }
func (e *stubEnvelope) GetCausalChain() []string    { return e.CausalChainVal }
func (e *stubEnvelope) GetReplaySequence() uint64   { return e.ReplaySequenceVal }
func (e *stubEnvelope) GetEvidenceHash() string     { return e.EvidenceHashVal }
func (e *stubEnvelope) GetTrustScore() float64      { return e.TrustScoreVal }
func (e *stubEnvelope) GetSignature() string        { return e.SignatureVal }
func (e *stubEnvelope) GetPayloadHash() string      { return e.PayloadHashVal }
func (e *stubEnvelope) GetSchemaVersion() string    { return e.SchemaVersionVal }
func (e *stubEnvelope) GetCreatedAt() time.Time     { return e.CreatedAtVal }
func (e *stubEnvelope) GetUpdatedAt() time.Time     { return e.UpdatedAtVal }
func (e *stubEnvelope) GetValidationHash() string   { return e.ValidationHashVal }

// TestFITEvent002 verifies the Envelope compatibility contract.
func TestFITEvent002(t *testing.T) {
	var env interface{} = &stubEnvelope{
		EventIDVal: "env_456",
	}

	if _, ok := env.(eventsv1.EventEnvelope); !ok {
		t.Fatal("stubEnvelope does not satisfy eventsv1.EventEnvelope contract interface")
	}
}

// deterministicSerializer implements eventsv1.Serializer
type deterministicSerializer struct{}

func (s *deterministicSerializer) Marshal(e eventsv1.Event) ([]byte, error) {
	return json.Marshal(struct {
		ID          string          `json:"event_id"`
		ParentID    string          `json:"parent_id"`
		AuthorityID string          `json:"authority_id"`
		IdentityID  string          `json:"identity_id"`
		LogicalTime uint64          `json:"logical_time"`
		Evidence    []string        `json:"evidence"`
		Signature   string          `json:"signature"`
		Payload     json.RawMessage `json:"payload"`
	}{
		ID:          e.GetEventID(),
		ParentID:    e.GetParentID(),
		AuthorityID: e.GetAuthorityID(),
		IdentityID:  e.GetIdentityID(),
		LogicalTime: e.GetLogicalTime(),
		Evidence:    e.GetEvidence(),
		Signature:   e.GetSignature(),
		Payload:     e.GetPayload(),
	})
}

func (s *deterministicSerializer) Unmarshal(data []byte) (eventsv1.Event, error) {
	var stub stubEvent
	if err := json.Unmarshal(data, &stub); err != nil {
		return nil, err
	}
	return &stub, nil
}

// TestFITEvent003 verifies serialization determinism.
func TestFITEvent003(t *testing.T) {
	serializer := &deterministicSerializer{}
	event1 := &stubEvent{
		ID:          "evt_789",
		ParentID:    "evt_000",
		AuthorityID: "auth_sys",
		IdentityID:  "id_user",
		LogicalTime: 42,
		Evidence:    []string{"sha256_hash1", "sha256_hash2"},
		Signature:   "sig_xxx",
		Payload:     json.RawMessage(`{"action":"sandbox"}`),
	}

	bytes1, err := serializer.Marshal(event1)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	bytes2, err := serializer.Marshal(event1)
	if err != nil {
		t.Fatalf("Marshal failed: %v", err)
	}

	if string(bytes1) != string(bytes2) {
		t.Fatal("Serialization is non-deterministic: outputs differ")
	}

	unmarshaled, err := serializer.Unmarshal(bytes1)
	if err != nil {
		t.Fatalf("Unmarshal failed: %v", err)
	}

	if unmarshaled.GetEventID() != event1.GetEventID() {
		t.Fatalf("Expected ID %s, got %s", event1.GetEventID(), unmarshaled.GetEventID())
	}
}
