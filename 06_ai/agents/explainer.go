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
func (e *Explainer) Explain(entry ledger.LedgerEntry) (string, float64) {
	confidence := 0.95 // Default confidence for MVP

	reason := ""
	reason = fmt.Sprintf("Action triggered by source %s. Event: %s, Cause: %s. Payload size: %d bytes.", 
		entry.Source, entry.EventID, entry.CauseID, len(entry.Payload))

	return reason, confidence
}
