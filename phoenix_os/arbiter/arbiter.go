package arbiter

import (
	"fmt"

	"phoenix/bus"
	"phoenix/monitor"
	"phoenix/warden"
)

// ActuationClass represents the risk tier of an action
type ActuationClass int

const (
	ClassObserve         ActuationClass = 0
	ClassLog             ActuationClass = 1
	ClassThrottle        ActuationClass = 2
	ClassLocalIsolate    ActuationClass = 3
	ClassClusterIsolate  ActuationClass = 4
	ClassKernelEmergency ActuationClass = 5
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
func (a *Arbiter) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, bool) {
	// 1. Determine Target Action Class based on Drift (Z-Score)
	targetState := warden.StateNormal
	requiredClass := ClassObserve

	if score.ZScore >= 5.0 {
		targetState = warden.StateContained
		requiredClass = ClassLocalIsolate
	} else if score.ZScore >= 2.0 {
		targetState = warden.StateSuspicious
		requiredClass = ClassLog
	}

	// 2. Payoff Calculation (Strategic Layer)
	benefit := score.ZScore * tcsScore
	cost := float64(requiredClass) * 1.5

	if benefit < cost && requiredClass > ClassLog {
		fmt.Printf("[ARBITER] STRATEGIC DENIAL: Benefit %.2f < Cost %.2f for Class %d\n", benefit, cost, requiredClass)
		return warden.StateNormal, false
	}

	// 3. TCS Gating
	if tcsScore < a.Policy.Thresholds[requiredClass] {
		fmt.Printf("[ARBITER] TCS DENIED: Action class %d requires TCS %.2f (Current: %.2f)\n", requiredClass, a.Policy.Thresholds[requiredClass], tcsScore)
		return warden.StateNormal, false
	}

	// 4. Budget Check
	if budget, ok := a.Policy.Budgets[requiredClass]; ok {
		if budget <= 0 {
			fmt.Printf("[ARBITER] BUDGET EXHAUSTED for Class %d\n", requiredClass)
			return warden.StateNormal, false
		}
		a.Policy.Budgets[requiredClass]--
	}

	return targetState, true
}
