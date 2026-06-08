/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package reasoning implements the strategic inference and explanation layers for PhoenixOS.
package reasoning

import (
	"context"

	"github.com/fallofpheonix/phoenix/cognition/knowledge"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// InferenceRequest encapsulates the goal, horizon, and context required for an AI reasoning cycle.
// It is designed to be model-agnostic, allowing translation to various LLM prompts.
type InferenceRequest struct {
	Goal      string                `json:"goal"`      // Primary objective of the reasoning cycle
	Horizon   int64                 `json:"horizon"`   // Maximum logical tick or time for prediction
	Beliefs   []*knowledge.Belief   `json:"beliefs"`   // Grounding beliefs relevant to the request
	Knowledge []*knowledge.Edge     `json:"knowledge"` // Causal graph edges for context
	Facts     []*ledger.FactPayload `json:"facts"`     // Direct evidence from the ledger
}

// InferenceResponse represents the structured output of a reasoning cycle.
type InferenceResponse struct {
	Logic      string  `json:"logic"`      // The chain of thought or reasoning steps
	Proposal   string  `json:"proposal"`   // The recommended system command or state change
	Confidence float64 `json:"confidence"` // Model's confidence in the proposal [0.0, 1.0]
}

// Provider defines the interface for interchangeable reasoning engines (LLMs).
// It ensures Cognitive Sovereignty by decoupling the OS world model from
// specific external or local reasoning backends.
type Provider interface {
	// Name returns the identifier of the reasoning provider.
	Name() string
	// Reason performs an inference cycle based on the provided request.
	// Inputs: ctx (context.Context), req (*InferenceRequest).
	// Returns: (*InferenceResponse, error).
	Reason(ctx context.Context, req *InferenceRequest) (*InferenceResponse, error)
}
