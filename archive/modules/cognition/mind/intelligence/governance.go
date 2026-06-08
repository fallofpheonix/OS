/**
 * FILE: governance.go
 * PATH: Phoenix.Cognition/intelligence/governance.go
 *
 * PURPOSE:
 * Implements PhoenixOS Evidence Governance (Phase 4B).
 * Tracks the delta between Replay and Reality to detect Determinism Leaks
 * and monitors the primary success metric (MTTE).
 *
 * SUBSYSTEM:
 * Cognition / Intelligence / Auditing
 */

package intelligence

import (
	"log"
	"runtime"
	"sync"
	"time"
)

// ResourceUsage tracks the OS overhead.
type ResourceUsage struct {
	RAMBytes  uint64 `json:"ram_bytes"`
	Timestamp int64  `json:"timestamp"`
}

// DeterminismLeak represents a conflict where Replay disagrees with Reality.
type DeterminismLeak struct {
	EventID    string  `json:"event_id"`
	RealityVal float64 `json:"reality_val"`
	ReplayVal  float64 `json:"replay_val"`
	Drift      float64 `json:"drift"`
	Timestamp  int64   `json:"timestamp"`
}

// SensorReputation tracks the reliability of a specific telemetry sensor.
type SensorReputation struct {
	ID         string  `json:"id"`
	Score      float64 `json:"score"` // 0.0 to 1.0
	LastUpdate int64   `json:"last_update"`
}

// GovernanceMonitor tracks OS performance and integrity metrics.
type GovernanceMonitor struct {
	mu sync.RWMutex

	// Integrity Metrics
	Leaks []DeterminismLeak

	// Operational Metrics (MTTE)
	DecisionStart        map[string]time.Time
	ExplanationLatencies []time.Duration

	// Sensor Reputation (Phase 4D)
	Sensors map[string]*SensorReputation

	// Resource Tracking (Phase 4G)
	ResourceHistory []ResourceUsage
}

// NewGovernanceMonitor initializes the auditing substrate.
func NewGovernanceMonitor() *GovernanceMonitor {
	gm := &GovernanceMonitor{
		DecisionStart: make(map[string]time.Time),
		Sensors:       make(map[string]*SensorReputation),
	}

	// Start background resource monitoring
	go gm.monitorResources()

	return gm
}

// monitorResources periodically samples RAM usage.
func (gm *GovernanceMonitor) monitorResources() {
	ticker := time.NewTicker(1 * time.Minute)
	for range ticker.C {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)

		gm.mu.Lock()
		usage := ResourceUsage{
			RAMBytes:  m.Alloc,
			Timestamp: time.Now().Unix(),
		}
		gm.ResourceHistory = append(gm.ResourceHistory, usage)

		// Budget Enforcement (Axiom Check)
		if m.Alloc > 1024*1024*100 { // Example: 100MB threshold
			log.Printf("[GOVERNANCE] WARNING: RAM Budget exceeded: %d bytes", m.Alloc)
		}
		gm.mu.Unlock()
	}
}

// RecordSensorClaim updates a sensor's reputation based on its accuracy.
// Employs EMA-based recovery logic (Exponential Moving Average).
func (gm *GovernanceMonitor) RecordSensorClaim(sensorID string, accurate bool) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	s, ok := gm.Sensors[sensorID]
	if !ok {
		s = &SensorReputation{ID: sensorID, Score: 1.0}
		gm.Sensors[sensorID] = s
	}

	const alpha = 0.01   // EMA coefficient for recovery
	const penalty = 0.05 // Immediate drop for inaccurate claims

	if accurate {
		s.Score = (1-alpha)*s.Score + alpha
	} else {
		// Penalty
		s.Score -= penalty
		if s.Score < 0.1 {
			s.Score = 0.1
		} // Minimum reputation floor
	}
	s.LastUpdate = time.Now().Unix()
}

// GetReputation returns the current score for a sensor.
func (gm *GovernanceMonitor) GetReputation(sensorID string) float64 {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if s, ok := gm.Sensors[sensorID]; ok {
		return s.Score
	}
	return 1.0 // Default for new sensors
}

// RecordDecisionStart marks the beginning of an anomaly detection event.
func (gm *GovernanceMonitor) RecordDecisionStart(decisionID string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()
	gm.DecisionStart[decisionID] = time.Now()
}

// RecordExplanationComplete calculates the MTTE for a specific decision.
func (gm *GovernanceMonitor) RecordExplanationComplete(decisionID string) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	if start, ok := gm.DecisionStart[decisionID]; ok {
		latency := time.Since(start)
		gm.ExplanationLatencies = append(gm.ExplanationLatencies, latency)
		delete(gm.DecisionStart, decisionID)

		log.Printf("[GOVERNANCE] MTTE Recorded for %s: %v", decisionID, latency)
	}
}

// AuditDeterminism compares a real sensor value against its replayed counterpart.
func (gm *GovernanceMonitor) AuditDeterminism(eventID string, reality, replay float64) {
	gm.mu.Lock()
	defer gm.mu.Unlock()

	drift := reality - replay
	if drift != 0 {
		leak := DeterminismLeak{
			EventID:    eventID,
			RealityVal: reality,
			ReplayVal:  replay,
			Drift:      drift,
			Timestamp:  time.Now().Unix(),
		}
		gm.Leaks = append(gm.Leaks, leak)

		// MASTER INVARIANT: Reality > Replay in case of leak.
		log.Printf("[SOVEREIGN ANOMALY] Determinism Leak detected in event %s! Drift: %.4f", eventID, drift)
	}
}

// CalculateMTTE returns the average time to explain across all recorded decisions.
func (gm *GovernanceMonitor) CalculateMTTE() time.Duration {
	gm.mu.RLock()
	defer gm.mu.RUnlock()

	if len(gm.ExplanationLatencies) == 0 {
		return 0
	}

	var total time.Duration
	for _, l := range gm.ExplanationLatencies {
		total += l
	}
	return total / time.Duration(len(gm.ExplanationLatencies))
}
