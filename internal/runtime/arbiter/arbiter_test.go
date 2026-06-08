/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package arbiter

import (
	"testing"
)

type MockPoAEngine struct {
	ShouldAuthorize bool
}

func (m *MockPoAEngine) RequestQuorum(proposalID, targetState string) (bool, error) {
	return m.ShouldAuthorize, nil
}

func TestArbiter_FullControlLoop(t *testing.T) {
	validator := NewPolicyValidator()
	translator := NewTranslator()

	mockPoA := &MockPoAEngine{ShouldAuthorize: true}
	bridge := NewConsensusBridge(mockPoA)

	t.Run("Valid Advisory Execution", func(t *testing.T) {
		adv := &AdvisoryEnvelope{
			AdvisoryID:         "adv-001",
			RecommendationType: "CONTAINMENT_PROPOSAL",
			BoundedScope:       "PROCESS:456",
			RiskScore:          0.8,
		}

		// 1. Validate
		if err := validator.Validate(adv); err != nil {
			t.Fatalf("Validation failed unexpectedly: %v", err)
		}

		// 2. Authorize
		auth, err := bridge.Authorize(adv)
		if err != nil || !auth {
			t.Fatalf("Authorization failed unexpectedly: %v", err)
		}

		// 3. Translate
		req, err := translator.Translate(adv)
		if err != nil {
			t.Fatalf("Translation failed unexpectedly: %v", err)
		}

		if req.Action != "BPF_MAP_UPDATE:blocked_pids" {
			t.Errorf("Expected BPF_MAP_UPDATE, got %s", req.Action)
		}
	})

	t.Run("Policy Violation - Forbidden Scope", func(t *testing.T) {
		adv := &AdvisoryEnvelope{
			AdvisoryID:         "adv-002",
			RecommendationType: "CONTAINMENT_PROPOSAL",
			BoundedScope:       "KERNEL_ROOT_ACCESS",
			RiskScore:          0.9,
		}

		if err := validator.Validate(adv); err == nil {
			t.Error("Expected policy violation for KERNEL_ROOT, got nil")
		}
	})

	t.Run("Consensus Rejection", func(t *testing.T) {
		mockPoA.ShouldAuthorize = false
		adv := &AdvisoryEnvelope{
			AdvisoryID:         "adv-003",
			RecommendationType: "NETWORK_ISOLATION",
			BoundedScope:       "NET_NS:789",
			RiskScore:          0.95,
		}

		auth, err := bridge.Authorize(adv)
		if err == nil || auth {
			t.Error("Expected quorum rejection, got authorization")
		}
	})
}
