/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package disorder

import (
	"math"
)

// CalculateSDI computes the Security Disorder Index using Shannon Entropy.
// In this context, probabilities p_s are derived from event category distributions.
func CalculateSDI(distribution map[string]float64) float64 {
	var sdi float64
	var total float64

	for _, count := range distribution {
		total += count
	}

	if total == 0 {
		return 0.0
	}

	for _, count := range distribution {
		p := count / total
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}

	return sdi
}
