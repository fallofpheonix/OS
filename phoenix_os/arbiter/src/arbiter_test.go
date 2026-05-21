package main

import (
	"testing"
)

func TestCalculateQuorum(t *testing.T) {
	arbiter := Arbiter{
		Nodes: []Node{
			{"node1", 1.0, true},
			{"node2", 2.0, true},
			{"node3", 1.0, false},
		},
		QuorumThreshold: 0.6,
	}

	votes := map[string]bool{
		"node1": true,
		"node2": false,
		"node3": true,
	}

	// Reputation: node1=1, node2=2, total=3
	// Positive: node1=1, node2=0, total=1
	// 1/3 = 0.33 < 0.6 => false
	if arbiter.CalculateQuorum(votes) {
		t.Errorf("Expected false for 0.33 threshold")
	}

	votes2 := map[string]bool{
		"node1": true,
		"node2": true,
	}
	// (1+2)/3 = 1.0 >= 0.6 => true
	if !arbiter.CalculateQuorum(votes2) {
		t.Errorf("Expected true for 1.0 threshold")
	}
}
