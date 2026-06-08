/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: genesis.go
 * PATH: Phoenix.Nucleus/ledger/genesis.go
 *
 * PURPOSE:
 * Implements the Genesis Block, the root of the PhoenixOS universe.
 * Defines the initial Identity, Authority, and Mission of the system.
 *
 * SUBSYSTEM:
 * Nucleus / Ledger Cycle / Genesis
 *
 * DEPENDENCIES:
 * crypto/sha256, encoding/json, time
 *
 * DEPENDENTS:
 * Phoenix.Nucleus/recovery, Phoenix.Nucleus/ledger
 *
 * SECURITY:
 * This is the ultimate Trust Anchor. The hash of this block is the root state.
 */

package ledger

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// BEGINNER EXPLANATION:
// This file is the "Big Bang" of PhoenixOS. It defines the very first
// information the system ever knows.

// INTERMEDIATE EXPLANATION:
// The GenesisBlock defines the version, root authority ID, and the initial
// mission statement. Its hash becomes the `ParentHash` for the second event.

// EXPERT EXPLANATION:
// Axiomatic root of the system. All subsequent state transitions must
// causally link back to this block. It satisfies Q580-Q585 by defining
// the "I AM" event and the root entropy.

/**
 * GenesisBlock
 *
 * Represents the root of the PhoenixOS Ledger.
 */
type GenesisBlock struct {
	Version           string    `json:"version"`
	Timestamp         time.Time `json:"timestamp"`
	Identity          string    `json:"identity"` // "I AM"
	RootAuthority     string    `json:"root_authority"`
	Mission           string    `json:"mission"`
	InitialEntropy    string    `json:"initial_entropy"`
	FixedPointDivisor uint64    `json:"fixed_point_divisor"` // Architecture-independent precision divisor
	Hash              string    `json:"hash"`
}

/**
 * CalculateHash
 *
 * Generates the SHA-256 hash of the Genesis Block.
 */
func (g *GenesisBlock) CalculateHash() string {
	// Create a copy without the hash to ensure idempotency
	temp := *g
	temp.Hash = ""
	data, _ := json.Marshal(temp)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

/**
 * NewGenesisBlock
 *
 * Creates a standard PhoenixOS Genesis Block.
 *
 * Return: A pointer to the initialized and hashed Genesis block.
 */
func NewGenesisBlock() *GenesisBlock {
	g := &GenesisBlock{
		Version:           "1.0.0",
		Timestamp:         time.Unix(1772323200, 0), // 2026-03-15
		Identity:          "PHOENIX-GENESIS-0",
		RootAuthority:     "AUTHORITY-ROOT-SUPER",
		Mission:           "CONSERVE-ORDER-QUENCH-DISORDER",
		InitialEntropy:    "0.0",
		FixedPointDivisor: 1_000_000, // 10^6 millionth-level precision
	}
	g.Hash = g.CalculateHash()
	return g
}
