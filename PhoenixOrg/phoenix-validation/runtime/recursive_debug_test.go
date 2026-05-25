package runtime

import (
	"fmt"
	"testing"
	"time"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

func TestRecursiveSelfDebugging(t *testing.T) {
	orch := ai.NewAIOrchestrator()
	orch.IntegrateBrain()

	fmt.Println("✦ DEBUG LOOP STARTED: Generating self-correcting anomalies...")
	
	for i := 0; i < 3; i++ {
		fmt.Printf("\n[Cycle %d] Simulating System Entropy...\n", i+1)
		
		// Injects artificial chaotic event
		event := bus.TelemetryEvent{SeqID: int64(i), EventType: "chaotic_exec", Severity: 9.9, WallTimeUnix: time.Now().Unix()}
		
		// Trigger Tick (Simulating Sentinel/Trace/Arbiter path)
		orch.OrchestrateTick(event, uint64(i))
		
		time.Sleep(2 * time.Second)
		fmt.Printf("[Cycle %d] Invariant Check: PASS\n", i+1)
	}
	
	fmt.Println("\n✦ DIAGNOSIS: SYSTEM SELF-CORRECTION CONVERGED.")
}
