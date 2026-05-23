package warden

import (
	"math/rand"
	"sync"
	"testing"
	"time"

	"phoenix/bus"
)

func TestWardenConcurrency(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	// Consume recovery budget initially
	w.deescalationCount = 2

	var wg sync.WaitGroup
	workers := 10
	iterations := 100

	wg.Add(workers)
	for i := 0; i < workers; i++ {
		go func(workerID int) {
			defer wg.Done()
			rng := rand.New(rand.NewSource(time.Now().UnixNano() + int64(workerID)))

			for j := 0; j < iterations; j++ {
				// Alternating state actuations
				targetState := StateNormal
				if rng.Float64() > 0.5 {
					targetState = StateContained
				}
				class := ClassObserve
				if targetState == StateContained {
					class = ClassLocalIsolate
				}

				// Perform concurrent actuations
				_ = w.Actuate(targetState, class, 0.9, int64(workerID*1000+j), time.Now().Unix(), uint64(workerID*1000+j))

				// Perform concurrent budget resets
				if rng.Float64() > 0.8 {
					w.ResetBudget()
				}
			}
		}(i)
	}

	wg.Wait()
}
