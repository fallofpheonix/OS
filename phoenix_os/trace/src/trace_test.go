package main

import "testing"

func TestTrace(t *testing.T) {
	g := NewGraph()
	g.AddNode("1", "process")
	g.AddNode("2", "process")
	g.AddEdge("1", "2")
	
	path := g.Trace("1")
	if len(path) != 2 {
		t.Errorf("Expected path length 2, got %d", len(path))
	}
}
