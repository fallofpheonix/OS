package replay

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// Replayer orchestrates the forensic reconstruction and validation of telemetry streams.
type Replayer struct {
	Events []bus.TelemetryEvent
}

// NewReplayer creates a new Replayer instance with a set of events.
func NewReplayer(events []bus.TelemetryEvent) *Replayer {
	return &Replayer{
		Events: events,
	}
}

// Execute performs the full replay pipeline: Verify -> Reconstruct -> Report.
func (r *Replayer) Execute(baseline []bus.TelemetryEvent) error {
	fmt.Printf("[REPLAY] Starting replay of %d events...\n", len(r.Events))

	// 1. Verify Hash Chain Integrity
	fmt.Println("[REPLAY] Step 1: Verifying Cryptographic Hash Chain...")
	if err := VerifyChain(r.Events); err != nil {
		return fmt.Errorf("integrity check failed: %v", err)
	}
	fmt.Println("[REPLAY] Integrity Verified.")

	// 2. Reconstruct Causal Timeline
	fmt.Println("[REPLAY] Step 2: Reconstructing Causal Timeline...")
	tl, err := ReconstructTimeline(r.Events)
	if err != nil {
		return fmt.Errorf("timeline reconstruction failed: %v", err)
	}
	fmt.Printf("[REPLAY] Timeline Reconstructed (%d processes found).\n", len(tl.ProcessMap))
	tl.PrintTimeline()

	// 3. Detect Divergence against Baseline (if provided)
	if len(baseline) > 0 {
		fmt.Println("[REPLAY] Step 3: Detecting Divergence against Baseline...")
		report := DetectDivergence(r.Events, baseline)
		report.PrintReport()
		if report.HasDiverged {
			return fmt.Errorf("replay diverged from baseline at %d points", len(report.Points))
		}
	} else {
		fmt.Println("[REPLAY] Step 3: No baseline provided, skipping divergence check.")
	}

	fmt.Println("[REPLAY] Replay completed successfully.")
	return nil
}

// ReplayFidelity calculates the match percentage between two runs.
func ReplayFidelity(actual, expected []bus.TelemetryEvent) float64 {
	report := DetectDivergence(actual, expected)
	if !report.HasDiverged {
		return 100.0
	}

	total := len(expected)
	if total == 0 {
		return 0.0
	}

	// Simple fidelity: (total - divergent points) / total
	// Note: Stream length mismatch is treated as divergence.
	matches := total - len(report.Points)
	if matches < 0 {
		matches = 0
	}

	return (float64(matches) / float64(total)) * 100.0
}
