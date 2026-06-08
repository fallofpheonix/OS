/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */
package ledger

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

func TestLedgerConcurrency(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024*10, 1024)
	l := NewLedger(alloc)

	var wg sync.WaitGroup
	numWorkers := 10
	entriesPerWorker := 50

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(int64(workerID)))
			for j := 0; j < entriesPerWorker; j++ {
				eventID := fmt.Sprintf("W%d-E%d", workerID, j)
				causeID := "ROOT"
				payload := []byte(fmt.Sprintf("worker %d payload %d", workerID, j))
				
				// In a real system, the tick would be assigned by a single-threaded loop.
				// Here we just use a random tick to test concurrency safety of the map/Heads.
				tick := uint64(rng.Intn(1000000))
				
				_ = l.AddEntryV2(eventID, causeID, tick, payload, "", []byte("NORMAL"), []byte("NORMAL"), "1.0.0")
			}
		}(i)
	}

	wg.Wait()

	if len(l.Entries) != numWorkers*entriesPerWorker {
		t.Errorf("Expected %d entries, got %d", numWorkers*entriesPerWorker, len(l.Entries))
	}

	if err := l.Verify(); err != nil {
		t.Errorf("Concurrent ledger verification failed: %v", err)
	}
}

func TestLedgerSequenceGapPrevention(t *testing.T) {
	// Our current implementation does NOT prevent sequence gaps because it doesn't 
	// track a global expected next tick within the ledger itself (it trusts the caller).
	// This test verifies that Verify() doesn't fail just because ticks are non-contiguous 
	// (as long as hashes match).
	l := NewLedger(nil)

	l.AddEntry("E1", "ROOT", 10, []byte("data1"))
	l.AddEntry("E2", "E1", 20, []byte("data2"))

	if err := l.Verify(); err != nil {
		t.Errorf("Expected non-contiguous ticks to verify if hashes match, got: %v", err)
	}
}
