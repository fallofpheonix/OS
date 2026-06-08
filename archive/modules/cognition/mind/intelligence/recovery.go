/**
 * FILE: recovery.go
 * PATH: Phoenix.Cognition/intelligence/recovery.go
 *
 * PURPOSE:
 * Implements Phase 4F Recovery Validation.
 * Manages the transition to "Known Good" states and enforces the max rollback
 * protection (Axiom 3).
 *
 * SUBSYSTEM:
 * Cognition / Intelligence / Recovery
 */

package intelligence

import (
	"fmt"
	"log"
	"sync"
)

// RecoveryManager handles system restoration after containment failures.
type RecoveryManager struct {
	mu sync.RWMutex

	RollbackAttempts map[string]int // WorkloadID -> Count
	MaxAttempts      int
}

// NewRecoveryManager initializes the recovery substrate.
func NewRecoveryManager() *RecoveryManager {
	return &RecoveryManager{
		RollbackAttempts: make(map[string]int),
		MaxAttempts:      3, // Mandate: Max 3 attempts
	}
}

// TriggerRollback attempts to restore a workload to its last known good state.
// Enforces "Rollback Protection" axiom.
func (rm *RecoveryManager) TriggerRollback(workloadID string) error {
	rm.mu.Lock()
	defer rm.mu.Unlock()

	rm.RollbackAttempts[workloadID]++
	attempts := rm.RollbackAttempts[workloadID]

	if attempts > rm.MaxAttempts {
		log.Printf("[RECOVERY] QUENCH TRIGGERED for workload %s: Max rollback attempts reached.", workloadID)
		return fmt.Errorf("MAX_ROLLBACK_EXCEEDED")
	}

	log.Printf("[RECOVERY] Attempting rollback %d/%d for workload %s...", attempts, rm.MaxAttempts, workloadID)

	// In a full implementation, this would:
	// 1. Terminate current workload.
	// 2. Fetch last "Known Good" Fact from SQLite.
	// 3. Restart process with injected state.

	return nil
}

// Reset attempts for a successful recovery.
func (rm *RecoveryManager) Reset(workloadID string) {
	rm.mu.Lock()
	defer rm.mu.Unlock()
	delete(rm.RollbackAttempts, workloadID)
}
