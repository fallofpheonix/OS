// Package arbiter provides policy enforcement and decision-making for PhoenixOS.
// Core Domain Logic: Implements the "Enforcement Translator" which maps abstract policy
// recommendations (e.g., "isolate network") into concrete, executable substrate actions.
package arbiter

import (
	"fmt"
	"time"
)

// EnforcementRequest represents a concrete action to be executed by the substrate.
// Internal State: Specific action, target, and mandatory rollback plan for atomic execution.
// API Scope: Public within the arbiter and executor domains.
// Concurrency: Thread-safe (immutable).
type EnforcementRequest struct {
	RequestID    string
	Target       string
	Action       string
	Timeout      time.Duration
	RollbackPlan string
}

// Translator maps high-level advisories to concrete enforcement requests.
// Internal State: Stateless mapper.
// API Scope: Public; used for converting policy decisions into actions.
// Concurrency: Thread-safe (stateless).
type Translator struct{}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewTranslator creates a new advisory-to-enforcement mapper.
// I/O: None.
// Complexity: O(1).
func NewTranslator() *Translator {
	return &Translator{}
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// Translate converts an approved AdvisoryEnvelope into a bounded EnforcementRequest with a mandatory rollback plan.
// I/O: None.
// Complexity: O(1).
func (t *Translator) Translate(adv *AdvisoryEnvelope) (*EnforcementRequest, error) {
	req := &EnforcementRequest{
		RequestID: fmt.Sprintf("enforce-%s", adv.AdvisoryID),
		Timeout:   30 * time.Second,
	}

	// Deterministic Mapping Matrix
	switch adv.RecommendationType {
	case "CONTAINMENT_PROPOSAL":
		req.Target = "PROCESS_BOUNDARY"
		req.Action = "BPF_MAP_UPDATE:blocked_pids"
		req.RollbackPlan = "BPF_MAP_DELETE:blocked_pids"
	case "NETWORK_ISOLATION":
		req.Target = "NET_NAMESPACE"
		req.Action = "NS_SEVER_CONNECT"
		req.RollbackPlan = "NS_RESTORE_CONNECT"
	case "RECOVERY_INITIATION":
		req.Target = "FSM_STATE"
		req.Action = "TRANSITION_TO_WATCH"
		req.RollbackPlan = "NO_OP" // Recovery is low-risk but needs tracking
	default:
		return nil, fmt.Errorf("TRANSLATION_ERROR: Unknown recommendation type: %s", adv.RecommendationType)
	}

	return req, nil
}
