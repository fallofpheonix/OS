/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package scheduler

import (
	"testing"
)

func TestScheduler(t *testing.T) {
	s := &Scheduler{
		Tasks: []Task{
			{ID: "T1", DependencyScore: 5.0, Risk: 0.1, Readiness: 0.9, Blocked: false}, // High impact, low risk
			{ID: "T2", DependencyScore: 1.0, Risk: 0.8, Readiness: 0.2, Blocked: false}, // Low impact, high risk
			{ID: "T3", DependencyScore: 10.0, Risk: 0.5, Readiness: 0.1, Blocked: true}, // Blocked
		},
	}

	next := s.GetNextTask()
	if next == nil || next.ID != "T1" {
		t.Errorf("Expected T1 as next task, got %v", next)
	}

	blocked := s.GetBlockedTasks()
	if len(blocked) != 1 || blocked[0].ID != "T3" {
		t.Errorf("Expected T3 in blocked tasks, got %v", blocked)
	}
}
