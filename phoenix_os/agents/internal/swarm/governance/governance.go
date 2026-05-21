package governance

import (
	"sync"
)

// Policy defines the rules for swarm node participation.
type Policy struct {
	MinReputation float64
	QuorumSize    int
}

// SwarmGovernor enforces governance policies on the node network.
type SwarmGovernor struct {
	mu     sync.RWMutex
	Policy Policy
}

// NewSwarmGovernor initializes the governor with a policy.
func NewSwarmGovernor(policy Policy) *SwarmGovernor {
	return &SwarmGovernor{
		Policy: policy,
	}
}

// ValidateProposal checks if a node is authorized to participate in consensus.
func (sg *SwarmGovernor) ValidateProposal(nodeReputation float64) bool {
	sg.mu.RLock()
	defer sg.mu.RUnlock()
	return nodeReputation >= sg.Policy.MinReputation
}
