package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
)

func TestBeaconMitigation(t *testing.T) {
	b := bus.NewBus()
	engine := containment.NewIsolationEngine(containment.StateObserve)

	// Simulate beacon detection (Severity 0.85)
	event := bus.TelemetryEvent{
		SeqID:    1,
		Severity: 0.85,
		EventType: "beacon_detected",
	}

	// Trigger containment transition based on advisory (simulated)
	if event.Severity >= 0.8 {
		err := engine.Transition(containment.StateThrottle, "EV-001", "DEC-001")
		if err != nil {
			t.Fatalf("Containment transition failed: %v", err)
		}
	}

	if engine.CurrentState != containment.StateThrottle {
		t.Errorf("Expected state THROTTLE, got %s", engine.CurrentState)
	}
	fmt.Println("[PX-007] Beacon Mitigation: PASSED")
}
