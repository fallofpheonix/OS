package replay

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"testing"
)

func TestReplayIndexingAndSearch(t *testing.T) {
	events := []bus.TelemetryEvent{
		{SeqID: 1, UID: 0, EventType: "process.exec"},
		{SeqID: 2, UID: 1000, EventType: "network.connect"},
		{SeqID: 3, UID: 0, EventType: "file.open"},
	}

	idx := NewReplayIndex(events)

	// Test Search
	rootEvents := idx.SearchUID(0)
	if len(rootEvents) != 2 {
		t.Errorf("Expected 2 root events, got %d", len(rootEvents))
	}

	// Test Index Coverage
	if len(idx.ByType["network.connect"]) != 1 {
		t.Errorf("Expected 1 network event, got %d", len(idx.ByType["network.connect"]))
	}
}

func TestReplayDiff(t *testing.T) {
	runA := []bus.TelemetryEvent{
		{SeqID: 1, UID: 0, EventType: "process.exec"},
		{SeqID: 2, UID: 1000, EventType: "network.connect"},
	}
	runB := []bus.TelemetryEvent{
		{SeqID: 1, UID: 0, EventType: "process.exec"},
		{SeqID: 2, UID: 1001, EventType: "network.connect"}, // Divergence here
	}

	idxA := NewReplayIndex(runA)
	idxB := NewReplayIndex(runB)

	pos, err := Diff(idxA, idxB)
	if err == nil {
		t.Fatal("Expected divergence error, got nil")
	}
	if pos != 1 {
		t.Errorf("Expected divergence at index 1, got %d", pos)
	}
}
