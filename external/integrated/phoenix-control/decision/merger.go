package decision

import (
	"time"

	"github.com/fallofpheonix/phoenix-control/arbiter"
	"github.com/fallofpheonix/phoenix-control/warden"
)

// Merger combines multi-source inputs into a single authorized decision.
type Merger struct {
	MinConfidence float64
}

func NewMerger(minConf float64) *Merger {
	return &Merger{MinConfidence: minConf}
}

// Merge integrates Arbiter policy and Replay evidence.
func (m *Merger) Merge(policy arbiter.Decision, evidenceScore float64) IntegratedDecision {
	// 1. Normalize Policy Score based on Attack Cost vs. Total Cost
	// If TotalCost == AttackCost, policy score is high (0.8).
	// If TotalCost is much lower than AttackCost, policy score is moderate.
	policyScore := 0.5
	if policy.CostFactors.AttackCost > 0 {
		// Profitability-based normalization
		policyScore = 1.0 - (policy.TotalCost / policy.CostFactors.AttackCost)
		if policyScore < 0 {
			policyScore = 0
		}
		if policyScore > 1.0 {
			policyScore = 1.0
		}
	}

	// 2. Dynamic Weighting
	policyWeight := 0.6
	evidenceWeight := 0.4

	// If evidence score is extremely high (1.0), it can override policy weight
	if evidenceScore > 0.95 {
		policyWeight = 0.3
		evidenceWeight = 0.7
	}

	// 3. Confidence calculation: Multi-source weighted blend
	confidence := (policyWeight * policyScore) + (evidenceWeight * evidenceScore)

	authorizedAction := policy.Recommended
	authorizedState := warden.StateSafe

	// Gating: If confidence is too low, downgrade the action.
	if confidence < m.MinConfidence && authorizedAction > warden.ClassLog {
		authorizedAction = warden.ClassLog
		authorizedState = warden.StateWatch
	} else {
		// Map ActuationClass to SystemState
		switch authorizedAction {
		case warden.ClassObserve:
			authorizedState = warden.StateSafe
		case warden.ClassLog:
			authorizedState = warden.StateWatch
		case warden.ClassThrottle, warden.ClassLocalIsolate:
			authorizedState = warden.StateAlert
		case warden.ClassClusterIsolate, warden.ClassKernelEmergency:
			authorizedState = warden.StateContain
		}
	}

	return IntegratedDecision{
		PID:              policy.TargetPID,
		PolicyDecision:   policy,
		EvidenceScore:    evidenceScore,
		Confidence:       confidence,
		AuthorizedAction: authorizedAction,
		AuthorizedState:  authorizedState,
		Timestamp:        time.Unix(policy.Timestamp, 0),
	}
}
