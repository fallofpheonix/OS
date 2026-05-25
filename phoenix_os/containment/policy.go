package containment

import (
	"fmt"
	"time"
)

// Transition enforces the observe-throttle-isolate-recover lifecycle.
func (e *IsolationEngine) Transition(next IsolationState, evidenceID string, decisionID string) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	if !isValidContainmentTransition(e.CurrentState, next) {
		return fmt.Errorf("illegal containment transition from %s to %s", e.CurrentState, next)
	}

	record := IsolationRecord{
		Timestamp:  time.Now(),
		Previous:   e.CurrentState,
		Current:    next,
		EvidenceID: evidenceID,
		DecisionID: decisionID,
	}

	e.CurrentState = next
	e.History = append(e.History, record)
	return nil
}

func isValidContainmentTransition(current, next IsolationState) bool {
	switch current {
	case StateObserve:
		return next == StateWatch || next == StateThrottle
	case StateWatch:
		return next == StateThrottle || next == StateIsolate
	case StateThrottle:
		return next == StateIsolate
	case StateIsolate:
		return next == StateRecover
	case StateRecover:
		return next == StateObserve
	default:
		return false
	}
}
