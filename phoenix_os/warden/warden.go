package warden

import (
	"fmt"
	"sync"

	"phoenix/bus"
)

type SystemState string

const (
	StateNormal     SystemState = "NORMAL"
	StateSuspicious SystemState = "SUSPICIOUS"
	StateContained  SystemState = "CONTAINED"
	StateRecovery   SystemState = "RECOVERY"
)

// ActuationClass represents the risk tier of an action (RFC-010)
type ActuationClass int

const (
	ClassObserve         ActuationClass = 0
	ClassLog             ActuationClass = 1
	ClassThrottle        ActuationClass = 2
	ClassLocalIsolate    ActuationClass = 3
	ClassClusterIsolate  ActuationClass = 4
	ClassKernelEmergency ActuationClass = 5
)

type Warden struct {
	mu                sync.Mutex
	State             SystemState
	outBus            *bus.Bus
	lastElevated      int64  // wall unix
	lastTick          uint64 // logical tick
	dwellTicks        uint64 // hysteresis duration
	cooldownTicks     uint64 // stabilization period lock
	recoveryBudget    int    // max de-escalation actions authorized
	deescalationCount int    // tracker for de-escalation occurrences
}

func NewWarden(outBus *bus.Bus) *Warden {
	return &Warden{
		State:          StateNormal,
		outBus:         outBus,
		dwellTicks:     30, // Default hysteresis: 30 logical ticks
		cooldownTicks:  10, // Stabilization cooldown: 10 logical ticks
		recoveryBudget: 3,  // Maximum of 3 automatic recoveries allowed before operator reset
	}
}

func (w *Warden) Actuate(targetState SystemState, class ActuationClass, tcs float64, seqID int64, wallTime int64, currentTick uint64) bool {
	w.mu.Lock()
	defer w.mu.Unlock()

	if w.State == targetState {
		return false
	}

	// RFC-010 GATING: Secondary TCS Check (Defense-in-Depth)
	if class >= ClassLocalIsolate && tcs < 0.85 {
		fmt.Printf("[WARDEN SHIELD] TCS Gating: Action class %d denied due to low confidence (%.2f)\n", class, tcs)
		return false
	}

	isEscalation := w.isElevated(targetState, w.State)

	// 1. Stabilization Cooldown check
	if currentTick < w.lastTick+w.cooldownTicks {
		// Only allow bypass of cooldown for high-severity escalations (ClassLocalIsolate or higher)
		if !isEscalation || class < ClassLocalIsolate {
			fmt.Printf("[WARDEN SHIELD] Cooldown Lock: Transition from %s to %s denied (Tick: %d, Last: %d, Cooldown: %d)\n",
				w.State, targetState, currentTick, w.lastTick, w.cooldownTicks)
			return false
		}
	}

	// 2. Dwell limits & Recovery budget checks on de-escalation
	if !isEscalation {
		if currentTick < w.lastTick+w.dwellTicks {
			fmt.Printf("[WARDEN SHIELD] Hysteresis Block: De-escalation from %s to %s denied (Tick: %d, Last: %d, Dwell: %d)\n",
				w.State, targetState, currentTick, w.lastTick, w.dwellTicks)
			return false
		}

		if w.deescalationCount >= w.recoveryBudget {
			fmt.Printf("[WARDEN SHIELD] Recovery Budget Exhausted (%d/%d): De-escalation to %s blocked\n",
				w.deescalationCount, w.recoveryBudget, targetState)
			return false
		}
		w.deescalationCount++
	} else {
		// Reset budget if system has been in SAFE / NORMAL state for a long duration (> 100 ticks)
		if w.State == StateNormal && currentTick > w.lastTick+100 {
			w.deescalationCount = 0
		}
	}

	oldState := w.State
	w.State = targetState
	w.lastElevated = wallTime
	w.lastTick = currentTick

	fmt.Printf("[WARDEN] Actuation: %s -> %s (Event: %d, Tick: %d)\n", oldState, w.State, seqID, currentTick)
	
	actionPayload := []byte(fmt.Sprintf(`{"action":"transition","state":"%s"}`, w.State))
	w.outBus.Publish("warden.action", bus.TelemetryEvent{
		SeqID:        seqID,
		WallTimeUnix: wallTime,
		Source:       "phoenix.warden",
		EventType:    "fsm.transition",
		Severity:     1.0,
		Payload:      actionPayload,
	})

	return true
}

func (w *Warden) isElevated(target, current SystemState) bool {
	scoreMap := map[SystemState]int{
		StateNormal:     0,
		StateRecovery:   1,
		StateSuspicious: 2,
		StateContained:  3,
	}
	return scoreMap[target] > scoreMap[current]
}

func (w *Warden) ResetBudget() {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.deescalationCount = 0
}
