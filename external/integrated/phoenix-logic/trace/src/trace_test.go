package main

import "testing"

func TestTraceWithTiers(t *testing.T) {
	g := NewGraph()
	// Immutable nodes (Critical Infrastructure)
	g.AddNode("1", "kernel", true)
	g.AddNode("2", "init", true)
	g.AddNode("3", "auth", true)

	// Ephemeral nodes (User Processes)
	g.AddNode("100", "bash", false)
	g.AddNode("101", "ls", false)

	g.AddEdge("1", "2")
	g.AddEdge("2", "100")
	g.AddEdge("100", "101")

	// 1. Verify Retention of Critical Nodes
	g.TransitionTier("1", TierWarm)
	if g.nodes["1"].Tier != TierHot {
		t.Errorf("Critical node 'kernel' should stay in HOT tier, got %v", g.nodes["1"].Tier)
	}

	g.TransitionTier("3", TierCold)
	if g.nodes["3"].Tier != TierHot {
		t.Errorf("Critical node 'auth' should stay in HOT tier even if COLD requested, got %v", g.nodes["3"].Tier)
	}

	// 2. Verify Transition of Ephemeral Nodes
	g.TransitionTier("101", TierWarm)
	if g.nodes["101"].Tier != TierWarm {
		t.Errorf("Ephemeral node 'ls' should transition to WARM tier, got %v", g.nodes["101"].Tier)
	}

	g.TransitionTier("101", TierCold)
	if g.nodes["101"].Tier != TierCold {
		t.Errorf("Ephemeral node 'ls' should transition to COLD tier, got %v", g.nodes["101"].Tier)
	}

	// 3. Verify Trace Integrity after Transitions
	path := g.Trace("1")
	if len(path) != 4 { // kernel, init, bash, ls
		t.Errorf("Expected path length 4, got %d: %v", len(path), path)
	}
}
