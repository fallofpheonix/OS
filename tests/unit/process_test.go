package unit

import "testing"

func TestProcessAudit(t *testing.T) {
	audit := NewProcessAudit()
	
	action := ProcessAction{
		PID:        123,
		Action:     ActionThrottle,
		Reason:     "high cpu",
		EvidenceID: "ev-1",
		DecisionID: "dec-1",
	}
	
	audit.LogAction(action)
	
	if len(audit.History) != 1 {
		t.Errorf("expected 1 history entry, got %d", len(audit.History))
	}
	
	if audit.History[0].PID != 123 {
		t.Errorf("expected pid 123, got %d", audit.History[0].PID)
	}
}
