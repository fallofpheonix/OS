package replay

import (
	"encoding/json"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"testing"
)

func TestFullReplayPipeline(t *testing.T) {
	// 1. Create a mock sequence of events
	p1 := PayloadProcessExec{PPID: 1, Comm: "bash", Args: []string{"/bin/bash"}}
	p1Bytes, _ := json.Marshal(p1)

	p2 := PayloadProcessExec{PPID: 100, Comm: "curl", Args: []string{"curl", "google.com"}}
	p2Bytes, _ := json.Marshal(p2)

	events := []bus.TelemetryEvent{
		{
			SeqID:       1,
			MonotonicNs: 1000,
			PID:         100,
			EventType:   "process.exec",
			Payload:     p1Bytes,
		},
		{
			SeqID:       2,
			MonotonicNs: 2000,
			PID:         200,
			EventType:   "process.exec",
			Payload:     p2Bytes,
		},
	}

	// Sign the events to form a valid hash chain
	SignEvent(&events[0], "genesis")
	SignEvent(&events[1], events[0].Hash)

	// 2. Run Replayer
	replayer := NewReplayer(events)

	// Test without baseline
	err := replayer.Execute(nil)
	if err != nil {
		t.Fatalf("Replay execution failed: %v", err)
	}

	// 3. Test Divergence
	baseline := make([]bus.TelemetryEvent, len(events))
	copy(baseline, events)

	// Success case: matches baseline
	err = replayer.Execute(baseline)
	if err != nil {
		t.Fatalf("Replay should match baseline: %v", err)
	}

	// Failure case: tamper with an event
	events[1].PID = 999
	SignEvent(&events[1], events[0].Hash) // Valid hash chain, but different content

	err = replayer.Execute(baseline)
	if err == nil {
		t.Fatal("Replay should have detected divergence")
	}
	t.Logf("Detected expected divergence: %v", err)
}

func TestReplayFidelity(t *testing.T) {
	ev := bus.TelemetryEvent{SeqID: 1, Hash: "hash1"}
	actual := []bus.TelemetryEvent{ev}
	expected := []bus.TelemetryEvent{ev}

	fidelity := ReplayFidelity(actual, expected)
	if fidelity != 100.0 {
		t.Errorf("Expected 100%% fidelity, got %f", fidelity)
	}

	actual[0].Hash = "tampered"
	fidelity = ReplayFidelity(actual, expected)
	if fidelity >= 100.0 {
		t.Errorf("Expected <100%% fidelity after tampering, got %f", fidelity)
	}
}
