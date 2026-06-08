// Package tcs implements Temporal Confidence Scoring (Layer 3).
// Domain Logic: Measures the reliability of telemetry streams by analyzing packet loss and jitter.
// Responsibility: Provides a confidence metric (0.0 to 1.0) to inform system escalation and degradation decisions.
package tcs

import (
	"sync"
	"time"
)

// TelemetryEvent represents a raw signal from the kernel or network.
// Concurrency: Instances are immutable once created and thread-safe for reading.
// State Management: Encapsulates temporal and sequence metadata for a single telemetry packet.
type TelemetryEvent struct {
	Timestamp  time.Time
	SequenceID uint64
	Payload    []byte
	JitterNS   uint64 // Delta from expected arrival
}

// SlidingWindow bounds telemetry memory to prevent unbounded growth.
// Concurrency: Thread-safe via sync.RWMutex.
// State Management: Maintains a time-bounded window of events and tracks cumulative loss/jitter metrics.
type SlidingWindow struct {
	mu          sync.RWMutex
	events      []TelemetryEvent
	WindowSize  time.Duration
	lastSeqID   uint64
	droppedPkts uint64
	latestTime  time.Time
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewSlidingWindow creates the time-bounded telemetry window.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewSlidingWindow(windowSize time.Duration) *SlidingWindow {
	return &SlidingWindow{
		WindowSize: windowSize,
	}
}

// LABEL: [MUTABLE] [TIME_BOUNDED] [STABLE]
// AddEvent injects a new telemetry payload into the sliding window.
// I/O: None.
// Side Effects: Modifies internal state (events, lastSeqID, droppedPkts, latestTime). Triggers pruning.
// Complexity: O(N) where N is the number of events to prune, typically O(1) amortized.
func (w *SlidingWindow) AddEvent(e TelemetryEvent) {
	w.mu.Lock()
	defer w.mu.Unlock()

	// SECURITY SHIELD: Logical Time Lock
	// Prevent "Time Warp" attacks by ignoring events from the far future
	if !w.latestTime.IsZero() && e.Timestamp.After(w.latestTime.Add(w.WindowSize)) {
		return
	}

	// SECURITY SHIELD: Gap Capping
	// Prevent massive sequence jumps from permanently destroying confidence.
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

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Evaluate calculates the Telemetry Confidence Score (0.0 to 1.0).
// I/O: None.
// Side Effects: None.
// Complexity: O(N) where N is the number of events in the window.
func (w *SlidingWindow) Evaluate() float64 {
	w.mu.RLock()
	defer w.mu.RUnlock()

	if len(w.events) == 0 {
		return 0.0
	}

	var minSeq, maxSeq uint64
	first := true
	for _, e := range w.events {
		if first {
			minSeq = e.SequenceID
			maxSeq = e.SequenceID
			first = false
		} else {
			if e.SequenceID < minSeq {
				minSeq = e.SequenceID
			}
			if e.SequenceID > maxSeq {
				maxSeq = e.SequenceID
			}
		}
	}

	totalExpected := maxSeq - minSeq
	var lossRate float64
	if totalExpected > 0 {
		lossRate = float64(w.droppedPkts) / float64(totalExpected)
	}
	if lossRate > 1.0 {
		lossRate = 1.0
	}

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

	// Weighted TCS: 70% loss, 30% jitter
	w1, w2 := 0.7, 0.3
	tcs := (w1 * (1.0 - lossRate)) + (w2 * (1.0 - jNorm))

	return tcs
}

// prune removes events older than the window size from the sliding window.
// This prevents unbounded memory growth and ensures the TCS reflects
// only recent telemetry quality.
//
// WORKFLOW: Called at the end of AddEvent() after appending the new event.
// Finds the first event within the cutoff time and discards all earlier events.
// Also decays historical loss data by halving droppedPkts.
//
// MEMORY MANAGEMENT: Without pruning, the events slice would grow
// indefinitely under high event rates. Pruning bounds memory to
// O(window_size) events.
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
		w.droppedPkts /= 2 // Decay historical loss
	}
}
