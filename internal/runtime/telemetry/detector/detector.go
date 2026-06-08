/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 3b — THREAT DETECTION (Layer 1.5)
//
// The Detector implements rule-based threat detection for telemetry events.
// It evaluates events against a set of rules and produces ThreatResults
// indicating whether an event is a potential threat.
//
// WORKFLOW:
//   Monitor produces DriftScore → Detector.Analyze(event, lineageGraph)
//     → Rule 1: High entropy (> 7.0) → score += 0.8 (ransomware indicator)
//     → Rule 2: Suspicious process (gpg/sh doing writes) → score += 0.1
//     → Rule 3: Deep lineage (> 5 ancestors) → score += 0.1
//     → If total score >= 0.65: IsThreat = true
//   → ThreatResult fed to Arbiter for decision-making
//
// FAST PATH: Detector.FastPath() provides O(1) check for critical threats.
// If entropy > 8.0, the event is immediately flagged as a threat.
// This bypasses the full analysis pipeline for urgent responses.
// =========================================================================
package detector

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/telemetry/events"

	lineage "github.com/fallofpheonix/phoenix/foundation/observability/engine/process_lineage"
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
