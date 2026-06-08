// Package warden implements formal invariant verification for authority escalation.
// Domain Logic: Defines the proof-gates ensuring cryptographic, causal, and weight-based validity of system state transitions.
// Responsibility: Mathematically verifies security invariants during runtime.
package security

import (
	"fmt"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// EvidenceWeightInvariant implements the "Conservation of Authority" law.
// Concurrency: Thread-safe for concurrent Verify calls.
// State Management: Stores static thresholds for state transitions.
type EvidenceWeightInvariant struct {
	StateThresholds map[SystemState]phxmath.FixedPoint
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Verify checks if the request's EvidenceWeight meets the threshold for the TargetState.
// I/O: None.
// Side Effects: None.
// Complexity: O(1) map lookup.
func (e *EvidenceWeightInvariant) Verify(req AuthorityEscalationRequest, snap PostureSnapshot) error {
	threshold, exists := e.StateThresholds[req.TargetState]
	if exists && req.EvidenceWeight.V < threshold.V {
		return fmt.Errorf("invariant violation: insufficient evidence weight to reach %s (%v < %v)",
			req.TargetState, req.EvidenceWeight, threshold)
	}
	return nil
}

// CertificateInvariant ensures evidence weight has a valid cryptographic proof from the Ledger.
// Concurrency: Dependent on Validator implementation (assumed thread-safe).
// State Management: Maintains a reference to the cryptographic CertificateValidator.
type CertificateInvariant struct {
	Validator CertificateValidator
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Verify invokes the validator to check the request's certificate against the event ID.
// I/O: Potential external validator call.
// Side Effects: None.
// Complexity: O(V) where V is the complexity of the validator's VerifyCertificate call.
func (c *CertificateInvariant) Verify(req AuthorityEscalationRequest, snap PostureSnapshot) error {
	if c.Validator == nil {
		return fmt.Errorf("certificate invariant error: validator not initialized")
	}
	if !c.Validator.VerifyCertificate(req.EventID, req.EvidenceWeight, req.Certificate) {
		return fmt.Errorf("invariant violation: invalid evidence certificate for event %s", req.EventID)
	}
	return nil
}

// ContextualInvariant validates the Oracle's reasoning path and lineage consistency.
// Concurrency: Dependent on Provider implementation (assumed thread-safe).
// State Management: Maintains a reference to the GraphProvider for causal verification.
type ContextualInvariant struct {
	Provider GraphProvider
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Verify checks causal path validity and ensures namespace consistency.
// I/O: Potential external provider call.
// Side Effects: None.
// Complexity: O(P) where P is the length of the causal path.
func (c *ContextualInvariant) Verify(req AuthorityEscalationRequest, snap PostureSnapshot) error {
	if req.TargetNsproxy == 0 || req.TargetTgid == 0 {
		return fmt.Errorf("lineage violation: uninitialized metadata (Nsproxy/Tgid cannot be 0)")
	}

	if req.GraphProof == nil {
		return fmt.Errorf("contextual invariant violation: missing GraphProof")
	}

	valid, err := c.Provider.VerifyPath(req.GraphProof.Path)
	if !valid || err != nil {
		return fmt.Errorf("causal path verification failed: %v", err)
	}

	if req.TargetNsproxy != req.GraphProof.ExpectedNsproxy {
		return fmt.Errorf("lineage violation: namespace drift detected (got %d, expected %d)",
			req.TargetNsproxy, req.GraphProof.ExpectedNsproxy)
	}

	return nil
}
