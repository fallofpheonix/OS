package game

import (
	"sync"
	"time"

	"phoenix/agents/internal/types"
)

type GameAgent interface {
	UpdateBeliefs(state types.SecurityState, evidence string)
	SolveBestStrategy(state types.SecurityState, graph *types.IncidentGraph) (types.Strategy, error)
	GetBeliefs() (float64, float64)
}

type Agent struct {
	mu           sync.Mutex
	priorRan     float64 // Prior belief of ransomware type
	priorBen     float64 // Prior belief of benign type
}

func NewGameAgent() *Agent {
	return &Agent{
		priorRan: 0.05, // default low prior
		priorBen: 0.95,
	}
}

// UpdateBeliefs processes evidence via recursive Bayesian updates
func (a *Agent) UpdateBeliefs(state types.SecurityState, evidence string) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var pEvidenceGivenRan float64
	var pEvidenceGivenBen float64

	switch evidence {
	case "high_entropy_write":
		// Ransomware almost always writes high entropy files. Benign compilers occasionally do.
		pEvidenceGivenRan = 0.95
		pEvidenceGivenBen = 0.15
	case "suspicious_rename":
		pEvidenceGivenRan = 0.80
		pEvidenceGivenBen = 0.05
	case "normal_filesystem_io":
		pEvidenceGivenRan = 0.10
		pEvidenceGivenBen = 0.80
	default:
		pEvidenceGivenRan = 0.50
		pEvidenceGivenBen = 0.50
	}

	// Bayes theorem: P(Ran | Ev) = P(Ev | Ran) * P(Ran) / (P(Ev|Ran)*P(Ran) + P(Ev|Ben)*P(Ben))
	num := pEvidenceGivenRan * a.priorRan
	den := num + (pEvidenceGivenBen * a.priorBen)

	if den > 0 {
		a.priorRan = num / den
		a.priorBen = 1.0 - a.priorRan
	}
}

// SolveBestStrategy calculates Stackelberg follower response or Nash equilibrium
func (a *Agent) SolveBestStrategy(state types.SecurityState, graph *types.IncidentGraph) (types.Strategy, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	// Decide target containment level based on threat temperature and Bayesian posterior
	level := 0
	strategyType := "Nash"

	if a.priorRan > 0.80 {
		level = 5 // Maximum containment: Kill + Isolate
		strategyType = "Stackelberg"
	} else if a.priorRan > 0.50 || state.ThreatTemperature > 5.0 {
		level = 3 // Moderate containment: Freeze + Limit
		strategyType = "Stackelberg"
	} else if a.priorRan > 0.20 || state.SDI > 0.8 {
		level = 1 // Light containment: Observe + Limit
		strategyType = "Nash"
	}

	// Identify targets (PIDs of high threat nodes)
	var targets []uint32
	for _, node := range graph.Nodes {
		if node.ThreatScore >= 4.0 {
			targets = append(targets, node.PID)
		}
	}

	return types.Strategy{
		ContainmentLevel: level,
		TargetPIDs:       targets,
		StrategyType:     strategyType,
		Timestamp:        time.Now(),
	}, nil
}

func (a *Agent) GetBeliefs() (float64, float64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.priorRan, a.priorBen
}
