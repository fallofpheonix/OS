package v1

import "time"

// EventEnvelope defines the canonical envelope interface.
type EventEnvelope interface {
	GetEventID() string
	GetEventVersion() string
	GetEventType() string
	GetTimestamp() time.Time
	GetMonotonicTime() uint64
	GetSourceRepo() string
	GetSourceComponent() string
	GetParentEvent() string
	GetCorrelationID() string
	GetCausalChain() []string
	GetReplaySequence() uint64
	GetEvidenceHash() string
	GetTrustScore() float64
	GetSignature() string
	GetPayloadHash() string
	GetSchemaVersion() string
	GetCreatedAt() time.Time
	GetUpdatedAt() time.Time
	GetValidationHash() string
}
