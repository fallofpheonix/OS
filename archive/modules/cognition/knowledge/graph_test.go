/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package knowledge

import (
	"encoding/json"
	"testing"

	"github.com/fallofpheonix/phoenix/cognition/memory"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

func TestGraph_BasicOperations(t *testing.T) {
	g := NewGraph()

	f1 := &memory.Fact{ID: "F1"}
	f2 := &memory.Fact{ID: "F2"}

	g.AddNode(f1)
	g.AddNode(f2)

	if len(g.Nodes) != 2 {
		t.Errorf("expected 2 nodes, got %d", len(g.Nodes))
	}

	g.AddEdge("F1", "F2", RelCausality, 0.95)

	if len(g.Edges) != 1 {
		t.Fatal("expected 1 edge, got 0")
	}

	edge := g.Edges[0]
	if edge.FromID != "F1" || edge.ToID != "F2" || edge.Weight != 0.95 {
		t.Errorf("edge property mismatch")
	}
}

func TestGraph_Reconstruction(t *testing.T) {
	graph := NewGraph()

	// Pre-populate with facts (usually done by memory cycle)
	fact1 := &memory.Fact{ID: "FACT_A"}
	fact2 := &memory.Fact{ID: "FACT_B"}
	graph.AddNode(fact1)
	graph.AddNode(fact2)

	// Simulate relationship event.
	relPayload := ledger.RelationshipPayload{
		FromID:   "FACT_A",
		ToID:     "FACT_B",
		Relation: string(RelCausality),
		Weight:   0.88,
	}
	payload, _ := json.Marshal(relPayload)

	events := []*ledger.Event{
		{
			Type:    ledger.EventRelationship,
			Payload: payload,
		},
	}

	// Simulate an ignored event type
	events = append([]*ledger.Event{{Type: ledger.EventFact}}, events...)

	if err := graph.ReconstructFromLedger(events); err != nil {
		t.Fatalf("reconstruction failed: %v", err)
	}

	// Simulate bad payload
	badEvents := []*ledger.Event{{Type: ledger.EventRelationship, Payload: []byte("{bad-json")}}
	if err := graph.ReconstructFromLedger(badEvents); err == nil {
		t.Errorf("expected error on bad json payload")
	}

	if len(graph.Edges) != 1 {
		t.Errorf("expected 1 edge, got %d", len(graph.Edges))
	}

	edge := graph.Edges[0]
	if edge.Relation != RelCausality {
		t.Errorf("relation mismatch: expected CAUSED_BY, got %s", edge.Relation)
	}
}
