package unit

import "testing"

func TestNetworkAudit(t *testing.T) {
	audit := NewNetworkAudit()
	
	action := NetworkAction{
		Src:        "10.0.0.1",
		Dst:        "192.168.1.1",
		Port:       80,
		Action:     ActionThrottle,
		Reason:     "high bandwidth",
		EvidenceID: "ev-1",
		DecisionID: "dec-1",
	}
	
	audit.LogAction(action)
	
	if len(audit.History) != 1 {
		t.Errorf("expected 1 history entry, got %d", len(audit.History))
	}
	
	if audit.History[0].Sequence != 1 {
		t.Errorf("expected sequence 1, got %d", audit.History[0].Sequence)
	}
}
