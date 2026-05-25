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
	mu      sync.RWMutex
	edges   map[string][]string
	parents map[string][]string
	nodes   map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		edges:   make(map[string][]string),
		parents: make(map[string][]string),
		nodes:   make(map[string]*Node),
	}
}

func (g *Graph) AddNode(id string, t NodeType, ts int64) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.nodes[id]; !ok {
		g.nodes[id] = &Node{ID: id, Type: t, Timestamp: ts}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.edges[from] = append(g.edges[from], to)
	g.parents[to] = append(g.parents[to], from)
}

func (g *Graph) GetCausalSubtree(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var lineage []string
	visited := make(map[string]bool)
	g.dfs(id, visited, &lineage)
	return lineage
}

func (g *Graph) GetAncestors(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

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
	for _, neighbor := range g.edges[id] {
		g.dfs(neighbor, visited, lineage)
	}
}

func (g *Graph) dfsReverse(id string, visited map[string]bool, ancestors *[]string) {
	if visited[id] {
		return
	}
	visited[id] = true
	*ancestors = append(*ancestors, id)
	for _, parent := range g.parents[id] {
		g.dfsReverse(parent, visited, ancestors)
	}
}

func (g *Graph) GetLineage(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()
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
