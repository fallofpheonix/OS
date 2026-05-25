package kernel

import (
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestKernelReplayFidelity(t *testing.T) {
	inBus := bus.NewBus()
	outBus := bus.NewBus()

	norm := NewNormalizer(inBus, outBus)
	norm.Start()

	// 1. Define a fixed set of raw kernel events (The Input)
	rawEvents := []KernelEvent{
		{EventID: "1", PID: 101, Syscall: "execve", Timestamp: 1000},
		{EventID: "2", PID: 101, Syscall: "open", Timestamp: 1100},
		{EventID: "3", PID: 101, Syscall: "write", Timestamp: 1200},
	}

	// 2. First Run: Process input and capture output
	output1 := runNormalization(inBus, outBus, rawEvents)

	// 3. Second Run: Process same input again
	output2 := runNormalization(inBus, outBus, rawEvents)

	// 4. Verify Determinism (The Replay Contract)
	if len(output1) != len(output2) {
		t.Fatalf("Replay length mismatch: Run 1 produced %d events, Run 2 produced %d", len(output1), len(output2))
	}

	for i := range output1 {
		var n1, n2 NormalizedEvent
		json.Unmarshal(output1[i].Payload, &n1)
		json.Unmarshal(output2[i].Payload, &n2)

		if n1.Type != n2.Type {
			t.Errorf("Divergence at event %d: Type %s != %s", i, n1.Type, n2.Type)
		}
	}
}

func runNormalization(in, out *bus.Bus, inputs []KernelEvent) []bus.TelemetryEvent {
	results := out.Subscribe("system.events.normalized")

	for _, input := range inputs {
		in.Publish("kernel.telemetry.raw", bus.TelemetryEvent{
			SeqID:   int64(input.Timestamp),
			Payload: input.ToJSON(),
		})
	}

	// Capture outputs
	var outputs []bus.TelemetryEvent
	for len(outputs) < len(inputs) {
		outputs = append(outputs, <-results)
	}
	return outputs
}
