package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/10_kernel/runtime"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestRuntimeRingPressure(t *testing.T) {
	b := bus.NewBus()
	monitor := runtime.NewRingMonitor(b)
	injector := runtime.NewProbeInjector(b)

	injector.StressBurst("kernel", 1000, 0.5)
	
	pressure := monitor.GetPressure("kernel")
	if pressure <= 0 {
		t.Error("Expected recorded ring pressure")
	}
	fmt.Printf("[PX-013] Runtime Ring Pressure: %f\n", pressure)
}
