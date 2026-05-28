package security

import (
	"sync"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/PheonixGuard"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"
)

// ConflictStressTest simulates high-frequency conflicting AI directives.
func TestConflictStressHarness(t *testing.T) {
	b := bus.NewBus()
	w := warden.NewWarden(b)
	w.EnableDiagnostics("debug/panic_audit.log")

	// Setup orchestrator
	orch := ai.NewAIOrchestrator()
	
	// Inject high frequency conflicts
	var wg sync.WaitGroup
	numRoutines := 10
	
	start := time.Now()
	for i := 0; i < numRoutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			
			// Inject conflicting telemetry
			for j := 0; j < 5; j++ {
				event := bus.TelemetryEvent{
					SeqID:   int64(id*100 + j),
					PID:     1000 + id,
					Tgid:    uint32(100 + id),
					Nsproxy: uint32(10000 + id), // Valid initial lineage
				}
				
				// Inject anomaly with drifting lineage (lineage mismatch)
				if j%2 == 0 {
					event.Nsproxy = 0 // Trigger invariant violation
				}
				
				orch.OrchestrateTick(event, uint64(j))
				time.Sleep(10 * time.Millisecond)
			}
		}(i)
	}
	wg.Wait()
	duration := time.Since(start)
	
	t.Logf("Conflict stress test completed in %v", duration)
}
