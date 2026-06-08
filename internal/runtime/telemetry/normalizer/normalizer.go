// Package normalizer implements ingestion and formatting logic for raw telemetry data.
// Domain Logic: Standardizes diverse telemetry formats into canonical TelemetryEvent objects.
// Responsibility: Bridges the gap between disparate data sources (probes, logs) and the central event bus.
package normalizer

import (
	"encoding/json"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

// RawEvent represents un-normalized telemetry data from an external probe.
type RawEvent struct {
	Timestamp int64           `json:"ts"`
	Source    string          `json:"src"`
	PID       int             `json:"pid"`
	Type      string          `json:"type"`
	Severity  float64         `json:"sev"`
	Data      json.RawMessage `json:"data"`
}

// Normalize transforms a RawEvent into a canonical TelemetryEvent.
// LABEL: [PURE] [DETERMINISTIC] [STABLE]
func Normalize(re RawEvent, seqID uint64) (*bus.TelemetryEvent, error) {
	return &bus.TelemetryEvent{
		SeqID:        int64(seqID),
		MonotonicNs:  re.Timestamp,
		WallTimeUnix: re.Timestamp,
		Source:       re.Source,
		PID:          re.PID,
		EventType:    re.Type,
		Severity:     phxmath.FixedPoint{V: int64(re.Severity * 1000000)},
		Payload:      re.Data,
	}, nil
}
