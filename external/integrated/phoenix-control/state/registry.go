package state

import (
	"fmt"
	"sync"
)

// Registry manages the authoritative system state.
type Registry struct {
	mu       sync.RWMutex
	current  RuntimeState
	history  []RuntimeState
	Detector *IllegalTransitionDetector
}

// NewRegistry creates a fresh state registry starting at Safe.
func NewRegistry() *Registry {
	return &Registry{
		current:  Safe,
		history:  []RuntimeState{Safe},
		Detector: NewIllegalTransitionDetector(),
	}
}

// Get returns the current state.
func (r *Registry) Get() RuntimeState {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.current
}

// Set attempts to transition to a new state.
func (r *Registry) Set(target RuntimeState, reason string, tick int64) error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if err := ValidateTransition(r.current, target); err != nil {
		r.Detector.RecordViolation(r.current, target, tick)
		return err
	}

	r.history = append(r.history, target)
	r.current = target
	return nil
}

// Rollback undoes the last state transition.
func (r *Registry) Rollback() error {
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.history) <= 1 {
		return fmt.Errorf("no history to rollback")
	}

	// Remove last state
	r.history = r.history[:len(r.history)-1]
	r.current = r.history[len(r.history)-1]
	return nil
}
