// Package event defines the canonical data structures for the PhoenixOS event substrate.
// Core Domain Logic: Formally defines the Event, ArtifactManifest, and Checkpoint structures
// which constitute the fundamental unit of reality and historical record within the system.
package event

import (
	"encoding/json"
	"time"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
)

// Event represents the atomic unit of reality in PhoenixOS.
// Internal State: Encapsulates identity, authority, causality (ParentID), and state transition (Payload).
// API Scope: Public; primary communication and persistence primitive.
// Concurrency: Read-only instances are thread-safe.
type Event struct {
	EventID     string          `json:"event_id"`
	ParentID    string          `json:"parent_id"`
	AuthorityID string          `json:"authority_id"`
	IdentityID  string          `json:"identity_id"`
	LogicalTime uint64          `json:"logical_time"`
	Evidence    []string        `json:"evidence"` // Artifact Hashes
	Signature   string          `json:"signature"`
	Payload     json.RawMessage `json:"payload"`
}

var _ eventsv1.EventEnvelope = Event{}

func (e Event) GetEventID() string            { return e.EventID }
func (e Event) GetEventVersion() string       { return "1.0.0" }
func (e Event) GetEventType() string          { return "GenericEvent" }
func (e Event) GetTimestamp() time.Time       { return time.Now() }
func (e Event) GetMonotonicTime() uint64      { return e.LogicalTime }
func (e Event) GetSourceRepo() string         { return e.AuthorityID }
func (e Event) GetSourceComponent() string    { return "Core" }
func (e Event) GetParentEvent() string        { return e.ParentID }
func (e Event) GetCorrelationID() string      { return e.IdentityID }
func (e Event) GetCausalChain() []string      {
	if e.ParentID != "" {
		return []string{e.ParentID}
	}
	return nil
}
func (e Event) GetReplaySequence() uint64      { return e.LogicalTime }
func (e Event) GetEvidenceHash() string        {
	if len(e.Evidence) > 0 {
		return e.Evidence[0]
	}
	return ""
}
func (e Event) GetTrustScore() float64        { return 1.0 }
func (e Event) GetSignature() string          { return e.Signature }
func (e Event) GetPayloadHash() string        { return "" }
func (e Event) GetSchemaVersion() string      { return "1.0.0" }
func (e Event) GetCreatedAt() time.Time       { return time.Now() }
func (e Event) GetUpdatedAt() time.Time       { return time.Now() }
func (e Event) GetValidationHash() string     { return "" }

// ArtifactManifest represents an immutable, signed piece of data or code.
// Internal State: Metadata describing a binary artifact, its dependencies, and retention policy.
// API Scope: Public; used for artifact discovery and verification.
// Concurrency: Thread-safe (immutable).
type ArtifactManifest struct {
	Hash         string   `json:"hash"`
	Version      string   `json:"version"`
	Dependencies []string `json:"dependencies"` // Hashes of other artifacts
	Signer       string   `json:"signer"`
	Retention    string   `json:"retention"` // e.g., "PERMANENT", "EPHEMERAL"
}

// Checkpoint represents a deterministic snapshot of the authoritative state.
// Internal State: Aggregated state hash and offset for ledger replay.
// API Scope: Public; used for fast-sync and recovery initialization.
// Concurrency: Thread-safe (immutable).
type Checkpoint struct {
	StateHash          string   `json:"state_hash"`
	ReplayOffset       uint64   `json:"replay_offset"`
	ArtifactReferences []string `json:"artifact_references"`
	Timestamp          time.Time `json:"timestamp"`
}
