/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package actuation

import (
	"errors"
	"testing"
	"time"
)

func TestExecutor_Execute_Success(t *testing.T) {
	ex := NewExecutor()
	actionCalled := false
	rollbackCalled := false

	action := func() error {
		actionCalled = true
		return nil
	}
	rollback := RollbackPlan{
		Target: "test-target",
		Action: func() error {
			rollbackCalled = true
			return nil
		},
	}

	err := ex.Execute(action, rollback)
	if err != nil {
		t.Errorf("Expected success, got error: %v", err)
	}
	if !actionCalled {
		t.Error("Action was not called")
	}
	if rollbackCalled {
		t.Error("Rollback was called on success")
	}
}

func TestExecutor_Execute_Timeout(t *testing.T) {
	ex := &EnforcementExecutor{Timeout: 10 * time.Millisecond}
	rollbackCalled := false

	action := func() error {
		time.Sleep(100 * time.Millisecond)
		return nil
	}
	rollback := RollbackPlan{
		Target: "test-target",
		Action: func() error {
			rollbackCalled = true
			return nil
		},
	}

	err := ex.Execute(action, rollback)
	if err == nil {
		t.Error("Expected timeout error, got nil")
	}
	if !rollbackCalled {
		t.Error("Rollback was not called on timeout")
	}
}

func TestExecutor_Execute_Failure(t *testing.T) {
	ex := NewExecutor()
	rollbackCalled := false

	action := func() error {
		return errors.New("hard fail")
	}
	rollback := RollbackPlan{
		Target: "test-target",
		Action: func() error {
			rollbackCalled = true
			return nil
		},
	}

	err := ex.Execute(action, rollback)
	if err == nil {
		t.Error("Expected failure error, got nil")
	}
	if !rollbackCalled {
		t.Error("Rollback was not called on failure")
	}
}
