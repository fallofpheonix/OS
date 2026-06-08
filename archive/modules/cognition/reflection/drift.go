/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package reflection implements the epistemic safety and reality drift auditing layer for PhoenixOS.
package reflection

import (
	"sync"
)

// RealityDriftAuditor tracks the cumulative divergence between internal world models and measured reality.
// It implements Axiom Q808, treating drift as a thermodynamic signal for system-level safety cycles.
// Internal state is synchronized via a RWMutex to support high-frequency telemetry auditing.
type RealityDriftAuditor struct {
	mu            sync.RWMutex
	Cumulative    float64 `json:"cumulative_drift"`
	SampleCount   uint64  `json:"sample_count"`
	Threshold     float64 `json:"threshold"` // Cumulative average threshold for quarantine
	IsQuarantined bool    `json:"is_quarantined"`
}

// NewRealityDriftAuditor initializes a new auditor with a specified error threshold.
func NewRealityDriftAuditor(threshold float64) *RealityDriftAuditor {
	return &RealityDriftAuditor{
		Threshold: threshold,
	}
}

// RecordError adds a new divergence sample to the auditor and evaluates the Quarantine condition.
// Average drift is calculated as (Cumulative / SampleCount).
// Side Effects: Sets IsQuarantined to true if average drift crosses the threshold (Axiom Q809).
// Complexity: O(1) time / O(1) space.
func (a *RealityDriftAuditor) RecordError(err *ReflectionError) {
	a.mu.Lock()
	defer a.mu.Unlock()

	a.Cumulative += err.Divergence
	a.SampleCount++

	// Q809: Evaluate average drift against safety boundary.
	if a.SampleCount > 0 {
		avgDrift := a.Cumulative / float64(a.SampleCount)
		if avgDrift > a.Threshold {
			a.IsQuarantined = true
		}
	}
}

// Reset clears the cumulative drift history and releases the quarantine state.
// Security: This method MUST only be invoked after a successful system recovery or manual policy recalibration.
// Complexity: O(1) time / O(1) space.
func (a *RealityDriftAuditor) Reset() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.Cumulative = 0
	a.SampleCount = 0
	a.IsQuarantined = false
}
