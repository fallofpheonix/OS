/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

func TestBus(t *testing.T) {
	b := bus.NewBus()
	ch := b.Subscribe("test")

	b.Publish("test", bus.TelemetryEvent{Payload: []byte("data")})
	msg := <-ch
	if string(msg.Payload) != "data" {
		t.Errorf("Expected data, got %s", string(msg.Payload))
	}
}

func BenchmarkBus(b *testing.B) {
	busObj := bus.NewBus()
	_ = busObj.Subscribe("bench")
	for i := 0; i < b.N; i++ {
		busObj.Publish("bench", bus.TelemetryEvent{Payload: []byte("payload")})
	}
}
