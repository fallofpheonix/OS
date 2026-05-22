package warden

import (
	"fmt"

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
	State        SystemState
	outBus       *bus.Bus
	lastElevated int64 // wall unix
	lastTick     uint64 // logical tick
	dwellTicks   uint64 // hysteresis duration
}

func NewWarden(outBus *bus.Bus) *Warden {
	return &Warden{
		State:      StateNormal,
		outBus:     outBus,
		dwellTicks: 30, // Default hysteresis: 30 logical ticks
	}
}

func (w *Warden) Actuate(targetState SystemState, class ActuationClass, tcs float64, seqID int64, wallTime int64, currentTick uint64) bool {
	if w.State == targetState {
		return false
	}

	// RFC-010 GATING: Secondary TCS Check (Defense-in-Depth)
	if class >= ClassLocalIsolate && tcs < 0.85 {
		fmt.Printf("[WARDEN SHIELD] TCS Gating: Action class %d denied due to low confidence (%.2f)\n", class, tcs)
		return false
	}

	// RED TEAM MITIGATION: Actuation Hysteresis
	// If the system is in an elevated state (Suspicious/Contained), 
	// prevent de-escalation until w.dwellTicks have passed.
	isEscalation := w.isElevated(targetState, w.State)
	if !isEscalation && currentTick < w.lastTick+w.dwellTicks {
		fmt.Printf("[WARDEN SHIELD] Hysteresis Block: De-escalation from %s to %s denied (Tick: %d, Last: %d)\n", 
			w.State, targetState, currentTick, w.lastTick)
		return false
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

