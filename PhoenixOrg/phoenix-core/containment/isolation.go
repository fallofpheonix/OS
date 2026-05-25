package containment

import (
	"sync"
	"time"
)

type IsolationState string

const (
	StateObserve IsolationState = "OBSERVE"
	StateWatch   IsolationState = "WATCH"
	StateThrottle IsolationState = "THROTTLE"
	StateIsolate  IsolationState = "ISOLATE"
	StateRecover  IsolationState = "RECOVER"
)

// IsolationRecord audits every containment state change.
type IsolationRecord struct {
	Timestamp  time.Time
	Previous   IsolationState
	Current    IsolationState
	EvidenceID string
	DecisionID string
}

// IsolationEngine monitors the containment lifecycle.
type IsolationEngine struct {
	mu           sync.RWMutex
	CurrentState IsolationState
	History      []IsolationRecord
}

func NewIsolationEngine(initial IsolationState) *IsolationEngine {
	return &IsolationEngine{
		CurrentState: initial,
		History: []IsolationRecord{
			{Timestamp: time.Now(), Current: initial, Previous: initial},
		},
	}
}
