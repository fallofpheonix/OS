/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 11 — STATE AUDIT TRAIL (Layer 5)
//
// The StateRegistry maintains the AUTHORITATIVE system state with a
// complete audit trail of all transitions. Every state change is recorded
// with timestamp, reason, and evidence ID.
//
// WORKFLOW:
//   Warden.Transition() → StateRegistry.Transition()
//     → Validate transition against FSM rules
//     → Record StateRecord{Previous, Current, Timestamp, Reason, EvidenceID}
//     → Update GlobalMetrics (transition count, latency, state entries)
//     → Ledger records the transition for forensic replay
//
// NOTE: This is a SEPARATE state tracking system from the Warden FSM.
// The Warden manages enforcement state; the StateRegistry manages
// audit state. They use DIFFERENT state definitions:
//   Warden: SAFE → WATCH → SUSPICIOUS → CRITICAL → COMPROMISED
//   Registry: SAFE → WATCH → ALERT → CONTAIN → RECOVERY
//
// This is an architectural inconsistency — the two systems are never reconciled.
// =========================================================================
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

// NewStateRegistry creates the state tracking system with an initial state.
// Called once during system startup with the initial safe state.
//
// WORKFLOW: startup → NewStateRegistry(StateSafe) → ready for Transition() calls
func NewStateRegistry(initial SystemState) *StateRegistry {
	return &StateRegistry{
		CurrentState: initial,
		History: []StateRecord{
			{ID: 0, Previous: initial, Current: initial, Timestamp: time.Now(), Reason: "init"},
		},
	}
}

// Transition attempts a state transition with full audit recording.
// This is the STATE CHANGE FUNCTION for the audit tracking system.
//
// WORKFLOW: Called from the Orchestrator or Warden during state changes:
//  1. Acquire exclusive lock (prevent concurrent transitions)
//  2. Validate transition against FSM rules (SAFE→WATCH→ALERT→CONTAIN→RECOVERY)
//  3. Build StateRecord with before/after state, timestamp, reason, evidence
//  4. Update CurrentState and append to History
//  5. Update GlobalMetrics (transition count, state entries, latency)
//
// AUDIT TRAIL: Every transition is recorded in the History slice.
// The History is used for:
//   - Deterministic replay (CYCLE 11)
//   - Rollback verification (CYCLE 10)
//   - Forensic analysis (post-incident)
//
// METRICS: GlobalMetrics tracks:
//   - TransitionCount: total transitions
//   - IllegalTransitions: rejected transitions (security events)
//   - Rollbacks: recovery transitions
//   - StateEntry counts: time spent in each state
//   - TransitionLatency: performance monitoring
func (r *StateRegistry) Transition(next SystemState, reason, evidenceID, decisionID string) error {
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
