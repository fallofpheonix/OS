package kernel

import (
	"testing"
	"github.com/fallofpheonix/phoenix-os/10_kernel/runtime"
)

func TestCPUAffinity(t *testing.T) {
	ar := &runtime.AffinityRunner{}
	err := ar.LockToCore(2)
	if err != nil {
		t.Fatalf("Failed to lock core: %v", err)
	}

	if ar.CurrentCore() != 2 {
		t.Errorf("Expected CoreID 2, got %d", ar.CurrentCore())
	}
	fmt.Println("[PX-013] CPU Affinity: PASSED")
}
