/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 8 — FSM ENFORCEMENT (Layer 3)
//
// The Warden is the DETERMINISTIC STATE MACHINE that governs system security.
// It enforces a strict state ladder: SAFE → WATCH → SUSPICIOUS → CRITICAL → COMPROMISED
// with defined transitions in both directions (escalation and de-escalation).
//
// WORKFLOW:
//   [CYCLE 6] Arbiter evaluates DriftScore + TCS → decides target state
//     → [CYCLE 7] Oracle provides directive (if authorized)
//       → [CYCLE 8] Warden.Transition(targetState) → FSM state change
//         → [CYCLE 9] State change triggers kernel enforcement
//           → eBPF map updates → LSM hooks enforce execution blocks
//
//   [CYCLE 10] On failure: KillSwitch.Engage() → Lock() → no more transitions
//     → RecoveryLoop attempts rollback to previous safe state
//
// STATE LADDER:
//   SAFE:       Normal operation, no threats detected
//   WATCH:      Minor anomaly detected, monitoring increased
//   SUSPICIOUS: Multiple anomalies confirmed, containment prepared
//   CRITICAL:   Active threat detected, process isolation initiated
//   COMPROMISED: Full breach, all non-essential processes frozen
//
// INVARIANT: Transitions MUST follow the ladder. No skipping states.
// This prevents a single event from jumping from SAFE to COMPROMISED.
// =========================================================================
package engine

import (
	"context"
	"fmt"
	"sync"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

type SystemState string

const (
	StateSafe        SystemState = "SAFE"
	StateWatch       SystemState = "WATCH"
	StateSuspicious  SystemState = "SUSPICIOUS"
	StateCritical    SystemState = "CRITICAL"
	StateCompromised SystemState = "COMPROMISED"
)

type ActuationClass string

const (
	ClassNone     ActuationClass = "NONE"
	ClassLog      ActuationClass = "LOG"
	ClassThrottle ActuationClass = "THROTTLE"
	ClassIsolate  ActuationClass = "ISOLATE"
)

type GraphProof struct {
	Path            []string
	ExpectedNsproxy uint32
}

type AuthorityEscalationRequest struct {
	EventID        string
	TargetPID      int
	TargetTgid     int
	TargetNsproxy  uint32
	TargetState    SystemState
	ActuationClass ActuationClass
	EvidenceWeight float64
	Certificate    []byte
	GraphProof     *GraphProof
}

// Warden represents the deterministic FSM enforcement engine.
type Warden struct {
	mu     sync.RWMutex
	state  SystemState
	locked bool
}

// NewWarden creates the FSM enforcement engine in SAFE state.
// Called once during system startup. The Warden starts in SAFE mode,
// meaning no threats are detected and the system operates normally.
//
// WORKFLOW: startup → NewWarden() → ready for Transition() calls
//
//	from ArbiterFeature.Evaluate() in CYCLE 6/8
func NewWarden() *Warden {
	return &Warden{
		state: StateSafe,
	}
}

// ActuateRequest is the ENTRY POINT for the AI Orchestrator to trigger state changes.
// Called from WardenFeature.Actuate() in CYCLE 8.
//
// WORKFLOW: AI Orchestrator → WardenFeature.Actuate()
//
//	→ Warden.ActuateRequest(req, seq, lamport) → Transition(req.TargetState)
//	  → FSM state changes → returns true if successful
//
// SECURITY NOTE: This method IGNORES all fields in the request except TargetState.
// The Certificate, GraphProof, EvidenceWeight, TargetPID, TargetTgid, and TargetNsproxy
// fields are DISCARDED. This is a critical authorization bypass.
// The method should validate the certificate, check evidence weight thresholds,
// and verify the graph proof before allowing the transition.
func (w *Warden) ActuateRequest(req AuthorityEscalationRequest, seq int, lamport uint64) bool {
	return w.Transition(req.TargetState) == nil
}

// Actuate implements securityv1.Actuator interface.
func (w *Warden) Actuate(ctx context.Context, action securityv1.Containment) error {
	var targetState SystemState
	switch action.Level() {
	case securityv1.LevelNone:
		targetState = StateSafe
	case securityv1.LevelMonitor:
		targetState = StateWatch
	case securityv1.LevelSandbox:
		targetState = StateSuspicious
	case securityv1.LevelIsolate:
		targetState = StateCritical
	case securityv1.LevelQuench:
		targetState = StateCompromised
	default:
		return fmt.Errorf("unknown containment level: %v", action.Level())
	}
	return w.Transition(targetState)
}

// Kill implements securityv1.Actuator interface.
func (w *Warden) Kill(ctx context.Context, pid int) error {
	w.Lock()
	return nil
}

// GetCurrentLevel implements securityv1.Actuator interface.
func (w *Warden) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	state := w.GetState()
	switch state {
	case StateSafe:
		return securityv1.LevelNone, nil
	case StateWatch:
		return securityv1.LevelMonitor, nil
	case StateSuspicious:
		return securityv1.LevelSandbox, nil
	case StateCritical:
		return securityv1.LevelIsolate, nil
	case StateCompromised:
		return securityv1.LevelQuench, nil
	default:
		return securityv1.LevelNone, fmt.Errorf("unknown warden state: %s", state)
	}
}

