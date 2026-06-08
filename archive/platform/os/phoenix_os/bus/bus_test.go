/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package bus

import (
	"testing"
)

func TestBusOrderingProof(t *testing.T) {
	b := NewBus()
	ch := b.Subscribe("test-topic")

	e1 := TelemetryEvent{EventID: "e1", Severity: 1.0}
	e2 := TelemetryEvent{EventID: "e2", Severity: 1.0}

	b.Publish("test-topic", e1)
	b.Publish("test-topic", e2)

	received1 := <-ch
	received2 := <-ch

	if received1.LamportClock >= received2.LamportClock {
		t.Errorf("Expected LamportClock to be monotonic, got %d then %d", received1.LamportClock, received2.LamportClock)
	}

	if received2.PrevHash != received1.Hash {
		t.Errorf("e2.PrevHash should be set to e1.Hash")
	}
}
