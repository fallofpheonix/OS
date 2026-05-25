package warden

import (
	"fmt"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix-contracts"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

type SystemState = contracts.SystemState

const (
	StateSafe     = contracts.StateSafe
	StateWatch    = contracts.StateWatch
	StateAlert    = contracts.StateAlert
	StateContain  = contracts.StateContain
	StateRecovery = contracts.StateRecovery
)

// TriggerType defines the types of telemetry events that can trigger a state transition.
type TriggerType string

const (
	TriggerForkBurst     TriggerType = "PROCESS_FORK_BURST"
	TriggerNetworkBeacon TriggerType = "NETWORK_BEACON"
	TriggerMassWrite     TriggerType = "MASS_WRITE"
	TriggerReverseShell  TriggerType = "REVERSE_SHELL"
	TriggerHumanOverride TriggerType = "HUMAN_OVERRIDE"
)

// Action interface for deterministic containment actions.
type Action interface {
	Observe(targetID int)
	Snapshot(targetID int)
	Isolate(targetID int)
	Recover(targetID int)
}

type Warden struct {
	mu                sync.RWMutex
	State             SystemState
	Bus               *bus.Bus
	RecoveryBudget    int
	DeescalationCount int
	LastTick          uint64
	DwellTicks        uint64 // Hysteresis duration
}

func NewWarden(b *bus.Bus) *Warden {
	return &Warden{
		State:          StateSafe,
		Bus:            b,
		RecoveryBudget: 3,
		DwellTicks:     30,
	}
}

// EvaluateTrigger processes a normalized event and determines if a state transition is required.
func (w *Warden) EvaluateTrigger(trigger TriggerType, targetPID int, currentTick uint64) {
	w.mu.Lock()
	defer w.mu.Unlock()

	oldState := w.State
	newState := w.State

	switch trigger {
	case TriggerReverseShell:
		newState = StateContain
	case TriggerMassWrite, TriggerNetworkBeacon:
		if w.State == StateSafe || w.State == StateWatch {
			newState = StateAlert
		}
	case TriggerForkBurst:
		if w.State == StateSafe {
			newState = StateWatch
		}
	case TriggerHumanOverride:
		newState = StateRecovery
	}

	if newState != oldState {
		w.transition(newState, trigger, targetPID, currentTick)
	}
}

// Actuate performs a forced state transition and returns true if the state changed.
// This is used by Arbiter, Decision Bus, and other higher-level modules.
func (w *Warden) Actuate(target SystemState, class ActuationClass, confidence float64, targetPID int, timestamp int64, tick uint64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	oldState := w.State

	// Confidence gating
	if confidence < 0.5 && class > ClassLog {
		fmt.Printf("[WARDEN] Actuation Denied: Confidence %.2f too low for Class %d\n", confidence, class)
		return false
	}

	if target != oldState {
		w.transition(target, "DIRECT_ACTUATION", targetPID, tick)
		return w.State != oldState
	}

	return false
}

// ActuateLegacy provides backward compatibility for callers using interface{} for targetPID.
func (w *Warden) ActuateLegacy(target SystemState, class ActuationClass, confidence float64, targetPID interface{}, timestamp int64, tick uint64) bool {
	pid := 0
	switch v := targetPID.(type) {
	case int:
		pid = v
	case int64:
		pid = int(v)
	}
	return w.Actuate(target, class, confidence, pid, timestamp, tick)
}

func (w *Warden) transition(newState SystemState, trigger TriggerType, pid int, tick uint64) {
	// [STABILITY] Budget replenishment: Require prolonged stability in StateSafe.
	// TODO: Replace with formal TLA+ validated quench logic (Stage B).
	if w.State == StateSafe && newState == StateSafe && tick > w.LastTick+(w.DwellTicks*10) {
		if w.DeescalationCount > 0 {
			fmt.Printf("[WARDEN] Stability Achieved: Recovery Budget Replenished (Tick %d)\n", tick)
			w.DeescalationCount = 0
		}
	}

	// [SECURITY] Detect oscillation: Rapid re-escalation after de-escalation
	if !w.isDeescalation(newState, w.State) && w.State == StateSafe && tick < w.LastTick+(w.DwellTicks*2) {
		fmt.Printf("[WARDEN] OSCILLATION WARNING: Rapid re-escalation to %s detected at Tick %d\n", newState, tick)
	}

	// Dwell limit check for de-escalation
	if w.isDeescalation(newState, w.State) {
		if tick < w.LastTick+w.DwellTicks {
			fmt.Printf("[WARDEN] Hysteresis Block: De-escalation to %s denied (Tick %d)\n", newState, tick)
			return
		}
		if w.DeescalationCount >= w.RecoveryBudget && trigger != TriggerHumanOverride {
			fmt.Printf("[WARDEN] Recovery Budget Exhausted: De-escalation to %s blocked\n", newState)
			return
		}
		w.DeescalationCount++
	}

	oldState := w.State
	fmt.Printf("[WARDEN] FSM Transition: %s -> %s (Trigger: %s, PID: %d, Tick: %d)\n", oldState, newState, trigger, pid, tick)
	w.State = newState
	w.LastTick = tick

	// [PHOENIX-F1] Publish state transition to the bus
	if w.Bus != nil {
		payload := []byte(fmt.Sprintf(`{"from":"%s","to":"%s","trigger":"%s","pid":%d,"tick":%d}`, oldState, newState, trigger, pid, tick))
		w.Bus.Publish("phoenix.sys.state", bus.TelemetryEvent{
			SeqID:        -int64(tick), // Use negative tick for system internal events
			LogicalTick:  tick,
			WallTimeUnix: time.Now().Unix(),
			Source:       "phoenix.warden",
			EventType:    "system.state_transition",
			Severity:     1.0,
			Payload:      payload,
		})
	}

	// Execute deterministic action based on new state
	w.executeAction(newState, pid)
}

// ProcessAdvice takes AI recommendations and translates them into Warden state transitions or actions.
func (w *Warden) ProcessAdvice(command string, confidence float64, pid int, tick uint64) {
	if confidence < 0.7 {
		fmt.Printf("[WARDEN] AI Advice Rejected: Low Confidence (%.2f)\n", confidence)
		return
	}

	fmt.Printf("[WARDEN] Processing AI Advice: %s (Confidence: %.2f)\n", command, confidence)

	switch command {
	case "ISOLATE_PID":
		w.transition(StateContain, "AI_ADVICE", pid, tick)
	case "THROTTLE_NETWORK":
		w.transition(StateAlert, "AI_ADVICE", pid, tick)
	case "REVOKE_UID":
		fmt.Printf("[WARDEN ACTION] REVOKING UID for PID %d\n", pid)
		w.transition(StateContain, "AI_ADVICE", pid, tick)
	case "LOG_ONLY":
		fmt.Printf("[WARDEN ACTION] AI Suggests Log Only. No state change.\n")
	}
}

func (w *Warden) executeAction(state SystemState, pid int) {
	switch state {
	case StateWatch:
		fmt.Printf("[WARDEN ACTION] Observe PID %d\n", pid)
	case StateAlert:
		fmt.Printf("[WARDEN ACTION] Snapshot PID %d\n", pid)
	case StateContain:
		fmt.Printf("[WARDEN ACTION] Isolate PID %d\n", pid)
	case StateRecovery:
		fmt.Printf("[WARDEN ACTION] Recover PID %d\n", pid)
	}
}

func (w *Warden) isDeescalation(target, current SystemState) bool {
	rank := map[SystemState]int{
		StateSafe:     0,
		StateWatch:    1,
		StateAlert:    2,
		StateContain:  3,
		StateRecovery: 0, // Recovery moves back towards Safe
	}
	return rank[target] < rank[current]
}

func (w *Warden) ResetBudget() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.DeescalationCount = 0
}
