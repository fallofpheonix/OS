package tcs

import (
	"testing"
	"time"
)

func TestTCSDynamicSequence(t *testing.T) {
	window := NewSlidingWindow(60 * time.Second)

	// Ingest out-of-order events
	window.AddEvent(TelemetryEvent{
		Timestamp:  time.Now(),
		SequenceID: 10,
		Payload:    []byte(`{}`),
	})

	window.AddEvent(TelemetryEvent{
		Timestamp:  time.Now().Add(1 * time.Second),
		SequenceID: 15,
		Payload:    []byte(`{}`),
	})

	window.AddEvent(TelemetryEvent{
		Timestamp:  time.Now().Add(-5 * time.Second), // Out-of-order past
		SequenceID: 5,
		Payload:    []byte(`{}`),
	})

	// Assert that dynamic min/max handles this without underflowing.
	// Minimum sequence = 5, Maximum sequence = 15. Expected range = 10.
	score := window.Evaluate()
	
	// Check that we got a valid score between 0.0 and 1.0 (no underflow/overflow bounds violation)
	if score < 0.0 || score > 1.0 {
		t.Errorf("TCS Evaluate returned out-of-bounds score: %f", score)
	}
}
