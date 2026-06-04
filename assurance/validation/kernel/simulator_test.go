/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package kernel

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/kernel/sandbox"
)

func TestKernelSimulator(t *testing.T) {
	ks := sandbox.NewKernelSimulator()

	// Test map limits
	for i := 0; i < 1024; i++ {
		err := ks.UpdateMap("test", i)
		if err != nil {
			t.Errorf("Failed to update map at entry %d: %v", i, err)
		}
	}
	err := ks.UpdateMap("overflow", 1025)
	if err == nil {
		t.Error("Expected map overflow error, got nil")
	}

	// Test stack depth
	err = ks.CheckStackDepth(256)
	if err != nil {
		t.Errorf("Stack depth 256 should be fine: %v", err)
	}
	err = ks.CheckStackDepth(1024)
	if err == nil {
		t.Error("Expected stack depth error, got nil")
	}

	// Test global energy budget
	err = ks.RequestEnergy(500.0)
	if err != nil {
		t.Errorf("Failed to request 500 energy: %v", err)
	}

	err = ks.RequestEnergy(600.0)
	if err == nil {
		t.Error("Expected energy budget to be exceeded")
	}

	if ks.ConsumedEnergy != 500.0 {
		t.Errorf("Expected 500 consumed energy, got %v", ks.ConsumedEnergy)
	}
}
