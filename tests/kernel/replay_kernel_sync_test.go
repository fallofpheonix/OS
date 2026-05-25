package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/kernel/sandbox"
)

func TestReplayKernelSync(t *testing.T) {
	// 1. Generate a sequence of events
	k1 := sandbox.NewKernelSimulator()
	events := []struct {
		data string
		size int
	}{
		{"event-1", 10},
		{"event-2", 20},
		{"event-3", 15},
	}
	
	for _, e := range events {
		_ = k1.SubmitToRingBuffer(e.data, e.size)
	}
	
	// 2. Replay the same sequence in a new simulator
	k2 := sandbox.NewKernelSimulator()
	for _, e := range events {
		_ = k2.SubmitToRingBuffer(e.data, e.size)
	}
	
	// 3. Verify state parity
	if k1.RingBufferUsed != k2.RingBufferUsed {
		t.Errorf("State mismatch: RingBufferUsed k1=%d, k2=%d", k1.RingBufferUsed, k2.RingBufferUsed)
	}
	
	if len(k1.Events) != len(k2.Events) {
		t.Errorf("State mismatch: Event count k1=%d, k2=%d", len(k1.Events), len(k2.Events))
	}
	
	for i := range k1.Events {
		if k1.Events[i].Data != k2.Events[i].Data {
			t.Errorf("Event %d data mismatch", i)
		}
		// Note: IDs will match because simulator starts at 0, 
		// but Timestamps might differ unless we mocked the clock.
	}
}
