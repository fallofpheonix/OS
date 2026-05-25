package graph

import (
	"fmt"
	"sync"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/types"
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
			LastSeen: ev.Timestamp,
		}
		
		// Set basic criticality based on comm
		if ev.Comm == "init" || ev.Comm == "systemd" || ev.Comm == "nginx" {
			n.Criticality = 0.9
		} else {
			n.Criticality = 0.2
		}
		
		a.nodes[nodeID] = n
	} else {
		n.LastSeen = ev.Timestamp
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
		parent, parentExists := a.nodes[parentID]
		if !parentExists {
			parent = &types.ProcessNode{
				PID:         ev.PPID,
				Comm:        "unknown",
				LastSeen:    ev.Timestamp,
				Criticality: 0.5,
			}
			a.nodes[parentID] = parent
		}
		
		// Set depth
		n.Depth = parent.Depth + 1
		
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

	// Update importance metrics
	a.recalculateMetrics()

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

func (a *Agent) recalculateMetrics() {
	inDegree := make(map[string]int)
	outDegree := make(map[string]int)

	for from, toList := range a.edges {
		outDegree[from] = len(toList)
		for _, to := range toList {
			inDegree[to]++
		}
	}

	for id, node := range a.nodes {
		// Centrality
		total := float64(inDegree[id] + outDegree[id])
		node.Centrality = total / float64(len(a.nodes))
		
		// Spread (out-degree fan-out)
		node.Spread = float64(outDegree[id]) / 10.0 // normalized
		if node.Spread > 1.0 {
			node.Spread = 1.0
		}
		
		// Multi-factor Importance Score
		// Importance = Centrality * 0.4 + Criticality * 0.3 + Spread * 0.2 + (1/Depth) * 0.1
		depthFactor := 1.0
		if node.Depth > 0 {
			depthFactor = 1.0 / node.Depth
		}
		
		node.Importance = (node.Centrality * 0.4) + (node.Criticality * 0.3) + (node.Spread * 0.2) + (depthFactor * 0.1)
		if node.Importance > 1.0 {
			node.Importance = 1.0
		}
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
