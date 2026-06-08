package ledger

import (
	"bytes"
	"fmt"
	"sort"
)

// PURPOSE: Reconstructs the in-memory Ledger state from a durable file.
// CONTRACT: Given a valid ledger file, produces a Ledger whose Head Hash
//           matches the value recorded at checkpoint.
// FAILURE: Returns error if any entry fails hash verification or the
//           file is structurally invalid.
// CONNECTS: foundation/ledger/src/persist.go (source)
//           foundation/ledger/src/ledger.go (target)

type Replayer struct {
	persistor *Persistor
}

func NewReplayer(path string) (*Replayer, error) {
	p, err := NewPersistor(path)
	if err != nil {
		return nil, err
	}
	return &Replayer{persistor: p}, nil
}

func (r *Replayer) Replay() (*Ledger, error) {
	entries, header, err := r.persistor.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read ledger file: %w", err)
	}

	if header.Version != "1.0" {
		return nil, fmt.Errorf("incompatible ledger schema version: %s", header.Version)
	}

	// [RECTIFIED]: Step A - Canonical Ordering.
	// Sort entries by (Tick, Sequence) to ensure deterministic application
	// regardless of the order they appear in the file.
	sort.Slice(entries, func(i, j int) bool {
		if entries[i].LogicalTick != entries[j].LogicalTick {
			return entries[i].LogicalTick < entries[j].LogicalTick
		}
		return entries[i].Sequence < entries[j].Sequence
	})

	// WHY: allocator is nil during replay because we assume the written entries already
	// passed resource bounds checks at the time they were accepted. Replay must succeed
	// to recover state, bounding happens on ingest, not recovery.
	l := NewLedger(nil)

	// WHY: WithPersistor is not called on replay ledger. We are reading from disk,
	// not writing to it. The in-memory ledger built here should not mirror its replayed
	// entries back to the persistor.

	for i, entry := range entries {
		expectedHash := entry.Hash

		// Recompute hash to ensure integrity
		actualHash := l.computeHash(entry)

		// WHY: hash mismatch causes immediate halt. A hash mismatch during replay means
		// either the disk data was tampered with, corrupted, or the hashing algorithm changed.
		// Neither is recoverable by continuing; the state is tainted.
		if !bytes.Equal(expectedHash, actualHash) {
			return nil, fmt.Errorf("hash verification failed at index %d: expected %x, got %x", i, expectedHash, actualHash)
		}

		l.Entries[fmt.Sprintf("%x", actualHash)] = entry
		l.Heads = [][]byte{actualHash}
		l.Counter = entry.LogicalTick + 1
		l.Sequence = entry.Sequence // Maintain sequence continuity
	}

	return l, nil
}
