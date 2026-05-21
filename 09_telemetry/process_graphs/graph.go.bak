// Package process_graphs implements RFC-006 Process Lineage DAG.
//
// It provides a thread-safe directed acyclic graph for tracking process
// relationships (fork/execve/exit) in real-time telemetry streams.
// Gate 2 requires lineage DAG queries to complete in <1ms at 10K nodes.
package process_graphs

import (
	"sync"
	"time"

	"sentinel/telemetry/bus/normalizer"
)

// ProcessNode represents a single process in the lineage DAG.
type ProcessNode struct {
	PID          uint32     `json:"pid"`
	PPID         uint32     `json:"ppid"`
	OriginalPPID uint32     `json:"original_ppid"` // Preserved across reparenting
	Comm         string     `json:"comm"`
	ExePath      string     `json:"exe_path"`
	Args         []string   `json:"args"`
	UID          uint32     `json:"uid"`
	GID          uint32     `json:"gid"`
	StartTime    time.Time  `json:"start_time"`
	ExitTime     *time.Time `json:"exit_time,omitempty"`
	ExitCode     int32      `json:"exit_code,omitempty"`
	IsActive     bool       `json:"is_active"`
	Children     []uint32   `json:"children"`
}

// ProcessGraph is a thread-safe process lineage DAG.
// It supports concurrent reads and exclusive writes via sync.RWMutex.
// Maximum capacity is enforced at MaxNodes (default 10,000 per RFC-006).
type ProcessGraph struct {
	mu       sync.RWMutex
	Nodes    map[uint32]*ProcessNode
	MaxNodes int
}

// NewProcessGraph creates a new ProcessGraph with a default capacity of 10,000 nodes.
func NewProcessGraph() *ProcessGraph {
	return &ProcessGraph{
		Nodes:    make(map[uint32]*ProcessNode),
		MaxNodes: 10000,
	}
}

// NewProcessGraphWithCapacity creates a new ProcessGraph with a specified capacity.
func NewProcessGraphWithCapacity(maxNodes int) *ProcessGraph {
	if maxNodes <= 0 {
		maxNodes = 10000
	}
	return &ProcessGraph{
		Nodes:    make(map[uint32]*ProcessNode),
		MaxNodes: maxNodes,
	}
}

// Update processes a normalized telemetry event and updates the graph.
// Handles fork, execve, and exit event types.
func (g *ProcessGraph) Update(evt *normalizer.Event) {
	g.mu.Lock()
	defer g.mu.Unlock()

	switch evt.EventType {
	case "fork":
		g.handleFork(evt)
	case "execve":
		g.handleExecve(evt)
	case "exit":
		g.handleExit(evt)
	}
}

func (g *ProcessGraph) handleFork(evt *normalizer.Event) {
	childPID := evt.PID
	ppid := evt.PPID

	// Enforce capacity limit — prune oldest inactive nodes if at capacity
	if len(g.Nodes) >= g.MaxNodes {
		g.pruneOldestInactive()
	}

	node := &ProcessNode{
		PID:          childPID,
		PPID:         ppid,
		OriginalPPID: ppid,
		Comm:         evt.Comm,
		ExePath:      evt.ExePath,
		UID:          evt.UID,
		GID:          evt.GID,
		StartTime:    evt.Timestamp,
		IsActive:     true,
		Children:     []uint32{},
	}
	g.Nodes[childPID] = node

	// Register child in parent
	if parent, exists := g.Nodes[ppid]; exists {
		parent.Children = append(parent.Children, childPID)
	}
}

func (g *ProcessGraph) handleExecve(evt *normalizer.Event) {
	pid := evt.PID
	node, exists := g.Nodes[pid]
	if !exists {
		// Process appeared without a fork (e.g., we started monitoring mid-lifecycle)
		if len(g.Nodes) >= g.MaxNodes {
			g.pruneOldestInactive()
		}
		node = &ProcessNode{
			PID:          pid,
			PPID:         evt.PPID,
			OriginalPPID: evt.PPID,
			StartTime:    evt.Timestamp,
			IsActive:     true,
			Children:     []uint32{},
		}
		g.Nodes[pid] = node
	}

	node.Comm = evt.Comm
	node.ExePath = evt.ExePath
	node.UID = evt.UID
	node.GID = evt.GID

	// Extract args from payload
	if argsVal, ok := evt.Payload["args"]; ok {
		if argsSlice, ok := argsVal.([]interface{}); ok {
			node.Args = make([]string, 0, len(argsSlice))
			for _, v := range argsSlice {
				if s, ok := v.(string); ok {
					node.Args = append(node.Args, s)
				}
			}
		}
	}
}

