package consensus

import (
	"github.com/fallofpheonix/phoenix/internal/contracts"
)

// PURPOSE: Defines the transactional buffer between network proposals and
//          the canonical ledger. Prevents unverified consensus rounds from
//          polluting the permanent forensic record.
// CONTRACT: Implementations must be safe for concurrent use.
// FAILURE: Commit on an uncommitted height/round returns error.
//          Rollback on an unknown height/round is a no-op, not an error.
// CONNECTS: internal/consensus/certificate.go (QC verification triggers Commit)
//           internal/ledger/src/ledger.go (Commit output feeds AddEntryV2)

type StagingArea interface {
	Propose(height uint64, round uint32, events []contracts.Event) error
	Commit(height uint64, round uint32) ([]contracts.Event, error)
	Rollback(height uint64, round uint32) error
	Abort(height uint64) error
	GetTentative(height uint64, round uint32) ([]contracts.Event, bool)
}
