package validation

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
)

// TestTelemetryJitter verifies that logical clock sequencing orders events correctly despite wall-clock ingestion jitter.
func TestTelemetryJitter(t *testing.T) {
	audit := containment.NewProcessAudit()

	// Log actions with dynamic delays (simulating telemetry jitter)
	audit.LogAction(containment.ProcessAction{PID: 10, Action: containment.ActionMonitor})
	time.Sleep(5 * time.Millisecond)
	audit.LogAction(containment.ProcessAction{PID: 20, Action: containment.ActionThrottle})

	// Verify that sequences increment monotonically, independent of host jitter delay
	if audit.History[0].Sequence != 1 || audit.History[1].Sequence != 2 {
		t.Errorf("telemetry jitter caused sequence ordering failure: seq1=%d, seq2=%d", audit.History[0].Sequence, audit.History[1].Sequence)
	}
}

// TestDuplicateStorm verifies the audit trail handles duplicate action storm logs safely.
func TestDuplicateStorm(t *testing.T) {
	audit := containment.NewProcessAudit()

	action1 := containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "event-1"}
	action2 := containment.ProcessAction{PID: 100, Action: containment.ActionMonitor, Reason: "event-1"} // duplicate payload

	audit.LogAction(action1)
	audit.LogAction(action2)

	// In PhoenixOS, every logged action is assigned a unique logical sequence number and serialized hash,
	// so duplicate payloads are safely logged as distinct historical events in the audit trail.
	if audit.History[0].Hash == audit.History[1].Hash {
		t.Errorf("duplicate payload generated identical hashes, but must have unique logical clock / sequence metadata")
	}
	if audit.History[0].Sequence == audit.History[1].Sequence {
		t.Error("expected different sequence numbers for subsequent events")
	}
}

// TestMissingFrames verifies that gaps in sequence numbers or history records can be audited.
func TestMissingFrames(t *testing.T) {
	audit := file.NewFileAudit()

	audit.LogAction(file.FileAction{Path: "/bin/sh", Action: file.ActionMonitor})
	audit.LogAction(file.FileAction{Path: "/bin/ls", Action: file.ActionThrottle})

	// Audit missing sequence check (assert no gaps in sequence numbers)
	for i := 0; i < len(audit.History); i++ {
		expectedSeq := i + 1
		if audit.History[i].Sequence != expectedSeq {
			t.Errorf("sequence gap detected at position %d: expected %d, got %d", i, expectedSeq, audit.History[i].Sequence)
		}
	}
}
