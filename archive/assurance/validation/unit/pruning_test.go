/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"fmt"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
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
