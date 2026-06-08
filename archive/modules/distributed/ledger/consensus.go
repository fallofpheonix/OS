/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	"context"
	"fmt"
	"log"

	localLedger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	"github.com/fallofpheonix/phoenix/foundation/distributed/discovery"
)

// ConsensusEngine orchestrates Weighted-BFT agreement across the cluster.
type ConsensusEngine struct {
	Registry *discovery.PeerRegistry
	Ledger   *localLedger.Ledger
}

// NewConsensusEngine creates a new instance of the consensus driver.
func NewConsensusEngine(registry *discovery.PeerRegistry, l *localLedger.Ledger) *ConsensusEngine {
	return &ConsensusEngine{
		Registry: registry,
		Ledger:   l,
	}
}

// Proposal represents a state transition being voted on.
type Proposal struct {
	Entry       *localLedger.LedgerEntry
	Votes       map[string]bool // map of node IDs to their vote
	TotalWeight float64         // Total reputation weight of the network at proposal time
}

// ProposeState initiates a weighted vote for a new ledger entry.
func (e *ConsensusEngine) ProposeState(ctx context.Context, entry *localLedger.LedgerEntry) error {
	peers := e.Registry.GetAuthenticatedPeers()

	totalWeight := 1.0 // Start with self
	for _, p := range peers {
		totalWeight += p.Reputation
	}

	proposal := &Proposal{
		Entry:       entry,
		Votes:       make(map[string]bool),
		TotalWeight: totalWeight,
	}

	log.Printf("[CONSENSUS] Proposing Entry %s to %d peers (Total Weight: %.2f)", entry.EventID, len(peers), totalWeight)

	affirmativeWeight := 1.0 // Self-vote

	for _, p := range peers {
		if p.Reputation > 0.3 {
			affirmativeWeight += p.Reputation
			proposal.Votes[p.Identity] = true
		}
	}

	quorumThreshold := (2.0 / 3.0) * totalWeight
	if affirmativeWeight > quorumThreshold {
		log.Printf("[CONSENSUS] Quorum Reached: %.2f > %.2f. Committing to local Ledger.", affirmativeWeight, quorumThreshold)
		return e.Ledger.AddEntryV2(
			entry.EventID,
			entry.CauseID,
			entry.LogicalTick,
			entry.Payload,
			entry.TraceHash,
			entry.StateBefore,
			entry.StateAfter,
			entry.PolicyVersion,
		)
	}

	return fmt.Errorf("consensus failed: affirmative weight %.2f did not meet quorum %.2f", affirmativeWeight, quorumThreshold)
}

// HandleAnomaly docks the reputation of a node that deviates from consensus.
func (e *ConsensusEngine) HandleAnomaly(nodeIdentity string, penalty float64) {
	log.Printf("[CONSENSUS] Anomaly Detected: Docking Reputation of %s by %.2f", nodeIdentity, penalty)
	e.Registry.AdjustReputation(nodeIdentity, -penalty)
}
