package detector

import (
	"testing"
	"time"
	"phoenix/telemetry/events"
	"phoenix/telemetry/process_lineage"
)

func TestDetector(t *testing.T) {
	d := NewDetector()
	g := lineage.NewLineageGraph()
	now := time.Now()

	// Case 1: Benign event
	evt1 := events.Event{
		PID:       100,
		Comm:      "ls",
		EventType: "execve",
		Payload:   map[string]any{"entropy_score": 3.0},
	}
	res1 := d.Analyze(evt1, g)
	if res1.IsThreat {
		t.Error("ls should not be a threat")
	}

	// Case 2: High entropy write from 'gpg'
	g.AddProcess(200, 1, "gpg", "/usr/bin/gpg", now)
	evt2 := events.Event{
		PID:       200,
		Comm:      "gpg",
		EventType: "write",
		Payload:   map[string]any{"entropy_score": 8.5},
	}
	res2 := d.Analyze(evt2, g)
	if !res2.IsThreat {
		t.Errorf("gpg high entropy write should be a threat. Score: %f", res2.ImportanceScore)
	}
	if res2.ImportanceScore < 0.8 {
		t.Errorf("Expected score >= 0.8, got %f", res2.ImportanceScore)
	}

	// Case 3: FastPath
	if !d.FastPath(evt2) {
		t.Error("FastPath should trigger for entropy 8.5")
	}
}
