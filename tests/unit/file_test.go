package unit

import (
	"testing"
)

func TestFileAudit(t *testing.T) {
	audit := NewFileAudit()
	
	action := FileAction{
		Path:       "/etc/config",
		Action:     ActionThrottle,
		Reason:     "high I/O",
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
