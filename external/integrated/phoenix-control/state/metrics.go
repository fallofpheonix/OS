package state

import "sync"

// TransitionMetrics tracks statistics about state transitions.
type TransitionMetrics struct {
	mu               sync.Mutex
	Counts           map[string]uint64
	LastTransition   int64
	TotalTransitions uint64
}

func NewTransitionMetrics() *TransitionMetrics {
	return &TransitionMetrics{
		Counts: make(map[string]uint64),
	}
}

func (m *TransitionMetrics) RecordTransition(from, to RuntimeState) {
	m.mu.Lock()
	defer m.mu.Unlock()

	key := string(from) + "->" + string(to)
	m.Counts[key]++
	m.TotalTransitions++
}

func (m *TransitionMetrics) GetStats() map[string]uint64 {
	m.mu.Lock()
	defer m.mu.Unlock()
	stats := make(map[string]uint64)
	for k, v := range m.Counts {
		stats[k] = v
	}
	stats["total"] = m.TotalTransitions
	return stats
}
