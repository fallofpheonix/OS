/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 9b — CONTAINMENT TRANSITION RULES (Layer 4)
//
// This file defines the valid transitions for the IsolationEngine.
// The containment lifecycle follows a strict escalation/de-escalation path.
//
// TRANSITION MAP:
//   OBSERVE  → [WATCH, THROTTLE]
//   WATCH    → [THROTTLE, ISOLATE]
//   THROTTLE → [ISOLATE]
//   ISOLATE  → [RECOVER]
//   RECOVER  → [OBSERVE]
//
// DESIGN RATIONALE: The lifecycle prevents premature recovery.
// A process must go through OBSERVE → WATCH → THROTTLE → ISOLATE → RECOVER → OBSERVE
// before it can return to normal monitoring. This ensures sufficient evidence
// is accumulated before de-escalation.
// =========================================================================
package containment

import (
	"fmt"
	"time"
)

// Transition enforces the observe-throttle-isolate-recover lifecycle.
// This is the CONTAINMENT STATE CHANGE function.
//
// WORKFLOW: Called from the Warden or SandboxWarden during containment operations:
//  1. Acquire exclusive lock
//  2. Validate transition against containment rules
//  3. Record IsolationRecord with timestamp, evidence, decision
//  4. Update CurrentState and append to History
//
// AUDIT: Every containment transition is recorded in the History slice.
// This provides a complete audit trail for forensic analysis.
func (e *IsolationEngine) Transition(next IsolationState, evidenceID, decisionID string) error {
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
