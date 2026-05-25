package kernel

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestAttackSimulation(t *testing.T) {
	b := bus.NewBus()
	norm := NewNormalizer(b, b)
	trace := NewTraceEngine(b)

	norm.Start()
	trace.Start()

	// Subscribe to normalized events
	events := b.Subscribe("system.events.normalized")

	// 1. Simulate Reverse Shell (execve bash -> connect)
	b.Publish("kernel.telemetry.raw", bus.TelemetryEvent{
		SeqID:   1,
		PID:     2000,
		Payload: KernelEvent{EventID: "rs1", PID: 2000, Syscall: "execve", Process: "bash"}.ToJSON(),
	})
	b.Publish("kernel.telemetry.raw", bus.TelemetryEvent{
		SeqID:   2,
		PID:     2000,
		Payload: KernelEvent{EventID: "rs2", PID: 2000, Syscall: "connect", Process: "bash"}.ToJSON(),
	})

	// 2. Simulate Fork Bomb
	for i := 0; i < 10; i++ {
		b.Publish("kernel.telemetry.raw", bus.TelemetryEvent{
			SeqID:   int64(10 + i),
			PID:     3000 + i,
			Payload: KernelEvent{EventID: "fb", PID: 3000 + i, Syscall: "fork"}.ToJSON(),
		})
	}

	// Capture and verify
	count := 0
	timeout := time.After(1 * time.Second)

	for count < 12 { // 2 from reverse shell + 10 from fork bomb
		select {
		case <-events:
			count++
		case <-timeout:
			t.Fatalf("Validation failed: only captured %d/12 events", count)
		}
	}

	t.Logf("Successfully validated %d events across attack simulations.", count)
}
