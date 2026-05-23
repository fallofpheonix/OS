package warden

import (
	"fmt"
	"sync"

	"phoenix/bus"
)

type SystemState string

const (
	StateSafe        SystemState = "SAFE"
	StateWatch       SystemState = "WATCH"
	StateAlert       SystemState = "ALERT"
	StateContain     SystemState = "CONTAIN"
	StateRecovery    SystemState = "RECOVERY"
)

// TriggerType defines the types of telemetry events that can trigger a state transition.
type TriggerType string

const (
	TriggerForkBurst    TriggerType = "PROCESS_FORK_BURST"
	TriggerNetworkBeacon TriggerType = "NETWORK_BEACON"
	TriggerMassWrite    TriggerType = "MASS_WRITE"
	TriggerReverseShell TriggerType = "REVERSE_SHELL"
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

func (w *Warden) transition(newState SystemState, trigger TriggerType, pid int, tick uint64) {
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

	fmt.Printf("[WARDEN] FSM Transition: %s -> %s (Trigger: %s, PID: %d, Tick: %d)\n", w.State, newState, trigger, pid, tick)
	w.State = newState
	w.LastTick = tick

	// Execute deterministic action based on new state
	w.executeAction(newState, pid)
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
