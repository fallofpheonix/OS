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
