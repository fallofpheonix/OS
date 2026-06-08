package consensus

import (
	"testing"

	"github.com/fallofpheonix/phoenix/internal/contracts"
)

func TestPropose_StoresEvents(t *testing.T) {
	s := NewInMemoryStagingArea()
	events := []contracts.Event{
		{Type: contracts.EventSpawn},
		{Type: contracts.EventMove},
		{Type: contracts.EventVerify},
	}

	err := s.Propose(1, 1, events)
	if err != nil {
		t.Fatalf("Propose failed: %v", err)
	}

	got, ok := s.GetTentative(1, 1)
	if !ok || len(got) != 3 {
		t.Errorf("Expected 3 events, got %d", len(got))
	}
}

func TestCommit_SealsThenRejectsNewPropose(t *testing.T) {
	s := NewInMemoryStagingArea()
	events := []contracts.Event{{Type: contracts.EventMove}}

	s.Propose(5, 1, events)
	_, err := s.Commit(5, 1)
	if err != nil {
		t.Fatalf("Commit failed: %v", err)
	}

	err = s.Propose(5, 2, events)
	if err == nil {
		t.Error("Expected error proposing to committed height, got nil")
	}
}

func TestRollback_UnknownRoundIsNoOp(t *testing.T) {
	s := NewInMemoryStagingArea()
	err := s.Rollback(5, 99)
	if err != nil {
		t.Errorf("Expected nil error for unknown rollback, got %v", err)
	}
}

func TestAbort_ClearsAllRoundsAndAllowsRepropose(t *testing.T) {
	s := NewInMemoryStagingArea()
	e := []contracts.Event{{Type: contracts.EventSpawn}}

	s.Propose(5, 1, e)
	s.Propose(5, 2, e)

	err := s.Abort(5)
	if err != nil {
		t.Fatalf("Abort failed: %v", err)
	}

	err = s.Propose(5, 1, e)
	if err != nil {
		t.Errorf("Propose after abort failed: %v", err)
	}
}
