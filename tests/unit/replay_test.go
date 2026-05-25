package unit

import (
	"testing"
)

func TestRollbackReplay(t *testing.T) {
	audit := NewRollbackAudit()
	
	record := RollbackRecord{
		Component:  ComponentNetwork,
		Sequence:   1,
		Hash:       "dummy-hash",
	}
	
	audit.LogRollback(record)
	
	// Simulate Replay metadata
	replay := RollbackReplay{
		Record:       audit.History[0],
		ReplayCursor: audit.History[0].Sequence,
		Sequence:     audit.History[0].Sequence,
		Hash:         audit.History[0].Hash,
	}
	
	if replay.ReplayCursor != 1 {
		t.Errorf("expected cursor 1, got %d", replay.ReplayCursor)
	}
	
	if replay.Hash != audit.History[0].Hash {
		t.Error("replay hash mismatch")
	}
}
