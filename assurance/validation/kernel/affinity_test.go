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

	"github.com/fallofpheonix/phoenix/foundation/runtime/kernel/runtime"
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
