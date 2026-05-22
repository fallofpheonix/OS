package ledger

import (
	"fmt"
	"testing"

	"phoenix/common/resource"
)

func TestDeterministicPruning(t *testing.T) {
	l := NewLedger(resource.NewBoundedAllocator(1024*1024, 100))

	// 1. Add 10 entries
	for i := 0; i < 10; i++ {
		l.AddEntry(fmt.Sprintf("E%d", i), "PREV", []byte("data"))
	}
	
	// 2. Prune depth 5 (keep only last 5)
	removed, err := l.Prune(5)
	if err != nil {
		t.Fatalf("Pruning failed: %v", err)
	}
	
	if removed != 5 {
		t.Errorf("Expected 5 entries removed, got %d", removed)
	}
	
	if len(l.Entries) != 5 {
		t.Errorf("Expected 5 entries remaining, got %d", len(l.Entries))
	}
	
	// 3. Verify that remaining entries are the latest ones (Ticks 5-9)
	for _, entry := range l.SortedEntries() {
		if entry.LogicalTick < 5 {
			t.Errorf("Stale entry found after pruning: Tick %d", entry.LogicalTick)
		}
	}
}
