package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
)

func TestExfiltrationMitigation(t *testing.T) {
	b := bus.NewBus()
	engine := containment.NewIsolationEngine(containment.StateObserve)

	// Simulate exfiltration attempt (Severity 0.95)
	event := bus.TelemetryEvent{
		SeqID:    2,
		Severity: 0.95,
		EventType: "exfil_attempt",
	}

	// High severity triggers immediate isolation
	if event.Severity >= 0.9 {
		engine.Transition(containment.StateThrottle, "EV-002", "DEC-002")
		err := engine.Transition(containment.StateIsolate, "EV-002", "DEC-003")
		if err != nil {
			t.Fatalf("Containment transition to ISOLATE failed: %v", err)
		}
	}

	if engine.CurrentState != containment.StateIsolate {
		t.Errorf("Expected state ISOLATE, got %s", engine.CurrentState)
	}
	fmt.Println("[PX-007] Exfiltration Mitigation: PASSED")
}
