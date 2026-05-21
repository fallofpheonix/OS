package graph

import (
	"fmt"
	"sync"
	"time"

	"phoenix/agents/internal/types"
)

type GraphAgent interface {
	UpdateGraph(ev types.TelemetryEvent) error
	GetIncidentGraph() (*types.IncidentGraph, error)
	GetAttackDAG() *types.IncidentGraph
}

type Agent struct {
	mu            sync.RWMutex
	nodes         map[string]*types.ProcessNode
	edges         map[string][]string
	incidentNodes map[string]*types.ProcessNode
	incidentEdges map[string][]string
}

func NewGraphAgent() *Agent {
	return &Agent{
		nodes:         make(map[string]*types.ProcessNode),
		edges:         make(map[string][]string),
		incidentNodes: make(map[string]*types.ProcessNode),
		incidentEdges: make(map[string][]string),
	}
}

func (a *Agent) UpdateGraph(ev types.TelemetryEvent) error {
	a.mu.Lock()
	defer a.mu.Unlock()

	nodeID := fmt.Sprintf("%d", ev.PID)
	
	// Create or update node
	n, exists := a.nodes[nodeID]
	if !exists {
		n = &types.ProcessNode{
			PID:      ev.PID,
			Comm:     ev.Comm,
			ExePath:  ev.ExePath,
			LastSeen: time.Now(),
		}
		a.nodes[nodeID] = n
	} else {
		n.LastSeen = time.Now()
	}

	// Calculate threat score based on classification or rules
	if ev.Comm == "malware" || ev.Category == "network" && ev.Network != nil && ev.Network.DPort == 4444 {
		n.ThreatScore = 9.0
	} else if ev.Comm == "bash" {
		n.ThreatScore = 4.0
	} else {
		n.ThreatScore = 1.0
	}

	// Add edge from parent to child
	if ev.PPID > 0 {
		parentID := fmt.Sprintf("%d", ev.PPID)
		if _, parentExists := a.nodes[parentID]; !parentExists {
			a.nodes[parentID] = &types.ProcessNode{
				PID:      ev.PPID,
				Comm:     "unknown",
				LastSeen: time.Now(),
			}
		}
		
		// Add edge if not already exists
		alreadyExists := false
		for _, child := range a.edges[parentID] {
			if child == nodeID {
				alreadyExists = true
				break
			}
		}
		if !alreadyExists {
			a.edges[parentID] = append(a.edges[parentID], nodeID)
		}
	}

	// Update centrality metrics (degree centrality as a simple proxy)
	a.recalculateCentrality()

	// If node threat score is high, promote to incident graph
	if n.ThreatScore >= 4.0 {
		a.incidentNodes[nodeID] = n
		if ev.PPID > 0 {
			parentID := fmt.Sprintf("%d", ev.PPID)
			if pNode, ok := a.nodes[parentID]; ok {
				a.incidentNodes[parentID] = pNode
				a.incidentEdges[parentID] = append(a.incidentEdges[parentID], nodeID)
			}
		}
	}

	return nil
}

func (a *Agent) recalculateCentrality() {
	inDegree := make(map[string]int)
	outDegree := make(map[string]int)

	for from, toList := range a.edges {
		outDegree[from] = len(toList)
		for _, to := range toList {
			inDegree[to]++
		}
	}

	for id, node := range a.nodes {
		total := float64(inDegree[id] + outDegree[id])
		node.Centrality = total / float64(len(a.nodes))
	}
}

func (a *Agent) GetIncidentGraph() (*types.IncidentGraph, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return &types.IncidentGraph{
		Nodes: a.incidentNodes,
		Edges: a.incidentEdges,
	}, nil
}

func (a *Agent) GetAttackDAG() *types.IncidentGraph {
	a.mu.RLock()
	defer a.mu.RUnlock()

	return &types.IncidentGraph{
		Nodes: a.nodes,
		Edges: a.edges,
	}
}
