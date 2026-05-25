package decision

import (
	"fmt"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix-control/arbiter"
	"github.com/fallofpheonix/phoenix-control/warden"
)

// DecisionBus aggregates outputs from Replay, Arbiter, and Warden.
// It acts as the final gate before actuation.
type DecisionBus struct {
	mu           sync.RWMutex
	Decisions    []IntegratedDecision
	Subscribers  []chan IntegratedDecision
	MaxHistory   int
	DroppedCount uint64
}

// IntegratedDecision combines policy (Arbiter), evidence (Replay), and state (Warden).
type IntegratedDecision struct {
	PID              int
	PolicyDecision   arbiter.Decision
	EvidenceScore    float64 // 0.0 - 1.0 from Replay
	Confidence       float64 // Merged confidence
	AuthorizedAction warden.ActuationClass
	AuthorizedState  warden.SystemState
	Timestamp        time.Time
}

func NewDecisionBus() *DecisionBus {
	return &DecisionBus{
		Decisions:   make([]IntegratedDecision, 0),
		Subscribers: make([]chan IntegratedDecision, 0),
		MaxHistory:  1000, // Deterministic bound
	}
}

// Publish routes an integrated decision to all subscribers (e.g., Warden for actuation).
func (b *DecisionBus) Publish(d IntegratedDecision) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Bounded history to prevent memory leaks
	if len(b.Decisions) >= b.MaxHistory {
		b.Decisions = b.Decisions[1:]
	}
	b.Decisions = append(b.Decisions, d)

	for i, sub := range b.Subscribers {
		select {
		case sub <- d:
		default:
			// Non-blocking send failure
			b.DroppedCount++
			fmt.Printf("[DECISION BUS] AUDIT: Decision dropped for subscriber %d (PID: %d, Action: %v). DroppedCount: %d\n",
				i, d.PID, d.AuthorizedAction, b.DroppedCount)
			// TODO: ReplayPersist() for forensic audit trail
		}
	}
}

// Subscribe allows modules to listen for authorized decisions.
func (b *DecisionBus) Subscribe() chan IntegratedDecision {
	b.mu.Lock()
	defer b.mu.Unlock()

	ch := make(chan IntegratedDecision, 10)
	b.Subscribers = append(b.Subscribers, ch)
	return ch
}
