// Package monitor implements sensory analysis for PhoenixOS telemetry.
// Domain Logic: Provides statistical utilities for anomaly detection.
// Responsibility: Calculates Shannon entropy and KL divergence to detect anomalous byte distributions in telemetry payloads.
package monitor

import (
	"math"
)

// Result represents the calculation output for entropy and divergence analysis.
// Concurrency: Instances are immutable and thread-safe.
// State Management: Encapsulates statistical metrics and anomaly status.
type Result struct {
	Entropy      float64 `json:"entropy"`
	KLDivergence float64 `json:"kl_divergence"`
	IsAnomaly    bool    `json:"is_anomaly"`
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Calculate computes Shannon Entropy and KL Divergence for the provided data.
// I/O: None.
// Side Effects: None.
// Complexity: O(N) where N is the length of the data.
func Calculate(data []byte, baseline []float64) Result {
	if len(data) == 0 {
		return Result{}
	}

	counts := make([]int, 256)
	for _, b := range data {
		counts[b]++
	}

	var entropy float64
	var klDiv float64
	n := float64(len(data))

	for i := 0; i < 256; i++ {
		p := float64(counts[i]) / n
		if p > 0 {
			entropy -= p * math.Log2(p)

			if baseline != nil && i < len(baseline) {
				q := baseline[i]
				if q > 0 {
					klDiv += p * math.Log2(p/q)
				} else {
					klDiv += p * 8.0
				}
			}
		}
	}

	isAnomaly := entropy > 7.5 || klDiv > 4.0

	return Result{
		Entropy:      entropy,
		KLDivergence: klDiv,
		IsAnomaly:    isAnomaly,
	}
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// CalculateEntropy is a convenience function that returns only the Shannon Entropy.
// I/O: None.
// Side Effects: None.
// Complexity: O(N) where N is the length of the data.
func CalculateEntropy(data []byte) float64 {
	res := Calculate(data, nil)
	return res.Entropy
}
