/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — SHANNON ENTROPY CALCULATION
//
// ShannonEntropy calculates the information entropy of a byte slice.
// High entropy (close to 8.0) indicates random/encrypted data.
// Low entropy (close to 0.0) indicates repetitive/predictable data.
//
// WORKFLOW:
//   Monitor.CalculateEntropy(data) → ShannonEntropy(data) → entropy value
//   Detector.Analyze(event) → check entropy_score in payload → threat detection
//   TCS.Evaluate() → jitter normalization → confidence score
//
// ALGORITHM: Standard Shannon entropy:
//   H = -Σ p(x) * log2(p(x)) for all byte values x
//   Where p(x) = count(x) / total_bytes
//
// COMPLEXITY: O(N) where N = length of data slice.
// Space: O(256) for frequency array (fixed size).
//
// USAGE: High entropy (> 7.5) is a strong indicator of encryption/ransomware.
// The Detector uses this as Rule 1 for threat detection.
// =========================================================================
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
	for i := range freq {
		c := freq[i]
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
