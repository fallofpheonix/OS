/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package state

import (
	"fmt"
	"time"
)

// Rollback restores the system to the immediately preceding state.
func (r *StateRegistry) Rollback(reason, evidenceID, decisionID string) error {
	start := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	if len(r.History) <= 1 {
		return fmt.Errorf("no history to rollback")
	}

	// Retrieve the previous transition record
	lastRecord := r.History[len(r.History)-1]
	previousState := lastRecord.Previous

	// S5.6: Verify rollback legality
	if !isValidTransition(r.CurrentState, previousState, true) {
		GlobalMetrics.IncIllegal()
		return fmt.Errorf("illegal rollback from %s to %s", r.CurrentState, previousState)
	}

	// Audit record
	record := StateRecord{
		ID:         len(r.History),
		Previous:   r.CurrentState,
		Current:    previousState,
		Timestamp:  time.Now(),
		Reason:     fmt.Sprintf("ROLLBACK: %s (Decision: %s)", reason, decisionID),
		EvidenceID: evidenceID,
	}

	r.CurrentState = previousState
	r.History = append(r.History, record)
	GlobalMetrics.IncRollback()
	GlobalMetrics.IncStateEntry(previousState)

	duration := time.Since(start).Nanoseconds()
	GlobalMetrics.UpdateRollbackLatency(int(duration))

	return nil
}
