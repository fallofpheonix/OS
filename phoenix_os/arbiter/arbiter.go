package arbiter

import (
	"log"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/PheonixGuard"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/monitor"
)

type Arbiter struct {
	Bus *bus.Bus
}

func NewArbiter(b *bus.Bus) *Arbiter {
	return &Arbiter{Bus: b}
}

func (a *Arbiter) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, warden.ActuationClass, bool) {
	// PHOENIX ARBITER: Strategic Decision Engine
	
	// 1. Initial Assessment based on Drift Score (L3/L6)
	target := warden.StateSafe
	class := warden.ClassLog
	
	switch {
	case score.Severity > 7.0:
		target = warden.StateCritical
		class = warden.ClassIsolate
	case score.Severity > 4.0:
		target = warden.StateSuspicious
		class = warden.ClassThrottle
	case score.Severity > 2.0:
		target = warden.StateWatch
		class = warden.ClassLog
	}

	// 2. Cross-check with TCS (L3)
	if tcsScore < 0.5 {
		log.Printf("[Arbiter] WARNING: Low TCS (%.2f). Downgrading decision.", tcsScore)
		return warden.StateSafe, warden.ClassLog, false
	}

	// 3. Authority Authorization
	// For Stage 1, we authorize everything to see the Oracle's reasoning,
	// but the Warden will still block if it violates the FSM.
	return target, class, true
}
