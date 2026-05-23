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
	return (t.DependencyScore * 0.5 + t.Readiness * 0.5) / (t.Risk + 0.1)
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
