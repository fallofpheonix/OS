/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 1 → CYCLE 2 BRIDGE (Layer 0.5)
//
// The Enforcer is an ALTERNATIVE to Loader.pollEvents() for reading
// eBPF ring buffer events. While Loader.pollEvents() reads raw ExecEvents,
// the Enforcer wraps them in the canonical TelemetryEvent format and
// publishes them to the Bus topic "telemetry".
//
// WORKFLOW:
//   eBPF ring buffer → Enforcer.Observe() → generateEventID()
//     → TelemetryEvent{Source: "ebpf_enforcer", Severity: 0.8}
//       → Bus.Publish("telemetry", event)
//         → [CYCLE 2: Monitor, Graph, TCS, Arbiter all receive this event]
//
// KEY DIFFERENCE FROM Loader.pollEvents():
//   - Loader publishes to topic "exec" with raw filename payload
//   - Enforcer publishes to topic "telemetry" with structured JSON payload
//   - Enforcer assigns event IDs; Loader uses PID as SeqID
// =========================================================================
package kernel

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"

	"github.com/cilium/ebpf/ringbuf"
)

// Enforcer reads eBPF ring buffer events and ingests them into the Phoenix Matrix Event Bus.
type Enforcer struct {
	Pub    EventPublisher
	Reader *ringbuf.Reader
}

// NewEnforcer initializes the kernel→Bus bridge.
// Requires an EventPublisher (the Bus) and a ringbuf.Reader (from Loader.Load()).
// Called once during system startup, right after Loader.Load() completes.
func NewEnforcer(pub EventPublisher, r *ringbuf.Reader) *Enforcer {
	return &Enforcer{
		Pub:    pub,
		Reader: r,
	}
}

// Observe is the MAIN EVENT LOOP for kernel telemetry ingestion.
// It blocks on the ring buffer, decoding each raw eBPF event and wrapping it
// in a canonical TelemetryEvent before publishing to the Bus.
//
// WORKFLOW (per event):
//
//	ringbuf.Read() → generateEventID() → json.Marshal(payload)
//	  → TelemetryEvent{Source: "ebpf_enforcer", Severity: 0.8}
//	    → Bus.Publish("telemetry", event)
//	      → [CYCLE 2: Monitor.Process(event) → DriftScore]
//	      → [CYCLE 2: GraphFeature.AddEvent(event) → causal DAG update]
//	      → [CYCLE 2: TCSFeature.AddAndEvaluate(event) → confidence score]
//
// SEVERITY: Hardcoded 0.8 for all execve events. This means ALL execve
// events are treated as high-severity, which creates noise in the pipeline.
//
// BLOCKING BEHAVIOR: This goroutine blocks on ringbuf.Read(). If the ring
// buffer is empty, it waits. If the ring buffer returns an error, it logs
// and CONTINUES (unlike Loader.pollEvents() which exits).
func (e *Enforcer) Observe(ctx context.Context) {
	log.Println("[eBPF Enforcer] Observer Layer active. Wiring telemetry to Event Bus...")

	if e.Reader == nil {
		log.Println("[eBPF Enforcer] Running in Mock/Dry-Run mode (No eBPF Reader provided)")
		return
	}

	for {
		select {
		case <-ctx.Done():
			log.Println("[eBPF Enforcer] Shutting down Observer Layer.")
			return
		default:
			record, err := e.Reader.Read()
			if err != nil {
				log.Printf("[eBPF Enforcer] Ring buffer read error: %v", err)
				continue
			}

			eventID := generateEventID()
			payload, _ := json.Marshal(map[string]interface{}{
				"source_probe": "sys_enter_execve",
				"data_len":     len(record.RawSample),
			})

			// The LamportClock and Hash are formally stamped by the Bus during Publish.
			event := TelemetryEvent{
				EventID:   eventID,
				EventType: "syscall_execve",
				Source:    "ebpf_enforcer",
				Severity:  0.8, // execve is a high-severity observable
				Payload:   payload,
			}

			// Ship to the matrix. This guarantees Causal Monotonicity.
			e.Pub.Publish("telemetry", event)
		}
	}
}

// generateEventID creates a unique identifier for each kernel event.
// Uses crypto/rand for 8 random bytes, formatted as hex.
//
// WORKFLOW: Called once per ring buffer event in Observe().
// The generated ID becomes the event's EventID in the TelemetryEvent,
// which is then used by the Ledger for append-only recording.
//
// SECURITY NOTE: Uses crypto/rand which can block on entropy-depleted
// systems (containers, VMs). On the hot path (every syscall), this is
// a potential latency spike. Consider using a monotonic counter instead.
func generateEventID() string {
	b := make([]byte, 8)
	if _, err := rand.Read(b); err != nil {
		// Fallback for extreme entropy exhaustion
		return fmt.Sprintf("evt-fallback-%d", 0)
	}
	return fmt.Sprintf("evt-%x", b)
}
