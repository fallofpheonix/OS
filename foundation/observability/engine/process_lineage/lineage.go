/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package lineage

import (
	"sync"
	"time"
)

// ProcessNode represents a single process in the lineage graph.
type ProcessNode struct {
	PID       uint32         `json:"pid"`
	PPID      uint32         `json:"ppid"`
	Comm      string         `json:"comm"`
	ExePath   string         `json:"exe_path"`
	Args      []string       `json:"args"`
	UID       uint32         `json:"uid"`
	GID       uint32         `json:"gid"`
	StartTime time.Time      `json:"start_time"`
	ExitTime  *time.Time     `json:"exit_time,omitempty"`
	IsActive  bool           `json:"is_active"`
	Children  []uint32       `json:"children"`
	Payload   map[string]any `json:"payload,omitempty"`
}

// LineageGraph tracks process relationships over time.
type LineageGraph struct {
	mu    sync.RWMutex
	Nodes map[uint32]*ProcessNode
}

// NewLineageGraph initializes a new process graph.
func NewLineageGraph() *LineageGraph {
	return &LineageGraph{
		Nodes: make(map[uint32]*ProcessNode),
	}
}

// AddProcess adds or updates a process in the graph.
func (g *LineageGraph) AddProcess(pid, ppid uint32, comm, exe string, timestamp time.Time) {
	g.mu.Lock()
	defer g.mu.Unlock()

	node, exists := g.Nodes[pid]
	if !exists {
		node = &ProcessNode{
			PID:       pid,
			PPID:      ppid,
			StartTime: timestamp,
			IsActive:  true,
			Children:  []uint32{},
		}
		g.Nodes[pid] = node
	}

	node.Comm = comm
	node.ExePath = exe

	if parent, ok := g.Nodes[ppid]; ok {
		// Avoid duplicates in children list
		found := false
		for _, childID := range parent.Children {
			if childID == pid {
				found = true
				break
			}
		}
		if !found {
			parent.Children = append(parent.Children, pid)
		}
	}
}

// ExitProcess marks a process as inactive.
func (g *LineageGraph) ExitProcess(pid uint32, timestamp time.Time) {
	g.mu.Lock()
	defer g.mu.Unlock()

	if node, ok := g.Nodes[pid]; ok {
		node.IsActive = false
		node.ExitTime = &timestamp
	}
}

// GetAncestors returns the list of PIDs in the parent chain.
func (g *LineageGraph) GetAncestors(pid uint32) []uint32 {
	g.mu.RLock()
	defer g.mu.RUnlock()

	ancestors := []uint32{}
	curr := pid
	for {
		node, ok := g.Nodes[curr]
		if !ok || node.PPID == 0 || node.PPID == curr {
			break
		}
		ancestors = append(ancestors, node.PPID)
		curr = node.PPID
		if len(ancestors) > 1024 { // Safety break
			break
		}
	}
	return ancestors
}

// GetDescendants returns all PIDs in the process subtree (BFS).
func (g *LineageGraph) GetDescendants(pid uint32) []uint32 {
	g.mu.RLock()
	defer g.mu.RUnlock()

	descendants := []uint32{}
	queue := []uint32{pid}
	visited := make(map[uint32]bool)

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]

		if visited[curr] {
			continue
		}
		visited[curr] = true

		if node, ok := g.Nodes[curr]; ok {
			for _, childID := range node.Children {
				descendants = append(descendants, childID)
				queue = append(queue, childID)
			}
		}

		if len(visited) > 100000 { // Safety limit
			break
		}
	}
	return descendants
}

// Size returns the number of nodes in the graph.
func (g *LineageGraph) Size() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.Nodes)
}
