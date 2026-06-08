/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ledger

import (
	"context"
	"time"
)

// Entry represents a single record in the distributed ledger.
type Entry struct {
	Index     uint64    `json:"index"`
	Term      uint64    `json:"term"`
	Timestamp time.Time `json:"timestamp"`
	Payload   []byte    `json:"payload"`
	EventID   string    `json:"event_id"`
	Signature []byte    `json:"signature"`
}

// ConsensusLedger defines the interface for a replicated, fault-tolerant evidence ledger.
type ConsensusLedger interface {
	// Propose submits a new entry to the cluster for consensus.
	// Returns the expected index or an error if the proposal fails.
	Propose(ctx context.Context, eventID string, payload []byte) (uint64, error)

	// Commit confirms that an entry has reached quorum and is permanently appended.
	Commit(ctx context.Context, index uint64) error

	// Get retrieves an entry by its index.
	Get(ctx context.Context, index uint64) (*Entry, error)

	// Subscribe returns a channel that receives new entries as they are committed.
	Subscribe(ctx context.Context) (<-chan Entry, error)

	// Verify checks the integrity of the local ledger shard.
	Verify(ctx context.Context) error
}
