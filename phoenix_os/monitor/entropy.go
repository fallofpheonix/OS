package entropy_engine

import (
	"math"
)

// Result represents the calculation output
type Result struct {
	Entropy      float64 `json:"entropy"`
	KLDivergence float64 `json:"kl_divergence"`
	IsAnomaly    bool    `json:"is_anomaly"`
}

// Calculate Shannon Entropy and KL Divergence
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

// CalculateEntropy is a convenience function that returns only the Shannon Entropy
func CalculateEntropy(data []byte) float64 {
	res := Calculate(data, nil)
	return res.Entropy
}
