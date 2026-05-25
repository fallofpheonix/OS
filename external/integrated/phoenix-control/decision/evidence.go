package decision

import (
	"github.com/fallofpheonix/phoenix-logic/replay"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// EvidenceAggregator gathers and scores evidence from Replay and other forensic sources.
type EvidenceAggregator struct{}

func (e *EvidenceAggregator) Score(r *replay.Replayer, baseline []bus.TelemetryEvent) float64 {
	if len(r.Events) == 0 {
		return 0.0
	}

	// Forensic Factors (0.0 - 1.0 each)

	// 1. Hash Integrity (Binary)
	hashScore := 0.0
	if err := replay.VerifyChain(r.Events); err == nil {
		hashScore = 1.0
	}

	// 2. Divergence Factor
	divergenceScore := 1.0
	if len(baseline) > 0 {
		report := replay.DetectDivergence(r.Events, baseline)
		if len(r.Events) > 0 {
			divergenceScore = 1.0 - (float64(len(report.Points)) / float64(len(r.Events)))
			if divergenceScore < 0 {
				divergenceScore = 0
			}
		}
	}

	// 3. Time Continuity (Entropy check placeholder)
	// Check for gaps in logical ticks
	continuityScore := 1.0
	if len(r.Events) > 1 {
		gaps := 0
		for i := 1; i < len(r.Events); i++ {
			if r.Events[i].LogicalTick != r.Events[i-1].LogicalTick+1 {
				gaps++
			}
		}
		continuityScore = 1.0 - (float64(gaps) / float64(len(r.Events)))
	}

	// 4. Causal Validation (Causal Graph density placeholder)
	// High density of process relationships increases confidence
	causalScore := 0.5
	if tl, err := replay.ReconstructTimeline(r.Events); err == nil {
		if len(tl.ProcessMap) > 0 {
			causalScore = 0.8
		}
	}

	// Weighted Aggregate Score
	// Hash integrity and Causal validation are heavily weighted.
	return (hashScore * 0.4) + (causalScore * 0.3) + (divergenceScore * 0.2) + (continuityScore * 0.1)
}
