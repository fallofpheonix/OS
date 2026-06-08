package contracts

// EventType defines the semantic category of an OS intent.
type EventType uint16

const (
	EventSpawn EventType = iota
	EventMove
	EventUpdateValidator
	EventVerify
)

// Event is the smallest semantic unit of intent.
// It is stripped of all authority and context metadata.
// INV-004: State mutations consume only semantic Events.
type Event struct {
	Version uint16    `json:"version"`
	Type    EventType `json:"type"`
	Payload []byte    `json:"payload"`
}
