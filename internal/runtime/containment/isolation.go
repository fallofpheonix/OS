/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9 — CONTAINMENT STATE MACHINE (Layer 4)
//
// The IsolationEngine tracks the CONTAINMENT LIFECYCLE for processes.
// It enforces a strict state machine: OBSERVE → WATCH → THROTTLE → ISOLATE → RECOVER
//
// WORKFLOW:
//   Warden detects threat → IsolationEngine.Transition(THROTTLE)
//     → Process gets rate-limited
//   → If threat persists: Transition(ISOLATE)
//     → Process gets fully isolated (cgroup freeze + namespace sever)
//   → After containment: Transition(RECOVER)
//     → Process gets restored
//   → Back to OBSERVE: monitoring resumes
//
// NOTE: This is a SEPARATE containment tracking system from the Warden FSM.
// The Warden tracks system-level security state; the IsolationEngine tracks
// per-process containment state. They use different state definitions.
// =========================================================================
package containment

import (
	"sync"
	"time"
)

type IsolationState string

const (
	StateObserve  IsolationState = "OBSERVE"
	StateWatch    IsolationState = "WATCH"
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

// NewIsolationEngine creates the containment lifecycle tracker.
// Called once during system startup with initial state OBSERVE.
func NewIsolationEngine(initial IsolationState) *IsolationEngine {
	return &IsolationEngine{
		CurrentState: initial,
		History: []IsolationRecord{
			{Timestamp: time.Now(), Current: initial, Previous: initial},
		},
	}
}
