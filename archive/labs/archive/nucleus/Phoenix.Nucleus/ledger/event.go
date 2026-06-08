/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: event.go
 * PATH: Phoenix.Nucleus/ledger/event.go
 *
 * PURPOSE:
 * Defines the immutable Event structure and payloads, the backbone of the PhoenixOS Ledger.
 * Every system state transition must be recorded here to ensure replayability and forensic integrity.
 *
 * SUBSYSTEM:
 * Nucleus / Ledger Cycle
 *
 * DEPENDENCIES:
 * crypto/sha256, encoding/json, time
 *
 * DEPENDENTS:
 * Phoenix.Nucleus/authority, Phoenix.Cognition/memory, Phoenix.Crucible/simulation
 *
 * SECURITY:
 * Events use SHA-256 hash chaining to detect history tampering.
 * All state transitions require a valid AuthorityRef to ensure traceability to Genesis.
 *
 * PERFORMANCE:
 * O(N) where N is payload size for hashing and serialization.
 */

package ledger

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// BEGINNER EXPLANATION:
// This file is like a "Journal" for PhoenixOS. Every time something important happens—
// like giving someone permission or observing a fact—we write it down as an "Event."
// Because each event is linked to the previous one with a secret code (a hash), 
// nobody can change the past without breaking the whole chain.

// INTERMEDIATE EXPLANATION:
// This module provides the Event DTO and EventType constants. It implements 
// SHA-256 hash-chaining to ensure ledger integrity. Other repositories (Cognition, 
// Crucible) consume these events to rebuild their internal states.

// EXPERT EXPLANATION:
// This is the implementation of the Linear Ledger Cycle. It enforces 
// bit-for-bit determinism in state representation. The AuthorityRef field 
// serves as the cryptographic link between the capability that authorized 
// the event and the event itself, satisfying the Authority Conservation engine.

// AuthorityAtom represents the smallest indivisible unit of authority (Satoshi equivalent).
type AuthorityAtom uint64

// EventType defines the category of a ledger event.
type EventType string

const (
	EventGenesis             EventType = "GENESIS"
	EventAuthorityIssue      EventType = "AUTHORITY_ISSUE"
	EventAuthorityRevoke     EventType = "AUTHORITY_REVOKE"
	EventAuthorityFreeze     EventType = "AUTHORITY_FREEZE"
	EventAuthorityQuarantine EventType = "AUTHORITY_QUARANTINE"
	EventDelegation          EventType = "AUTHORITY_DELEGATE"
	EventCapabilityMint      EventType = "CAPABILITY_MINT"
	EventCapabilityRevoke    EventType = "CAPABILITY_REVOKE"
	EventFact                EventType = "FACT_RECORD"
	EventFactUpdate          EventType = "FACT_UPDATE"
	EventRelationship        EventType = "RELATIONSHIP_RECORD"
	EventBelief              EventType = "BELIEF_RECORD"
	EventPrediction          EventType = "PREDICTION_RECORD"
	EventInference           EventType = "INFERENCE_TRACE"
	EventHypothesis          EventType = "HYPOTHESIS_RECORD"
	EventSimulationResult    EventType = "SIMULATION_RESULT"
	EventEnforce             EventType = "ENFORCEMENT_ACTION"
	EventFactRecord          EventType = "FACT_RECORD"
	EventDoctrineProposal    EventType = "DOCTRINE_PROPOSAL"
	EventDoctrineActivated   EventType = "DOCTRINE_ACTIVATED"
	EventEngineVersionCommitted EventType = "ENGINE_VERSION_COMMITTED"
	EventAdjudication        EventType = "ADJUDICATION_ACTION"
	EventAnomalySpawned      EventType = "ANOMALY_SPAWNED"
	EventInstitutionCreated  EventType = "INSTITUTION_CREATED"
	EventPopulationGrowth    EventType = "POPULATION_GROWTH"
	EventRecovery            EventType = "RECOVERY_SIGNAL"
)

// Payload definitions for deterministic unmarshaling during replay.

type InstitutionCreatedPayload struct {
	SchemaVersion int    `json:"schema_version"`
	ID            string `json:"id"`
	Type          string `json:"type"`
}

type PopulationGrowthPayload struct {
	SchemaVersion int    `json:"schema_version"`
	InstitutionID string `json:"institution_id"`
	GrowthAmount  int64  `json:"growth_amount"`
}


type HypothesisPayload struct {
	SchemaVersion   int             `json:"schema_version"`
	HypothesisID    string          `json:"hypothesis_id"`
	BeliefID        string          `json:"belief_id"`
	ExpectedOutcome string          `json:"expected_outcome"`
	Confidence      ConfidenceScore `json:"confidence"`
}

type SimulationResultPayload struct {
	SchemaVersion int     `json:"schema_version"`
	SimulationID  string  `json:"simulation_id"`
	HypothesisID  string  `json:"hypothesis_id"`
	Scenario      string  `json:"scenario"`
	Success       bool    `json:"success"`
	Divergence    float64 `json:"divergence"`
}

type FactPayload struct {
	SchemaVersion   int                        `json:"schema_version"`
	FactID          string                     `json:"fact_id"`
	Observations    map[string]Observation     `json:"observations"`
	ConfidenceScore ConfidenceScore            `json:"confidence_score"`
	Timestamp       int64                      `json:"timestamp"`
}

type FactUpdatePayload struct {
	SchemaVersion   int             `json:"schema_version"`
	FactID          string          `json:"fact_id"`
	ConfidenceScore ConfidenceScore `json:"confidence_score"`
	Reason          string          `json:"reason"`
}

