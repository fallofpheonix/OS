package v1

// PURPOSE: Defines the write contract for all ledger implementations.
// CONTRACT: Any type implementing LedgerWriter may serve as the durable
//
//	forensic record. The Applier depends on this interface, never
//	on a concrete ledger type.
//
// CONNECTS: foundation/runtime/bus/applier.go (consumer)
//
//	foundation/ledger/src (current implementation)
//	foundation/ledger/core (future implementation)
type LedgerWriter interface {
	// AddEntry appends one event to the immutable ledger.
	// PURPOSE: Authoritative record of a single system transition.
	// CONTRACT: eventID must be unique. causeID may be empty for root events.
	//           The tick must be provided by the caller to ensure replay
	//           reconstructs the exact same sequence.
	// FAILURE: Returns error on disk full or resource exhaustion.
	//           Panics on hash chain violation or internal corruption.
	AddEntry(eventID, causeID string, tick uint64, payload []byte) error
}
