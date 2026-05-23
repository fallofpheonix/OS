package game

import (
	"phoenix/agents/internal/types"
	"testing"
	"time"
)

func TestGameAgent(t *testing.T) {
	agent := NewGameAgent()
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	state := types.SecurityState{
		ThreatTemperature: 0.2,
		SDI:               0.1,
	}

	graph := &types.IncidentGraph{
		Nodes: make(map[string]*types.ProcessNode),
	}

	// Test belief update
	priorRan, _ := agent.GetBeliefs()
	agent.UpdateBeliefs(state, "high_entropy_write")
	postRan, _ := agent.GetBeliefs()

	if postRan <= priorRan {
		t.Errorf("expected ransomware belief to increase, got %f -> %f", priorRan, postRan)
	}

	// Test strategy solving
	graph.Nodes["200"] = &types.ProcessNode{
		PID:         200,
		ThreatScore: 9.0,
	}

	now := time.Now()
	strategy, err := agent.SolveBestStrategy(state, graph, now)
	if err != nil {
		t.Fatalf("failed to solve strategy: %v", err)
	}

	if strategy.ContainmentLevel == 0 {
		t.Error("expected non-zero containment level for high threat node")
	}

	foundTarget := false
	for _, pid := range strategy.TargetPIDs {
		if pid == 200 {
			foundTarget = true
			break
		}
	}
	if !foundTarget {
		t.Error("PID 200 should be in target list")
	}
}
