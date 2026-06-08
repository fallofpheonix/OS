package consensus

import (
	"fmt"
	"sync"

	"github.com/fallofpheonix/phoenix/internal/contracts"
)

type InMemoryStagingArea struct {
	mu        sync.RWMutex
	rounds    map[uint64]map[uint32][]contracts.Event
	committed map[uint64]bool
}

func NewInMemoryStagingArea() *InMemoryStagingArea {
	return &InMemoryStagingArea{
		rounds:    make(map[uint64]map[uint32][]contracts.Event),
		committed: make(map[uint64]bool),
	}
}

// Propose
// PURPOSE: Adds a set of events to a tentative block at the given height/round.
// CONTRACT: Returns error if the height has already been committed (sealed).
// FAILURE: Returns error if height is committed.
func (s *InMemoryStagingArea) Propose(height uint64, round uint32, events []contracts.Event) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.committed[height] {
		return fmt.Errorf("height %d is sealed", height)
	}

	if _, ok := s.rounds[height]; !ok {
		s.rounds[height] = make(map[uint32][]contracts.Event)
	}
	s.rounds[height][round] = events
	return nil
}

// Commit
// PURPOSE: Finalizes a tentative block and seals the height for further proposals.
// CONTRACT: Returns the events if found, else error.
// FAILURE: Returns error if height/round coordinate is unknown.
func (s *InMemoryStagingArea) Commit(height uint64, round uint32) ([]contracts.Event, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	r, ok := s.rounds[height]
	if !ok {
		return nil, fmt.Errorf("unknown height %d", height)
	}
	events, ok := r[round]
	if !ok {
		return nil, fmt.Errorf("unknown round %d at height %d", round, height)
	}

	s.committed[height] = true
	return events, nil
}

// Rollback
// PURPOSE: Discards a specific round's proposal at a given height.
// CONTRACT: No-op if round/height is unknown.
// FAILURE: Always returns nil (no-op safety).
func (s *InMemoryStagingArea) Rollback(height uint64, round uint32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	if r, ok := s.rounds[height]; ok {
		delete(r, round)
	}
	return nil
}

// Abort
// PURPOSE: Clears all rounds at a height and removes the commitment seal.
// CONTRACT: Deletes all round data and allow height to be re-proposed.
// FAILURE: Always returns nil.
func (s *InMemoryStagingArea) Abort(height uint64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.rounds, height)
	delete(s.committed, height)
	return nil
}

// GetTentative
// PURPOSE: Retrieves proposed events without triggering a commitment seal.
// CONTRACT: Returns (events, true) if found, (nil, false) otherwise.
// FAILURE: Thread-safe read-only access.
func (s *InMemoryStagingArea) GetTentative(height uint64, round uint32) ([]contracts.Event, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	if r, ok := s.rounds[height]; ok {
		if events, ok := r[round]; ok {
			return events, true
		}
	}
	return nil, false
}
