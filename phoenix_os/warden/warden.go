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

type Warden struct {
	State        SystemState
	outBus       *bus.Bus
	lastElevated int64 // unix seconds
}

func NewWarden(outBus *bus.Bus) *Warden {
	return &Warden{
		State:  StateNormal,
		outBus: outBus,
	}
}

func (w *Warden) Actuate(targetState SystemState, seqID int64, wallTime int64) bool {
	if w.State == targetState {
		return false
	}

	oldState := w.State
	w.State = targetState
	w.lastElevated = wallTime

	fmt.Printf("[WARDEN] Actuation: %s -> %s (Event: %d)\n", oldState, w.State, seqID)
	
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
