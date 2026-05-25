package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestForkBombThrottling(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	b.Subscribe(topic)

	// Simulate fork bomb by flooding the bus with low-severity events
	for i := 0; i < bus.QueueCapacity + 100; i++ {
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), Severity: 0.1})
	}

	// Bus should throttle (drop) the excess events
	if b.Dropped == 0 {
		t.Error("Expected events to be dropped under flood")
	}
	fmt.Printf("[PX-007] Fork Bomb Throttling: Dropped=%d\n", b.Dropped)
}
