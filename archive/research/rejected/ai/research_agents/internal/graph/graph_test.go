package graph

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/types"
	"testing"
)

func TestGraphAgent(t *testing.T) {
	agent := NewGraphAgent()
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	ev := types.TelemetryEvent{
		PID:      100,
		PPID:     1,
		Comm:     "test",
		ExePath:  "/bin/test",
		Category: "process",
	}

	err := agent.UpdateGraph(ev)
	if err != nil {
		t.Fatalf("failed to update graph: %v", err)
	}

	dag := agent.GetAttackDAG()
	if len(dag.Nodes) < 2 { // 100 and its parent 1
		t.Errorf("expected at least 2 nodes, got %d", len(dag.Nodes))
	}

	if _, ok := dag.Nodes["100"]; !ok {
		t.Error("node 100 missing from DAG")
	}
}
