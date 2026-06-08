/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package drift

import "testing"

func TestTrendDetector(t *testing.T) {
	// Rising trend
	rising := []float64{0.1, 0.2, 0.3, 0.4, 0.5}
	if trend := DetectTrend(rising); trend != Rising {
		t.Errorf("Expected Rising, got %s", trend)
	}

	// Falling trend
	falling := []float64{0.5, 0.4, 0.3, 0.2, 0.1}
	if trend := DetectTrend(falling); trend != Falling {
		t.Errorf("Expected Falling, got %s", trend)
	}

	// Stable trend
	stable := []float64{0.2, 0.201, 0.199, 0.202, 0.2}
	if trend := DetectTrend(stable); trend != Stable {
		t.Errorf("Expected Stable, got %s", trend)
	}
    
    // Oscillating trend
	oscillating := []float64{0.1, 0.5, 0.1, 0.5, 0.1, 0.5}
	if trend := DetectTrend(oscillating); trend != Oscillating {
		t.Errorf("Expected Oscillating, got %s", trend)
	}
}
