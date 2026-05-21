package physics

import (
	"math"
	"sync"
	"time"

	"phoenix/agents/internal/types"
	"phoenix/monitor" // imports phoenix/monitor (entropy_engine)
)

type PhysicsAgent interface {
	CalculateSDI(states []int8) float64
	GetSecurityState(graph *types.IncidentGraph, now time.Time) (types.SecurityState, error)
}

type Agent struct {
	mu           sync.RWMutex
	lastEntropy  float64
	lastSDI      float64
	threatTemp   float64
	tempFilter   *KalmanFilter
}

func NewPhysicsAgent() *Agent {
	return &Agent{
		threatTemp: 0.1, // baseline normal temp
		tempFilter: NewKalmanFilter(0.1, 1.0, 1.0, 0.1),
	}
}

// CalculateSDI computes Security Disorder Index using Shannon-like calculation
func (a *Agent) CalculateSDI(states []int8) float64 {
	if len(states) == 0 {
		return 0
	}
	counts := make(map[int8]int)
	for _, s := range states {
		counts[s]++
	}
	var sdi float64
	n := float64(len(states))
	for _, count := range counts {
		p := float64(count) / n
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}
	return sdi
}

func (a *Agent) GetSecurityState(graph *types.IncidentGraph, now time.Time) (types.SecurityState, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Gather node threat scores as states
	var states []int8
	var maxThreat float64
	var avgImportance float64

	if len(graph.Nodes) > 0 {
		var totalImportance float64
		for _, node := range graph.Nodes {
			// Convert threat scores 0..10 to states -1 (benign), 0 (unknown), 1 (malicious)
			if node.ThreatScore >= 7.0 {
				states = append(states, 1)
			} else if node.ThreatScore >= 4.0 {
				states = append(states, 0)
			} else {
				states = append(states, -1)
			}

			if node.ThreatScore > maxThreat {
				maxThreat = node.ThreatScore
			}
			totalImportance += node.Importance
		}
		avgImportance = totalImportance / float64(len(graph.Nodes))
	} else {
		states = append(states, -1)
	}

	// Calculate SDI
	sdi := a.CalculateSDI(states)
	a.lastSDI = sdi

	// Threat Temperature theta_T increases with high threat scores and high node importance
	targetTemp := 0.1 + (maxThreat * 0.5) + (avgImportance * 2.0)
	if targetTemp > 10.0 {
		targetTemp = 10.0
	}
	
	// Apply Kalman filter to smooth and predict threat temperature
	a.threatTemp = a.tempFilter.Update(targetTemp)

	// Retrieve Shannon entropy from monitor package
	testData := make([]byte, 256)
	if maxThreat >= 7.0 {
		// High disorder data
		for i := range testData {
			testData[i] = byte(i)
		}
	} else {
		// Low disorder uniform data
		for i := range testData {
			testData[i] = 'A'
		}
	}
	entropy := entropy_engine.CalculateEntropy(testData)
	a.lastEntropy = entropy

	isAnomaly := sdi > 1.0 || entropy > 7.5 || a.threatTemp > 4.0

	return types.SecurityState{
		Timestamp:         now,
		Entropy:           entropy,
		KLDivergence:      0.0, // baseline
		ThreatTemperature: a.threatTemp,
		SDI:               sdi,
		IsAnomaly:         isAnomaly,
	}, nil
}
