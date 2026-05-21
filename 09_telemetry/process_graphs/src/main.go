package main

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
	ID   string
	Type NodeType
}

type Graph struct {
	mu    sync.RWMutex
	edges map[string][]string
	nodes map[string]*Node
}

func NewGraph() *Graph {
	return &Graph{
		edges: make(map[string][]string),
		nodes: make(map[string]*Node),
	}
}

func (g *Graph) AddNode(id string, t NodeType) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.nodes[id]; !ok {
		g.nodes[id] = &Node{ID: id, Type: t}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.edges[from] = append(g.edges[from], to)
}

func (g *Graph) GetLineage(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var lineage []string
	visited := make(map[string]bool)
	g.dfs(id, visited, &lineage)
	return lineage
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

func main() {
	g := NewGraph()
	g.AddNode("1", Process)
	g.AddNode("2", Process)
	g.AddEdge("1", "2")
	fmt.Printf("Lineage of 1: %v\n", g.GetLineage("1"))
}
