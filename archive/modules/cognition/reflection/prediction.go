/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package reflection implements the epistemic safety and reality drift auditing layer for PhoenixOS.
package reflection

import (
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// Prediction models an expected system outcome in reality.
// Every prediction is linked to an action ID and carries an expectation payload (usually a hash).
type Prediction struct {
	ID        string `json:"id"`
	ActionID  string `json:"action_id"`
	Expects   []byte `json:"expects"`   // Expected data hash or canonical state value
	Timestamp int64  `json:"timestamp"` // Prediction horizon (Unix epoch)
}

// ReflectionError measures the mathematical divergence between a prediction and a verified fact.
// Implements Axiom Q807 (Continuous Divergence).
type ReflectionError struct {
	PredictionID string                 `json:"prediction_id"`
	FactID       string                 `json:"fact_id"`
	Divergence   float64                `json:"divergence"` // Score [0.0 (Perfect) to 1.0 (Total Failure)]
	Confidence   ledger.ConfidenceScore `json:"confidence"`
}

// Engine manages the generation, storage, and verification of system predictions.
// It acts as the "Forecasting" component of the cognitive cycle.
type Engine struct {
	Predictions map[string]*Prediction
	Metrics     *AccuracyMetric
}

// NewEngine initializes a prediction engine with an empty prediction set and reset metrics.
func NewEngine() *Engine {
	return &Engine{
		Predictions: make(map[string]*Prediction),
		Metrics:     &AccuracyMetric{},
	}
}

// Predict records a new speculative outcome for a given system action.
// Inputs: id (Prediction ID), actionID (Link to actuation), expects (Expected payload), horizon (Expiration).
// Complexity: O(1) time / O(1) space.
func (e *Engine) Predict(id, actionID string, expects []byte, horizon int64) {
	e.Predictions[id] = &Prediction{
		ID:        id,
		ActionID:  actionID,
		Expects:   expects,
		Timestamp: horizon,
	}
}

// Verify matches a verified reality Fact against a previously recorded Prediction.
// It calculates continuous divergence using a Hamming distance metric (Axiom Q807).
// Inputs: pID (Prediction ID), fID (Fact ID), actual (Observed reality payload).
// Returns: (*ReflectionError) containing the divergence score.
// Complexity: O(L) where L is the length of the expected payload.
func (e *Engine) Verify(pID, fID string, actual []byte) *ReflectionError {
	p, ok := e.Predictions[pID]
	if !ok {
		return nil
	}

	e.Metrics.TotalPredictions++

	// Continuous Divergence (Q807): Hamming Distance normalized (0.0 to 1.0)
	divergence := calculateDivergence(p.Expects, actual)

	if divergence == 0.0 {
		e.Metrics.CorrectCount++
	}
	e.Metrics.AggregateError += divergence

	return &ReflectionError{
		PredictionID: pID,
		FactID:       fID,
		Divergence:   divergence,
	}
}

// calculateDivergence performs a normalized Hamming distance calculation between two byte slices.
// Complexity: O(N) where N is max(len(expected), len(actual)).
func calculateDivergence(expected, actual []byte) float64 {
	if len(expected) == 0 || len(actual) == 0 {
		if len(expected) == len(actual) {
			return 0.0
		}
		return 1.0
	}

	diff := 0
	length := len(expected)
	if len(actual) < length {
		length = len(actual)
	}

	for i := 0; i < length; i++ {
		if expected[i] != actual[i] {
			diff++
		}
	}

	// Calculate absolute length difference.
	lenDiff := len(expected) - len(actual)
	if lenDiff < 0 {
		lenDiff = -lenDiff
	}
	diff += lenDiff

	maxLen := len(expected)
	if len(actual) > maxLen {
		maxLen = len(actual)
	}

	return float64(diff) / float64(maxLen)
}

// AccuracyMetric tracks the long-term aggregate performance of the prediction engine.
type AccuracyMetric struct {
	TotalPredictions uint64  `json:"total_predictions"`
	CorrectCount     uint64  `json:"correct_count"`
	AggregateError   float64 `json:"aggregate_error"`
}

// CalculateAccuracy returns the current prediction accuracy ratio [0.0, 1.0].
func (am *AccuracyMetric) CalculateAccuracy() float64 {
	if am.TotalPredictions == 0 {
		return 1.0
	}
	return float64(am.CorrectCount) / float64(am.TotalPredictions)
}

// ReconstructFromLedger rebuilds prediction metrics by replaying historical events.
// Implements the Prediction Replay Theorem (Q668).
// Inputs: events ([]*ledger.Event) - Chronological event stream.
// Complexity: O(N) where N is number of prediction events.
func (am *AccuracyMetric) ReconstructFromLedger(events []*ledger.Event) error {
	for _, e := range events {
		if e.Type != ledger.EventPrediction {
			continue
		}
		am.TotalPredictions++
	}
	return nil
}
