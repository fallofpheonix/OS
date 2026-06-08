// Package consensus implements Proof-of-Authority voting logic.
// Domain Logic: Ensures critical state transitions are authorized by a 2/3 weight quorum of participating nodes.
// Responsibility: Handles distributed voting and quorum verification for the Nexus cluster.
package consensus

import (
	"crypto/ed25519"
	"crypto/sha256"
	"fmt"
	"log"
	"sync"

	"github.com/fallofpheonix/phoenix/foundation/distributed/identity"
)

// QuorumProof represents the aggregate authorization for a state transition.
// Concurrency: Read-only instances are thread-safe.
// State Management: Stores the immutable result of a consensus round.
type QuorumProof struct {
	ProposalID        string
	AffirmativeWeight float64
	TotalWeight       float64
	Authorized        bool
}

// PoAEngine handles voting and quorum verification.
// Concurrency: Thread-safe via sync.Mutex.
// State Management: Orchestrates consensus using node weights from the registry.
type PoAEngine struct {
	Registry *identity.NodeRegistry
	mu       sync.Mutex
}

// LABEL: [CREATIONAL] [UNCONSTRAINED] [STABLE]
// NewPoAEngine initializes a new PoA consensus engine.
// I/O: None.
// Side Effects: None.
// Complexity: O(1).
func NewPoAEngine(r *identity.NodeRegistry) *PoAEngine {
	return &PoAEngine{Registry: r}
}

// LABEL: [MUTABLE] [DETERMINISTIC] [STABLE]
// RequestQuorum calculates the aggregate vote weight against the total cluster.
// I/O: Logs authorization status.
// Side Effects: None (beyond logging).
// Complexity: O(N) where N is the number of provided votes.
func (p *PoAEngine) RequestQuorum(proposalID, targetState string, votes map[string][]byte) (*QuorumProof, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	totalClusterWeight := p.Registry.GetTotalWeight()
	if totalClusterWeight == 0 {
		return nil, fmt.Errorf("CONSENSUS_ERROR: no authenticated nodes in registry")
	}

	affirmativeWeight := 0.0
	msg := []byte(fmt.Sprintf("%s:%s", proposalID, targetState))
	h := sha256.Sum256(msg)

	for nodeID, signature := range votes {
		cert, ok := p.Registry.GetPeer(nodeID)
		if !ok {
			log.Printf("[CONSENSUS] Warning: Vote from unknown node %s ignored", nodeID)
			continue
		}

		if ed25519.Verify(cert.PublicKey, h[:], signature) {
			affirmativeWeight += cert.Weight
		} else {
			log.Printf("[CONSENSUS] Warning: Invalid signature from node %s", nodeID)
		}
	}

	threshold := (2.0 / 3.0) * totalClusterWeight
	authorized := affirmativeWeight >= threshold

	proof := &QuorumProof{
		ProposalID:        proposalID,
		AffirmativeWeight: affirmativeWeight,
		TotalWeight:       totalClusterWeight,
		Authorized:        authorized,
	}

// ChooseFork selects the winning chain from a set of alternatives.
// Minimum viable rule: highest sequence number wins on reconnect.
// CONNECTS: T4.2 implementation.
func (p *PoAEngine) ChooseFork(localSeq, remoteSeq uint64, localHash, remoteHash []byte) (useRemote bool) {
	// Rule: Highest sequence number wins. 
	// If sequences are tied, we currently stick with local (liveness preference).
	if remoteSeq > localSeq {
		return true
	}
	return false
}
