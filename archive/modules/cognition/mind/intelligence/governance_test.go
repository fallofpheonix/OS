package intelligence

import (
	"testing"
	"time"
)

func TestGovernanceMonitor_MTTE(t *testing.T) {
	gm := NewGovernanceMonitor()

	gm.RecordDecisionStart("dec-1")
	time.Sleep(10 * time.Millisecond)
	gm.RecordExplanationComplete("dec-1")

	mtte := gm.CalculateMTTE()
	if mtte < 10*time.Millisecond {
		t.Errorf("expected MTTE >= 10ms, got %v", mtte)
	}
}

func TestGovernanceMonitor_DeterminismLeak(t *testing.T) {
	gm := NewGovernanceMonitor()

	// Test no leak
	gm.AuditDeterminism("E1", 0.5, 0.5)
	if len(gm.Leaks) != 0 {
		t.Error("leak recorded incorrectly for matching values")
	}

	// Test leak
	gm.AuditDeterminism("E2", 0.9, 0.1)
	if len(gm.Leaks) != 1 {
		t.Fatal("failed to record determinism leak")
	}

	leak := gm.Leaks[0]
	if leak.Drift != 0.8 {
		t.Errorf("expected drift 0.8, got %f", leak.Drift)
	}
}

func TestGovernanceMonitor_SensorReputation(t *testing.T) {
	gm := NewGovernanceMonitor()

	sensor := "s1"

	// Initial reputation
	rep := gm.GetReputation(sensor)
	if rep != 1.0 {
		t.Errorf("expected initial reputation 1.0, got %f", rep)
	}

	// Penalty
	gm.RecordSensorClaim(sensor, false)
	rep = gm.GetReputation(sensor)
	if rep >= 1.0 {
		t.Errorf("reputation failed to drop after penalty: %f", rep)
	}

	// Recovery
	prevRep := rep
	gm.RecordSensorClaim(sensor, true)
	rep = gm.GetReputation(sensor)
	if rep <= prevRep {
		t.Errorf("reputation failed to recover after accurate claim: %f", rep)
	}
}
