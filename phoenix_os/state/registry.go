package state

import (
	"fmt"
	"sync"
	"time"
)

// StateRegistry maintains the authoritative system state.
type StateRegistry struct {
	mu           sync.RWMutex
	CurrentState SystemState
	History      []StateRecord
}

// NewStateRegistry initializes a registry.
func NewStateRegistry(initial SystemState) *StateRegistry {
	return &StateRegistry{
		CurrentState: initial,
		History: []StateRecord{
			{ID: 0, Previous: initial, Current: initial, Timestamp: time.Now(), Reason: "init"},
		},
	}
}

// Transition attempts a state transition with audit.
func (r *StateRegistry) Transition(next SystemState, reason string, evidenceID string, decisionID string) error {
	start := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	if !isValidTransition(r.CurrentState, next, false) {
		GlobalMetrics.IncIllegal()
		return fmt.Errorf("illegal transition from %s to %s", r.CurrentState, next)
	}

	record := StateRecord{
		ID:         len(r.History),
		Previous:   r.CurrentState,
		Current:    next,
		Timestamp:  time.Now(),
		Reason:     fmt.Sprintf("%s (Decision: %s)", reason, decisionID),
		EvidenceID: evidenceID,
	}

	r.CurrentState = next
	r.History = append(r.History, record)
	GlobalMetrics.IncTransition()
	GlobalMetrics.IncStateEntry(next)
	
	duration := time.Since(start).Nanoseconds()
	GlobalMetrics.UpdateTransitionLatency(int(duration))
	
	return nil
}
