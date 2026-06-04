/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 5b — DEGRADATION DETECTION (Layer 3.5)
//
// The DegradationMonitor tracks TCS score transitions to detect when
// the telemetry stream degrades below acceptable quality.
//
// WORKFLOW:
//   TCSFeature.EvaluateDegradation(score) → DegradationMonitor.Evaluate(score)
//     → If score drops below 0.85: state → DEGRADED
//     → If score recovers above 0.85: state → NORMAL
//     → Transition recorded to Ledger for audit trail
//
// THRESHOLD: 0.85 (configurable). Below this, the system is considered
// to have degraded telemetry quality, which may trigger more conservative
// behavior in the Arbiter.
//
// TRANSITION COUNT: Tracked for metrics and health monitoring.
// Rapid transitions (flapping) may indicate an unstable telemetry source.
// =========================================================================
package tcs

import (
	"log"
)

type ActuationPayload struct {
	ActionID string
	CauseID  string
	TargetIP uint32
	Action   string
}

type DegradationMonitor struct {
	window          *SlidingWindow
	payloadChan     chan<- ActuationPayload
	threshold       float64
	stateDegraded   bool
	transitionCount int
}

// NewDegradationMonitor creates the telemetry quality tracker.
// Called once during TCSFeature initialization with a 0.85 threshold.
func NewDegradationMonitor(window *SlidingWindow, payloadChan chan<- ActuationPayload) *DegradationMonitor {
	return &DegradationMonitor{
		window:          window,
		payloadChan:     payloadChan,
		threshold:       0.85,
		stateDegraded:   false,
		transitionCount: 0,
	}
}

// IsDegraded returns whether the telemetry stream is currently degraded.
// Used by the Arbiter to adjust decision thresholds during degraded mode.
func (d *DegradationMonitor) IsDegraded() bool {
	return d.stateDegraded
}

// Evaluate checks the TCS score against the degradation threshold.
// Called from TCSFeature.EvaluateDegradation() after every TCS calculation.
//
// WORKFLOW: TCSFeature.EvaluateDegradation(score) → DegradationMonitor.Evaluate(score)
//
//	→ If score < 0.85 AND not already degraded:
//	  → Set stateDegraded = true
//	  → Increment transitionCount
//	  → Log warning: "ENTERING DEGRADED STATE"
//	→ If score >= 0.85 AND currently degraded:
//	  → Set stateDegraded = false
//	  → Increment transitionCount
//	  → Log recovery: "Resuming NORMAL state"
//
// SIDE EFFECTS: Updates stateDegraded and transitionCount.
// No lock required — called from single-threaded OrchestrateTick().
func (d *DegradationMonitor) Evaluate(score float64) {
	if score < d.threshold && !d.stateDegraded {
		d.stateDegraded = true
		d.transitionCount++
		log.Printf("[TCS] WARNING: Score dropped to %.2f. ENTERING DEGRADED STATE.", score)
		return
	}
	if score >= d.threshold && d.stateDegraded {
		d.stateDegraded = false
		d.transitionCount++
		log.Printf("[TCS] RECOVERY: Score restored to %.2f. Resuming NORMAL state.", score)
	}
}
