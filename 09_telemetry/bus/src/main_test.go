package main

import (
	"testing"
)

func TestEventBus(t *testing.T) {
	bus := NewEventBus()
	ch := bus.Subscribe("test")

	expected := "data"
	bus.Publish("test", expected)

	received := <-ch
	if received != expected {
		t.Errorf("Expected %v, got %v", expected, received)
	}
}

func BenchmarkEventBus(b *testing.B) {
	bus := NewEventBus()
	_ = bus.Subscribe("bench")

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		bus.Publish("bench", "payload")
	}
}

func BenchmarkEventBusParallel(b *testing.B) {
	bus := NewEventBus()
	_ = bus.Subscribe("bench")

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			bus.Publish("bench", "payload")
		}
	})
}
