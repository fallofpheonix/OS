package security

import "sync"

// ReputationManager tracks the trustworthiness of nodes in the swarm.
type ReputationManager struct {
	mu         sync.RWMutex
	Reputation map[string]float64
}

// NewReputationManager initializes with default scores.
func NewReputationManager() *ReputationManager {
	return &ReputationManager{
		Reputation: make(map[string]float64),
	}
}

// Deduct reduces a node's reputation score for suspicious behavior.
func (rm *ReputationManager) Deduct(nodeID string, amount float64) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	rm.Reputation[nodeID] -= amount
}

// GetScore returns the current reputation of a node.
func (rm *ReputationManager) GetScore(nodeID string) float64 {
	rm.mu.RLock()
	defer rm.mu.RUnlock()
	return rm.Reputation[nodeID]
}
