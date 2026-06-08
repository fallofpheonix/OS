/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/*
 * REPOSITORY: PhoenixKernel
 * ARCHITECTURAL JUSTIFICATION: Telemetry bridge from eBPF ring buffer to PhoenixCore Event Bus.
 * DEPENDENCY BOUNDARY: Depends on PhoenixCore/bus. Telemetry only.
 * DETERMINISTIC CONSIDERATIONS: Non-blocking dispatch, monotonic time anchoring.
 */

package bridge

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"log"
	"time"

	"github.com/cilium/ebpf/ringbuf"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// SyscallEvent must match the C struct in syscalls.c
type SyscallEvent struct {
	Pid         uint32
	Timestamp   uint64
	SyscallID   uint32
	EntropyFlag uint32
}

// TelemetryBridge consumes raw eBPF events and republishes to the OS bus.
type TelemetryBridge struct {
	Reader *ringbuf.Reader
	Bus    *bus.Bus
}

// NewTelemetryBridge initializes the kernel-to-user bridge.
func NewTelemetryBridge(reader *ringbuf.Reader, b *bus.Bus) *TelemetryBridge {
	return &TelemetryBridge{
		Reader: reader,
		Bus:    b,
	}
}

// Run starts the continuous event translation loop.
func (t *TelemetryBridge) Run() {
	for {
		record, err := t.Reader.Read()
		if err != nil {
			log.Printf("[PhoenixKernel] eBPF ring buffer error: %v", err)
			continue
		}

		var event SyscallEvent
		if err := binary.Read(bytes.NewReader(record.RawSample), binary.LittleEndian, &event); err != nil {
			log.Printf("[PhoenixKernel] Failed to decode kernel event: %v", err)
			continue
		}

		// Convert to canonical TelemetryEvent
		payload, _ := json.Marshal(event)
		eventType := "SYSCALL_STABLE"
		severity := phxmath.FixedPoint{V: 100000} // 0.1
		if event.EntropyFlag == 1 {
			eventType = "SYSCALL_HIGH_ENTROPY"
			severity = phxmath.FixedPoint{V: 800000} // 0.8
		}

		telem := bus.TelemetryEvent{
			SeqID:        int64(event.Timestamp),
			MonotonicNs:  int64(event.Timestamp),
			WallTimeUnix: time.Now().Unix(),
			Source:       "PhoenixKernel:SyscallBoundary",
			PID:          int(event.Pid),
			EventType:    eventType,
			Severity:     severity,
			Payload:      payload,
		}

		// Non-blocking publish to the central bus
		go t.Bus.Publish("phoenix.events.normal", telem)
	}
}
