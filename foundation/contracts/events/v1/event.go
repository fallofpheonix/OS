package v1

import "encoding/json"

// Event defines the behavior of a canonical system event.
type Event interface {
	GetEventID() string
	GetParentID() string
	GetAuthorityID() string
	GetIdentityID() string
	GetLogicalTime() uint64
	GetEvidence() []string
	GetSignature() string
	GetPayload() json.RawMessage
}

// Serializer defines the interface for event serialization and deserialization.
type Serializer interface {
	Marshal(e Event) ([]byte, error)
	Unmarshal(data []byte) (Event, error)
}
