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
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"time"
)

// SecurityState represents the Phoenix Matrix security states as a deterministic enum.
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
		return "UNKNOWN"
	}
}

// ThreatScore is a deterministic fixed-point value used for state evaluation.
type ThreatScore = phxmath.FixedPoint

var (
	ThresholdCompromised = phxmath.NewFixedPointRaw(950000) // 0.95
	ThresholdCritical    = phxmath.NewFixedPointRaw(800000) // 0.80
	ThresholdSuspicious  = phxmath.NewFixedPointRaw(600000) // 0.60
	ThresholdWatch       = phxmath.NewFixedPointRaw(300000) // 0.30
)

// Controller manages security state transitions and actuation.
type Controller struct {
	CurrentState   SecurityState
	LastAction     string
	ReactionTime   time.Duration
	StochasticMode bool
}

// NewController initializes a new finite-state controller.
func NewController() *Controller {
	return &Controller{
		CurrentState:   StateSafe,
		ReactionTime:   50 * time.Millisecond,
		StochasticMode: false,
	}
}

// UpdateState transitions the controller based on a deterministic threat score.
func (c *Controller) UpdateState(score ThreatScore) (state SecurityState, action string) {
	var nextState SecurityState
	var nextAction string

	val := score.V
	switch {
	case val >= ThresholdCompromised.V:
		nextState = StateCompromised
		nextAction = "ISOLATE"
	case val >= ThresholdCritical.V:
		nextState = StateCritical
		nextAction = "FREEZE"
	case val >= ThresholdSuspicious.V:
		nextState = StateSuspicious
		nextAction = "LIMIT"
	case val >= ThresholdWatch.V:
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
