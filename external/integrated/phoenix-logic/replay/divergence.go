package replay

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// DivergencePoint identifies where two event streams differ.
type DivergencePoint struct {
	Index         int
	ActualEvent   bus.TelemetryEvent
	ExpectedEvent bus.TelemetryEvent
	Reason        string
}

// DivergenceReport summarizes the differences between two replay runs.
type DivergenceReport struct {
	HasDiverged bool
	Points      []DivergencePoint
}

// DetectDivergence compares two event streams and returns a report of differences.
func DetectDivergence(actual, expected []bus.TelemetryEvent) *DivergenceReport {
	report := &DivergenceReport{
		HasDiverged: false,
		Points:      []DivergencePoint{},
	}

	lenActual := len(actual)
	lenExpected := len(expected)
	minLen := lenActual
	if lenExpected < minLen {
		minLen = lenExpected
	}

	for i := 0; i < minLen; i++ {
		a := actual[i]
		e := expected[i]

		// Check for core field mismatches
		if a.EventType != e.EventType {
			report.addPoint(i, a, e, fmt.Sprintf("EventType mismatch: actual %s, expected %s", a.EventType, e.EventType))
		} else if a.PID != e.PID {
			report.addPoint(i, a, e, fmt.Sprintf("PID mismatch: actual %d, expected %d", a.PID, e.PID))
		} else if a.Hash != e.Hash {
			report.addPoint(i, a, e, "Hash mismatch (Integrity violation or payload change)")
		}
		// Note: We don't necessarily check MonotonicNs for exact match as small jitter might exist,
		// but sequence order should be identical in a deterministic replay.
	}

	if lenActual != lenExpected {
		report.HasDiverged = true
		// We could add a special point for length mismatch
		reason := fmt.Sprintf("Stream length mismatch: actual %d, expected %d", lenActual, lenExpected)
		var lastA, lastE bus.TelemetryEvent
		if lenActual > 0 {
			lastA = actual[lenActual-1]
		}
		if lenExpected > 0 {
			lastE = expected[lenExpected-1]
		}
		report.Points = append(report.Points, DivergencePoint{
			Index:         minLen,
			ActualEvent:   lastA,
			ExpectedEvent: lastE,
			Reason:        reason,
		})
	}

	return report
}

func (r *DivergenceReport) addPoint(idx int, a, e bus.TelemetryEvent, reason string) {
	r.HasDiverged = true
	r.Points = append(r.Points, DivergencePoint{
		Index:         idx,
		ActualEvent:   a,
		ExpectedEvent: e,
		Reason:        reason,
	})
}

// PrintReport prints the divergence findings to stdout.
func (r *DivergenceReport) PrintReport() {
	if !r.HasDiverged {
		fmt.Println("[DIVERGENCE] No divergence detected. Replay matches baseline.")
		return
	}

	fmt.Printf("[DIVERGENCE] Found %d divergence points:\n", len(r.Points))
	for _, p := range r.Points {
		fmt.Printf("  - Point %d: %s\n", p.Index, p.Reason)
		fmt.Printf("    Actual:   Seq %d, Type %s, PID %d\n", p.ActualEvent.SeqID, p.ActualEvent.EventType, p.ActualEvent.PID)
		fmt.Printf("    Expected: Seq %d, Type %s, PID %d\n", p.ExpectedEvent.SeqID, p.ExpectedEvent.EventType, p.ExpectedEvent.PID)
	}
}
