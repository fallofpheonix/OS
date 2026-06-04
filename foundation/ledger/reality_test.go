/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	"testing"
)

func TestCalculateConfidence(t *testing.T) {
	weights := map[string]float64{
		"sensor1": 1.0,
		"sensor2": 2.0,
		"sensor3": 1.0,
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
			expected: 1.0,
		},
		{
			name: "Majority Agreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("true")},
				"sensor2": {Value: []byte("true")},
				"sensor3": {Value: []byte("false")},
			},
			expected: 0.75, // (1+2)/4
		},
		{
			name: "Heavily Weighted Disagreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("true")},
				"sensor2": {Value: []byte("false")},
				"sensor3": {Value: []byte("true")},
			},
			expected: 0.5, // max(1+1, 2)/4 = 2/4
		},
		{
			name: "Complete Disagreement",
			observations: map[string]Observation{
				"sensor1": {Value: []byte("A")},
				"sensor2": {Value: []byte("B")},
				"sensor3": {Value: []byte("C")},
			},
			expected: 0.5, // max(1, 2, 1)/4 = 2/4
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mse := &MultiSensorEvidence{Observations: tt.observations}
			score := mse.CalculateConfidence(weights)
			if score != tt.expected {
				t.Errorf("Expected confidence %f, got %f", tt.expected, score)
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
		ConfidenceScore: 0.8,
	}

	if !mse.IsValid(2, 0.7) {
		t.Error("Expected true for 2 sensors and 0.8 confidence")
	}

	if mse.IsValid(3, 0.7) {
		t.Error("Expected false for 3 sensors")
	}

	if mse.IsValid(2, 0.9) {
		t.Error("Expected false for 0.9 confidence")
	}
}
