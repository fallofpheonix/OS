package telemetry

import (
	"testing"
)

func TestTelemetryAgent(t *testing.T) {
	agent := NewTelemetryAgent(10)
	if agent == nil {
		t.Fatal("failed to create agent")
	}

	err := agent.Start()
	if err != nil {
		t.Fatalf("failed to start agent: %v", err)
	}

	ev := agent.GenerateMockEvent()
	if ev.EventID == "" {
		t.Error("mock event has no ID")
	}

	lineage, err := agent.GetLineage(ev.PID)
	if err != nil {
		t.Fatalf("failed to get lineage: %v", err)
	}

	if len(lineage) == 0 {
		t.Error("lineage should not be empty for a recorded event")
	}

	err = agent.Stop()
	if err != nil {
		t.Fatalf("failed to stop agent: %v", err)
	}
}
