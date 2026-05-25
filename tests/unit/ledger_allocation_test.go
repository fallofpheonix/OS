package unit

import (
	"fmt"
	"testing"
)

type mockAllocator struct {
	count  uint64
	failAt uint64
}

func (m *mockAllocator) Allocate(bytes uint64) error {
	m.count++
	if m.failAt > 0 && m.count == m.failAt {
		return fmt.Errorf("simulated allocation failure")
	}
	return nil
}

func (m *mockAllocator) Deallocate(bytes uint64) {}

func TestLedgerSequenceGapPrevention(t *testing.T) {
	alloc := &mockAllocator{failAt: 3}
	l := NewLedger(alloc)

	// Insertion 1: Success
	if err := l.AddEntry("EV1", "SRC1", []byte("payload1")); err != nil {
		t.Fatalf("Expected success on first insertion, got %v", err)
	}

	// Insertion 2: Success
	if err := l.AddEntry("EV2", "SRC2", []byte("payload2")); err != nil {
		t.Fatalf("Expected success on second insertion, got %v", err)
	}

	// Insertion 3: Failure (Mocked)
	err := l.AddEntry("EV3", "SRC3", []byte("payload3"))
	if err == nil {
		t.Fatal("Expected failure on third insertion, but it succeeded")
	}

	// Insertion 4: Success
	if err := l.AddEntry("EV4", "SRC4", []byte("payload4")); err != nil {
		t.Fatalf("Expected success on fourth insertion, got %v", err)
	}

	// Verify Sequence Continuity
	entries := l.SortedEntries()
	if len(entries) != 3 {
		t.Errorf("Expected 3 successful entries in ledger, got %d", len(entries))
	}

	// Check logical ticks: 0, 1, 2
	expectedTicks := []uint64{0, 1, 2}
	for i, entry := range entries {
		if entry.LogicalTick != expectedTicks[i] {
			t.Errorf("Sequence gap detected! Entry at pos %d has logical tick %d, expected %d", i, entry.LogicalTick, expectedTicks[i])
		}
	}
}
