package tcs

import (
	"sync"
	"time"
)

// TelemetryEvent represents a raw signal from the kernel or network
type TelemetryEvent struct {
	Timestamp  time.Time
	SequenceID uint64
	Payload    []byte
	JitterNS   uint64 // Delta from expected arrival
}

// SlidingWindow bounds telemetry memory to prevent unbounded growth
type SlidingWindow struct {
	mu          sync.RWMutex
	events      []TelemetryEvent
	WindowSize  time.Duration
	lastSeqID   uint64
	droppedPkts uint64
	latestTime  time.Time
}

// NewSlidingWindow creates a bounded telemetry window
func NewSlidingWindow(windowSize time.Duration) *SlidingWindow {
	return &SlidingWindow{
		WindowSize: windowSize,
	}
}

// AddEvent injects a new telemetry payload into the window
func (w *SlidingWindow) AddEvent(e TelemetryEvent) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// 1. SECURITY SHIELD: Logical Time Lock
	// Prevent "Time Warp" attacks by ignoring events from the far future
	// relative to current system time.
	if !w.latestTime.IsZero() && e.Timestamp.After(w.latestTime.Add(w.WindowSize)) {
		// Ignore events more than 1 window-size into the future
		return
	}

	// 2. SECURITY SHIELD: Gap Capping
	// Prevent massive sequence jumps from permanently destroying confidence.
	// Cap perceived loss at 100 packets per event.
	if w.lastSeqID > 0 && e.SequenceID > w.lastSeqID+1 {
		gap := e.SequenceID - w.lastSeqID - 1
		if gap > 100 {
			gap = 100
		}
		w.droppedPkts += gap
	}
	w.lastSeqID = e.SequenceID

	if e.Timestamp.After(w.latestTime) {
		w.latestTime = e.Timestamp
	}
	w.events = append(w.events, e)
	w.prune(w.latestTime)
}

// Evaluate calculates the Telemetry Confidence Score (0.0 to 1.0)
// Formula: T_C = w1*(1 - P_loss) + w2*(1 - J_norm)
func (w *SlidingWindow) Evaluate() float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if len(w.events) == 0 {
		return 0.0 // No data = zero confidence
	}

	// 1. Loss Rate (P_loss)
	totalExpected := w.events[len(w.events)-1].SequenceID - w.events[0].SequenceID
	var lossRate float64
	if totalExpected > 0 {
		lossRate = float64(w.droppedPkts) / float64(totalExpected)
	}
	if lossRate > 1.0 {
		lossRate = 1.0
	}

	// 2. Normalized Jitter (J_norm)
	var totalJitter uint64
	for _, e := range w.events {
		totalJitter += e.JitterNS
	}
	avgJitter := float64(totalJitter) / float64(len(w.events))

	// Normalize: 50ms jitter = fully degraded
	jNorm := avgJitter / 50_000_000.0
	if jNorm > 1.0 {
		jNorm = 1.0
	}

	// 3. Weighted TCS: 70% loss, 30% jitter
	w1, w2 := 0.7, 0.3
	tcs := (w1 * (1.0 - lossRate)) + (w2 * (1.0 - jNorm))

	return tcs
}

func (w *SlidingWindow) prune(now time.Time) {
	if len(w.events) == 0 {
		return
	}

	cutoff := now.Add(-w.WindowSize)
	splitIdx := 0
	foundValid := false

	for i, e := range w.events {
		if e.Timestamp.After(cutoff) {
			splitIdx = i
			foundValid = true
			break
		}
	}

	if !foundValid {
		w.events = nil
		w.droppedPkts = 0
		return
	}

	if splitIdx > 0 {
		w.events = w.events[splitIdx:]
		w.droppedPkts = w.droppedPkts / 2 // Decay historical loss
	}
}
