package governance

import "testing"

func TestGovernance(t *testing.T) {
	policy := Policy{MinReputation: 0.7, QuorumSize: 3}
	gov := NewSwarmGovernor(policy)

	if gov.ValidateProposal(0.8) == false {
		t.Error("Proposal should be accepted for compliant reputation")
	}

	if gov.ValidateProposal(0.5) == true {
		t.Error("Proposal should be rejected for low reputation")
	}
}
