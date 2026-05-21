package physics

import (
	"phoenix/agents/internal/types"
	"testing"
)

func TestPhysicsAgent(t *testing.T) {
	agent := NewPhysicsAgent()
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	// Test SDI calculation
	states := []int8{1, 1, -1, -1}
	sdi := agent.CalculateSDI(states)
	if sdi <= 0 {
		t.Errorf("expected positive SDI, got %f", sdi)
	}

	// Test Security State
	graph := &types.IncidentGraph{
		Nodes: make(map[string]*types.ProcessNode),
	}
	graph.Nodes["100"] = &types.ProcessNode{
		PID:         100,
		ThreatScore: 8.0,
		Centrality:  0.5,
	}

	state, err := agent.GetSecurityState(graph)
	if err != nil {
		t.Fatalf("failed to get security state: %v", err)
	}

	if state.ThreatTemperature <= 0.1 {
		t.Errorf("expected increased threat temperature, got %f", state.ThreatTemperature)
	}

	if !state.IsAnomaly {
		t.Error("expected state to be marked as anomaly")
	}
}
