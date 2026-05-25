package arbiter

import (
	"github.com/fallofpheonix/phoenix-control/warden"
	"github.com/fallofpheonix/phoenix-logic/monitor"
)

// SystemContext represents the inputs for the Arbiter's policy reasoning.
// It aggregates telemetry, replay history, and system state.
type SystemContext struct {
	ThreatScore   monitor.DriftScore `json:"threat_score"`
	ReplayHistory []string           `json:"replay_history"` // Evidence from Track B (Replay)
	RiskLevel     float64            `json:"risk_level"`     // 0.0 - 1.0 (Overall system risk)
	SystemLoad    float64            `json:"system_load"`    // 0.0 - 1.0 (Current CPU/Mem pressure)
}

// CostModel represents the multi-dimensional cost factors for an actuation decision.
type CostModel struct {
	AttackCost              float64 `json:"attack_cost"`
	ContainmentCost         float64 `json:"containment_cost"`
	RecoveryCost            float64 `json:"recovery_cost"`
	ResidualRisk            float64 `json:"residual_risk"`
	ReplayConfidencePenalty float64 `json:"replay_confidence_penalty"`
}

// DecisionScore represents the final weighted confidence for a specific action.
type DecisionScore struct {
	ReplayConfidence float64 `json:"replay_confidence"` // [0.0 - 1.0] Based on hash-match history
	EvidenceWeight   float64 `json:"evidence_weight"`   // [0.0 - 1.0] Based on source reliability
	ThreatScore      float64 `json:"threat_score"`      // [0.0 - 1000.0] Raw drift score from monitor
	PolicyFactor     float64 `json:"policy_factor"`     // [0.5 - 2.0] Multiplier from Arbiter policy
}

// Calculate returns the final weighted score.
// Formula: Score = ReplayConfidence * EvidenceWeight * ThreatScore * PolicyFactor
func (s DecisionScore) Calculate() float64 {
	return s.ReplayConfidence * s.EvidenceWeight * s.ThreatScore * s.PolicyFactor
}

// Decision represents the final recommendation from the Arbiter.
type Decision struct {
	TargetPID   int                   `json:"target_pid"`
	Score       DecisionScore         `json:"score"`
	CostFactors CostModel             `json:"cost_factors"`
	TotalCost   float64               `json:"total_cost"`
	Recommended warden.ActuationClass `json:"recommended_action"`
	Reasoning   string                `json:"reasoning"`
	Timestamp   int64                 `json:"timestamp"`
}
