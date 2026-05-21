package main

import (
	"testing"
	"time"
)

func TestNexusConsensus(t *testing.T) {
	nexus := NewNexus("node-01")
	
	s1 := NodeState{NodeID: "node-01", SDI: 0.2, Timestamp: time.Now()}
	s2 := NodeState{NodeID: "node-02", SDI: 0.8, Timestamp: time.Now()}
	
	nexus.UpdateCluster(s1)
	nexus.UpdateCluster(s2)
	
	avg := nexus.GetGlobalSDI()
	if avg != 0.5 {
		t.Errorf("Expected average SDI 0.5, got %f", avg)
	}
}
