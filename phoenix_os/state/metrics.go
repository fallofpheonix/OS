package state

import (
	"encoding/json"
	"sync"
)

// Metrics captures state transition counters and system telemetry.
type Metrics struct {
	mu                 sync.Mutex
	TransitionCount    int
	IllegalTransitions int
	Rollbacks          int

	SafeEntries     int
	WatchEntries    int
	AlertEntries    int
	ContainEntries  int
	RecoveryEntries int

	ReplayMismatch   int
	RollbackFailures int

	SnapshotCreates  int
	SnapshotRestores int

	RecoveryCount int

	MeanTransitionLatency int
	MeanRollbackLatency   int
}

// GlobalMetrics provides a simple implementation of S6.
var GlobalMetrics Metrics

func (m *Metrics) IncTransition() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.TransitionCount++
}

func (m *Metrics) IncIllegal() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.IllegalTransitions++
}

func (m *Metrics) IncRollback() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.Rollbacks++
	m.RecoveryCount++
}

func (m *Metrics) IncStateEntry(s SystemState) {
	m.mu.Lock()
	defer m.mu.Unlock()
	switch s {
	case StateSafe:
		m.SafeEntries++
	case StateWatch:
		m.WatchEntries++
	case StateAlert:
		m.AlertEntries++
	case StateContain:
		m.ContainEntries++
	case StateRecovery:
		m.RecoveryEntries++
	}
}

func (m *Metrics) IncReplayMismatch() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.ReplayMismatch++
}

func (m *Metrics) UpdateTransitionLatency(ns int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	// Simple running average
	if m.TransitionCount > 0 {
		m.MeanTransitionLatency = (m.MeanTransitionLatency*(m.TransitionCount-1) + ns) / m.TransitionCount
	} else {
		m.MeanTransitionLatency = ns
	}
}

func (m *Metrics) UpdateRollbackLatency(ns int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	if m.Rollbacks > 0 {
		m.MeanRollbackLatency = (m.MeanRollbackLatency*(m.Rollbacks-1) + ns) / m.Rollbacks
	} else {
		m.MeanRollbackLatency = ns
	}
}

func (m *Metrics) IncSnapshotCreates() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.SnapshotCreates++
}

func (m *Metrics) IncSnapshotRestores() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.SnapshotRestores++
}

func (m *Metrics) ExportMetrics() ([]byte, error) {
	m.mu.Lock()
	defer m.mu.Unlock()
	return json.Marshal(m)
}

func (m *Metrics) Reset() {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.TransitionCount = 0
	m.IllegalTransitions = 0
	m.Rollbacks = 0
	m.SafeEntries = 0
	m.WatchEntries = 0
	m.AlertEntries = 0
	m.ContainEntries = 0
	m.RecoveryEntries = 0
	m.ReplayMismatch = 0
	m.RollbackFailures = 0
	m.SnapshotCreates = 0
	m.SnapshotRestores = 0
	m.RecoveryCount = 0
	m.MeanTransitionLatency = 0
	m.MeanRollbackLatency = 0
}
