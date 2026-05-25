package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/10_kernel/runtime"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
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
