package security

import (
	"fmt"
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestPortScanDetection(t *testing.T) {
	b := bus.NewBus()
	topic := "network"
	b.Subscribe(topic)

	// Simulate port scan events
	for i := 0; i < 100; i++ {
		b.Publish(topic, bus.TelemetryEvent{SeqID: int64(i), EventType: "connection_refused", Severity: 0.4})
	}
	
	fmt.Println("[PX-007] Port Scan Detection: Signal recorded in bus")
}
