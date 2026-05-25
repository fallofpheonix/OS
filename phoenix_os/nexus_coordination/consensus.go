package nexus

type Node struct {
	ID         string
	Reputation float64
}

// Propose evaluates a proposal based on reputation-weighted consensus
func Propose(proposalID string, nodes []Node) bool {
	var totalReputation float64
	var consensusScore float64

	for _, node := range nodes {
		totalReputation += node.Reputation
		if node.Reputation > 0.5 { // Simple PoA criteria
			consensusScore += node.Reputation
		}
	}

	return consensusScore / totalReputation > 0.6 // 60% threshold
}
