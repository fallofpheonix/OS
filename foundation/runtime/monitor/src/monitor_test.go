/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package main

import (
	"math"
	"testing"
)

func TestEntropy(t *testing.T) {
	data := make([]byte, 256)
	for i := 0; i < 256; i++ {
		data[i] = byte(i)
	}
	res := CalculateEntropy(data)
	if res.Entropy < 7.9 {
		t.Errorf("Expected max entropy, got %f", res.Entropy)
	}
}

func TestKalman(t *testing.T) {
	f := NewKalmanFilter(0.1, 0.1, 1.0, 0.0)
	v := f.Filter(10.0)
	if v == 0 || v == 10.0 {
		t.Errorf("Filter did not smooth signal: %f", v)
	}
}

func TestImportanceScore(t *testing.T) {
	score := CalculateImportanceScore(1.0, 1.0, 1.0, 1.0, 1.0)
	if math.Abs(score-1.0) > 1e-9 {
		t.Errorf("Expected score 1.0 for normalized inputs, got %f", score)
	}

	criticalScore := CalculateImportanceScore(0.1, 1.0, 0.1, 0.1, 0.1)
	lowImportanceScore := CalculateImportanceScore(1.0, 0.1, 1.0, 1.0, 1.0)

	if criticalScore <= lowImportanceScore {
		t.Errorf("Criticality should outweigh other factors: %f vs %f", criticalScore, lowImportanceScore)
	}
}