type RelationshipPayload struct {
	SchemaVersion int     `json:"schema_version"`
	FromID        string  `json:"from_id"`
	ToID          string  `json:"to_id"`
	Relation      string  `json:"relation"`
	Weight        float64 `json:"weight"`
}

type BeliefPayload struct {
	SchemaVersion int             `json:"schema_version"`
	BeliefID      string          `json:"belief_id"`
	FactIDs       []string        `json:"fact_ids"`
	Confidence    ConfidenceScore `json:"confidence"`
	Statement     string          `json:"statement"`
}

type PredictionPayload struct {
	SchemaVersion int    `json:"schema_version"`
	PredictionID  string `json:"prediction_id"`
	ActionID      string `json:"action_id"`
	Expects       []byte `json:"expects"`
	Timestamp     int64  `json:"timestamp"`
}

type InferencePayload struct {
	SchemaVersion int     `json:"schema_version"`
	RequestID     string  `json:"request_id"`
	Provider      string  `json:"provider_name"`
	Logic         string  `json:"logic"`
	Confidence    float64 `json:"confidence"`
}

type AuthorityIssuePayload struct {
	SchemaVersion int           `json:"schema_version"`
	ID            string        `json:"id"`
	Atoms         AuthorityAtom `json:"atoms"`
}

type AuthorityRevokePayload struct {
	SchemaVersion int    `json:"schema_version"`
	ID            string `json:"id"`
	Reason        string `json:"reason"`
}

type FactRecordPayload struct {
	SchemaVersion int    `json:"schema_version"`
	ArchiveID     string `json:"archive_id"`
	FactData      []byte `json:"fact_data"`
	Evidence      []byte `json:"evidence"`
}

type DoctrineProposalPayload struct {
	SchemaVersion int    `json:"schema_version"`
	InstitutionID string `json:"institution_id"`
	DoctrineID    string `json:"doctrine_id"`
	Content       []byte `json:"content"`
}

type DoctrineActivatedPayload struct {
	SchemaVersion int    `json:"schema_version"`
	DoctrineID    string `json:"doctrine_id"`
	Sequence      uint64 `json:"sequence"`
}

type EngineVersionCommittedPayload struct {
	SchemaVersion          int    `json:"schema_version"`
	Version                string `json:"version"`
	SubsystemManifestHash  string `json:"subsystem_manifest_hash"`
}

type AdjudicationEventPayload struct {
	SchemaVersion int    `json:"schema_version"`
	CourtID       string `json:"court_id"`
	TargetID      string `json:"target_id"`
	Verdict       string `json:"verdict"`
	EvidenceRef   string `json:"evidence_ref"`
}

type AuthorityStatusPayload struct {
	SchemaVersion int    `json:"schema_version"`
	ID            string `json:"id"`
	Status        string `json:"status"` // ACTIVE, FROZEN, QUARANTINED, REVOKED
	Reason        string `json:"reason"`
}

type AnomalySpawnedPayload struct {
	SchemaVersion int     `json:"schema_version"`
	AnomalyID     string  `json:"anomaly_id"`
	TargetID      string  `json:"target_id"`
	InitialDrift  float64 `json:"initial_drift"`
}

type AuthorityDelegatePayload struct {
	ParentID   string                   `json:"parent_id"`
	ChildAtoms map[string]AuthorityAtom `json:"child_atoms"`
}

type CapabilityMintPayload struct {
	TokenID     string        `json:"token_id"`
	AuthorityID string        `json:"authority_id"`
	Atoms       AuthorityAtom `json:"atoms"`
	Scope       []string      `json:"scope"`
}

type CapabilityRevokePayload struct {
	TokenID string `json:"token_id"`
	Reason  string `json:"reason"`
}

/**
 * Event
 *
 * Represents a single, immutable record in the Phoenix Ledger.
 *
 * Lifecycle:
 * Created by authorized components, appended to the linear chain, and 
 * persisted indefinitely. NEVER modified after creation.
 */
type Event struct {
	Sequence      uint64    `json:"sequence"`
	Type          EventType `json:"type"`
	Timestamp     time.Time `json:"timestamp"`
	Payload       []byte    `json:"payload"`
	ParentHash    string    `json:"parent_hash"`
	AuthorityRef  string    `json:"authority_ref"` // Reference to the issuing authority
	Hash          string    `json:"hash"`
}

/**
 * CalculateHash
 *
 * Generates the SHA-256 hash of the Event content.
 *
 * This ensures the content of the event cannot be changed after creation.
 * It excludes the Hash field itself to prevent circular reference.
 *
 * Complexity: O(PayloadSize)
 * Thread Safety: Read-only access to Event.
 */
func (e *Event) CalculateHash() string {
	// We exclude the Hash field itself from the calculation
	type tempEvent Event
	t := tempEvent(*e)
	t.Hash = ""
	data, _ := json.Marshal(t)
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash)
}

/**
 * NewEvent
 *
 * Creates a new event, linking it to its parent and authorizing it.
 *
 * Input:
 * - seq: The next logical sequence number in the chain.
 * - etype: Category of the transition.
 * - payload: Structured data for the event.
 * - parentHash: The hash of the immediately preceding event.
 * - authRef: The ID of the authority authorizing this event.
 *
 * Output:
 * - An initialized and internally hashed Event pointer.
 *
 * Side Effects:
 * - Captures the current UTC time.
 */
func NewEvent(seq uint64, etype EventType, payload []byte, parentHash string, authRef string) *Event {
	e := &Event{
		Sequence:     seq,
		Type:         etype,
		Timestamp:    time.Now().UTC(),
		Payload:      payload,
		ParentHash:   parentHash,
		AuthorityRef: authRef,
	}
	e.Hash = e.CalculateHash()
	return e
}
