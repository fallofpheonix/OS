/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package bus

import (
	"sync"
	"testing"
	"time"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

func TestBusPriorityLaneAndEvidenceReserve(t *testing.T) {
	b := NewBus()
	ch := b.Subscribe("test.telemetry")

	// Fill the subscriber channel up to 90% (which is past HighWatermark = 85%)
	targetSize := QueueCapacity * 90 / 100
	for i := 0; i < targetSize; i++ {
		ch <- TelemetryEvent{SeqID: int64(i), Severity: phxmath.FixedPoint{V: 100000}}
	}

	// Verify that a low-priority event (Severity = 0.5) is dropped under pressure
	droppedBefore := b.Dropped
	b.Publish("test.telemetry", TelemetryEvent{SeqID: 99999, Severity: phxmath.FixedPoint{V: 500000}})
	if b.Dropped <= droppedBefore {
		t.Error("Expected low-priority event to be dropped when fill ratio >= 85%")
	}

	// Verify that a high-priority critical event (Severity = 0.9) is processed successfully
	// Clear one slot so it doesn't hit 100% full queue limit and trigger pre-emption
	<-ch
	b.Publish("test.telemetry", TelemetryEvent{SeqID: 88888, Severity: phxmath.FixedPoint{V: 900000}})

	// Read until we find the high-priority event or empty channel
	found := false
	for len(ch) > 0 {
		e := <-ch
		if e.SeqID == 88888 {
			found = true
			break
		}
	}
	if !found {
		t.Error("Expected high-priority event to bypass the reserve gating and enter the queue")
	}
}

func TestBusOverflowSnapshotTrigger(t *testing.T) {
	b := NewBus()
	ch := b.Subscribe("test.overflow")

	var wg sync.WaitGroup
	wg.Add(1)

	var overflowTriggered bool
	var triggeredTopic string
	var triggeredPressure float64

	b.OnOverflow = func(topic string, pressure float64, event TelemetryEvent) {
		overflowTriggered = true
		triggeredTopic = topic
		triggeredPressure = pressure
		wg.Done()
	}

	// Fill the subscriber channel up to 96% (past CriticalWatermark = 95%)
	targetSize := QueueCapacity * 96 / 100
	for i := 0; i < targetSize; i++ {
		ch <- TelemetryEvent{SeqID: int64(i), Severity: phxmath.FixedPoint{V: 900000}}
	}

	// Publish one critical event (Severity >= 0.9) to trigger the watermark check
	b.Publish("test.overflow", TelemetryEvent{SeqID: 100000, Severity: phxmath.FixedPoint{V: 950000}})

	// Wait with a timeout
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
	case <-time.After(1 * time.Second):
		t.Fatal("Timeout waiting for overflow callback")
	}

	if !overflowTriggered {
		t.Error("Expected OnOverflow callback to be triggered when queue fill ratio >= 95%")
	}
	if triggeredTopic != "test.overflow" {
		t.Errorf("Expected topic 'test.overflow', got %s", triggeredTopic)
	}
	if triggeredPressure < 0.95 {
		t.Errorf("Expected pressure >= 0.95, got %f", triggeredPressure)
	}
}
