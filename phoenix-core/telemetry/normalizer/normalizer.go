package normalizer

import (
	"encoding/json"
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix-core/telemetry/bus"
)

// RawEvent represents un-normalized telemetry received from Linux kernel.
type RawEvent struct {
	Timestamp int64           `json:"ts"`
	Source    string          `json:"src"`
	PID       int             `json:"pid"`
	Type      string          `json:"type"`
	Data      json.RawMessage `json:"data"`
}

// Normalizer transforms raw input into authoritative bus TelemetryEvents.
type Normalizer struct{}

func NewNormalizer() *Normalizer {
	return &Normalizer{}
}

// Normalize converts a RawEvent into a canonical TelemetryEvent.
func (n *Normalizer) Normalize(raw []byte) (*bus.TelemetryEvent, error) {
	var re RawEvent
	if err := json.Unmarshal(raw, &re); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw event: %w", err)
	}

	// Transform to canonical TelemetryEvent
	return &bus.TelemetryEvent{
		MonotonicNs:  re.Timestamp, // Simplified mapping
		WallTimeUnix: re.Timestamp,
		Source:       re.Source,
		PID:          re.PID,
		EventType:    re.Type,
		Payload:      re.Data,
	}, nil
}
