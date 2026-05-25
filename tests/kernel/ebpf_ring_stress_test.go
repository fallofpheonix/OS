package kernel

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestRingOverflow(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	ch := b.Subscribe(topic)

	overflowTriggered := false
	b.OnOverflow = func(t string, pressure float64, event bus.TelemetryEvent) {
		overflowTriggered = true
	}

	// Fill queue to HighWatermark
	count := int(float64(bus.QueueCapacity) * 0.96)
	for i := 0; i < count; i++ {
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), Severity: 0.1})
	}

	pressure := b.QueuePressure(topic)
	if pressure < bus.CriticalWatermark {
		t.Errorf("Expected pressure above critical watermark, got %f", pressure)
	}

	if !overflowTriggered {
		t.Error("Expected OnOverflow callback to be triggered")
	}

	if b.Dropped == 0 {
		t.Error("Expected dropped events due to severity filtering in critical zone")
	}

	fmt.Printf("[PX-001] Ring Overflow: Pressure=%f, Dropped=%d\n", pressure, b.Dropped)
	
	// Flush channel
	for len(ch) > 0 {
		<-ch
	}
}

func TestDroppedFrames(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	b.Subscribe(topic) // capacity is 65536

	// Publish more than capacity with low severity
	total := bus.QueueCapacity + 1000
	for i := 0; i < total; i++ {
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), Severity: 0.1})
	}

	if b.Dropped == 0 {
		t.Error("Expected dropped events when exceeding capacity")
	}
	fmt.Printf("[PX-001] Dropped Frames: Total=%d, Dropped=%d\n", total, b.Dropped)
}

func TestKernelOrdering(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	ch := b.Subscribe(topic)

	// Publish events with specific sequence IDs
	b.Publish(topic, bus.TelemetryEvent{SeqID: 1, LogicalTick: 10})
	b.Publish(topic, bus.TelemetryEvent{SeqID: 2, LogicalTick: 11})

	e1 := <-ch
	e2 := <-ch

	if e1.SeqID != 1 || e2.SeqID != 2 {
		t.Errorf("Expected events in order 1, 2; got %d, %d", e1.SeqID, e2.SeqID)
	}
	if e1.LogicalTick >= e2.LogicalTick {
		t.Errorf("Expected monotonic ticks, got %d then %d", e1.LogicalTick, e2.LogicalTick)
	}
	fmt.Println("[PX-001] Kernel Ordering: PASSED")
}

func TestKernelBurst(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	b.Subscribe(topic)

	// Rapid burst of high priority events
	for i := 0; i < 1000; i++ {
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), Severity: 0.95})
	}

	if b.Dropped > 0 {
		t.Errorf("Critical events dropped during burst: %d", b.Dropped)
	}
	fmt.Println("[PX-001] Kernel Burst: PASSED")
}
