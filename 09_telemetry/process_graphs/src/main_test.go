package main

import (
	"testing"
)

func TestGraphLineage(t *testing.T) {
	g := NewGraph()
	g.AddNode("root", Process)
	g.AddNode("child", Process)
	g.AddNode("grandchild", Process)

	g.AddEdge("root", "child")
	g.AddEdge("child", "grandchild")

	lineage := g.GetLineage("root")
	if len(lineage) != 3 {
		t.Errorf("Expected lineage length 3, got %d", len(lineage))
	}
}

func BenchmarkGraphInsertion(b *testing.B) {
	g := NewGraph()
	for i := 0; i < b.N; i++ {
		id := string(rune(i))
		g.AddNode(id, Process)
	}
}

func BenchmarkGraphTraversal(b *testing.B) {
	g := NewGraph()
	g.AddNode("0", Process)
	for i := 1; i <= 15; i++ {
		from := string(rune(i - 1))
		to := string(rune(i))
		g.AddNode(to, Process)
		g.AddEdge(from, to)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = g.GetLineage("0")
	}
}
