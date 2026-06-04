// Package warden implements the tactical enforcement FSM for PhoenixOS.
// Domain Logic: Orchestrates state transition validation against formal invariants and triggers physical actuators.
// Responsibility: Acts as the central enforcement engine to maintain system integrity and containment.
package warden

import (
	"context"
	"fmt"
	"log"
	"sync"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/config"
	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// Actuator defines the canonical security actuator contract.
type Actuator = securityv1.Actuator

type SystemState string
type ActuationClass string

const (
	StateSafe        SystemState = "SAFE"
	StateWatch       SystemState = "WATCH"
	StateSuspicious  SystemState = "SUSPICIOUS"
	StateCritical    SystemState = "CRITICAL"
	StateCompromised SystemState = "COMPROMISED"

	ClassLog      ActuationClass = "LOG"
	ClassWarn     ActuationClass = "WARN"
	ClassThrottle ActuationClass = "THROTTLE"
	ClassFreeze   ActuationClass = "FREEZE"
	ClassIsolate  ActuationClass = "ISOLATE"
	ClassKill     ActuationClass = "KILL"
)

// Warden is the primary enforcement controller.
// Concurrency: Thread-safe via sync.RWMutex.
// State Management: Manages system state transitions, invariant registries, and actuator orchestration.
type Warden struct {
	mu               sync.RWMutex
	Bus              *bus.Bus
	State            SystemState
	LastAction       string
	History          []string
	Policies         *config.RedLines
	Invariants       []Invariant
	Actuators        []Actuator
	DiagnosticLogger *log.Logger
	ShadowMode       bool // P7.2: Shadow Mode enforcement
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewWarden initializes a new Warden FSM instance.
// I/O: None.
// Side Effects: Initializes registries and sets default ShadowMode.
// Complexity: O(1).
// Deprecated: new codebase should use contracts or initialize via NewWarden.
func NewWarden(b *bus.Bus) *Warden {
	return &Warden{
		Bus:        b,
		State:      StateSafe,
		Invariants: make([]Invariant, 0),
		Actuators:  make([]Actuator, 0),
		ShadowMode: true,
	}
}

// LABEL: [MUTABLE] [UNCONSTRAINED] [STABLE]
// ActuateRequest is the core decision gate for system escalation and containment.
// I/O: Potential logging and actuator interactions.
// Side Effects: Updates system state, triggers physical actuators, and appends to history.
// Complexity: O(I + A) where I is the number of invariants and A is the number of actuators.
func (w *Warden) ActuateRequest(req AuthorityEscalationRequest, seq int, lamportClock uint64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 1. FORMAL INVARIANT VERIFICATION
	for _, inv := range w.Invariants {
		if err := inv.Verify(req, w.State); err != nil {
			w.logViolation("[Warden Panic] Invariant Violation: %v", err)
			w.emergencyHaltLocked(req, err)
			return false
		}
	}

	// 2. POLICY MODE GATING
	if w.Policies != nil {
		if w.Policies.Actuation.Mode == "LOG_ONLY" {
			log.Printf("[Warden Policy] LOG_ONLY mode active. Suppressing %s actuation.", req.ActuationClass)
			w.History = append(w.History, fmt.Sprintf("%d:%s->%s(SUPPRESSED)", lamportClock, w.State, req.TargetState))
			return true
		}
	}

	// 3. DETERMINISTIC GUARD: Transition Validation
	if !w.isValidTransition(w.State, req.TargetState) {
		w.logViolation("[Warden Violation] Illegal state transition attempted: %s -> %s", w.State, req.TargetState)
		return false
	}

	// 4. SHADOW MODE ENFORCEMENT
	if w.ShadowMode {
		log.Printf("[Warden ShadowMode] WouldHaveContained: %s (PID: %d, Class: %s)", req.TargetState, req.TargetPID, req.ActuationClass)
		w.History = append(w.History, fmt.Sprintf("%d:%s->%s(SHADOW)", lamportClock, w.State, req.TargetState))
		w.State = req.TargetState
		return true
	}

	stateBefore := w.State
	w.State = req.TargetState
	w.LastAction = string(req.ActuationClass)

	// 5. PHYSICAL ACTUATION (The Containment Ladder)
	if req.TargetPID > 0 {
		ctx := context.Background()
		for _, act := range w.Actuators {
			var err error
			name := "UnknownActuator"
			if na, ok := act.(interface{ Name() string }); ok {
				name = na.Name()
			}
			if req.ActuationClass == ClassKill {
				err = act.Kill(ctx, req.TargetPID)
			} else if req.ActuationClass != ClassLog {
				err = act.Actuate(ctx, req)
			} else {
				log.Printf("[Warden Log] PID %d observed under %s", req.TargetPID, req.TargetState)
			}
			if err != nil {
				w.logViolation("[Warden Actuator Error] %s failed %s on PID %d: %v", name, req.ActuationClass, req.TargetPID, err)
			}
		}
	}

	log.Printf("[Warden Actuation] SUCCESS: %s -> %s | Action: %s | LC: %d", stateBefore, req.TargetState, req.ActuationClass, lamportClock)
	w.History = append(w.History, fmt.Sprintf("%d:%s->%s", lamportClock, stateBefore, req.TargetState))

	return true
}

type emergencyContainment struct {
	target string
	level  securityv1.ContainmentLevel
	reason string
}

func (e emergencyContainment) Target() string                      { return e.target }
func (e emergencyContainment) Level() securityv1.ContainmentLevel  { return e.level }
func (e emergencyContainment) Reason() string                      { return e.reason }

func (w *Warden) emergencyHaltLocked(req AuthorityEscalationRequest, err error) {
	w.logViolation("[Warden EMERGENCY HALT] Quenching System: %v", err)
	w.State = StateCompromised
	if req.TargetPID > 0 {
		ctx := context.Background()
		action := emergencyContainment{
			target: fmt.Sprintf("PID:%d", req.TargetPID),
			level:  securityv1.LevelIsolate,
			reason: fmt.Sprintf("emergency halt: %v", err),
		}
		for _, act := range w.Actuators {
			_ = act.Actuate(ctx, action)
		}
	}
}

func (w *Warden) isValidTransition(current, target SystemState) bool {
	if current == target {
		return true
	}
	// Strict state ladder
	transitions := map[SystemState][]SystemState{
		StateSafe:        {StateWatch},
		StateWatch:       {StateSafe, StateSuspicious},
		StateSuspicious:  {StateWatch, StateCritical},
		StateCritical:    {StateSuspicious, StateCompromised},
		StateCompromised: {StateCritical},
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

// LABEL: [READ_ONLY] [UNCONSTRAINED] [STABLE]
// GetStatus returns a formatted status string of the Warden's current state.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func (w *Warden) GetStatus() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return fmt.Sprintf("State: %s, LastAction: %s", w.State, w.LastAction)
}

func mapLevelToSystemState(level securityv1.ContainmentLevel) SystemState {
	switch level {
	case securityv1.LevelNone:
		return StateSafe
	case securityv1.LevelMonitor:
		return StateWatch
	case securityv1.LevelSandbox:
		return StateSuspicious
	case securityv1.LevelIsolate:
		return StateCritical
	case securityv1.LevelQuench:
		return StateCompromised
	default:
		return StateSafe
	}
}

func mapLevelToActuationClass(level securityv1.ContainmentLevel) ActuationClass {
	switch level {
	case securityv1.LevelNone:
		return ClassLog
	case securityv1.LevelMonitor:
		return ClassWarn
	case securityv1.LevelSandbox:
		return ClassThrottle
	case securityv1.LevelIsolate:
		return ClassIsolate
	case securityv1.LevelQuench:
		return ClassKill
	default:
		return ClassLog
	}
}

// Actuate implements securityv1.Actuator.
func (w *Warden) Actuate(ctx context.Context, action securityv1.Containment) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	targetState := mapLevelToSystemState(action.Level())
	pid := 0
	_, _ = fmt.Sscanf(action.Target(), "PID:%d", &pid)

	req := AuthorityEscalationRequest{
		EventID:        "CONTRACT-ACTUATION",
		TargetPID:      pid,
		TargetState:    targetState,
		ActuationClass: mapLevelToActuationClass(action.Level()),
		EvidenceWeight: 1.0,
	}

	// 1. FORMAL INVARIANT VERIFICATION
	for _, inv := range w.Invariants {
		if err := inv.Verify(req, w.State); err != nil {
			w.logViolation("[Warden Panic] Invariant Violation: %v", err)
			w.emergencyHaltLocked(req, err)
			return err
		}
	}

	// 2. POLICY MODE GATING
	if w.Policies != nil {
		if w.Policies.Actuation.Mode == "LOG_ONLY" {
			log.Printf("[Warden Policy] LOG_ONLY mode active. Suppressing %s actuation.", req.ActuationClass)
			w.History = append(w.History, fmt.Sprintf("0:%s->%s(SUPPRESSED)", w.State, req.TargetState))
			return nil
		}
	}

	// 3. DETERMINISTIC GUARD: Transition Validation
	if !w.isValidTransition(w.State, req.TargetState) {
		w.logViolation("[Warden Violation] Illegal state transition attempted: %s -> %s", w.State, req.TargetState)
		return fmt.Errorf("illegal state transition: %s -> %s", w.State, req.TargetState)
	}

	// 4. SHADOW MODE ENFORCEMENT
	if w.ShadowMode {
		log.Printf("[Warden ShadowMode] WouldHaveContained: %s (PID: %d, Class: %s)", req.TargetState, req.TargetPID, req.ActuationClass)
		w.History = append(w.History, fmt.Sprintf("0:%s->%s(SHADOW)", w.State, req.TargetState))
		w.State = req.TargetState
		return nil
	}

	stateBefore := w.State
	w.State = req.TargetState
	w.LastAction = string(req.ActuationClass)

	// 5. PHYSICAL ACTUATION (The Containment Ladder)
	if req.TargetPID > 0 {
		for _, act := range w.Actuators {
			var err error
			name := "UnknownActuator"
			if na, ok := act.(interface{ Name() string }); ok {
				name = na.Name()
			}
			if req.ActuationClass == ClassKill {
				err = act.Kill(ctx, req.TargetPID)
			} else if req.ActuationClass != ClassLog {
				err = act.Actuate(ctx, req)
			} else {
				log.Printf("[Warden Log] PID %d observed under %s", req.TargetPID, req.TargetState)
			}
			if err != nil {
				w.logViolation("[Warden Actuator Error] %s failed %s on PID %d: %v", name, req.ActuationClass, req.TargetPID, err)
			}
		}
	}

	log.Printf("[Warden Actuation] SUCCESS: %s -> %s | Action: %s | LC: 0", stateBefore, req.TargetState, req.ActuationClass)
	w.History = append(w.History, fmt.Sprintf("0:%s->%s", stateBefore, req.TargetState))

	return nil
}

// Kill implements securityv1.Actuator.
func (w *Warden) Kill(ctx context.Context, pid int) error {
	w.mu.Lock()
	defer w.mu.Unlock()

	for _, act := range w.Actuators {
		if err := act.Kill(ctx, pid); err != nil {
			w.logViolation("[Warden Kill Error] Actuator failed to kill PID %d: %v", pid, err)
		}
	}
	return nil
}

// GetCurrentLevel implements securityv1.Actuator.
func (w *Warden) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return mapSystemStateToLevel(w.State), nil
}

// Target implements securityv1.Containment.
func (w *Warden) Target() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return w.LastAction
}

// Level implements securityv1.Containment.
func (w *Warden) Level() securityv1.ContainmentLevel {
	w.mu.RLock()
	defer w.mu.RUnlock()
	return mapSystemStateToLevel(w.State)
}

// Reason implements securityv1.Containment.
func (w *Warden) Reason() string {
	w.mu.RLock()
	defer w.mu.RUnlock()
	if len(w.History) > 0 {
		return w.History[len(w.History)-1]
	}
	return "No history"
}

// CurrentLevel implements securityv1.Escalation.
func (w *Warden) CurrentLevel() securityv1.ContainmentLevel {
	return w.Level()
}

// TargetLevel implements securityv1.Escalation.
func (w *Warden) TargetLevel() securityv1.ContainmentLevel {
	return w.Level()
}

// TriggerReason implements securityv1.Escalation.
func (w *Warden) TriggerReason() string {
	return w.Reason()
}
