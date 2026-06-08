/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package drift

import (
	"math"
)

// Baseline defines the expected drift thresholds per module.
type Baseline struct {
	Modules map[string]float64
}

// ComputeDrift calculates the aggregate system drift from historical records against a baseline.
func ComputeDrift(history []DriftRecord, baseline Baseline) float64 {
	if len(history) == 0 {
		return 0.0
	}

	var data []float64
	for _, rec := range history {
		data = append(data, rec.Drift)
	}

	// Calculate current statistical drift relative to history
	statDrift := CalculateStatisticalDrift(data)

	// Normalize statistical drift (Z-score)
	normalizedStatDrift := math.Min(1.0, math.Abs(statDrift)/3.0)

	// Calculate aggregate baseline deviation
	var totalBaseline float64
	for _, val := range baseline.Modules {
		totalBaseline += val
	}
	avgBaseline := 0.0
	if len(baseline.Modules) > 0 {
		avgBaseline = totalBaseline / float64(len(baseline.Modules))
	}

	return (normalizedStatDrift + avgBaseline) / 2.0
}
