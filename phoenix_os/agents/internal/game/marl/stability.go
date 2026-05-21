package marl

import (
	"sync"
	"time"
)

// StabilityController enforces action limits and cooling periods for MARL agents.
type StabilityController struct {
	mu            sync.RWMutex
	ActionDebt    float64
	LastAction    time.Time
	Cooldown      time.Duration
	MaxContainment float64
}

// NewStabilityController initializes a controller with defined limits.
func NewStabilityController(cooldown time.Duration, maxContainment float64) *StabilityController {
	return &StabilityController{
		Cooldown:       cooldown,
		MaxContainment: maxContainment,
	}
}

// CanAct checks if an agent is allowed to perform a containment action.
func (sc *StabilityController) CanAct(cost float64) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	// Check Cooldown
	if time.Since(sc.LastAction) < sc.Cooldown {
		return false
	}

	// Check Containment Rate (Action Debt)
	if sc.ActionDebt+cost > sc.MaxContainment {
		return false
	}

	return true
}

// RecordAction updates the stability state after a successful action.
func (sc *StabilityController) RecordAction(cost float64) {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.ActionDebt += cost
	sc.LastAction = time.Now()
}

// ResetDebt clears accumulated action debt for the next cycle.
func (sc *StabilityController) ResetDebt() {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.ActionDebt = 0
}
