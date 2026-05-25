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

func NewDegradationMonitor(window *SlidingWindow, payloadChan chan<- ActuationPayload) *DegradationMonitor {
	return &DegradationMonitor{
		window:          window,
		payloadChan:     payloadChan,
		threshold:       0.85,
		stateDegraded:   false,
		transitionCount: 0,
	}
}

func (d *DegradationMonitor) IsDegraded() bool {
	return d.stateDegraded
}

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
