/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
package ledger

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

func TestDeterministicPruning(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 1024)
	l := NewLedger(alloc)

	for i := 0; i < 100; i++ {
		l.AddEntry("E", "ROOT", uint64(i), []byte("data"))
	}

	// Prune anything older than the last 10 entries
	// Note: We need to update l.Counter for Prune to work if we want it to
	// use ticks. Actually, the current Prune implementation uses l.Counter.
	l.Counter = 100

	removed, err := l.Prune(10)
	if err != nil {
		t.Fatalf("Prune failed: %v", err)
	}

	if removed != 90 {
		t.Errorf("Expected 90 removed entries, got %d", removed)
	}

	if len(l.Entries) != 10 {
		t.Errorf("Expected 10 entries remaining, got %d", len(l.Entries))
	}
}

func TestLedgerAllocationBounding(t *testing.T) {
	// Total 1000 bytes, max 100 objects
	alloc := resource.NewBoundedAllocator(1000, 100)
	l := NewLedger(alloc)

	// Add entries until we hit the limit
	if err := l.AddEntry("EV1", "SRC1", 0, []byte("payload1")); err != nil {
		t.Errorf("Initial append failed: %v", err)
	}

	if err := l.AddEntry("EV2", "SRC2", 1, []byte("payload2")); err != nil {
		t.Errorf("Second append failed: %v", err)
	}

	// Large payload to trigger OOM
	largePayload := make([]byte, 2000)
	err := l.AddEntry("EV3", "SRC3", 2, largePayload)
	if err == nil {
		t.Error("Expected OOM error for large payload, but got nil")
	}

	// Verify we can still append small payloads
	if err := l.AddEntry("EV4", "SRC4", 3, []byte("payload4")); err != nil {
		t.Errorf("Small append failed after OOM: %v", err)
	}
}
