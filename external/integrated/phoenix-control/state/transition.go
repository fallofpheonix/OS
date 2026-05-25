package state

import (
	"fmt"
	"sync"
)

// IllegalTransitionDetector tracks attempted invalid state changes.
type IllegalTransitionDetector struct {
	mu             sync.Mutex
	attempts       []TransitionAttempt
	violationCount int
}

type TransitionAttempt struct {
	From RuntimeState
	To   RuntimeState
	Time int64 // Logical clock tick
}

func NewIllegalTransitionDetector() *IllegalTransitionDetector {
	return &IllegalTransitionDetector{
		attempts: make([]TransitionAttempt, 0),
	}
}

// RecordViolation logs an illegal transition attempt.
func (d *IllegalTransitionDetector) RecordViolation(from, to RuntimeState, tick int64) {
	d.mu.Lock()
	defer d.mu.Unlock()
	d.attempts = append(d.attempts, TransitionAttempt{From: from, To: to, Time: tick})
	d.violationCount++
}

func (d *IllegalTransitionDetector) GetViolations() []TransitionAttempt {
	d.mu.Lock()
	defer d.mu.Unlock()
	return d.attempts
}

// ValidateTransition checks if a move from current to next is architecturally sound.
func ValidateTransition(current, next RuntimeState) error {
	if current == next {
		return nil // NOP transition is valid
	}
	allowed, ok := ValidTransitions[current]
	if !ok {
		return fmt.Errorf("invalid current state: %s", current)
	}
	for _, a := range allowed {
		if a == next {
			return nil
		}
	}
	return fmt.Errorf("illegal state transition: %s -> %s", current, next)
}
