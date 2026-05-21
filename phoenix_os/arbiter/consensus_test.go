package arbiter

import (
	"phoenix/agents/internal/swarm"
	"testing"
)

func TestConsensusEngine(t *testing.T) {
	rep := swarm.NewReputationStore()
	rep.UpdateReputation("node1", 10.0) // High trust
	rep.UpdateReputation("node2", 1.0)  // Low trust

	engine := NewConsensusEngine(rep, 0.5)

	// High trust node votes true
	engine.SubmitVote(Vote{"node1", true, 0.9})
	// Low trust node votes false
	engine.SubmitVote(Vote{"node2", false, 0.9})

	if !engine.EvaluateQuorum() {
		t.Error("Expected quorum to be true based on high trust node")
	}
}
