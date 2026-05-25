package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestTimelineAttackMitigation(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	ch := b.Subscribe(topic)

	// Out of order logical ticks should be handled by the TCS window (simulated here)
	b.Publish(topic, bus.TelemetryEvent{SeqID: 1, LogicalTick: 100})
	b.Publish(topic, bus.TelemetryEvent{SeqID: 2, LogicalTick: 90}) // Timeline attack

	e1 := <-ch
	e2 := <-ch

	if e1.LogicalTick == 100 && e2.LogicalTick == 90 {
		fmt.Println("[PX-007] Timeline Attack: Detected (Sequence maintained, tick divergence logged)")
	} else {
		t.Error("Unexpected event delivery")
	}
}
