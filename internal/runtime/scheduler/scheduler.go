/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — RISK/READINESS TASK SCHEDULER
//
// The Scheduler prioritizes tasks based on dependency impact, readiness,
// and risk. It's used to determine which tasks should be executed first
// in the hardening and verification pipeline.
//
// WORKFLOW:
//   Scheduler.GetNextTask() → sort by Score → return highest-scoring task
//   Scheduler.GetBlockedTasks() → return all blocked tasks
//
// SCORING FORMULA:
//   Score = (DependencyScore * 0.5 + Readiness * 0.5) / (Risk + 0.1)
//   Higher score = higher priority
//
// ALGORITHM: O(N log N) per GetNextTask() call (sorts all tasks).
// Space: O(N) for the task list.
//
// STATUS: Currently not used by any other module. Orphaned functionality.
// =========================================================================
package scheduler

import (
	"sort"
)

type Task struct {
	ID              string
	DependencyScore float64 // Number of tasks depending on this
	Risk            float64 // 0.0 to 1.0
	Readiness       float64 // 0.0 to 1.0 (e.g. telemetry availability)
	Blocked         bool
}

type Scheduler struct {
	Tasks []Task
}

func (s *Scheduler) Score(t Task) float64 {
	if t.Blocked {
		return -1.0
	}
	// Formula: (Dependency Impact * Readiness) / (Risk + 0.1)
	return (t.DependencyScore*0.5 + t.Readiness*0.5) / (t.Risk + 0.1)
}

func (s *Scheduler) GetNextTask() *Task {
	if len(s.Tasks) == 0 {
		return nil
	}

	sort.Slice(s.Tasks, func(i, j int) bool {
		return s.Score(s.Tasks[i]) > s.Score(s.Tasks[j])
	})

	if s.Score(s.Tasks[0]) < 0 {
		return nil // All blocked
	}

	return &s.Tasks[0]
}

func (s *Scheduler) GetBlockedTasks() []Task {
	var blocked []Task
	for _, t := range s.Tasks {
		if t.Blocked {
			blocked = append(blocked, t)
		}
	}
	return blocked
}
