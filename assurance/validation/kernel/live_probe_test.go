/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package kernel

import (
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/kernel/runtime"
)

func TestLiveProbeInjection(t *testing.T) {
	b := bus.NewBus()
	injector := runtime.NewProbeInjector(b)
	ch := b.Subscribe("kernel")

	injector.InjectEvent("kernel", bus.TelemetryEvent{SeqID: 100})

	e := <-ch
	if e.SeqID != 100 {
		t.Errorf("Expected SeqID 100, got %d", e.SeqID)
	}
	fmt.Println("[PX-013] Live Probe Injection: PASSED")
}
