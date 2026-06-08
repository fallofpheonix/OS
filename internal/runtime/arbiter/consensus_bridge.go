// Package arbiter provides policy enforcement and decision-making for PhoenixOS.
// Core Domain Logic: Implements the "Consensus Bridge" which ensures that high-risk
// security actions are authorized by a distributed quorum of nodes (Proof of Authority).
package arbiter

import (
	"fmt"
	"log"
)

// PoAEngineInterface abstracts the consensus engine for distributed authorization.
// API Scope: Internal interface for consensus providers.
type PoAEngineInterface interface {
	RequestQuorum(proposalID string, targetState string) (bool, error)
}

// ConsensusBridge handles the interaction with the distributed cluster for authorization.
// Internal State: Reference to the PoA consensus engine.
// API Scope: Public; critical safety gate for distributed system state changes.
// Concurrency: Thread-safe if the underlying PoA engine is thread-safe.
type ConsensusBridge struct {
	Engine PoAEngineInterface
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewConsensusBridge initializes a new distributed consensus connector.
// I/O: None.
// Complexity: O(1).
func NewConsensusBridge(engine PoAEngineInterface) *ConsensusBridge {
	return &ConsensusBridge{Engine: engine}
}

// LABEL: [IO_BOUND] [PUBLIC_API] [STABLE]
// Authorize requests a QuorumProof for high-risk advisories from the distributed cluster.
// I/O: Network IO (via PoA engine).
// Complexity: O(1) + network latency of the consensus protocol.
func (b *ConsensusBridge) Authorize(adv *AdvisoryEnvelope) (bool, error) {
	// Only request quorum for high-risk proposals
	if adv.RiskScore < 0.7 {
		log.Printf("[Arbiter] Low risk advisory %s bypasses cluster quorum.", adv.AdvisoryID)
		return true, nil
	}

	log.Printf("[Arbiter] High risk advisory %s requires cluster quorum. Risk Score: %.2f", adv.AdvisoryID, adv.RiskScore)
	authorized, err := b.Engine.RequestQuorum(adv.AdvisoryID, adv.RecommendationType)
	if err != nil {
		return false, fmt.Errorf("CONSENSUS_BRIDGE_ERROR: %w", err)
	}

	if !authorized {
		return false, fmt.Errorf("QUORUM_DENIED: Cluster rejected proposal %s", adv.AdvisoryID)
	}

	return true, nil
}
