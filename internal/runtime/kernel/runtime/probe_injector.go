/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package runtime

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// ProbeInjector simulates the injection of events into the eBPF ring buffer.
type ProbeInjector struct {
	targetBus *bus.Bus
}

func NewProbeInjector(b *bus.Bus) *ProbeInjector {
	return &ProbeInjector{targetBus: b}
}

func (pi *ProbeInjector) InjectEvent(topic string, event bus.TelemetryEvent) {
	pi.targetBus.Publish(topic, event)
}

func (pi *ProbeInjector) StressBurst(topic string, count int, severity float64) {
	fpSeverity := phxmath.NewFixedPointRaw(int64(severity * 1000000))
	for i := 0; i < count; i++ {
		pi.targetBus.Publish(topic, bus.TelemetryEvent{
			SeqID:    int64(i),
			Severity: fpSeverity,
		})
	}
}
