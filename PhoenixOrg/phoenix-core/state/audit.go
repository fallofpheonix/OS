package state

import "time"

// StateRecord captures a single transition with audit metadata.
type StateRecord struct {
	ID          int
	Previous    SystemState
	Current     SystemState
	Timestamp   time.Time
	Reason      string
	EvidenceID  string
}
