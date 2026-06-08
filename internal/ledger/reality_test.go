/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"testing"
)

func TestCalculateConfidence(t *testing.T) {
	weights := map[string]phxmath.FixedPoint{
		"sensor1": phxmath.NewFixedPoint(1),
		"sensor2": phxmath.NewFixedPoint(2),
		"sensor3": phxmath.NewFixedPoint(1),
	}

	tests := []struct {
		name         string
		observations map[string]Observation
		expected     ConfidenceScore
	}{
		{
			name: "Perfect Consensus",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("true")},
				"sensor2": {Value: []byte("true")},
				"sensor3": {Value: []byte("true")},
			},
			expected: ConfidenceScore(phxmath.NewFixedPoint(1)),
		},
		{
			name: "Majority Agreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("true")},
				"sensor2": {Value: []byte("true")},
				"sensor3": {Value: []byte("false")},
			},
			expected: ConfidenceScore(phxmath.NewFixedPointRaw(750000)), // 0.75
		},
		{
			name: "Heavily Weighted Disagreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("true")},
				"sensor2": {Value: []byte("false")},
				"sensor3": {Value: []byte("true")},
			},
			expected: ConfidenceScore(phxmath.NewFixedPointRaw(500000)), // 0.5
		},
		{
			name: "Complete Disagreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("A")},
				"sensor2": {Value: []byte("B")},
				"sensor3": {Value: []byte("C")},
			},
			expected: ConfidenceScore(phxmath.NewFixedPointRaw(500000)), // 0.5
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mse := &MultiSensorEvidence{Observations: tt.observations}
			score := mse.CalculateConfidence(weights)
			if phxmath.FixedPoint(score).V != phxmath.FixedPoint(tt.expected).V {
				t.Errorf("Expected confidence %v, got %v", tt.expected, score)
			}
		})
	}
}

func TestIsValid(t *testing.T) {
	mse := &MultiSensorEvidence{
		Observations: map[string]Observation{
			"s1": {Value: []byte("v")},
			"s2": {Value: []byte("v")},
		},
		ConfidenceScore: ConfidenceScore(phxmath.NewFixedPointRaw(800000)), // 0.8
	}

	if !mse.IsValid(2, ConfidenceScore(phxmath.NewFixedPointRaw(700000))) {
		t.Error("Expected true for 2 sensors and 0.8 confidence")
	}

	if mse.IsValid(3, ConfidenceScore(phxmath.NewFixedPointRaw(700000))) {
		t.Error("Expected false for 3 sensors")
	}

	if mse.IsValid(2, ConfidenceScore(phxmath.NewFixedPointRaw(900000))) {
		t.Error("Expected false for 0.9 confidence")
	}
}
