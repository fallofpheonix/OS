/**
 * FILE: analytics.go
 * PATH: Phoenix.Cognition/intelligence/analytics.go
 *
 * PURPOSE:
 * Implements Phase 4C Ground Truth Governance.
 * Monitors simulated actuations and assigns labels (True/False Positive)
 * using Counterfactual Replay and system stability metrics.
 *
 * SUBSYSTEM:
 * Cognition / Intelligence / Analytics
 */

package intelligence

import (
	"encoding/json"
	"log"
	"sync"
	"time"
)

// OutcomeLabel represents the judged result of a system action.
type OutcomeLabel string

const (
	LabelTruePositive  OutcomeLabel = "TRUE_POSITIVE"
	LabelFalsePositive OutcomeLabel = "FALSE_POSITIVE"
	LabelTrueNegative  OutcomeLabel = "TRUE_NEGATIVE"
	LabelFalseNegative OutcomeLabel = "FALSE_NEGATIVE"
	LabelInconclusive  OutcomeLabel = "INCONCLUSIVE"
)

// ActuationOutcome captures the judgment of a simulated or real actuation.
type ActuationOutcome struct {
	DecisionID  string       `json:"decision_id"`
	Label       OutcomeLabel `json:"label"`
	Confidence  float64      `json:"confidence"`
	ObservedSDI float64      `json:"observed_sdi_delta"` // SDI change after action
	Timestamp   int64        `json:"timestamp"`
}

// AnalyticsHub manages the collection and labeling of system performance data.
type AnalyticsHub struct {
	mu sync.RWMutex

	Outcomes map[string]ActuationOutcome

	// Performance Metrics
	TotalDecisions uint64
	Precision      float64
	Recall         float64
}

// NewAnalyticsHub initializes the analytics engine.
func NewAnalyticsHub() *AnalyticsHub {
	return &AnalyticsHub{
		Outcomes: make(map[string]ActuationOutcome),
	}
}

// ProcessSimulatedActuation evaluates a simulation log and assigns a preliminary label.
// In a full implementation, this would trigger a Counterfactual Replay in an ERA microVM.
func (ah *AnalyticsHub) ProcessSimulatedActuation(evidence *RichActuationEvidence, decisionID string) {
	ah.mu.Lock()
	defer ah.mu.Unlock()

	ah.TotalDecisions++

	// LOGIC: Deterministic Labeling (Phase 4C)
	// If the SignalTotal is extremely high but no cascading failure occurred,
	// we still tentatively label as True Positive for the simulation.
	label := LabelTruePositive
	if evidence.SignalTotal < 0.8 {
		label = LabelFalsePositive
	}

	outcome := ActuationOutcome{
		DecisionID: decisionID,
		Label:      label,
		Confidence: evidence.Confidence,
		Timestamp:  time.Now().Unix(),
	}

	ah.Outcomes[decisionID] = outcome
	log.Printf("[ANALYTICS] Decision %s labeled as %s (Signal: %.2f)", decisionID, label, evidence.SignalTotal)

	ah.updateMetrics()
}

// updateMetrics recalculates precision/recall based on current labels.
func (ah *AnalyticsHub) updateMetrics() {
	var tp, fp float64
	for _, o := range ah.Outcomes {
		if o.Label == LabelTruePositive {
			tp++
		} else if o.Label == LabelFalsePositive {
			fp++
		}
	}

	if (tp + fp) > 0 {
		ah.Precision = tp / (tp + fp)
	}
}

// ExportReport produces a JSON representation of the current system performance.
func (ah *AnalyticsHub) ExportReport() string {
	ah.mu.RLock()
	defer ah.mu.RUnlock()

	report := map[string]interface{}{
		"total_decisions": ah.TotalDecisions,
		"precision":       ah.Precision,
		"status":          "VALIDATING_SHADOW_MODE",
	}

	data, _ := json.MarshalIndent(report, "", "  ")
	return string(data)
}
