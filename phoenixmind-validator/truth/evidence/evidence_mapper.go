package evidence

import "time"

// TruthState represents the assessed state of an entity.
type TruthState string

const (
	UNKNOWN   TruthState = "UNKNOWN"
	OBSERVED  TruthState = "OBSERVED"
	VALIDATED TruthState = "VALIDATED"
	WARNING   TruthState = "WARNING"
	ESCALATED TruthState = "ESCALATED"
	BLOCKED   TruthState = "BLOCKED"
	REJECTED  TruthState = "REJECTED"
)

// Evidence is a single piece of data about an entity from a specific source.
type Evidence struct {
	ID        string
	EntityID  string
	Source    string
	Timestamp time.Time
	Payload   map[string]interface{}
	State     TruthState
}

// Mapper translates raw data into a structured Evidence object.
type Mapper func(rawData interface{}) (Evidence, error)
