package main

import (
	"fmt"
	"sync"
	"time"
)

type StorageTier uint8

const (
	TierHot  StorageTier = 0 // In-Memory
	TierWarm StorageTier = 1 // Local Compressed (Simulated)
	TierCold StorageTier = 2 // Skeleton Chain (Simulated)
)

type Node struct {
	ID        string
	Type      string
	Tier      StorageTier
	Immutable bool
	LastSeen  time.Time
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

func (g *Graph) AddNode(id, t string, immutable bool) {
	g.mu.Lock()
	defer g.mu.Unlock()
	if _, ok := g.nodes[id]; !ok {
		g.nodes[id] = &Node{
			ID:        id,
			Type:      t,
			Tier:      TierHot,
			Immutable: immutable,
			LastSeen:  time.Now(),
		}
	}
}

func (g *Graph) AddEdge(from, to string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.edges[from] = append(g.edges[from], to)
}

// TransitionTier moves nodes between tiers based on age and importance
func (g *Graph) TransitionTier(id string, target StorageTier) {
	g.mu.Lock()
	defer g.mu.Unlock()
	
	node, ok := g.nodes[id]
	if !ok { return }

	// Retention rule: Never prune critical nodes from HOT if immutable
	if node.Immutable && target > TierHot {
		fmt.Printf("[TRACE] Retention Rule: Node %s (%s) preserved in HOT tier\n", id, node.Type)
		return
	}

	fmt.Printf("[TRACE] Transitioning Node %s from Tier %d to %d\n", id, node.Tier, target)
	node.Tier = target
	
	if target == TierCold {
		// Simulate compression: in real impl, remove fields not in "skeleton"
		fmt.Printf("[TRACE] Node %s compacted to Skeleton Chain in COLD tier\n", id)
	}
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
	fmt.Println("Phoenix Trace starting with 3-Tier Storage...")
	g := NewGraph()

	// 1. Add critical nodes (Immutable)
	g.AddNode("1", "kernel", true)
	g.AddNode("2", "init", true)
	g.AddEdge("1", "2")

	// 2. Add ephemeral process nodes
	g.AddNode("5001", "process-short", false)
	g.AddEdge("2", "5001")

	// 3. Simulate lifecycle
	time.Sleep(100 * time.Millisecond)
	
	fmt.Println("Performing Lifecycle Management...")
	// Critical nodes stay HOT
	g.TransitionTier("1", TierWarm)
	
	// Ephemeral nodes move to COLD
	g.TransitionTier("5001", TierCold)

	path := g.Trace("1")
	fmt.Printf("Full Causal Trace for Kernel: %v\n", path)
}
