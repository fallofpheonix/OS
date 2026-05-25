package decision

import "math"

// ConfidenceEngine refines the confidence score using statistical and historical data.
type ConfidenceEngine struct{}

func (c *ConfidenceEngine) Refine(baseConfidence float64, systemRisk float64) float64 {
	// Adjust confidence based on overall system risk.
	// High system risk increases the threshold for certainty.
	if systemRisk > 0.8 {
		return baseConfidence * 0.8
	}

	// Apply a sigmoid-like smoothing
	return 1.0 / (1.0 + math.Exp(-10.0*(baseConfidence-0.5)))
}
