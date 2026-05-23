package guard

import (
	"os"
	"testing"

	"phoenix/bus"
)

func TestReplayEquivalenceAndSequenceAllocation(t *testing.T) {
	// Create a temporary JSONL events file with sample data out of order
	tempFile, err := os.CreateTemp("", "test_events_*.jsonl")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	eventsData := []string{
		`{"timestamp": "2026-05-21T11:03:19.818000Z", "event_id": "e002", "category": "process", "event_type": "execve", "host_id": "h1", "pid": 1002, "uid": 501, "gid": 20, "comm": "sh", "exe_path": "/bin/sh", "payload": {"entropy_score": 3.2}}`,
		`{"timestamp": "2026-05-21T11:03:19.817000Z", "event_id": "e001", "category": "process", "event_type": "execve", "host_id": "h1", "pid": 1001, "uid": 501, "gid": 20, "comm": "sh", "exe_path": "/bin/sh", "payload": {"entropy_score": 3.2}}`,
		`{"timestamp": "2026-05-21T11:03:19.819000Z", "event_id": "e003", "category": "process", "event_type": "execve", "host_id": "h1", "pid": 1003, "uid": 501, "gid": 20, "comm": "sh", "exe_path": "/bin/sh", "payload": {"entropy_score": 7.9}}`,
	}

	for _, line := range eventsData {
		if _, err := tempFile.WriteString(line + "\n"); err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}
	}
	tempFile.Close()

	b := bus.NewBus()
	// Replay 1 (Natural order from temp file)
	adapter1 := NewGuardAdapter(b, tempFile.Name(), ModeExact, 1.0, 42)
	evs1, err := adapter1.FetchEvents()
	if err != nil {
		t.Fatalf("FetchEvents failed: %v", err)
	}

	// Replay 2 (Create another temp file but with shuffled lines)
	tempFileShuffled, err := os.CreateTemp("", "test_events_shuffled_*.jsonl")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFileShuffled.Name())

	shuffledData := []string{eventsData[1], eventsData[2], eventsData[0]} // different line order
	for _, line := range shuffledData {
		if _, err := tempFileShuffled.WriteString(line + "\n"); err != nil {
			t.Fatalf("failed to write to temp file: %v", err)
		}
	}
	tempFileShuffled.Close()

	adapter2 := NewGuardAdapter(b, tempFileShuffled.Name(), ModeExact, 1.0, 42)
	evs2, err := adapter2.FetchEvents()
	if err != nil {
		t.Fatalf("FetchEvents failed: %v", err)
	}

	// Verify both runs produce the same deterministic sequence
	if len(evs1) != len(evs2) {
		t.Fatalf("Mismatch in event count: %d vs %d", len(evs1), len(evs2))
	}

	for i := range evs1 {
		if evs1[i].EventID != evs2[i].EventID {
			t.Errorf("Mismatch at index %d: EventID %s vs %s", i, evs1[i].EventID, evs2[i].EventID)
		}
		if evs1[i].SeqID != evs2[i].SeqID {
			t.Errorf("Mismatch at index %d: SeqID %d vs %d", i, evs1[i].SeqID, evs2[i].SeqID)
		}
		if evs1[i].SequenceNo != evs2[i].SequenceNo {
			t.Errorf("Mismatch at index %d: SequenceNo %d vs %d", i, evs1[i].SequenceNo, evs2[i].SequenceNo)
		}
	}
}

func TestDeterministicSequenceAllocationFaultMode(t *testing.T) {
	// Verify that fault mode (dropping events randomly) respects the deterministic random seed
	tempFile, err := os.CreateTemp("", "test_events_fault_*.jsonl")
	if err != nil {
		t.Fatalf("failed to create temp file: %v", err)
	}
	defer os.Remove(tempFile.Name())

	// Write 50 events
	for i := 0; i < 50; i++ {
		line := `{"timestamp": "2026-05-21T11:03:19.818Z", "event_id": "e", "category": "process", "event_type": "execve", "host_id": "h1", "pid": 1, "uid": 501, "gid": 20, "comm": "sh", "exe_path": "/bin/sh", "payload": {"entropy_score": 3.2}}`
		tempFile.WriteString(line + "\n")
	}
	tempFile.Close()

	b := bus.NewBus()
	// Two adapters with the exact same seed in ModeFault
	adapter1 := NewGuardAdapter(b, tempFile.Name(), ModeFault, 1.0, 999)
	adapter2 := NewGuardAdapter(b, tempFile.Name(), ModeFault, 1.0, 999)

	evs1, _ := adapter1.FetchEvents()
	evs2, _ := adapter2.FetchEvents()

	if len(evs1) != len(evs2) {
		t.Fatalf("Fault mode output length mismatch: %d vs %d", len(evs1), len(evs2))
	}

	for i := range evs1 {
		if evs1[i].SeqID != evs2[i].SeqID {
			t.Errorf("SeqID mismatch at %d: %d vs %d", i, evs1[i].SeqID, evs2[i].SeqID)
		}
	}
}
