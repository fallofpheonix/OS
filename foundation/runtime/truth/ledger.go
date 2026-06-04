/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/* =========================================================================
 * WORKFLOW POSITION: CYCLE 11d — TRUTH LEDGER (Layer 5)
 *
 * The TruthLedger maintains a sequence-to-hash mapping that blocks
 * tampering. Once a hash is recorded for a sequence number, it cannot
 * be changed without detection.
 *
 * WORKFLOW:
 *   AddEntry(EvidenceWrapper{Sequence, Hash})
 *     → If sequence exists with same hash: accept (idempotent)
 *     → If sequence exists with different hash: TAMPER DETECTED
 *     → If sequence is new: record hash
 *   → Snapshot() → return copy of all entries
 *
 * SECURITY: The ledger is append-only by design.
 * Attempting to overwrite an existing hash returns an error.
 * This prevents an attacker from modifying historical evidence.
 *
 * THREAD SAFETY: Uses sync.Mutex for concurrent access.
 * ========================================================================= */
package truth

import (
	"fmt"
	"sync"
)

// EvidenceWrapper contains a simple event sequence and hash.
type EvidenceWrapper struct {
	Sequence int
	Hash     string
}

// TruthLedger maintains sequence to hash mapping and blocks tampering.
type TruthLedger struct {
	mu      sync.Mutex
	entries map[int]string
}

// NewTruthLedger initializes a TruthLedger.
func NewTruthLedger() *TruthLedger {
	return &TruthLedger{
		entries: make(map[int]string),
	}
}

// AddEntry stores a sequence and hash, returning an error on duplication with a different hash.
func (l *TruthLedger) AddEntry(entry EvidenceWrapper) error {
	l.mu.Lock()
	defer l.mu.Unlock()

	if existingHash, ok := l.entries[entry.Sequence]; ok {
		if existingHash != entry.Hash {
			return fmt.Errorf("hash tamper detected: sequence %d has hash %s, tried to overwrite with %s", entry.Sequence, existingHash, entry.Hash)
		}
		return nil
	}

	l.entries[entry.Sequence] = entry.Hash
	return nil
}

// Snapshot returns a point-in-time copy of the truth ledger entries.
func (l *TruthLedger) Snapshot() map[int]string {
	l.mu.Lock()
	defer l.mu.Unlock()
	snap := make(map[int]string)
	for k, v := range l.entries {
		snap[k] = v
	}
	return snap
}
