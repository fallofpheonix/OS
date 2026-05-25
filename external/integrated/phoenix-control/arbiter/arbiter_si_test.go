package arbiter_test

import (
	"github.com/fallofpheonix/phoenix-control/arbiter"
	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-logic/monitor"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"testing"
)

func TestAdaptiveSIModel(t *testing.T) {
	b := bus.NewBus()
	arb := arbiter.NewArbiter(b)

	// Case 1: High Criticality Node (Root), Rare Threat mode (Adaptive triggered by high Z)
	// UID=0 (CN=1.0), GS=0.8, FA=0.01 (Rare), Z=5.1
	// M(FA) = exp(-2.0 * 0.01) = 0.98
	// SI = (1.0 * 0.8) * 0.98 = 0.784
	// Benefit = 5.1 * 0.9 * 0.784 = 3.598
	// Cost (ClassLocalIsolate) = 3 * 1.5 = 4.5
	// Expected: Strategic Denial (Benefit < Cost)
	scoreRare := monitor.DriftScore{
		UID:       0,
		Severity:  0.8,
		Frequency: 0.01,
		ZScore:    5.1,
	}

	state1, class1, auth1 := arb.Evaluate(scoreRare, 0.9)
	if auth1 {
		t.Errorf("Expected Strategic Denial for high criticality rare threat, got %s/%d/Authorized", state1, class1)
	}

	// Case 2: Frequent Threat (Brute Force Simulation)
	// UID=100 (CN=0.8), GS=0.5, FA=0.8 (Frequent), Z=2.1
	// FrequentThreat Mode (Z < 3)
	// M(FA) = exp(2.0 * 0.8) = 4.95 -> SI capped at 1.0
	// Benefit = 2.1 * 0.9 * 1.0 = 1.89
	// Cost (ClassLog) = 1 * 1.5 = 1.5
	// Expected: Authorized
	scoreFrequent := monitor.DriftScore{
		UID:       100,
		Severity:  0.5,
		Frequency: 0.8,
		ZScore:    2.1,
	}

	state2, class2, auth2 := arb.Evaluate(scoreFrequent, 0.9)
	if state2 != warden.StateSuspicious || !auth2 || class2 != warden.ClassLog {
		t.Errorf("Expected SUSPICIOUS/Log/Authorized for frequent threat, got %s/%d/%v", state2, class2, auth2)
	}

	// Case 3: Stealth Attack (Low Frequency, High Z) with higher CN
	// UID=0 (CN=1.0), GS=1.0, FA=0.001, Z=10.0
	// M(FA) = exp(-2.0 * 0.001) = 0.998
	// SI = 1.0 * 1.0 * 0.998 = 0.998
	// Benefit = 10.0 * 0.9 * 0.998 = 8.982
	// Cost = 4.5
	// Expected: Authorized/CONTAINED
	scoreStealth := monitor.DriftScore{
		UID:       0,
		Severity:  1.0,
		Frequency: 0.001,
		ZScore:    10.0,
	}
	state3, class3, auth3 := arb.Evaluate(scoreStealth, 0.9)
	if state3 != warden.StateContained || !auth3 || class3 != warden.ClassLocalIsolate {
		t.Errorf("Expected CONTAINED/Isolate/Authorized for stealth attack, got %s/%d/%v", state3, class3, auth3)
	}
}
