package main

import (
	"fmt"
	"sync"
	"time"
)

type NodeState struct {
	NodeID    string    `json:"node_id"`
	SDI       float64   `json:"sdi"`
	Timestamp time.Time `json:"timestamp"`
}

type Nexus struct {
	mu        sync.RWMutex
	Cluster   map[string]NodeState
	SelfID    string
}

func NewNexus(id string) *Nexus {
	return &Nexus{
		Cluster: make(map[string]NodeState),
		SelfID:  id,
	}
}

func (n *Nexus) UpdateCluster(state NodeState) {
	n.mu.Lock()
	defer n.mu.Unlock()
	
	// CRDT-like logic: only update if the new state is fresher
	if current, exists := n.Cluster[state.NodeID]; !exists || state.Timestamp.After(current.Timestamp) {
		n.Cluster[state.NodeID] = state
	}
}

func (n *Nexus) GetGlobalSDI() float64 {
	n.mu.RLock()
	defer n.mu.RUnlock()
	
	var total float64
	for _, s := range n.Cluster {
		total += s.SDI
	}
	if len(n.Cluster) == 0 { return 0 }
	return total / float64(len(n.Cluster))
}

func main() {
	fmt.Println("Phoenix Nexus Online")
	nexus := NewNexus("node-alpha")
	
	// Simulate receiving state from node-beta
	nexus.UpdateCluster(NodeState{
		NodeID:    "node-beta",
		SDI:       0.85,
		Timestamp: time.Now(),
	})
	
	fmt.Printf("Global Cluster SDI: %f\n", nexus.GetGlobalSDI())
}
