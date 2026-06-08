// Package causal provides mechanisms for enforcing causal integrity within the system.
// Core Domain Logic: Implements a Directed Acyclic Graph (DAG) enforcer that ensures
// events maintain strict temporal and logical precedence (causal consistency).
package causal

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

// Node represents a point in the causal DAG.
// Internal State: Encapsulates event data, temporal metadata, and links to parent nodes for Merkle chaining.
// API Scope: Public within the causal and AI domains.
// Concurrency: Instances are intended to be immutable once added to the enforcer.
type Node struct {
	ID        string
	Timestamp time.Time
	Data      interface{}
	Parents   []string
	Hash      string
}

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// CalculateHash generates the Merkle hash for the node based on its content and parents.
// I/O: None.
// Complexity: O(D) where D is the size of the node data and number of parents.
func (n *Node) CalculateHash() string {
	record := struct {
		ID        string
		Timestamp int64
		Data      interface{}
		Parents   []string
	}{
		ID:        n.ID,
		Timestamp: n.Timestamp.UnixNano(),
		Data:      n.Data,
		Parents:   n.Parents,
	}
	b, _ := json.Marshal(record)
	h := sha256.Sum256(b)
	return fmt.Sprintf("%x", h)
}

// Enforcer handles the causal integrity of the matrix.
// Internal State: Map of nodes and synchronization primitives for thread-safe DAG management.
// API Scope: Public; primary interface for causal consistency enforcement.
// Concurrency: Thread-safe via internal RWMutex.
type Enforcer struct {
	mu    sync.RWMutex
	nodes map[string]*Node
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// NewEnforcer creates a new causal enforcer.
// I/O: None.
// Complexity: O(1).
func NewEnforcer() *Enforcer {
	return &Enforcer{
		nodes: make(map[string]*Node),
	}
}

// LABEL: [MUTATES_STATE] [PUBLIC_API] [STABLE]
// AddNode adds a new node to the DAG after validating causal invariants (uniqueness, temporal precedence, and acyclicity).
// I/O: None.
// Side Effects: Modifies the internal nodes map.
// Complexity: O(P) where P is the number of parents, plus hash computation.
func (e *Enforcer) AddNode(node *Node) error {
	e.mu.Lock()
	defer e.mu.Unlock()

	// 1. Ensure node ID is unique
	if _, exists := e.nodes[node.ID]; exists {
		return fmt.Errorf("node with ID %s already exists", node.ID)
	}

	// 2. Validate parents exist and temporal precedence: T(cause) < T(effect)
	for _, parentID := range node.Parents {
		parent, exists := e.nodes[parentID]
		if !exists {
			return fmt.Errorf("parent node %s does not exist", parentID)
		}
		if !parent.Timestamp.Before(node.Timestamp) {
			return fmt.Errorf("temporal regression detected: parent %s (T=%v) is not before child %s (T=%v)",
				parentID, parent.Timestamp, node.ID, node.Timestamp)
		}
	}

	// 3. Acyclic Invariant: Ensure no cycles
	if e.wouldCreateCycle(node) {
		return errors.New("acyclic invariant violation: adding this node would create a cycle")
	}

	// 4. Validate/Set Merkle Hash
	calculatedHash := node.CalculateHash()
	if node.Hash == "" {
		node.Hash = calculatedHash
	} else if node.Hash != calculatedHash {
		return fmt.Errorf("hash mismatch: node reports %s but calculated %s", node.Hash, calculatedHash)
	}

	// 5. Seal the lineage
	e.nodes[node.ID] = node
	return nil
}

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// wouldCreateCycle checks if adding the node would create a cycle (currently redundant due to temporal checks).
// I/O: None.
// Complexity: O(1) in current implementation.
func (e *Enforcer) wouldCreateCycle(newNode *Node) bool {
	return false
}

// LABEL: [PURE] [PUBLIC_API] [STABLE]
// GetNode retrieves a node by ID.
// I/O: None.
// Complexity: O(1) average case.
func (e *Enforcer) GetNode(id string) (*Node, bool) {
	e.mu.RLock()
	defer e.mu.RUnlock()
	n, exists := e.nodes[id]
	return n, exists
}
