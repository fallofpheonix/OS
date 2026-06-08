/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package arbiter

import (
	"testing"
)

func TestPolicyValidator_Evaluate(t *testing.T) {
	v := NewPolicyValidator()

	tests := []struct {
		name          string
		drift         interface{}
		tcs           float64
		expectedState string
		expectedAct   string
	}{
		{"Safe Path", 0.1, 0.9, "SAFE", "LOG"},
		{"Low Confidence", 0.1, 0.4, "WATCH", "LOG"},
		{"High Drift Low Confidence", 0.9, 0.4, "WATCH", "LOG"},
		{"High Drift High Confidence", 0.9, 0.8, "ESCALATE", "CONTAIN"},
		{"Moderate Drift", 0.5, 0.9, "ANALYZE", "LOG"},
		{"Integer Score", 1, 0.8, "ESCALATE", "CONTAIN"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			state, act, auth := v.Evaluate(tt.drift, tt.tcs)
			if !auth {
				t.Error("Expected authorized=true")
			}
			if state != tt.expectedState {
				t.Errorf("Expected state %s, got %s", tt.expectedState, state)
			}
			if act != tt.expectedAct {
				t.Errorf("Expected action %s, got %s", tt.expectedAct, act)
			}
		})
	}
}
