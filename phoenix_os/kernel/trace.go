package kernel

import (
	"encoding/json"
	"fmt"
	"sync"

	"phoenix/bus"
)

// TraceNode represents a single entity in the causal lineage graph.
type TraceNode struct {
	PID        int      `json:"pid"`
	ParentPID  int      `json:"ppid"`
	Binary     string   `json:"binary"`
	Children   []int    `json:"children"`
	Files      []string `json:"files"`
	Sockets    []string `json:"sockets"`
	Confidence float64  `json:"confidence"`
}

// TraceEngine manages the causal lineage graph (L4).
type TraceEngine struct {
	mu    sync.RWMutex
	Nodes map[int]*TraceNode
	Bus   *bus.Bus
}

func NewTraceEngine(b *bus.Bus) *TraceEngine {
	return &TraceEngine{
		Nodes: make(map[int]*TraceNode),
		Bus:   b,
	}
}

func (t *TraceEngine) Start() {
	events := t.Bus.Subscribe("system.events.normalized")
	go func() {
		for event := range events {
			t.handleEvent(event)
		}
	}()
}

func (t *TraceEngine) handleEvent(event bus.TelemetryEvent) {
	var ne NormalizedEvent
	if err := json.Unmarshal(event.Payload, &ne); err != nil {
		return
	}

	t.mu.Lock()
	defer t.mu.Unlock()

	node, exists := t.Nodes[event.PID]
	if !exists {
		node = &TraceNode{
			PID:        event.PID,
			Confidence: ne.Confidence,
		}
		t.Nodes[event.PID] = node
	}

	switch ne.Type {
	case "PROCESS_START":
		// Binary name could be extracted from a richer payload in the future
		node.Binary = "unknown" 
	case "PROCESS_FORK":
		// Child PID would be in the payload in a real implementation
	case "FILE_ACCESS":
		node.Files = append(node.Files, "access")
	case "NETWORK_EVENT":
		node.Sockets = append(node.Sockets, "active")
	}

	// In a real system, this would trigger an update to the 3-tier Trace Storage.
	fmt.Printf("[TRACE] Updated lineage for PID %d (Type: %s)\n", event.PID, ne.Type)
}
