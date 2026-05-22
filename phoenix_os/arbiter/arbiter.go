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

// Payoff represents the calculated benefit of an action vs its cost
type Payoff struct {
	Benefit float64
	Cost    float64
}

// Arbiter implements L5.5 Strategic Policy using a Game-Theoretic approach
type Arbiter struct {
	bus           *bus.Bus
	tcsThresholds map[ActuationClass]float64
	budgets       map[ActuationClass]int // Remaning actions allowed in current window
}

func NewArbiter(b *bus.Bus) *Arbiter {
	return &Arbiter{
		bus: b,
		tcsThresholds: map[ActuationClass]float64{
			ClassObserve:         0.0,
			ClassLog:             0.0,
			ClassThrottle:        0.60,
			ClassLocalIsolate:    0.85,
			ClassClusterIsolate:  0.95,
			ClassKernelEmergency: 0.99,
		},
		budgets: map[ActuationClass]int{
			ClassThrottle:        100,
			ClassLocalIsolate:    50,
			ClassClusterIsolate:  10,
			ClassKernelEmergency: 1,
		},
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
	benefit := score.ZScore * tcsScore // Scale benefit by telemetry confidence
	cost := float64(requiredClass) * 1.5

	if benefit < cost && requiredClass > ClassLog {
		fmt.Printf("[ARBITER] STRATEGIC DENIAL: Benefit %.2f < Cost %.2f for Class %d\n", benefit, cost, requiredClass)
		return warden.StateNormal, false
	}

	// 3. TCS Gating
	if tcsScore < a.tcsThresholds[requiredClass] {
		fmt.Printf("[ARBITER] TCS DENIED: Action class %d requires TCS %.2f (Current: %.2f)\n", requiredClass, a.tcsThresholds[requiredClass], tcsScore)
		return warden.StateNormal, false
	}

	// 4. Budget Check
	if budget, ok := a.budgets[requiredClass]; ok {
		if budget <= 0 {
			fmt.Printf("[ARBITER] BUDGET EXHAUSTED for Class %d\n", requiredClass)
			return warden.StateNormal, false
		}
		a.budgets[requiredClass]--
	}

	return targetState, true
}
