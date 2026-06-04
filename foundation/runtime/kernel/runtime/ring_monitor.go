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
)

// RingMonitor tracks pressure and drops in the eBPF ring buffer.
type RingMonitor struct {
	bus *bus.Bus
}

func NewRingMonitor(b *bus.Bus) *RingMonitor {
	return &RingMonitor{bus: b}
}

func (rm *RingMonitor) GetPressure(topic string) float64 {
	return rm.bus.QueuePressure(topic)
}

func (rm *RingMonitor) GetDroppedCount() int64 {
	return rm.bus.Dropped
}
