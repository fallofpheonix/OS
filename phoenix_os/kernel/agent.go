package kernel

import (
	"fmt"
	"runtime"
	"time"

	"phoenix/bus"
)

// TelemetryAgent coordinates the ingestion of kernel events into the Phoenix Bus.
type TelemetryAgent struct {
	Bus *bus.Bus
}

func NewTelemetryAgent(b *bus.Bus) *TelemetryAgent {
	return &TelemetryAgent{Bus: b}
}

// Start initiates telemetry collection. On non-Linux systems, it falls back to simulation.
func (a *TelemetryAgent) Start() {
	if runtime.GOOS != "linux" {
		fmt.Printf("[KERNEL AGENT] Non-Linux environment (%s) detected. Falling back to Mock Simulation.\n", runtime.GOOS)
		go a.runSimulation()
		return
	}

	fmt.Println("[KERNEL AGENT] Linux environment detected. Initializing eBPF Probes...")
	// In a real implementation, we would load the eBPF ELF here and start reading from the ring buffer.
	// go a.runEBPF()
}

func (a *TelemetryAgent) runSimulation() {
	mock := NewMockGenerator()
	ticker := time.NewTicker(100 * time.Millisecond) // 10 events per second for visibility
	defer ticker.Stop()

	for range ticker.C {
		event := mock.Generate()
		
		// Publish to Bus for Stage 2 processing
		a.Bus.Publish("kernel.telemetry.raw", bus.TelemetryEvent{
			SeqID:        time.Now().UnixNano(),
			WallTimeUnix: time.Now().Unix(),
			Source:       "phoenix.kernel.ebpf",
			EventType:    "kernel.syscall",
			Severity:     0.1,
			Payload:      event.ToJSON(),
		})
	}
}