func (g *ProcessGraph) handleExit(evt *normalizer.Event) {
	pid := evt.PID
	node, exists := g.Nodes[pid]
	if !exists {
		return
	}

	node.IsActive = false
	now := evt.Timestamp
	node.ExitTime = &now
	if codeVal, ok := evt.Payload["exit_code"]; ok {
		if f, ok := codeVal.(float64); ok {
			node.ExitCode = int32(f)
		}
	}

	// Reparent children to init (PID 1) — preserving OriginalPPID
	for _, childPID := range node.Children {
		if child, childExists := g.Nodes[childPID]; childExists && child.IsActive {
			child.PPID = 1
		}
	}
}

// pruneOldestInactive removes the oldest inactive node to make room.
func (g *ProcessGraph) pruneOldestInactive() {
	var oldestPID uint32
	var oldestTime time.Time
	found := false

	for pid, node := range g.Nodes {
		if !node.IsActive {
			if !found || node.StartTime.Before(oldestTime) {
				oldestPID = pid
				oldestTime = node.StartTime
				found = true
			}
		}
	}

	if found {
		// Remove from parent's children list
		if parent, exists := g.Nodes[g.Nodes[oldestPID].PPID]; exists {
			for i, childPID := range parent.Children {
				if childPID == oldestPID {
					parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
					break
				}
			}
		}
		delete(g.Nodes, oldestPID)
	}
}

// GetLineage returns the ancestor chain for a given PID (BFS upward).
// Performance target: <1ms for 10K node graph (Gate 2).
func (g *ProcessGraph) GetLineage(pid uint32) []*ProcessNode {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var lineage []*ProcessNode
	visited := make(map[uint32]bool)
	current := pid

	for {
		node, exists := g.Nodes[current]
		if !exists || visited[current] {
			break
		}
		visited[current] = true
		lineage = append(lineage, node)

		// Stop at init or self-parent
		if node.PPID == 0 || node.PPID == current {
			break
		}
		current = node.PPID
	}

	return lineage
}

// GetDescendants returns all descendant processes of a given PID (BFS downward).
func (g *ProcessGraph) GetDescendants(pid uint32) []*ProcessNode {
	g.mu.RLock()
	defer g.mu.RUnlock()

	var descendants []*ProcessNode
	visited := make(map[uint32]bool)
	queue := []uint32{pid}

	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]

		if visited[current] {
			continue
		}
		visited[current] = true

		node, exists := g.Nodes[current]
		if !exists {
			continue
		}

		// Don't include the root PID itself
		if current != pid {
			descendants = append(descendants, node)
		}

		queue = append(queue, node.Children...)
	}

	return descendants
}

// GetNode returns a process node by PID, or nil if not found.
func (g *ProcessGraph) GetNode(pid uint32) *ProcessNode {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return g.Nodes[pid]
}

// NodeCount returns the number of nodes in the graph.
func (g *ProcessGraph) NodeCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	return len(g.Nodes)
}

// ActiveCount returns the number of active (non-exited) processes.
func (g *ProcessGraph) ActiveCount() int {
	g.mu.RLock()
	defer g.mu.RUnlock()
	count := 0
	for _, node := range g.Nodes {
		if node.IsActive {
			count++
		}
	}
	return count
}

// Prune removes all inactive nodes older than the given duration.
func (g *ProcessGraph) Prune(olderThan time.Duration) int {
	g.mu.Lock()
	defer g.mu.Unlock()

	cutoff := time.Now().Add(-olderThan)
	pruned := 0
	toDelete := make([]uint32, 0)

	for pid, node := range g.Nodes {
		if !node.IsActive && node.ExitTime != nil && node.ExitTime.Before(cutoff) {
			toDelete = append(toDelete, pid)
		}
	}

	for _, pid := range toDelete {
		// Remove from parent's children list
		if parent, exists := g.Nodes[g.Nodes[pid].PPID]; exists {
			for i, childPID := range parent.Children {
				if childPID == pid {
					parent.Children = append(parent.Children[:i], parent.Children[i+1:]...)
					break
				}
			}
		}
		delete(g.Nodes, pid)
		pruned++
	}

	return pruned
}
