package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/kernel/sandbox"
)

func TestEventLossMetrics(t *testing.T) {
	k := sandbox.NewKernelSimulator()
	k.RingBufferSize = 100
	
	// Intentionally cause loss
	for i := 0; i < 20; i++ {
		_ = k.SubmitToRingBuffer("data", 10)
	}
	
	if k.DroppedEvents != 10 {
		t.Errorf("Expected 10 dropped events, got %d", k.DroppedEvents)
	}
	
	// Verify remaining events are still valid
	count := 0
	for {
		_, err := k.ConsumeFromRingBuffer()
		if err != nil {
			break
		}
		count++
	}
	
	if count != 10 {
		t.Errorf("Expected 10 events in buffer, got %d", count)
	}
}
