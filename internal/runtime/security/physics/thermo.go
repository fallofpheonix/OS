/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package physics

import (
	"math"

	"github.com/fallofpheonix/phoenix/foundation/runtime/security/physics/disorder"
)

// StateVector is a lightweight compatibility alias for the integrated model demo.
type StateVector []float64

// SystemState represents the thermodynamic state of the system.
type SystemState struct {
	Entropy     float64 `json:"entropy"`
	Temperature float64 `json:"temperature"`
	Energy      float64 `json:"energy"`
	Disorder    float64 `json:"disorder"`
}

// ComputeState calculates the full thermodynamic state from event distributions.
func ComputeState(counts map[string]float64, threatScore float64) SystemState {
	sdi := disorder.CalculateSDI(counts)

	// Abstractions:
	// Energy is the total volume of activity (normalized).
	var energy float64
	for _, v := range counts {
		energy += v
	}

	// Temperature is an abstraction of the Threat Level.
	// High temperature = high threat, leading to "phase transitions" (state changes).
	temperature := threatScore * 100.0

	return SystemState{
		Entropy:     sdi, // Simplified: SDI is our proxy for system entropy
		Temperature: temperature,
		Energy:      energy,
		Disorder:    sdi,
	}
}

// CalculateSDI computes a simple disorder score from a discrete state vector.
func CalculateSDI(states StateVector) float64 {
	if len(states) == 0 {
		return 0
	}
	counts := make(map[float64]float64)
	for _, state := range states {
		counts[state]++
	}
	var sdi float64
	total := float64(len(states))
	for _, count := range counts {
		p := count / total
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}
	return sdi
}

// CalculateEnergy returns a tiny deterministic energy estimate used by the demo.
func CalculateEnergy(states StateVector, j, h float64) float64 {
	var sum float64
	for _, state := range states {
		sum += math.Abs(state)
	}
	return (sum * j) + (float64(len(states)) * h)
}
