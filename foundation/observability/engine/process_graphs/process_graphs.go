/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 4c — PROCESS GRAPH DATA STRUCTURES (Layer 1.5)
//
// This file defines the Graph data structure used by the Orchestrator's
// GraphFeature to track causal relationships between events.
//
// WORKFLOW:
//   GraphFeature.AddEvent(event) → Graph.AddNode() + Graph.AddEdge()
//     → Node created with event ID, type, timestamp
//     → Edge created from parent (CausalID) to child (event ID)
//   → Graph used by Oracle for GraphProof generation
//   → Graph used by Arbiter for lineage-based threat detection
//
// DATA STRUCTURES:
//   Node: ID, Type (process/file/network), Timestamp, Metadata
//   Graph: Nodes map, Edges map, Parents map, Mu (RWMutex)
//
// THREAD SAFETY: All graph operations acquire the mutex.
// Read operations (GetCausalSubtree, GetAncestors) use RLock.
// Write operations (AddNode, AddEdge) use full Lock.
// =========================================================================
package process_graphs

import (
	"fmt"
	"sync"
)

type NodeType string

const (
	Process NodeType = "process"
	File    NodeType = "file"
	Network NodeType = "network"
)

type Node struct {
	ID        string
	Type      NodeType
	Timestamp int64
	Metadata  map[string]interface{}
}

type Graph struct {
	Mu      sync.RWMutex
	Edges   map[string][]string
	Parents map[string][]string
	Nodes   map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		Edges:   make(map[string][]string),
		Parents: make(map[string][]string),
		Nodes:   make(map[string]*Node),
	}
}

func (g *Graph) AddNode(id string, t NodeType, ts int64) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	if _, ok := g.Nodes[id]; !ok {
		g.Nodes[id] = &Node{ID: id, Type: t, Timestamp: ts}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Edges[from] = append(g.Edges[from], to)
	g.Parents[to] = append(g.Parents[to], from)
}

func (g *Graph) GetCausalSubtree(id string) []string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()

	var lineage []string
	visited := make(map[string]bool)
	g.dfs(id, visited, &lineage)
	return lineage
}

func (g *Graph) GetAncestors(id string) []string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()

	var ancestors []string
	visited := make(map[string]bool)
	g.dfsReverse(id, visited, &ancestors)
	return ancestors
}

func (g *Graph) dfs(id string, visited map[string]bool, lineage *[]string) {
	if visited[id] {
		return
	}
	visited[id] = true
	*lineage = append(*lineage, id)
	for _, neighbor := range g.Edges[id] {
		g.dfs(neighbor, visited, lineage)
	}
}

func (g *Graph) dfsReverse(id string, visited map[string]bool, ancestors *[]string) {
	if visited[id] {
		return
	}
	visited[id] = true
	*ancestors = append(*ancestors, id)
	for _, parent := range g.Parents[id] {
		g.dfsReverse(parent, visited, ancestors)
	}
}

func (g *Graph) GetLineage(id string) []string {
	g.Mu.RLock()
	defer g.Mu.RUnlock()
	visited := make(map[string]bool)
	var lineage []string
	g.dfs(id, visited, &lineage)
	return lineage
}

func Example() {
	g := NewGraph()
	g.AddNode("1", Process, 0)
	g.AddNode("2", Process, 0)
	g.AddEdge("1", "2")
	fmt.Printf("Lineage of 1: %v\n", g.GetLineage("1"))
}
