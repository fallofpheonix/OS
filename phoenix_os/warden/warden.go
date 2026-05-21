package warden

import (
	"fmt"

	"github.com/fallofpheonix/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix_os/monitor"
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
	SuspThresh   float64
	ContThresh   float64
	RecovThresh  float64
	DwellTime    int64 // seconds
	inCh         chan monitor.DriftScore
	outBus       *bus.Bus
	lastElevated int64 // unix seconds
}

func NewWarden(inCh chan monitor.DriftScore, outBus *bus.Bus) *Warden {
	return &Warden{
		State:       StateNormal,
		SuspThresh:  2.0,
		ContThresh:  4.0,
		RecovThresh: 1.0,
		DwellTime:   30, // 30 seconds
		inCh:        inCh,
		outBus:      outBus,
	}
}

func (w *Warden) Start() {
	go func() {
		for score := range w.inCh {
			w.Evaluate(score)
		}
	}()
}

func (w *Warden) Evaluate(score monitor.DriftScore) {
	now := score.WallTimeUnix
	z := score.ZScore
	oldState := w.State

	// Escalate immediately
	if z >= w.ContThresh {
		w.State = StateContained
		w.lastElevated = now
	} else if z >= w.SuspThresh && w.State == StateNormal {
		w.State = StateSuspicious
		w.lastElevated = now
	}

	// De-escalate with hysteresis (dwell time)
	if w.State == StateContained && z < w.RecovThresh {
		if now - w.lastElevated > w.DwellTime {
			w.State = StateRecovery
			w.lastElevated = now
		}
	} else if w.State == StateRecovery && z < w.SuspThresh {
		if now - w.lastElevated > w.DwellTime {
			w.State = StateNormal
		}
	} else if w.State == StateSuspicious && z < w.SuspThresh {
		if now - w.lastElevated > w.DwellTime {
			w.State = StateNormal
		}
	} else if z >= w.SuspThresh {
		w.lastElevated = now
	}

	if oldState != w.State {
		fmt.Printf("[WARDEN] FSM Transition: %s -> %s (Z-Score: %.2f)\n", oldState, w.State, z)
		actionPayload := []byte(fmt.Sprintf(`{"action":"transition","state":"%s"}`, w.State))
		w.outBus.Publish("warden.action", bus.TelemetryEvent{
			SeqID:     score.EventID,
			Source:    "phoenix.warden",
			EventType: "fsm.transition",
			Severity:  1.0,
			Payload:   actionPayload,
		})
	}
}
