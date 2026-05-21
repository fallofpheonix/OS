package swarm

import (
	"sync"
)

type ReputationStore struct {
	mu         sync.RWMutex
	reputation map[string]float64
}

func NewReputationStore() *ReputationStore {
	return &ReputationStore{
		reputation: make(map[string]float64),
	}
}

func (s *ReputationStore) UpdateReputation(nodeID string, delta float64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.reputation[nodeID] += delta
	if s.reputation[nodeID] < 0 {
		s.reputation[nodeID] = 0
	}
}

func (s *ReputationStore) GetReputation(nodeID string) float64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.reputation[nodeID]
}
