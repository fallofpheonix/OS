package swarm

import (
	"sync"
)

// ConsensusModule handles PoA and Weighted Quorum for node agreement.
type ConsensusModule struct {
	mu            sync.RWMutex
	NodeReputation map[string]float64
	QuorumWeight  float64
}

// NewConsensusModule initializes consensus with default weight.
func NewConsensusModule(quorumWeight float64) *ConsensusModule {
	return &ConsensusModule{
		NodeReputation: make(map[string]float64),
		QuorumWeight:   quorumWeight,
	}
}

// Propose evaluates a proposal based on weighted reputation quorum.
func (cm *ConsensusModule) Propose(nodeID string, reputation float64) bool {
	cm.mu.Lock()
	defer cm.mu.Unlock()

	cm.NodeReputation[nodeID] = reputation

	// Calculate total weight
	var totalWeight float64
	for _, r := range cm.NodeReputation {
		totalWeight += r
	}

	// Agreement reached if proposer has sufficient reputation and total quorum is met
	if reputation > 0.5 && totalWeight >= cm.QuorumWeight {
		return true
	}
	return false
}
