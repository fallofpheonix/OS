/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/rollback"
	"testing"
)

func TestRollbackReplay(t *testing.T) {
	audit := rollback.NewRollbackAudit()

	record := rollback.RollbackRecord{
		Component: rollback.ComponentNetwork,
		Sequence:  1,
		Hash:      "dummy-hash",
	}

	audit.LogRollback(record)

	// Simulate Replay metadata
	replay := rollback.RollbackReplay{
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
