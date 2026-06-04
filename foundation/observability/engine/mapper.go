/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 4 — CAUSAL LINEAGE TRACKING (Layer 1.5)
//
// The GraphEngine maps process relationships as a directed acyclic graph (DAG).
// It tracks parent-child process relationships and maintains lineage chains
// for forensic analysis and threat detection.
//
// WORKFLOW:
//   eBPF captures fork/execve → GraphEngine.HandleFork(pid, ppid)
//     → Create node for new process
//     → Create edge from parent to child
//     → If parent missing: re-parent to 0xDEADBEEF (ORPHAN_MONITOR)
//
// DATA STRUCTURES:
//   ProcessNode: PID, PPID, Binary, StartTime, Status, MerkleHash, Children
//   GraphEngine: map[PID]*ProcessNode (in-memory DAG)
//
// ORPHAN HANDLING: Processes whose parent is not tracked are re-parented
// to a virtual node 0xDEADBEEF (ORPHAN_MONITOR). This ensures no process
// is lost from the lineage graph.
// =========================================================================
package engine

import (
	"sync"
	"time"
)

type ProcessNode struct {
	PID        int
	PPID       int
	Binary     string
	StartTime  time.Time
	IsIsolated bool
	Status     string
	MerkleHash string
	Children   []int
}

type GraphEngine struct {
	mu    sync.RWMutex
	nodes map[int]*ProcessNode
}

func NewGraphEngine() *GraphEngine {
	e := &GraphEngine{
		nodes: make(map[int]*ProcessNode),
	}
	// Initialize the ORPHAN_MONITOR virtual node
	e.nodes[0xDEADBEEF] = &ProcessNode{
		PID:    0xDEADBEEF,
		Binary: "ORPHAN_MONITOR",
		Status: "ACTIVE",
	}
	return e
}

func (e *GraphEngine) GetNode(pid int) (*ProcessNode, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	n, exists := e.nodes[pid]
	return n, exists
}

func (e *GraphEngine) HandleFork(pid, ppid int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.nodes[ppid]; !exists {
		ppid = 0xDEADBEEF
	}

	child := &ProcessNode{
		PID:       pid,
		PPID:      ppid,
		StartTime: time.Now(),
		Status:    "ACTIVE",
		Children:  []int{},
	}
	e.nodes[pid] = child

	if parent, ok := e.nodes[ppid]; ok {
		parent.Children = append(parent.Children, pid)
	}
}

func (e *GraphEngine) HandleIsolation(pid int) {
	e.mu.Lock()
	defer e.mu.Unlock()

	node, exists := e.nodes[pid]
	if !exists {
		return
	}

	node.IsIsolated = true
	node.Status = "ISOLATED"

	// virtualPID for isolated child
	virtualPID := pid*1000 + 1
	virtualChild := &ProcessNode{
		PID:       virtualPID,
		PPID:      pid,
		StartTime: time.Now(),
		Status:    "ACTIVE",
		Children:  []int{},
	}
	e.nodes[virtualPID] = virtualChild
	node.Children = append(node.Children, virtualPID)
}

func (e *GraphEngine) GetCausalChain(pid int) ([]*ProcessNode, error) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	var chain []*ProcessNode
	current := pid

	for current != 0 {
		node, exists := e.nodes[current]
		if !exists {
			break
		}
		chain = append(chain, node)
		current = node.PPID
	}

	return chain, nil
}
