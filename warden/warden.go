package warden

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-os/telemetry/bus"
)

type SystemState string

const (
	StateNormal     SystemState = "NORMAL"
	StateSuspicious SystemState = "SUSPICIOUS"
	StateContained  SystemState = "CONTAINED"
)

type Warden struct {
	State  SystemState
	outBus *bus.Bus
}

func NewWarden(outBus *bus.Bus) *Warden {
	return &Warden{State: StateNormal, outBus: outBus}
}

func (w *Warden) Actuate(targetState SystemState, seqID uint64, wallTime int64) bool {
	if w.State == targetState {
		return false
	}
	w.State = targetState
	fmt.Printf("[WARDEN] Actuation to %s (Event %d)
", w.State, seqID)
	
	actionPayload := []byte(fmt.Sprintf(`{"action":"transition","state":"%s"}`, w.State))
	w.outBus.Publish("warden.action", bus.TelemetryEvent{
		SeqID:        seqID,
		WallTimeUnix: wallTime,
		Source:       "phoenix.warden",
		EventType:    "fsm.transition",
		Payload:      actionPayload,
	})
	return true
}
