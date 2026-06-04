/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package kernel

import (
	"fmt"
	"sync/atomic"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

func TestRingOverflow(t *testing.T) {
	b := bus.NewBus()
	topic := "kernel"
	ch := b.Subscribe(topic)

	var overflowTriggered int32
	b.OnOverflow = func(t string, pressure float64, event bus.TelemetryEvent) {
		atomic.StoreInt32(&overflowTriggered, 1)
	}

	// Fill queue to HighWatermark
	count := int(bus.QueueCapacity * 96 / 100)
	for i := 0; i < count; i++ {
		sev := 0.95
		if i > int(bus.QueueCapacity*94/100) {
			sev = 0.85
		}
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), Severity: sev})
	}

	pressure := b.QueuePressure(topic)
	if pressure < bus.CriticalWatermark {
		t.Errorf("Expected pressure above critical watermark, got %f", pressure)
	}

	// Wait up to 500ms for overflow to be triggered since it runs in a goroutine
	for i := 0; i < 50; i++ {
		if atomic.LoadInt32(&overflowTriggered) == 1 {
			break
		}
		time.Sleep(10 * time.Millisecond)
	}

	if atomic.LoadInt32(&overflowTriggered) == 0 {
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
