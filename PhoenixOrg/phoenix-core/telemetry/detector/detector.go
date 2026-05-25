package detector

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/telemetry/events"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/telemetry/process_lineage"
)

// ThreatResult represents the outcome of a detection analysis.
type ThreatResult struct {
	ImportanceScore float64 `json:"importance_score"`
	IsThreat        bool    `json:"is_threat"`
	Reason          string  `json:"reason"`
}

// Detector implements the rule-based detection logic for Stage 4.
type Detector struct {
	EntropyThreshold float64
	ScoreThreshold   float64
}

// NewDetector initializes a new rule-based detector.
func NewDetector() *Detector {
	return &Detector{
		EntropyThreshold: 7.0,
		ScoreThreshold:   0.65,
	}
}

// Analyze evaluates an event and its context in the lineage graph.
func (d *Detector) Analyze(evt events.Event, g *lineage.LineageGraph) ThreatResult {
	score := 0.0
	reason := ""

	// Rule 1: High Entropy (Potential Encryption/Ransomware)
	if entropy, ok := evt.Payload["entropy_score"].(float64); ok {
		if entropy >= d.EntropyThreshold {
			score += 0.8 // Immediate trigger above threshold
			reason = "High entropy detected"
		}
	}

	// Rule 2: Suspicious Process (e.g., 'gpg' or 'sh' doing lots of writes)
	if evt.EventType == "write" || evt.EventType == "rename" {
		if evt.Comm == "gpg" || evt.Comm == "sh" {
			score += 0.1
			if reason != "" {
				reason += ", "
			}
			reason += "Suspicious process activity"
		}
	}

	// Rule 3: Lineage Depth (Deep process trees can be suspicious)
	ancestors := g.GetAncestors(evt.PID)
	if len(ancestors) > 5 {
		score += 0.1
	}

	return ThreatResult{
		ImportanceScore: score,
		IsThreat:        score >= d.ScoreThreshold,
		Reason:          reason,
	}
}

// FastPath provides a low-latency check for critical threats.
func (d *Detector) FastPath(evt events.Event) bool {
	if entropy, ok := evt.Payload["entropy_score"].(float64); ok {
		return entropy > 8.0 // Immediate trigger for extremely high entropy
	}
	return false
}
