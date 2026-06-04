/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package security

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/security/physics"
	"testing"
)

func TestPhysicsState(t *testing.T) {
	counts := map[string]float64{
		"process":    50,
		"filesystem": 50,
	}

	state := physics.ComputeState(counts, 0.5)

	// ln(2) approx 0.693
	if state.Entropy < 0.69 || state.Entropy > 0.70 {
		t.Errorf("Expected entropy around 0.693, got %f", state.Entropy)
	}

	if state.Temperature != 50.0 {
		t.Errorf("Expected temperature 50.0, got %f", state.Temperature)
	}
}
