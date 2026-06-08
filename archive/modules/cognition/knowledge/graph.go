/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package knowledge implements the deterministic world model for PhoenixOS,
// encompassing causal graph mapping and high-confidence semantic belief management.
package knowledge

import (
	"encoding/json"
	"sync"

	"github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

// Relationship defines the semantic link type between two facts in the knowledge graph.
type Relationship string

const (
	// RelCausality indicates that one fact is a direct consequence of another.
	RelCausality Relationship = "CAUSED_BY"
	// RelIdentity indicates that two facts refer to the same logical entity.
	RelIdentity Relationship = "IS_SAME_AS"
	// RelAssociation indicates a non-causal statistical correlation between facts.
	RelAssociation Relationship = "ASSOCIATED_WITH"
	// RelSequence indicates a chronological ordering without proven causality.
	RelSequence Relationship = "FOLLOWED_BY"
)

// Edge represents a directional relationship in the Knowledge Graph with an associated confidence weight.
type Edge struct {
	FromID   string             `json:"from_id"`
	ToID     string             `json:"to_id"`
	Relation Relationship       `json:"relation"`
	Weight   phxmath.FixedPoint `json:"weight"` // Confidence score [0.0, 1.0]
}

// Graph implements an interconnected web of verified facts and their causal relationships.
// State is synchronized via a RWMutex to support concurrent graph traversal and updates.
type Graph struct {
	mu    sync.RWMutex
	Nodes map[string]*memory.Fact
	Edges []*Edge
}

// NewGraph initializes an empty knowledge graph with allocated internal structures.
func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]*memory.Fact),
		Edges: make([]*Edge, 0),
	}
}

// AddNode integrates a verified fact into the graph as a node.
// Complexity: O(1) time / O(1) space.
func (g *Graph) AddNode(f *memory.Fact) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Nodes[f.ID] = f
}

// AddEdge creates a semantic relationship between two fact nodes.
// Inputs: fromID, toID (Fact IDs), rel (Relationship type), weight (Confidence).
// Complexity: O(1) time / O(1) space.
func (g *Graph) AddEdge(fromID, toID string, rel Relationship, weight phxmath.FixedPoint) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.Edges = append(g.Edges, &Edge{
		FromID:   fromID,
		ToID:     toID,
		Relation: rel,
		Weight:   weight,
	})
}

// ReconstructFromLedger rebuilds the knowledge graph edges from a ledger event stream.
// This ensures the cognitive understanding is reproducible from the immutable history.
// Inputs: events ([]*ledger.Event) - Slice of chronological ledger events.
// Returns: error if JSON unmarshaling fails.
// Complexity: O(N) where N is the number of relationship events.
func (g *Graph) ReconstructFromLedger(events []*ledger.Event) error {
	for _, e := range events {
		if e.Type != ledger.EventRelationship {
			continue
		}
		var p ledger.RelationshipPayload
		if err := json.Unmarshal(e.Payload, &p); err != nil {
			return err
		}
		g.AddEdge(p.FromID, p.ToID, Relationship(p.Relation), p.Weight)
	}
	return nil
}
