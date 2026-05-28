package ordering

import (
	"fmt"
	"sync"
	"testing"
)

func TestEventOrderer_RecordAndGetEvent(t *testing.T) {
	eo := NewEventOrderer()

	// Test initial state
	if eo.LastSequenceID() != 0 {
		t.Errorf("Expected initial LastSequenceID to be 0, got %d", eo.LastSequenceID())
	}

	// Record first event
	payload1 := []byte("event one")
	seq1, err := eo.RecordEvent(payload1)
	if err != nil {
		t.Fatalf("RecordEvent failed for payload1: %v", err)
	}
	if seq1 != 1 {
		t.Errorf("Expected first sequence ID to be 1, got %d", seq1)
	}
	if eo.LastSequenceID() != 1 {
		t.Errorf("Expected LastSequenceID to be 1, got %d", eo.LastSequenceID())
	}

	event1, err := eo.GetEvent(1)
	if err != nil {
		t.Fatalf("GetEvent(1) failed: %v", err)
	}
	if event1.SequenceID != 1 || string(event1.Payload) != string(payload1) {
		t.Errorf("Retrieved event 1 mismatch: %+v", event1)
	}

	// Record second event
	payload2 := []byte("event two")
	seq2, err := eo.RecordEvent(payload2)
	if err != nil {
		t.Fatalf("RecordEvent failed for payload2: %v", err)
	}
	if seq2 != 2 {
		t.Errorf("Expected second sequence ID to be 2, got %d", seq2)
	}
	if eo.LastSequenceID() != 2 {
		t.Errorf("Expected LastSequenceID to be 2, got %d", eo.LastSequenceID())
	}

	event2, err := eo.GetEvent(2)
	if err != nil {
		t.Fatalf("GetEvent(2) failed: %v", err)
	}
	if event2.SequenceID != 2 || string(event2.Payload) != string(payload2) {
		t.Errorf("Retrieved event 2 mismatch: %+v", event2)
	}

	// Test RecordEvent with nil payload
	_, err = eo.RecordEvent(nil)
	if err == nil {
		t.Error("Expected error for nil payload, got nil")
	} else if err.Error() != "event payload cannot be nil" {
		t.Errorf("Expected 'event payload cannot be nil' error, got: %v", err)
	}

	// Test GetEvent for non-existent event
	_, err = eo.GetEvent(99)
	if err == nil {
		t.Error("Expected error for non-existent event, got nil")
	} else if err.Error() != "event with sequence ID 99 not found" {
		t.Errorf("Expected 'event with sequence ID 99 not found' error, got: %v", err)
	}
}

func TestEventOrderer_Concurrency(t *testing.T) {
	eo := NewEventOrderer()
	numGoroutines := 100
	eventsPerGoroutine := 100

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	for i := 0; i < numGoroutines; i++ {
		go func(gid int) {
			defer wg.Done()
			for j := 0; j < eventsPerGoroutine; j++ {
				payload := []byte(fmt.Sprintf("event-%d-%d", gid, j))
				_, _ = eo.RecordEvent(payload)
			}
		}(i)
	}
	wg.Wait()

	finalSequenceID := eo.LastSequenceID()
	expectedTotalEvents := uint64(numGoroutines * eventsPerGoroutine)

	if finalSequenceID != expectedTotalEvents {
		t.Errorf("Expected final sequence ID %d, got %d", expectedTotalEvents, finalSequenceID)
	}

	// Verify that all events from 1 to finalSequenceID exist and are unique
	// This implicitly checks for correct ordering and no skipped IDs.
	foundEvents := make(map[uint64]bool)
	for i := uint64(1); i <= finalSequenceID; i++ {
		event, err := eo.GetEvent(i)
		if err != nil {
			t.Fatalf("Failed to retrieve event with sequence ID %d: %v", i, err)
		}
		if foundEvents[event.SequenceID] {
			t.Errorf("Duplicate event found for sequence ID %d", event.SequenceID)
		}
		foundEvents[event.SequenceID] = true
	}
	if uint64(len(foundEvents)) != finalSequenceID {
		t.Errorf("Expected to find %d unique events, found %d", finalSequenceID, len(foundEvents))
	}
}
