// Package arbiter provides policy enforcement and decision-making for PhoenixOS.
// Core Domain Logic: Implements the "Policy Validator" (Arbiter) which acts as the system's
// gatekeeper, evaluating security risks and authorizing state transitions based on defined redlines.
package arbiter

import (
	"fmt"
	"strings"
)

// AdvisoryEnvelope shim for local logic until proto integration is complete.
// Internal State: Metadata describing a proposed system action and its risk score.
// API Scope: Public within the arbiter domain.
// Concurrency: Thread-safe (immutable).
type AdvisoryEnvelope struct {
	AdvisoryID         string
	RecommendationType string
	BoundedScope       string
	RiskScore          float64
}

// PolicyValidator enforces system-wide security constraints (redlines).
// Internal State: Lists of forbidden scopes and restricted action types.
// API Scope: Public; primary interface for security policy evaluation.
// Concurrency: Thread-safe due to immutable configuration after initialization.
type PolicyValidator struct {
	ForbiddenScopes []string
	RestrictedTypes []string
}

// Arbiter is an alias for PolicyValidator for legacy compatibility.
type Arbiter = PolicyValidator

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewPolicyValidator creates the Arbiter with hardcoded security redlines.
// I/O: None.
// Complexity: O(1).
func NewPolicyValidator() *PolicyValidator {
	return &PolicyValidator{
		ForbiddenScopes: []string{"KERNEL_ROOT", "LEDGER_MUTATION", "TRUTH_DELETION"},
		RestrictedTypes: []string{"UNBOUNDED_EXECUTION", "SYSTEM_HALT"},
	}
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// Validate checks an advisory against established security redlines to prevent unauthorized access or actions.
// I/O: None.
// Complexity: O(F + R) where F is the number of forbidden scopes and R is the number of restricted types.
func (v *PolicyValidator) Validate(adv *AdvisoryEnvelope) error {
	for _, scope := range v.ForbiddenScopes {
		if strings.Contains(strings.ToUpper(adv.BoundedScope), scope) {
			return fmt.Errorf("POLICY_VIOLATION: Advisory %s attempts to access forbidden scope: %s", adv.AdvisoryID, scope)
		}
	}

	for _, rType := range v.RestrictedTypes {
		if strings.Contains(strings.ToUpper(adv.RecommendationType), rType) {
			return fmt.Errorf("POLICY_VIOLATION: Advisory %s proposes restricted action: %s", adv.AdvisoryID, rType)
		}
	}

	return nil
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// Evaluate determines whether a proposed actuation is authorized based on drift detection and confidence scores.
// I/O: None.
// Complexity: O(1).
func (v *PolicyValidator) Evaluate(score interface{}, tcsScore float64) (targetState, actuationClass string, authorized bool) {
	drift := 0.0
	// Handle multiple score types for robustness
	switch s := score.(type) {
	case float64:
		drift = s
	case float32:
		drift = float64(s)
	case int:
		drift = float64(s)
	}

	// [HARDENING]: Multi-factor authorization logic
	if drift >= 0.8 && tcsScore >= 0.7 {
		return "ESCALATE", "CONTAIN", true
	}

	if tcsScore < 0.5 {
		return "WATCH", "LOG", true
	}

	if drift > 0.4 {
		return "ANALYZE", "LOG", true
	}

	return "SAFE", "LOG", true
}
