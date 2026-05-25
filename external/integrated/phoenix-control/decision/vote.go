package decision

import "github.com/fallofpheonix/phoenix-control/warden"

// Voter implements a consensus-based decision mechanism for multi-agent or multi-node environments.
// This supports the Phoenix Nexus (L7) swarm coordination.
type Voter struct {
	RequiredVotes int
}

func NewVoter(required int) *Voter {
	return &Voter{RequiredVotes: required}
}

// Vote aggregates multiple integrated decisions to reach a final consensus using weighted confidence.
func (v *Voter) Vote(decisions []IntegratedDecision) (IntegratedDecision, bool) {
	if len(decisions) == 0 {
		return IntegratedDecision{}, false
	}

	// 1. Tally weighted votes for each Action/State pair
	type VoteKey struct {
		Action warden.ActuationClass
		State  warden.SystemState
	}
	tally := make(map[VoteKey]float64)
	var totalWeight float64

	for _, d := range decisions {
		key := VoteKey{d.AuthorizedAction, d.AuthorizedState}
		tally[key] += d.Confidence
		totalWeight += d.Confidence
	}

	// 2. Find the winner (majority of weighted confidence)
	var winner VoteKey
	var maxWeight float64
	for key, weight := range tally {
		if weight > maxWeight {
			maxWeight = weight
			winner = key
		}
	}

	// 3. Threshold Check: Winner must represent a clear consensus (e.g., > 50% of total confidence)
	if maxWeight < (totalWeight * 0.5) {
		return IntegratedDecision{}, false
	}

	// 4. Return a merged decision representing the consensus
	// We use the first decision matching the winning key as the template.
	for _, d := range decisions {
		if d.AuthorizedAction == winner.Action && d.AuthorizedState == winner.State {
			consensusDec := d
			consensusDec.Confidence = maxWeight / totalWeight // Normalized consensus confidence
			return consensusDec, true
		}
	}

	return IntegratedDecision{}, false
}
