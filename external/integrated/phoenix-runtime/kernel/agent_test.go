package kernel

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestTelemetryAgentMock(t *testing.T) {
	b := bus.NewBus()
	agent := NewTelemetryAgent(b)

	// Subscribe to raw telemetry channel
	events := b.Subscribe("kernel.telemetry.raw")

	// Start agent (mock mode on macOS)
	go agent.Start()

	// Wait for at least 3 events
	count := 0
	timeout := time.After(2 * time.Second)

	for count < 3 {
		select {
		case <-events:
			count++
		case <-timeout:
			t.Fatalf("Timed out waiting for mock events. Received %d", count)
		}
	}

	if count < 3 {
		t.Errorf("Expected at least 3 events, got %d", count)
	}
}
