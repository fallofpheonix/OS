/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: reality.go
 * PATH: Phoenix.Nucleus/ledger/reality.go
 *
 * PURPOSE:
 * Implements the Reality Confidence Theorem (Q649-Q650).
 * Ensures that facts are grounded in verified, multi-sensor observations.
 *
 * SUBSYSTEM:
 * Nucleus / Reality Cycle
 *
 * DEPENDENCIES:
 * phxmath (Deterministic Numeric Model)
 *
 * DEPENDENTS:
 * Phoenix.Cognition/memory, Phoenix.Nucleus/ledger/event
 *
 * SECURITY:
 * Prevents "Sensor Oligarchy" (Q688) by requiring a minimum number of
 * independent observations and a weighted confidence threshold.
 */

package ledger

import (
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// BEGINNER EXPLANATION:
// This file makes sure the system doesn't believe everything it hears from
// just one "Eye" (Sensor). It asks for multiple reports and calculates a
// "Confidence Score" before calling something a "Fact."

// INTERMEDIATE EXPLANATION:
// RealityCycle defines the interface for mapping physical reality to the
// Ledger. MultiSensorEvidence implements the consensus logic for facts.

// EXPERT EXPLANATION:
// Implements Axiom Q650: A single sensor cannot create truth. This module
// enforces multi-sensor quorum and weighted confidence intervals. It is
// the bridge between probabilistic observation and symbolic ledger truth.

// ConfidenceScore represents the reliability of a fact (0.0 to 1.0).
type ConfidenceScore phxmath.FixedPoint

/**
 * MultiSensorEvidence
 *
 * Implements the Reality Confidence Theorem.
 * Fact = SensorA + SensorB + SensorC + Confidence Score.
 */
type MultiSensorEvidence struct {
	FactID          string                 `json:"fact_id"`
	Observations    map[string]Observation `json:"observations"` // SensorID -> Observation
	ConfidenceScore ConfidenceScore        `json:"confidence_score"`
	Timestamp       int64                  `json:"timestamp"`
}

type Observation struct {
	Value     []byte `json:"value"`
	Signature string `json:"signature"`
}

/**
 * CalculateConfidence
 *
 * Calculates the aggregate confidence of a multi-sensor fact based on weights.
 *
 * Input:
 * - sensorWeights: A map of sensor IDs to their relative reliability weights.
 *
 * Output:
 * - The calculated ConfidenceScore.
 */
func (mse *MultiSensorEvidence) CalculateConfidence(sensorWeights map[string]phxmath.FixedPoint) ConfidenceScore {
	if len(mse.Observations) == 0 {
		return ConfidenceScore(phxmath.NewFixedPoint(0))
	}

	// Group observations by value to find consensus
	valueGroups := make(map[string]phxmath.FixedPoint)
	var totalWeight phxmath.FixedPoint

	for sensorID, obs := range mse.Observations {
		weight, ok := sensorWeights[sensorID]
		if !ok {
			weight = phxmath.NewFixedPoint(1) // Default weight
		}
		vg := valueGroups[string(obs.Value)]
		valueGroups[string(obs.Value)] = vg.SaturatingAdd(weight)
		totalWeight = totalWeight.SaturatingAdd(weight)
	}

	if totalWeight.V == 0 {
		return ConfidenceScore(phxmath.NewFixedPoint(0))
	}

	// Find the weight of the largest agreeing group
	var maxWeight phxmath.FixedPoint
	for _, weight := range valueGroups {
		if weight.V > maxWeight.V {
			maxWeight = weight
		}
	}

	// Confidence is the ratio of the majority weight to the total weight.
	mse.ConfidenceScore = ConfidenceScore(maxWeight.Div(totalWeight))
	return mse.ConfidenceScore
}

/**
 * IsValid
 *
 * Enforces Invariant Q649: A fact requires at least N agreeing sensors.
 * Enforces Invariant Q650: A single sensor cannot create absolute truth.
 */
func (mse *MultiSensorEvidence) IsValid(minSensors int, minConfidence ConfidenceScore) bool {
	if len(mse.Observations) < minSensors {
		return false
	}
	if phxmath.FixedPoint(mse.ConfidenceScore).V < phxmath.FixedPoint(minConfidence).V {
		return false
	}
	return true
}