// Transition attempts a deterministic state transition based on the FSM ladder.
// This is the CORE ENFORCEMENT FUNCTION of the entire system.
//
// WORKFLOW: Called from Actuate() or directly by the Arbiter.
//  1. Acquire exclusive lock (prevents concurrent transitions)
//  2. Check if FSM is locked (KillSwitch engaged)
//  3. Validate transition against the state ladder
//  4. If valid: update state, return nil
//  5. If invalid: return error, state unchanged
//
// SIDE EFFECTS:
//   - State change is IMMEDIATE and IRREVERSIBLE without rollback
//   - No audit trail is recorded here — the caller must record the transition
//   - The Ledger records the transition AFTER this function returns
//
// INVARIANT: This function is the ONLY way to change the Warden's state.
// All other code paths (KillSwitch.Engage, etc.) go through this function
// or Lock().
func (w *Warden) Transition(target SystemState) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.locked {
		return fmt.Errorf("FSM LOCKOUT: Warden is in emergency preservation mode")
	}

	if !isValidTransition(w.state, target) {
		return fmt.Errorf("INVALID TRANSITION: %s -> %s violation", w.state, target)
	}

	w.state = target
	return nil
}

// isValidTransition enforces the strict state ladder.
// This is the MATHEMATICAL HEART of the security model.
//
// TRANSITION MAP:
//
//	SAFE       → [WATCH]
//	WATCH      → [SAFE, SUSPICIOUS]
//	SUSPICIOUS → [WATCH, CRITICAL]
//	CRITICAL   → [SUSPICIOUS, COMPROMISED]
//	COMPROMISED → [CRITICAL]
//
// DESIGN RATIONALE: The ladder prevents catastrophic jumps.
// A single event cannot go from SAFE to COMPROMISED — it must pass through
// WATCH, SUSPICIOUS, and CRITICAL first. This forces the system to accumulate
// evidence before escalating, preventing false-positive catastrophes.
//
// DE-ESCALATION: Each state can de-escalate to the previous state.
// COMPROMISED → CRITICAL (requires verification), but NOT COMPROMISED → SAFE.
// Full recovery from COMPROMISED requires going through CRITICAL → SUSPICIOUS
// → WATCH → SAFE, each step requiring its own evidence.
func isValidTransition(current, target SystemState) bool {
	if current == target {
		return true
	}

	// Strict state ladder transitions
	transitions := map[SystemState][]SystemState{
		StateSafe:        {StateWatch},
		StateWatch:       {StateSafe, StateSuspicious},
		StateSuspicious:  {StateWatch, StateCritical},
		StateCritical:    {StateSuspicious, StateCompromised},
		StateCompromised: {StateCritical}, // Recovery path requires verification
	}

	allowed, ok := transitions[current]
	if !ok {
		return false
	}

	for _, a := range allowed {
		if a == target {
			return true
		}
	}

	return false
}

// Lock freezes the FSM in place (Emergency Override).
// Called by KillSwitch.Engage() during CYCLE 10 (Emergency Halt).
//
// WORKFLOW: KillSwitch.Engage() → Warden.Lock() → all future Transition() calls fail
//
//	→ System enters "preservation mode" — no state changes allowed
//	→ RecoveryLoop attempts to identify the cause and restore
//
// INVARIANT: Once locked, the FSM CANNOT be unlocked without process restart.
// This is by design — the KillSwitch is a one-way operation.
func (w *Warden) Lock() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.locked = true
}

// GetState returns the current FSM state.
// Called by the Arbiter to check current security posture.
// Called by the Game Server to display current state to operators.
// Called by the Ledger to record state before/after transitions.
//
// WORKFLOW: Read-only query, no side effects.
// Used in CYCLE 6 (Arbiter checks current state before deciding next state)
// and CYCLE 8 (Ledger records state before transition for audit trail).
func (w *Warden) GetState() SystemState {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.state
}
