package kernel

import (
	"encoding/json"
	"phoenix/bus"
)

// NormalizedEvent represents the standardized PhoenixOS event format (L3).
type NormalizedEvent struct {
	Type       string  `json:"type"`
	Source     string  `json:"source"`
	Confidence float64 `json:"confidence"`
}

// Normalizer converts raw kernel events into high-level system events.
type Normalizer struct {
	InBus  *bus.Bus
	OutBus *bus.Bus
}

func NewNormalizer(in, out *bus.Bus) *Normalizer {
	return &Normalizer{InBus: in, OutBus: out}
}

func (n *Normalizer) Start() {
	rawEvents := n.InBus.Subscribe("kernel.telemetry.raw")
	go func() {
		for event := range rawEvents {
			n.processEvent(event)
		}
	}()
}

func (n *Normalizer) processEvent(raw bus.TelemetryEvent) {
	var ke KernelEvent
	if err := json.Unmarshal(raw.Payload, &ke); err != nil {
		return
	}

	normalized := NormalizedEvent{
		Source:     "kernel",
		Confidence: 1.0, // Kernel-sourced events have absolute confidence
	}

	// Map syscalls to normalized event types
	switch ke.Syscall {
	case "execve":
		normalized.Type = "PROCESS_START"
	case "fork", "clone":
		normalized.Type = "PROCESS_FORK"
	case "exit":
		normalized.Type = "PROCESS_EXIT"
	case "open":
		normalized.Type = "FILE_ACCESS"
	case "read":
		normalized.Type = "DATA_READ"
	case "write":
		normalized.Type = "DATA_WRITE"
	case "connect", "accept", "bind":
		normalized.Type = "NETWORK_EVENT"
	default:
		normalized.Type = "GENERIC_SYSCALL"
	}

	payload, _ := json.Marshal(normalized)
	
	n.OutBus.Publish("system.events.normalized", bus.TelemetryEvent{
		SeqID:        raw.SeqID,
		WallTimeUnix: raw.WallTimeUnix,
		Source:       "phoenix.normalizer",
		EventType:    normalized.Type,
		Severity:     raw.Severity,
		Payload:      payload,
		PID:          ke.PID,
	})
}
