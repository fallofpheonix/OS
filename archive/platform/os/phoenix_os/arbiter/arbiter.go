/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: arbiter.go
 *
 * Purpose:
 * Implements the strategic decision engine for PhoenixOS. Reconciles drift 
 * scores and trust values to determine system state transitions.
 *
 * Subsystem:
 * Terminus-Arbiter
 *
 * Dependencies:
 * - PhoenixCore/bus
 * - PhoenixGuard
 * - phoenix_os/monitor
 *
 * Security:
 * - Critical: Determines the actuation class (Log/Throttle/Isolate).
 * - Risk: Miscalibration could lead to false positives (DoS) or false negatives.
 *
 * Performance:
 * - O(1) decision logic. Minimal overhead.
 *
 * @labels arbiter, strategic-decision, phase-2-complete
 */
package arbiter

import (
	"log"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/assurance/security"
	"github.com/fallofpheonix/phoenix/platform/os/phoenix_os/monitor"
)

/*
 * @class Arbiter
 * @description The strategic decision hub of the OS.
 * @responsibilities Reconciling security metrics into actionable state transitions.
 */
type Arbiter struct {
	Bus *bus.Bus
}

/**
 * NewArbiter initializes a new strategic arbiter.
 * @param b The system event bus.
 * @return *Arbiter A pointer to the initialized arbiter.
 */
func NewArbiter(b *bus.Bus) *Arbiter {
	return &Arbiter{Bus: b}
}

/**
 * Evaluate analyzes scores and determines the required system state.
 * @param score The drift score from the monitor service.
 * @param tcsScore The trusted computing score.
 * @return warden.SystemState The target state.
 * @return warden.ActuationClass The actuation severity.
 * @return bool Whether the decision is authorized.
 * @complexity O(1)
 */
func (a *Arbiter) Evaluate(score monitor.DriftScore, tcsScore float64) (warden.SystemState, warden.ActuationClass, bool) {
	// PHOENIX ARBITER: Strategic Decision Engine

	
	// 1. Initial Assessment based on Drift Score (L3/L6)
	target := warden.StateSafe
	class := warden.ClassLog
	
	switch {
	case score.ZScore > 7.0:
		target = warden.StateCritical
		class = warden.ClassIsolate
	case score.ZScore > 4.0:
		target = warden.StateSuspicious
		class = warden.ClassThrottle
	case score.ZScore > 2.0:
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
