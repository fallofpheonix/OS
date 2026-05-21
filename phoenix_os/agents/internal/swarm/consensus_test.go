package swarm

import (
	"testing"
)

func TestConsensus(t *testing.T) {
	cm := NewConsensusModule(1.5) // Need 1.5 total reputation

	// Node 1 proposes
	if cm.Propose("node-1", 0.6) {
		t.Error("Proposal should fail due to insufficient quorum")
	}

	// Node 2 joins
	if !cm.Propose("node-2", 1.0) {
		t.Error("Proposal should pass with sufficient quorum")
	}
}
