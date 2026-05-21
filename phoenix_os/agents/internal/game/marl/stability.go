package marl

import (
	"sync"
	"time"
)

// StabilityController enforces action limits and cooling periods for MARL agents.
type StabilityController struct {
	mu             sync.RWMutex
	actionDebt     float64
	lastAction     time.Time
	cooldown       time.Duration
	maxContainment float64
	decayRate      float64 // Debt reduction per second
	lastDecay      time.Time
}

// NewStabilityController initializes a controller with defined limits and decay rate.
func NewStabilityController(cooldown time.Duration, maxContainment float64, decayRate float64) *StabilityController {
	now := time.Now()
	return &StabilityController{
		cooldown:       cooldown,
		maxContainment: maxContainment,
		decayRate:      decayRate,
		lastDecay:      now,
	}
}

// applyDecay reduces the action debt based on elapsed time. Must be called with lock held.
func (sc *StabilityController) applyDecay(now time.Time) {
	elapsed := now.Sub(sc.lastDecay).Seconds()
	if elapsed > 0 {
		reduction := elapsed * sc.decayRate
		sc.actionDebt -= reduction
		if sc.actionDebt < 0 {
			sc.actionDebt = 0
		}
		sc.lastDecay = now
	}
}

// TryRecordAction atomically checks if an action is allowed and records it if so.
// It also applies time-based decay to the action debt.
func (sc *StabilityController) TryRecordAction(cost float64, now time.Time) bool {
	sc.mu.Lock()
	defer sc.mu.Unlock()

	sc.applyDecay(now)

	// Check Cooldown
	if now.Sub(sc.lastAction) < sc.cooldown {
		return false
	}

	// Check Containment Rate (Action Debt)
	if sc.actionDebt+cost > sc.maxContainment {
		return false
	}

	sc.actionDebt += cost
	sc.lastAction = now
	return true
}

// GetActionDebt returns the current action debt (mainly for testing/monitoring).
func (sc *StabilityController) GetActionDebt(now time.Time) float64 {
	sc.mu.Lock()
	defer sc.mu.Unlock()
	sc.applyDecay(now)
	return sc.actionDebt
}

