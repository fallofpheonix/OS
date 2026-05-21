package control

import (
	"fmt"
	"time"
)

// SecurityState represents the Phoenix Matrix security states.
type SecurityState string

const (
	StateSafe       SecurityState = "SAFE"
	StateWatch      SecurityState = "WATCH"
	StateSuspicious SecurityState = "SUSPICIOUS"
	StateCritical   SecurityState = "CRITICAL"
	StateCompromised SecurityState = "COMPROMISED"
)

// Controller manages security state transitions and actuation.
type Controller struct {
	CurrentState SecurityState
	LastAction   string
	ReactionTime time.Duration
}

// NewController initializes a new finite-state controller.
func NewController() *Controller {
	return &Controller{
		CurrentState: StateSafe,
	}
}

// UpdateState transitions the controller based on an importance score.
func (c *Controller) UpdateState(score float64) (SecurityState, string) {
	start := time.Now()
	var nextState SecurityState
	var action string

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
		c.ReactionTime = time.Since(start)
		return nextState, action
	}

	return c.CurrentState, "NONE"
}

// GetStatus returns the current controller status.
func (c *Controller) GetStatus() string {
	return fmt.Sprintf("State: %s, Last Action: %s, Reaction: %v", c.CurrentState, c.LastAction, c.ReactionTime)
}
