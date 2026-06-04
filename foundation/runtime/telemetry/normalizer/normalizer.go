/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 1c — TELEMETRY NORMALIZATION (Layer 0.5)
//
// The Normalizer transforms raw kernel telemetry into canonical
// TelemetryEvents with monotonic sequence IDs from the logical clock.
//
// WORKFLOW:
//   Raw kernel event → Normalizer.Normalize(raw)
//     → json.Unmarshal(raw) → RawEvent{Timestamp, Source, PID, Type, Severity}
//     → clock.Tick() → assign monotonic SeqID
//     → Transform to bus.TelemetryEvent with canonical fields
//     → Return normalized event for Bus distribution
//
// PURPOSE: Raw kernel events have inconsistent formats and no sequence IDs.
// The Normalizer ensures all events entering the Bus have:
//   - Monotonically increasing SeqID (from logical clock)
//   - Standardized field names
//   - Consistent timestamp format
//
// THREAD SAFETY: The logical clock uses atomic operations for
// concurrent-safe sequence ID generation.
// =========================================================================
package normalizer

import (
	"encoding/json"
	"fmt"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	clock "github.com/fallofpheonix/phoenix/foundation/runtime/common/logical_clock"
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
	clock *clock.Clock
}

func NewNormalizer(c *clock.Clock) *Normalizer {
	return &Normalizer{clock: c}
}

// Normalize converts a RawEvent into a canonical TelemetryEvent with an assigned SeqID.
func (n *Normalizer) Normalize(raw []byte) (*bus.TelemetryEvent, error) {
	var re RawEvent
	if err := json.Unmarshal(raw, &re); err != nil {
		return nil, fmt.Errorf("failed to unmarshal raw event: %w", err)
	}

	// Assign monotonic SeqID from the clock
	seqID := n.clock.Tick()

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
