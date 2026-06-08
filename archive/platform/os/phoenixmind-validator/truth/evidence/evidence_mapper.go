/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
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
