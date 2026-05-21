package main

import (
	"fmt"
	"math"
)

// StateVector represents nodes as +1 (benign) or -1 (compromised)
type StateVector []int8

// CalculateSDI computes the Shannon entropy of the state distribution
func CalculateSDI(states StateVector) float64 {
	if len(states) == 0 {
		return 0
	}

	counts := make(map[int8]int)
	for _, s := range states {
		counts[s]++
	}

	var sdi float64
	n := float64(len(states))
	for _, count := range counts {
		p := float64(count) / n
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}
	return sdi
}

// CalculateEnergy computes the Hamiltonian for the system
// H = -J * sum(sigma_i * sigma_j) - h * sum(sigma_i)
// For simplicity, we assume a fully connected graph for this primitive
func CalculateEnergy(states StateVector, J float64, h float64) float64 {
	var interaction float64
	var external float64

	for i := 0; i < len(states); i++ {
		external += float64(states[i])
		for j := i + 1; j < len(states); j++ {
			interaction += float64(states[i]) * float64(states[j])
		}
	}

	return -J*interaction - h*external
}

func main() {
	states := StateVector{1, 1, -1, 1, -1}
	fmt.Printf("SDI: %f\n", CalculateSDI(states))
	fmt.Printf("Energy: %f\n", CalculateEnergy(states, 1.0, 0.5))
}
