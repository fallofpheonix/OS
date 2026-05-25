package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/10_kernel/runtime"
)

func TestClockSkewDetection(t *testing.T) {
	cs := &runtime.ClockSkew{}
	cs.RecordSkew(1000, 1050)
	
	if cs.GetDrift() != -50 {
		t.Errorf("Expected -50ns drift, got %dns", cs.GetDrift())
	}
	fmt.Println("[PX-013] Clock Skew Detection: PASSED")
}
