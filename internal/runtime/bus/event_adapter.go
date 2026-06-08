package bus

import (
	"encoding/json"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// PURPOSE: Translates high-level Bus events into low-level Ledger entries.
// CONTRACT: Encapsulates all mapping logic between the telemetry layer
//           and the forensic layer. The Applier depends on this translation
//           to ensure no float64 or unvalidated data enters the ledger.
// FAILURE: Returns an error if the event payload is malformed or
//           mandatory metadata (EventID) is missing.
// CONNECTS: foundation/runtime/bus/applier.go (consumer)

// ToLedgerParams extracts the necessary fields from a TelemetryEvent
// for the LedgerWriter.AddEntry call.
// PURPOSE: Isolates the field mapping to prevent TelemetryEvent schema
//
//	leaking into the Applier's core loop.
//
// CONTRACT: The returned tick is the authoritative monotonic sequence.
func ToLedgerParams(event TelemetryEvent) (eventID string, causeID string, tick uint64, payload []byte) {
	// WHY: EventID is the primary key for global tracing.
	// If missing, tracing is broken.
	eventID = event.EventID

	// WHY: CausalID links events in the Merkle DAG.
	// Root events may have an empty CauseID.
	causeID = event.CausalID

	// WHY: LogicalTick is the monotonic clock tick assigned by the
	// producer (e.g. eBPF probe) or the bus. It defines the "When".
	tick = event.LogicalTick

	// WHY: Severity is now a FixedPoint. We map it into the payload.
	// BLOCKER-005: Severity and other semantic fields are not yet semantically validated.
	wrapper := struct {
		Severity phxmath.FixedPoint `json:"severity"`
		Raw      json.RawMessage    `json:"raw"`
	}{
		Severity: event.Severity,
		Raw:      event.Payload,
	}
	payload, _ = json.Marshal(wrapper)

	return eventID, causeID, tick, payload
}
