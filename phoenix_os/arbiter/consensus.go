package arbiter

import (
	"sync"
	"time"

	"phoenix/agents/internal/swarm"
)

type Vote struct {
	NodeID     string
	Decision   bool
	Confidence float64
}

type ConsensusEngine struct {
	mu           sync.Mutex
	reputation   *swarm.ReputationStore
	votes        []Vote
	threshold    float64
}

func NewConsensusEngine(rep *swarm.ReputationStore, threshold float64) *ConsensusEngine {
	return &ConsensusEngine{
		reputation: rep,
		votes:      make([]Vote, 0),
		threshold:  threshold,
	}
}

func (e *ConsensusEngine) SubmitVote(v Vote) {
	e.mu.Lock()
	defer e.mu.Unlock()
	e.votes = append(e.votes, v)
}

func (e *ConsensusEngine) EvaluateQuorum() bool {
	e.mu.Lock()
	defer e.mu.Unlock()

	var weightedTotal float64
	var totalReputation float64

	for _, v := range e.votes {
		rep := e.reputation.GetReputation(v.NodeID)
		if rep == 0 { rep = 1.0 } // Default minimal rep

		if v.Decision {
			weightedTotal += rep * v.Confidence
		}
		totalReputation += rep
	}

	if totalReputation == 0 {
		return false
	}

	return (weightedTotal / totalReputation) > e.threshold
}
