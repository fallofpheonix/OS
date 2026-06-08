package consensus

import (
	"crypto/ed25519"

	"github.com/fallofpheonix/phoenix/internal/contracts"
)

// Re-export core types from contracts to maintain compatibility within the consensus package.
type EventType = contracts.EventType
type Event = contracts.Event

const (
	EventSpawn           = contracts.EventSpawn
	EventMove            = contracts.EventMove
	EventVerify          = contracts.EventVerify
	EventUpdateValidator = contracts.EventUpdateValidator
)

// ParseHash re-exports the parser.
func ParseHash(s string) (Hash, error) {
	return contracts.ParseHash(s)
}

// SignEvent is deprecated. Use SignEnvelope.
func SignEvent(e *Event, priv ed25519.PrivateKey) ([]byte, error) {
	return nil, nil
}
