// Package warden implements formal invariant verification for authority escalation.
// Domain Logic: Defines the proof-gates ensuring cryptographic, causal, and weight-based validity of system state transitions.
// Responsibility: Mathematically verifies security invariants during runtime.
package warden

import (
	"fmt"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// Invariant defines the interface for formal proof-gates.
type Invariant interface {
	Verify(req AuthorityEscalationRequest, currentState SystemState) error
}

// AuthorityEscalationRequest carries the proof and telemetry required to transition system state.
// Concurrency: Read-only instances are thread-safe.
// State Management: Encapsulates transient request metadata and evidence proofs.
type AuthorityEscalationRequest struct {
	EventID        string
	TargetPID      int
	TargetTgid     int
	TargetNsproxy  uint32
	TargetState    SystemState
	ActuationClass ActuationClass
	EvidenceWeight float64
	Certificate    []byte
	GraphProof     *GraphProof
}

func mapSystemStateToLevel(state SystemState) securityv1.ContainmentLevel {
	switch state {
	case StateSafe:
		return securityv1.LevelNone
	case StateWatch:
		return securityv1.LevelMonitor
	case StateSuspicious:
		return securityv1.LevelSandbox
	case StateCritical:
		return securityv1.LevelIsolate
	case StateCompromised:
		return securityv1.LevelQuench
	default:
		return securityv1.LevelNone
	}
}

// Target implements securityv1.Containment.
func (req AuthorityEscalationRequest) Target() string {
	return fmt.Sprintf("PID:%d", req.TargetPID)
}

// Level implements securityv1.Containment.
func (req AuthorityEscalationRequest) Level() securityv1.ContainmentLevel {
	return mapSystemStateToLevel(req.TargetState)
}

// Reason implements securityv1.Containment.
func (req AuthorityEscalationRequest) Reason() string {
	return string(req.ActuationClass)
}

// CurrentLevel implements securityv1.Escalation.
func (req AuthorityEscalationRequest) CurrentLevel() securityv1.ContainmentLevel {
	return securityv1.LevelNone
}

// TargetLevel implements securityv1.Escalation.
func (req AuthorityEscalationRequest) TargetLevel() securityv1.ContainmentLevel {
	return req.Level()
}

// TriggerReason implements securityv1.Escalation.
func (req AuthorityEscalationRequest) TriggerReason() string {
	return req.Reason()
}

// EvidenceWeightInvariant implements the "Conservation of Authority" law.
// Concurrency: Thread-safe for concurrent Verify calls.
// State Management: Stores static thresholds for state transitions.
type EvidenceWeightInvariant struct {
	StateThresholds map[SystemState]float64
}

// LABEL: [READ_ONLY] [DETERMINISTIC] [STABLE]
// Verify checks if the request's EvidenceWeight meets the threshold for the TargetState.
// I/O: None.
// Side Effects: None.
// Complexity: O(1) map lookup.
func (e *EvidenceWeightInvariant) Verify(req AuthorityEscalationRequest, currentState SystemState) error {
	threshold, exists := e.StateThresholds[req.TargetState]
	if exists && req.EvidenceWeight < threshold {
		return fmt.Errorf("invariant violation: insufficient evidence weight to reach %s (%.2f < %.2f)",
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
func (c *CertificateInvariant) Verify(req AuthorityEscalationRequest, currentState SystemState) error {
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
func (c *ContextualInvariant) Verify(req AuthorityEscalationRequest, currentState SystemState) error {
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
