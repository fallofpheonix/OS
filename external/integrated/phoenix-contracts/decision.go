package contracts

import (
	"github.com/fallofpheonix/phoenix-control/state"
	"time"
)

// Decision defines the contract for actuation authorization.
type Decision interface {
	GetTargetPID() int
	GetAuthorizedAction() int // ActuationClass mapping
	GetAuthorizedState() state.RuntimeState
	GetConfidence() float64
	GetTimestamp() time.Time
	GetReasoning() string
}
