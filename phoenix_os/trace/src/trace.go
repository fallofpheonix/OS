package main

import (
	"sync"
)

type Node struct {
	ID   string
	Type string
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

func (g *Graph) AddNode(id, t string) {
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

func (g *Graph) Trace(id string) []string {
	g.mu.RLock()
	defer g.mu.RUnlock()
	var path []string
	visited := make(map[string]bool)
	g.dfs(id, visited, &path)
	return path
}

func (g *Graph) dfs(id string, visited map[string]bool, path *[]string) {
	if visited[id] { return }
	visited[id] = true
	*path = append(*path, id)
	for _, n := range g.edges[id] {
		g.dfs(n, visited, path)
	}
}

func main() {
	// Boot check
	NewGraph()
}
