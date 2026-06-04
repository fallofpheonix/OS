/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package control

import (
	"fmt"
	"math/rand"
	"time"
)

// SecurityState represents the Phoenix Matrix security states.
type SecurityState string

const (
	StateSafe        SecurityState = "SAFE"
	StateWatch       SecurityState = "WATCH"
	StateSuspicious  SecurityState = "SUSPICIOUS"
	StateCritical    SecurityState = "CRITICAL"
	StateCompromised SecurityState = "COMPROMISED"
)

// Controller manages security state transitions and actuation.
type Controller struct {
	CurrentState   SecurityState
	LastAction     string
	ReactionTime   time.Duration
	StochasticMode bool
	rng            *rand.Rand
}

// NewController initializes a new finite-state controller.
func NewController() *Controller {
	return &Controller{
		CurrentState:   StateSafe,
		ReactionTime:   50 * time.Millisecond,
		StochasticMode: true,
		rng:            rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// UpdateState transitions the controller based on an importance score.
func (c *Controller) UpdateState(score float64) (state SecurityState, action string) {
	var nextState SecurityState
	var nextAction string

	// 1. Stochastic Nuance for Uncertainty Zone (0.4 - 0.8)
	if c.StochasticMode && score >= 0.4 && score < 0.8 {
		// Use probability to decide if we should escalate early
		if c.rng.Float64() < score {
			score += 0.1 // Probabilistic escalation
		}
	}

	switch {
	case score >= 0.95:
		nextState = StateCompromised
		nextAction = "ISOLATE"
	case score >= 0.8:
		nextState = StateCritical
		nextAction = "FREEZE"
	case score >= 0.6:
		nextState = StateSuspicious
		nextAction = "LIMIT"
	case score >= 0.3:
		nextState = StateWatch
		nextAction = "OBSERVE"
	default:
		nextState = StateSafe
		nextAction = "NONE"
	}
	if nextState != c.CurrentState {
		c.CurrentState = nextState
		c.LastAction = nextAction
		return nextState, nextAction
	}

	return c.CurrentState, "NONE"
}

// GetStatus returns the current controller status.
func (c *Controller) GetStatus() string {
	return fmt.Sprintf("State: %s, Last Action: %s", c.CurrentState, c.LastAction)
}
