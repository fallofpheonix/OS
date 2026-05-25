package arbiter

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-control/warden"
)

// EvaluatePolicy evaluates the best containment action by comparing Attack Cost vs. Containment Cost.
// It follows the Research Lead mandate: policy reasoning only, no autonomous action.
// This function determines the highest severity containment that is still "profitable" (CC <= AC).
func EvaluatePolicy(ctx SystemContext) Decision {
	attackCost := CalculateAttackCost(ctx)

	// Default recommendation is Observe (zero cost)
	bestAction := warden.ClassObserve
	bestScore := DecisionScore{
		ReplayConfidence: 1.0,
		EvidenceWeight:   1.0,
		ThreatScore:      ctx.ThreatScore.Severity * 1000.0,
		PolicyFactor:     1.0,
	}
	bestModel := CostModel{AttackCost: attackCost}
	bestTotalCost := attackCost // Initial cost is just the threat persisting

	// Replay Confidence calculation (Track C2)
	replayConf := 1.0
	if len(ctx.ReplayHistory) == 0 {
		replayConf = 0.5 // 50% penalty for acting without evidence
	}

	// Candidate actions in increasing order of severity/cost.
	candidates := []warden.ActuationClass{
		warden.ClassLog,
		warden.ClassThrottle,
		warden.ClassLocalIsolate,
		warden.ClassClusterIsolate,
		warden.ClassKernelEmergency,
	}

	for _, class := range candidates {
		cc := GetContainmentCost(class)

		// Recovery Cost: Estimated cost to return to StateSafe (25% of CC)
		rc := cc * 0.25

		// Evidence Weighting (Track C2)
		evidenceWeight := 1.0 // TODO: Pull from truth layer/source reliability (Stage B)

		score := DecisionScore{
			ReplayConfidence: replayConf,
			EvidenceWeight:   evidenceWeight,
			ThreatScore:      ctx.ThreatScore.Severity * 1000.0,
			PolicyFactor:     1.0,
		}

		// Effectiveness depends on class severity (Non-linear model)
		effectiveness := 0.0
		switch class {
		case warden.ClassLog:
			effectiveness = 0.01
		case warden.ClassThrottle:
			effectiveness = 0.4
		case warden.ClassLocalIsolate:
			effectiveness = 0.7
		case warden.ClassClusterIsolate:
			effectiveness = 0.99
		case warden.ClassKernelEmergency:
			effectiveness = 1.0
		}

		// Residual Risk: Remaining attack cost after containment
		rr := attackCost * (1.0 - (effectiveness * score.Calculate() / 1000.0))

		// Replay Confidence Penalty (Track C2): Increases cost if evidence is missing
		confPenalty := 0.0
		if replayConf < 1.0 {
			confPenalty = cc * (1.0 - replayConf)
		}

		// Track C4: Counterfactual Analysis
		// TotalCost = Containment + Recovery + ResidualRisk + ConfidencePenalty
		totalCost := cc + rc + rr + confPenalty

		if totalCost < bestTotalCost {
			bestAction = class
			bestTotalCost = totalCost
			bestScore = score
			bestModel = CostModel{
				AttackCost:              attackCost,
				ContainmentCost:         cc,
				RecoveryCost:            rc,
				ResidualRisk:            rr,
				ReplayConfidencePenalty: confPenalty,
			}
		}
	}

	reasoning := fmt.Sprintf("Arbiter optimized to %v: TotalCost(%.1f) < PersistentCost(%.1f) | Score: %.2f | Load: %.2f",
		bestAction, bestTotalCost, attackCost, bestScore.Calculate(), ctx.SystemLoad)

	if len(ctx.ReplayHistory) > 0 {
		reasoning += fmt.Sprintf(" | Evidence: %d frames", len(ctx.ReplayHistory))
	} else {
		reasoning += " | WARNING: NO REPLAY EVIDENCE (Confidence Penalty Applied)"
	}

	return Decision{
		TargetPID:   ctx.ThreatScore.PID,
		Score:       bestScore,
		CostFactors: bestModel,
		TotalCost:   bestTotalCost,
		Recommended: bestAction,
		Reasoning:   reasoning,
		Timestamp:   ctx.ThreatScore.WallTimeUnix,
	}
}
