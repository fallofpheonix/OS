package arbiter

import (
	"reflect"
	"testing"

	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-logic/monitor"
)

// TestDecisionReplayDeterministic (Track C3) verifies that given the same
// SystemContext, the Arbiter consistently produces the exact same Decision.
func TestDecisionReplayDeterministic(t *testing.T) {
	ctx := SystemContext{
		ThreatScore: monitor.DriftScore{
			PID:          1234,
			DriftScore:   750.0,
			Severity:     0.75,
			ZScore:       12.0, // Critical
			WallTimeUnix: 1622000000,
		},
		ReplayHistory: []string{"hash1", "hash2", "hash3"},
		RiskLevel:     0.3,
		SystemLoad:    0.5,
	}

	decision1 := EvaluatePolicy(ctx)

	// Re-run the evaluation multiple times
	for i := 0; i < 100; i++ {
		decision2 := EvaluatePolicy(ctx)

		if !reflect.DeepEqual(decision1, decision2) {
			t.Errorf("Non-deterministic decision at iteration %d:\nFirst: %+v\nSecond: %+v", i, decision1, decision2)
		}
	}
}

// TestCounterfactualOptimization (Track C4) verifies that the Arbiter
// chooses the action that minimizes TotalCost.
func TestCounterfactualOptimization(t *testing.T) {
	// Scenario: High threat (AC=800) but no evidence (Confidence Penalty).
	// The Arbiter should avoid ClassClusterIsolate (CC=800) because RC/RR will push it over AC.
	ctx := SystemContext{
		ThreatScore: monitor.DriftScore{
			PID:        999,
			DriftScore: 800.0,
			Severity:   0.8,
			ZScore:     8.0, // High
		},
		ReplayHistory: []string{}, // NO EVIDENCE
		SystemLoad:    0.1,
	}

	decision := EvaluatePolicy(ctx)

	// Without evidence, it should pick a lighter action like Throttle or LocalIsolate
	// rather than a heavy one, even if the threat is high.
	if decision.Recommended == warden.ClassClusterIsolate || decision.Recommended == warden.ClassKernelEmergency {
		t.Errorf("Arbiter overreacted without evidence: Recommended %v", decision.Recommended)
	}

	if decision.TotalCost >= CalculateAttackCost(ctx) {
		t.Errorf("Arbiter recommendation is not profitable: TotalCost %.1f >= AttackCost %.1f",
			decision.TotalCost, CalculateAttackCost(ctx))
	}
}
