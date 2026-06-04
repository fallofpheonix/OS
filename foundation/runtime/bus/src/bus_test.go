/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package main

import "testing"

func TestBus(t *testing.T) {
	bus := NewBus()
	ch := bus.Subscribe("test")

	bus.Publish("test", "data")
	msg := <-ch
	if msg.Data != "data" {
		t.Errorf("Expected data, got %v", msg.Data)
	}
}

func BenchmarkBus(b *testing.B) {
	bus := NewBus()
	_ = bus.Subscribe("bench")
	for i := 0; i < b.N; i++ {
		bus.Publish("bench", "payload")
	}
}
