package arbiter

import (
	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-logic/monitor"
	"testing"
)

func TestEvaluatePolicy_CostInvariants(t *testing.T) {
	tests := []struct {
		name        string
		zScore      float64
		load        float64
		history     []string
		expectedAct warden.ActuationClass
	}{
		{
			name:        "Critical Threat Low Load",
			zScore:      15.0,
			load:        0.0,
			history:     []string{"h1", "h2", "h3"},
			expectedAct: warden.ClassLocalIsolate, // CC=200, RC=50, RR=1000*(1-0.6)=400, RP=0 => TC=650 <= AC=1000
		},
		{
			name:        "Critical Threat High Load",
			zScore:      15.0,
			load:        1.0,
			history:     []string{"h1", "h2", "h3"},
			expectedAct: warden.ClassLocalIsolate, // CC=200, RC=50, RR=2000*(1-0.7)=600, RP=0 => TC=850 <= AC=2000
		},
		{
			name:        "Medium Threat",
			zScore:      4.0,
			load:        0.0,
			expectedAct: warden.ClassObserve, // CC=1, RC=0.25, RR=100*(1-0.2)=80, RP=50 => TC=131.25 > AC=100
		},
		{
			name:        "Info Level",
			zScore:      0.5,
			load:        0.0,
			expectedAct: warden.ClassObserve,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := SystemContext{
				ThreatScore: monitor.DriftScore{
					ZScore:   tt.zScore,
					Severity: 1.0,
					PID:      1234,
				},
				SystemLoad:    tt.load,
				ReplayHistory: tt.history,
			}

			decision := EvaluatePolicy(ctx)
			if decision.Recommended != tt.expectedAct {
				t.Errorf("%s: expected action class %v, got %v (AttackCost: %.1f, ContainmentCost: %.1f)",
					tt.name, tt.expectedAct, decision.Recommended, decision.CostFactors.AttackCost, decision.CostFactors.ContainmentCost)
			}
		})
	}
}

func TestEvaluatePolicy_WithEvidence(t *testing.T) {
	ctx := SystemContext{
		ThreatScore: monitor.DriftScore{
			ZScore:   5.0,
			Severity: 1.0,
			PID:      555,
		},
		ReplayHistory: []string{"frame1", "frame2", "frame3"},
		SystemLoad:    0.2,
	}

	decision := EvaluatePolicy(ctx)
	// AC = 500 * 1.2 = 600
	// CC for LocalIsolate = 200 (<= 600)
	// CC for ClusterIsolate = 800 (> 600)
	// Expected: LocalIsolate

	if decision.Recommended != warden.ClassThrottle {
		t.Errorf("Expected Throttle, got %v", decision.Recommended)
	}

	expectedSubstr := "Evidence: 3 frames"
	if !contains(decision.Reasoning, expectedSubstr) {
		t.Errorf("Expected reasoning to contain '%s', got '%s'", expectedSubstr, decision.Reasoning)
	}
}

func contains(s, substr string) bool {
	return len(s) >= len(substr) && func() bool {
		for i := 0; i <= len(s)-len(substr); i++ {
			if s[i:i+len(substr)] == substr {
				return true
			}
		}
		return false
	}()
}
