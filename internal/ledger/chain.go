/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: chain.go
 * PATH: Phoenix.Nucleus/ledger/chain.go
 *
 * PURPOSE:
 * Implements the linear Event Chain for the PhoenixOS Ledger.
 * Provides the core logic for appending and verifying hashed events.
 *
 * SUBSYSTEM:
 * Nucleus / Ledger Cycle
 *
 * DEPENDENCIES:
 * errors, fmt, sync
 *
 * DEPENDENTS:
 * Phoenix.Nucleus/recovery, Phoenix.Terminus/explanations
 *
 * SECURITY:
 * Enforces cryptographic link verification (ParentHash) and sequence continuity.
 * Detection of content corruption or sequence breaks results in an immediate Error.
 *
 * PERFORMANCE:
 * O(1) for appends. O(N) for full chain verification.
 */

package ledger

import (
	"errors"
	"fmt"
	"sync"
)

// BEGINNER EXPLANATION:
// This file is like the "Binding" of the system's journal. It makes sure every page
// (Event) is in the right order and that no pages have been ripped out or changed.

// INTERMEDIATE EXPLANATION:
// The Chain manager maintains an in-memory sequence of events and an index for
// fast lookup. It enforces two critical invariants: Sequence continuity and
// Parent-Hash matching.

// EXPERT EXPLANATION:
// Implements the Linear Ledger Invariant. It serves as the authoritative
// source of state transitions. Verification is performed recursively (or iteratively
// over the slice) to prove the integrity of the Merkle-linear-chain.

/**
 * Chain
 *
 * Represents an in-memory, verifiable sequence of Ledger events.
 *
 * Responsibilities:
 * - Event appending with integrity checks.
 * - Sequence verification.
 * - Full chain audit.
 *
 * Thread Safety:
 * Thread-safe via sync.RWMutex.
 */
type Chain struct {
	mu                sync.RWMutex
	events            []*Event
	index             map[string]*Event // Hash to Event
	lastSequenceIndex map[string]uint64 // AuthorityRef to Last Sequence
}

/**
 * NewChain
 *
 * Initializes a new Ledger chain.
 */
func NewChain() *Chain {
	return &Chain{
		events:            make([]*Event, 0),
		index:             make(map[string]*Event),
		lastSequenceIndex: make(map[string]uint64),
	}
}

/**
 * Append
 *
 * Adds a new event to the chain after verifying its integrity and sequence.
 *
 * Input:
 * - e: The Event to append.
 *
 * Security:
 * - Detects invalid sequence numbers.
 * - Verifies hash chain continuity.
 * - Detects internal content corruption.
 *
 * Complexity: O(1)
 */
func (c *Chain) Append(e *Event) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	// 1. Verify Sequence
	expectedSeq := uint64(len(c.events))
	if e.Sequence != expectedSeq {
		return fmt.Errorf("invalid sequence: expected %d, got %d", expectedSeq, e.Sequence)
	}

	// 2. Verify Parent Hash
	if expectedSeq == 0 {
		if e.Type != EventGenesis {
			return errors.New("first event must be GENESIS")
		}
		if e.ParentHash != "" {
			return errors.New("GENESIS event cannot have a parent hash")
		}
	} else {
		parent := c.events[len(c.events)-1]
		if e.ParentHash != parent.Hash {
			return fmt.Errorf("parent hash mismatch: expected %s, got %s", parent.Hash, e.ParentHash)
		}
	}

	// 3. Verify Self Hash
	if e.Hash != e.CalculateHash() {
		return errors.New("event hash corruption detected")
	}

	// 4. Append
	c.events = append(c.events, e)
	c.index[e.Hash] = e
	c.lastSequenceIndex[e.AuthorityRef] = e.Sequence
	return nil
}

/**
 * GetHead
 *
 * Returns the latest event in the chain.
 *
 * Complexity: O(1)
 */
func (c *Chain) GetHead() *Event {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if len(c.events) == 0 {
		return nil
	}
	return c.events[len(c.events)-1]
}

/**
 * GetEventByHash
 *
 * Retrieves an event by its SHA-256 hash.
 *
 * Input:
 * - hash: Hex-encoded string.
 *
 * Complexity: O(1)
 */
func (c *Chain) GetEventByHash(hash string) (*Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	e, ok := c.index[hash]
	return e, ok
}

/**
 * GetBySequence
 *
 * Retrieves an event by its sequence number.
 *
 * Input:
 * - seq: Sequence number.
 *
 * Complexity: O(1)
 */
func (c *Chain) GetBySequence(seq uint64) (*Event, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	if seq >= uint64(len(c.events)) {
		return nil, false
	}
	return c.events[seq], true
}

/**
 * RecoverSequence
 *
 * Returns the last known sequence number for a given authority.
 * This allows for efficient replay prevention without scanning the whole ledger.
 */
func (c *Chain) RecoverSequence(authRef string) uint64 {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.lastSequenceIndex[authRef]
}

/**
 * VerifyChain
 *
 * Performs a full integrity check of the entire chain history.
 *
 * Security:
 * - Iteratively verifies every hash link in the chain.
 *
 * Complexity: O(N) where N is chain length.
 */
func (c *Chain) VerifyChain() error {
	c.mu.RLock()
	defer c.mu.RUnlock()

	var lastHash string
	for i, e := range c.events {
		if e.Sequence != uint64(i) {
			return fmt.Errorf("sequence break at index %d", i)
		}
		if i > 0 && e.ParentHash != lastHash {
			return fmt.Errorf("hash chain broken at index %d", i)
		}
		if e.Hash != e.CalculateHash() {
			return fmt.Errorf("content corruption at index %d", i)
		}
		lastHash = e.Hash
	}
	return nil
}
