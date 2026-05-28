package process_graphs

import (
	"sync"
)

type NodeType string

const (
	Process NodeType = "PROCESS"
	File    NodeType = "FILE"
	Network NodeType = "NETWORK"
)

type Node struct {
	ID        string
	Type      NodeType
	Timestamp int64
}

type Graph struct {
	Mu    sync.RWMutex
	Nodes map[string]Node
	Edges map[string][]string
}

func NewGraph() *Graph {
	return &Graph{
		Nodes: make(map[string]Node),
		Edges: make(map[string][]string),
	}
}

func (g *Graph) AddNode(id string, t NodeType, ts int64) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Nodes[id] = Node{ID: id, Type: t, Timestamp: ts}
}

func (g *Graph) AddEdge(from, to string) {
	g.Mu.Lock()
	defer g.Mu.Unlock()
	g.Edges[from] = append(g.Edges[from], to)
}
