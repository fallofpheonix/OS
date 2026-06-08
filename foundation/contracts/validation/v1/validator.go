package v1

// SemanticValidator defines the contract for logical verification of event data.
// PURPOSE: Prevents malformed or logically impossible events from entering the ledger.
// CONTRACT: Implementation MUST be deterministic and side-effect free.
// FAILURE: Returns error if the event violates system-level invariants.
// CONNECTS: foundation/runtime/bus/applier.go (consumer)
//
//	foundation/ledger/src/physics_validator.go (implementation)
type SemanticValidator interface {
	// Validate checks the semantic integrity of an event.
	Validate(eventID, causeID string, tick uint64, payload []byte) error
}
