/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package builder

import (
	"sync"
)

// TelemetryBridge acts as the sensory input for the Builder, 
// aggregating panic logs into actionable failure clusters.
type TelemetryBridge struct {
	mu sync.RWMutex
	// FailureClusters maps a failure pattern (e.g., "invariant_violation")
	// to a slice of associated TelemetryEventIDs.
	FailureClusters map[string][]string
}

// NewTelemetryBridge initializes the aggregation engine.
func NewTelemetryBridge() *TelemetryBridge {
	return &TelemetryBridge{
		FailureClusters: make(map[string][]string),
	}
}

// AggregatePanic ingests raw panic logs and maps them to failure clusters.
func (tb *TelemetryBridge) AggregatePanic(pattern string, eventID string) {
	tb.mu.Lock()
	defer tb.mu.Unlock()
	tb.FailureClusters[pattern] = append(tb.FailureClusters[pattern], eventID)
}

// GetTopFailurePattern returns the pattern with the highest failure frequency.
func (tb *TelemetryBridge) GetTopFailurePattern() string {
	tb.mu.RLock()
	defer tb.mu.RUnlock()
	
	topPattern := ""
	maxCount := 0
	
	for pattern, events := range tb.FailureClusters {
		if len(events) > maxCount {
			maxCount = len(events)
			topPattern = pattern
		}
	}
	
	return topPattern
}
