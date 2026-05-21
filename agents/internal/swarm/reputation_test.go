package swarm

import "testing"

func TestReputationStore(t *testing.T) {
	store := NewReputationStore()
	store.UpdateReputation("node1", 10.0)
	
	if store.GetReputation("node1") != 10.0 {
		t.Errorf("Expected 10.0, got %f", store.GetReputation("node1"))
	}
}
