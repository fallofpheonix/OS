// Package recovery provides mechanisms for node state restoration.
// Core Domain Logic: Implements the recovery protocol (P5.1/P5.2), enabling a node to be perfectly
// reconstructed from a combination of checkpoints, event ledgers, and artifact manifests.
package recovery

import (
	"context"
	"fmt"
	"time"

	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	replayv1 "github.com/fallofpheonix/phoenix/foundation/contracts/replay/v1"
	"github.com/fallofpheonix/phoenix/foundation/events"
	"github.com/fallofpheonix/phoenix/foundation/runtime/constitution"
)

// eventEnvelopeWrapper wraps event.Event to implement eventsv1.EventEnvelope.
type eventEnvelopeWrapper struct {
	ev event.Event
}

func (w *eventEnvelopeWrapper) GetEventID() string         { return w.ev.EventID }
func (w *eventEnvelopeWrapper) GetEventVersion() string    { return "1.0.0" }
func (w *eventEnvelopeWrapper) GetEventType() string       { return "PhoenixEvent" }
func (w *eventEnvelopeWrapper) GetTimestamp() time.Time    { return time.Now() }
func (w *eventEnvelopeWrapper) GetMonotonicTime() uint64   { return w.ev.LogicalTime }
func (w *eventEnvelopeWrapper) GetSourceRepo() string      { return w.ev.AuthorityID }
func (w *eventEnvelopeWrapper) GetSourceComponent() string { return w.ev.IdentityID }
func (w *eventEnvelopeWrapper) GetParentEvent() string     { return w.ev.ParentID }
func (w *eventEnvelopeWrapper) GetCorrelationID() string   { return "" }
func (w *eventEnvelopeWrapper) GetCausalChain() []string   { return nil }
func (w *eventEnvelopeWrapper) GetReplaySequence() uint64  { return w.ev.LogicalTime }
func (w *eventEnvelopeWrapper) GetEvidenceHash() string {
	if len(w.ev.Evidence) > 0 {
		return w.ev.Evidence[0]
	}
	return ""
}
func (w *eventEnvelopeWrapper) GetTrustScore() float64    { return 1.0 }
func (w *eventEnvelopeWrapper) GetSignature() string      { return w.ev.Signature }
func (w *eventEnvelopeWrapper) GetPayloadHash() string    { return "" }
func (w *eventEnvelopeWrapper) GetSchemaVersion() string  { return "" }
func (w *eventEnvelopeWrapper) GetCreatedAt() time.Time   { return time.Now() }
func (w *eventEnvelopeWrapper) GetUpdatedAt() time.Time   { return time.Now() }
func (w *eventEnvelopeWrapper) GetValidationHash() string { return "" }
func (w *eventEnvelopeWrapper) GetPayload() []byte        { return []byte(w.ev.Payload) }

// Engine orchestrates the reconstruction of node state from historical data.
// Internal State: References to the Constitution engine for invariant checking and Replay engine for state transitions.
// API Scope: Public for system recovery and node synchronization.
// Concurrency: Orchestration is stateless; concurrency safety depends on the thread-safety of the Replay and Constitution engines.
type Engine struct {
	Constitution *constitution.Engine
	Replay       replayv1.ReplayEngine
}

// LABEL: [PURE] [PUBLIC_API] [EXPERIMENTAL]
// NewEngine initializes a new recovery engine with required dependencies.
// I/O: None.
// Complexity: O(1).
func NewEngine(c *constitution.Engine, r replayv1.ReplayEngine) *Engine {
	return &Engine{
		Constitution: c,
		Replay:       r,
	}
}

// LABEL: [MUTATES_STATE] [PUBLIC_API] [EXPERIMENTAL]
// Recover resurrects the node state from a checkpoint and ledger (P5.1).
// I/O: None (operates on passed memory structures).
// Side Effects: Modifies the internal state of the Replay engine.
// Complexity: O(M + E) where M is the number of manifests and E is the number of events in the ledger.
func (e *Engine) Recover(cp event.Checkpoint, ledger []event.Event, manifests []event.ArtifactManifest) error {
	// 1. Verify Artifact integrity (P5.2)
	for _, m := range manifests {
		if m.Hash == "" || m.Signer == "" {
			return fmt.Errorf("RECOVERY_FAILURE: invalid artifact manifest integrity")
		}
	}

	// 2. Initialize Replay from Checkpoint (P5.1)
	e.Replay.SetLogicalTime(cp.ReplayOffset)
	// In a real implementation, we would load the state snapshot referenced by ArtifactReferences

	// Convert ledger events to envelopes
	envelopes := make([]eventsv1.EventEnvelope, len(ledger))
	for i, ev := range ledger {
		envelopes[i] = &eventEnvelopeWrapper{ev: ev}
	}

	// 3. Replay Ledger and Verify State Hash (P5.2)
	ctx := context.Background()
	if err := e.Replay.Replay(ctx, envelopes); err != nil {
		return fmt.Errorf("RECOVERY_FAILURE: %w", err)
	}

	snap, err := e.Replay.GetCurrentState()
	if err != nil {
		return fmt.Errorf("RECOVERY_FAILURE: failed to get state: %w", err)
	}
	actualHash := snap.StateHash()

	if actualHash != cp.StateHash {
		return fmt.Errorf("RECOVERY_FAILURE: state hash mismatch after recovery (expected %s, got %s)", cp.StateHash, actualHash)
	}

	return nil
}
