/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 *
 * FILE: fsm.go
 * PATH: assurance/security/control/fsm.go
 */

package control

import (
	"fmt"
	"time"
)

type SecurityState uint8

const (
	StateSafe SecurityState = iota
	StateWatch
	StateSuspicious
	StateCritical
	StateCompromised
)

func (s SecurityState) String() string {
	switch s {
	case StateSafe:
		return "SAFE"
	case StateWatch:
		return "WATCH"
	case StateSuspicious:
		return "SUSPICIOUS"
	case StateCritical:
		return "CRITICAL"
	case StateCompromised:
		return "COMPROMISED"
	default:
		return fmt.Sprintf("UNKNOWN(%d)", s)
	}
}

// Controller implements a tactical security enforcement state machine.
type Controller struct {
	CurrentState SecurityState
	ReactionTime time.Duration
}

// NewController initializes a new finite-state controller.
func NewController() *Controller {
	return &Controller{
		CurrentState: StateSafe,
		ReactionTime: 50 * time.Millisecond,
	}
}

// UpdateState transitions the controller based on an importance score.
func (c *Controller) UpdateState(score float64) (state SecurityState, action string) {
	// WHY: Stochastic behavior removed. math/rand seeded from wall clock time
	// breaks replay determinism (BLOCKER-006). The system cannot produce two
	// identical audit trails if security decisions depend on time-seeded randomness.
	// RESOLUTION PATH: Replace with ledger-derived entropy once the ledger is
	// fully wired into the security path (T3.4 prerequisite).

	var nextState SecurityState
	var nextAction string

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

	c.CurrentState = nextState
	return nextState, nextAction
}
