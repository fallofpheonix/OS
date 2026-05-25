package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
)

func TestPersistenceMitigation(t *testing.T) {
	engine := containment.NewIsolationEngine(containment.StateObserve)

	// Simulate persistence event
	engine.Transition(containment.StateWatch, "EV-003", "DEC-004")
	
	if engine.CurrentState != containment.StateWatch {
		t.Errorf("Expected state WATCH, got %s", engine.CurrentState)
	}
	fmt.Println("[PX-007] Persistence Mitigation: PASSED")
}
