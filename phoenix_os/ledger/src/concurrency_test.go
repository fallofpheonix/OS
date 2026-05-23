package ledger

import (
	"fmt"
	"math/rand"
	"sync"
	"testing"
	"time"

	"phoenix/common/resource"
)

func TestLedgerConcurrency(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 1000)
	l := NewLedger(alloc)

	var wg sync.WaitGroup
	workers := 10
	iterations := 50

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(workerID int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)))

			for j := 0; j < iterations; j++ {
				// Concurrent additions
				eventID := fmt.Sprintf("event-%d-%d", workerID, j)
				causeID := fmt.Sprintf("cause-%d-%d", workerID, j)
				payload := []byte(fmt.Sprintf(`{"worker":%d,"iter":%d}`, workerID, j))
				
				_ = l.AddEntryV2(eventID, causeID, payload, "NORMAL", "SUSPICIOUS", "1.0.0")

				// Concurrent reads/verifications
				if rng.Float64() > 0.7 {
					_ = l.Verify()
				}
				if rng.Float64() > 0.8 {
					_, _ = l.Checkpoint()
				}
				if rng.Float64() > 0.9 {
					_ = l.SortedEntries()
				}
			}
		}(i)
	}

	wg.Wait()
}
