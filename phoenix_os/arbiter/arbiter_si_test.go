package arbiter_test

import (
	"testing"
	"phoenix/arbiter"
	"phoenix/bus"
	"phoenix/monitor"
	"phoenix/warden"
)

func TestMultiplicativeSIModel(t *testing.T) {
	b := bus.NewBus()
	arb := arbiter.NewArbiter(b)

	// Case 1: Low SI (Normal user, high Z-score)
	// Z=4.0, TCS=0.9, SI=1.0 -> Benefit = 3.6
	// ClassLog (Class 1) -> Cost = 1.5. Benefit > Cost.
	// ClassLocalIsolate (Class 3) -> Cost = 4.5. Benefit < Cost. (Denied)
	scoreLowSI := monitor.DriftScore{
		ZScore:          4.0,
		ImportanceScore: 1.0,
	}
	
	state1, _, auth1 := arb.Evaluate(scoreLowSI, 0.9)
	if state1 != warden.StateSuspicious || !auth1 {
		t.Errorf("Expected SUSPICIOUS/Authorized for low SI, got %s/%v", state1, auth1)
	}

	// Case 2: High SI (Root user/System process, high Z-score)
	// Z=4.0, TCS=0.9, SI=2.0 -> Benefit = 7.2
	scoreHighSI := monitor.DriftScore{
		ZScore:          4.0,
		ImportanceScore: 2.0,
	}

	_, _, _ = arb.Evaluate(scoreHighSI, 0.9)
	
	// Case 3: Strategic Denial due to low SI
	// Z=5.1 (Class 3), TCS=0.86, SI=1.0 -> Benefit = 4.386. Cost = 4.5.
	// Benefit < Cost. Denied.
	scoreZ51LowSI := monitor.DriftScore{
		ZScore:          5.1,
		ImportanceScore: 1.0,
	}
	state3, _, auth3 := arb.Evaluate(scoreZ51LowSI, 0.86)
	if auth3 {
		t.Errorf("Expected Strategic Denial for Z=5.1/SI=1.0, but was authorized for %s", state3)
	}

	// Case 4: Strategic Authorization due to high SI
	// Z=5.1, TCS=0.86, SI=1.5 -> Benefit = 6.579. Cost = 4.5.
	// Benefit > Cost. Authorized.
	scoreZ51HighSI := monitor.DriftScore{
		ZScore:          5.1,
		ImportanceScore: 1.5,
	}
	state4, _, auth4 := arb.Evaluate(scoreZ51HighSI, 0.86)
	if state4 != warden.StateContained || !auth4 {
		t.Errorf("Expected CONTAINED/Authorized for Z=5.1/SI=1.5, got %s/%v", state4, auth4)
	}
}
