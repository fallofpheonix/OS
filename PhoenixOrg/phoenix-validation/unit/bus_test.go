package unit

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
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
