package entropy_engine

import (
	"math"
)

// ShannonEntropy calculates the Shannon entropy in bits for a byte slice.
func ShannonEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var freq [256]int
	for _, b := range data {
		freq[b]++
	}
	var entropy float64
	n := float64(len(data))
	for _, c := range freq {
		if c == 0 {
			continue
		}
		p := float64(c) / n
		entropy -= p * math.Log2(p)
	}
	return entropy
}

// EntropyResult is the minimal result shape used by the integrated model.
type EntropyResult struct {
	Entropy   float64
	IsAnomaly bool
}

// CalculateEntropy wraps ShannonEntropy and applies a simple anomaly threshold.
func CalculateEntropy(data []byte, _ interface{}) EntropyResult {
	entropy := ShannonEntropy(data)
	return EntropyResult{
		Entropy:   entropy,
		IsAnomaly: entropy > 7.5,
	}
}
