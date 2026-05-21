package physics

import (
	"phoenix/security/physics/disorder"
)

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
