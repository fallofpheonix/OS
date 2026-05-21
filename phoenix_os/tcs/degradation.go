package tcs

import (
	"fmt"
	"log"
)

// ActuationPayload represents an evidence payload for the Ledger
type ActuationPayload struct {
	ActionID string
	CauseID  string
	TargetIP uint32
	Action   string
}

// DegradationMonitor watches TCS and triggers state changes via circuit breaker
type DegradationMonitor struct {
	window      *SlidingWindow
	payloadChan chan<- ActuationPayload
	threshold   float64
	isDegraded      bool
	transitionCount int
}

// NewDegradationMonitor creates the circuit breaker
func NewDegradationMonitor(window *SlidingWindow, payloadChan chan<- ActuationPayload) *DegradationMonitor {
	return &DegradationMonitor{
		window:      window,
		payloadChan: payloadChan,
		threshold:   0.85,
		isDegraded:      false,
		transitionCount: 0,
	}
}

// Evaluate checks the current score and triggers transitions synchronously
func (d *DegradationMonitor) Evaluate(score float64) {
	// Circuit Breaker: Trip into DEGRADED state
	if score < d.threshold && !d.isDegraded {
		d.isDegraded = true
		d.transitionCount++
		log.Printf("[TCS] WARNING: Score dropped to %.2f. ENTERING DEGRADED STATE. Suspending autonomous enforcement.", score)

		select {
		case d.payloadChan <- ActuationPayload{
			ActionID: fmt.Sprintf("STATE-DEGRADE-%d", d.transitionCount),
			CauseID:  "TCS-THRESHOLD-VIOLATION",
			TargetIP: 0,
			Action:   "ENTER_DEGRADED_MODE",
		}:
		default:
			log.Println("[TCS] CRITICAL: Ledger channel saturated during DEGRADED transition.")
		}

		return
	}

	// Circuit Breaker: Recover to NORMAL state
	if score >= d.threshold && d.isDegraded {
		d.isDegraded = false
		d.transitionCount++
		log.Printf("[TCS] RECOVERY: Score restored to %.2f. Resuming NORMAL state.", score)

		select {
		case d.payloadChan <- ActuationPayload{
			ActionID: fmt.Sprintf("STATE-RECOVER-%d", d.transitionCount),
			CauseID:  "TCS-THRESHOLD-RESTORED",
			TargetIP: 0,
			Action:   "ENTER_NORMAL_MODE",
		}:
		default:
		}
	}
}
