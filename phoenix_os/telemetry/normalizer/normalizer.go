package normalizer

import (
	"encoding/json"
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/clock"
)

// RawEvent represents un-normalized telemetry received from Linux kernel.
type RawEvent struct {
	Timestamp int64           `json:"ts"`
	Source    string          `json:"src"`
	PID       int             `json:"pid"`
	Type      string          `json:"type"`
	Severity  float64         `json:"sev"`
	Data      json.RawMessage `json:"data"`
}

// Normalizer transforms raw input into authoritative bus TelemetryEvents.
type Normalizer struct {
	allocator *clock.SequenceAllocator
}

func NewNormalizer(alloc *clock.SequenceAllocator) *Normalizer {
	return &Normalizer{allocator: alloc}
}

// Normalize converts a RawEvent into a canonical TelemetryEvent with an assigned SeqID.
func (n *Normalizer) Normalize(raw []byte) (*bus.TelemetryEvent, error) {
	var re RawEvent
	if err := json.Unmarshal(raw, &re); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw event: %w", err)
	}

	// Assign monotonic SeqID from the allocator
	seqID := n.allocator.Allocate()

	// Transform to canonical TelemetryEvent
	return &bus.TelemetryEvent{
		SeqID:        int64(seqID),
		MonotonicNs:  re.Timestamp,
		WallTimeUnix: re.Timestamp,
		Source:       re.Source,
		PID:          re.PID,
		EventType:    re.Type,
		Severity:     re.Severity,
		Payload:      re.Data,
	}, nil
}
