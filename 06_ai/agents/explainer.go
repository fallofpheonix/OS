package agents

import (
	"fmt"
	"phoenix/ledger/src"
)

// Explainer AI provides natural language reasoning for Phoenix actions.
type Explainer struct {
	ModelName string
}

// NewExplainer initializes a new explainer agent.
func NewExplainer() *Explainer {
	return &Explainer{
		ModelName: "Phoenix-Explain-v1",
	}
}

// Explain returns a reasoning string and confidence score based on ledger evidence.
func (e *Explainer) Explain(evidence ledger.Evidence) (string, float64) {
	confidence := evidence.Confidence

	// Adjust confidence based on physics (SDI) and game utility if available
	// For this mock, we use the evidence.Confidence as a baseline.
	
	reason := ""
	switch evidence.Action {
	case "FREEZE":
		reason = fmt.Sprintf("Action 'FREEZE' triggered due to High SDI (%.2f) and low Game Utility in context %s.", 
			evidence.SDI, evidence.TraceHash)
	case "ISOLATE":
		reason = fmt.Sprintf("Action 'ISOLATE' triggered as system temperature exceeded critical thresholds. SDI: %.2f.", 
			evidence.SDI)
	default:
		reason = fmt.Sprintf("Standard monitoring in state %s.", evidence.Action)
	}

	return reason, confidence
}
