package arbiter

import (
	"fmt"

	"phoenix/bus"
	"phoenix/monitor"
	"phoenix/warden"
)

// Arbiter implements L5.5 Strategic Policy using a Game-Theoretic approach
type Arbiter struct {
	bus    *bus.Bus
	Policy Policy
}

func NewArbiter(b *bus.Bus) *Arbiter {
	return &Arbiter{
		bus:    b,
		Policy: DefaultPolicy(),
	}
}

// Evaluate decides whether an action is authorized based on Payoff and TCS
func (a *Arbiter) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, warden.ActuationClass, bool) {
	// 1. Determine Target Action Class based on Drift (Z-Score)
	targetState := warden.StateNormal
	requiredClass := warden.ClassObserve

	if score.ZScore >= 5.0 {
		targetState = warden.StateContained
		requiredClass = warden.ClassLocalIsolate
	} else if score.ZScore >= 2.0 {
		targetState = warden.StateSuspicious
		requiredClass = warden.ClassLog
	}

	// 2. Payoff Calculation (Strategic Layer)
	// benefit = Drift (Z) * Confidence (TCS) * Importance (SI)
	benefit := score.ZScore * tcsScore * score.ImportanceScore
	cost := float64(requiredClass) * 1.5

	if benefit < cost && requiredClass > warden.ClassLog {
		fmt.Printf("[ARBITER] STRATEGIC DENIAL: Benefit %.2f (Z:%.2f, TCS:%.2f, SI:%.2f) < Cost %.2f for Class %d\n", 
			benefit, score.ZScore, tcsScore, score.ImportanceScore, cost, requiredClass)
		return warden.StateNormal, requiredClass, false
	}

	// 3. TCS Gating
	if tcsScore < a.Policy.Thresholds[requiredClass] {
		fmt.Printf("[ARBITER] TCS DENIED: Action class %d requires TCS %.2f (Current: %.2f)\n", requiredClass, a.Policy.Thresholds[requiredClass], tcsScore)
		return warden.StateNormal, requiredClass, false
	}

	// 4. Budget Check
	if budget, ok := a.Policy.Budgets[requiredClass]; ok {
		// RED TEAM MITIGATION: Critical Bypass
		// Extremely high Z-Score (existential threat) with high confidence 
		// should bypass remaining budget to prevent DoS-via-noise.
		isCritical := score.ZScore >= 50.0 && tcsScore >= 0.95
		
		if budget <= 0 && !isCritical {
			fmt.Printf("[ARBITER] BUDGET EXHAUSTED for Class %d\n", requiredClass)
			return warden.StateNormal, requiredClass, false
		}
		
		if budget > 0 {
			a.Policy.Budgets[requiredClass]--
		}
	}

	return targetState, requiredClass, true
}
