package ledger

import (
	"bytes"
	"fmt"
	"os"
	"testing"
)

type MockApplier struct {
	FailAtTick uint64
}

func (m *MockApplier) Apply(stateBefore []byte, payload []byte) ([]byte, error) {
	// Simple mock behavior: stateAfter = payload
	// If it's the failure tick, return something else
	if m.FailAtTick > 0 && bytes.Contains(payload, []byte(fmt.Sprintf("%d", m.FailAtTick))) {
		return []byte("CORRUPT_STATE"), nil
	}
	return payload, nil
}

func TestReplayerSemanticVerification(t *testing.T) {
	path := "test_semantic.log"
	defer os.Remove(path)

	p, err := NewPersistor(path)
	if err != nil { t.Fatal(err) }
	p.WriteHeader(LedgerFileHeader{Version: "1.0"})

	l := NewLedger(nil).WithPersistor(p)
	
	// Add 3 valid entries where StateAfter == Payload
	for i := 1; i <= 3; i++ {
		payload := []byte(fmt.Sprintf("STATE_%d", i))
		err := l.AddEntryV2(fmt.Sprintf("E%d", i), "CAUSE", uint64(i), payload, "", nil, payload, "1.0")
		if err != nil { t.Fatal(err) }
	}

	replayer, _ := NewReplayer(path)

	// 1. Replay with matching applier
	applier := &MockApplier{}
	_, err = replayer.Replay(applier)
	if err != nil {
		t.Errorf("Replay failed with valid applier: %v", err)
	}

	// 2. Replay with diverging applier (fails at tick 2)
	divergingApplier := &MockApplier{FailAtTick: 2}
	_, err = replayer.Replay(divergingApplier)
	if err == nil {
		t.Error("Expected semantic divergence error, got nil")
	} else {
		t.Logf("Got expected semantic error: %v", err)
	}
}
