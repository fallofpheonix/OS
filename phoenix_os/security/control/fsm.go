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
	CurrentState    SecurityState
	LastAction      string
	ReactionTime    time.Duration
	StochasticMode  bool
	rng             *rand.Rand
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
func (c *Controller) UpdateState(score float64) (SecurityState, string) {
	var nextState SecurityState
	var action string

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
		action = "ISOLATE"
	case score >= 0.8:
		nextState = StateCritical
		action = "FREEZE"
	case score >= 0.6:
		nextState = StateSuspicious
		action = "LIMIT"
	case score >= 0.3:
		nextState = StateWatch
		action = "OBSERVE"
	default:
		nextState = StateSafe
		action = "NONE"
	}
	if nextState != c.CurrentState {
		c.CurrentState = nextState
		c.LastAction = action
		return nextState, action
	}

	return c.CurrentState, "NONE"
}

// GetStatus returns the current controller status.
func (c *Controller) GetStatus() string {
	return fmt.Sprintf("State: %s, Last Action: %s", c.CurrentState, c.LastAction)
}
