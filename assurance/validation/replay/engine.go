package replay

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/hash"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/serialization"
	eventsv1 "github.com/fallofpheonix/phoenix/foundation/contracts/events/v1"
	replayv1 "github.com/fallofpheonix/phoenix/foundation/contracts/replay/v1"
)

// STATUS: ACTIVE
// The Replay Engine implements ReplayEngine contract directly.

// State represents a deterministic snapshot of the system.
type State struct {
	LogicalTime uint64            `json:"logical_time"`
	Values      map[string]string `json:"values"`
}

// Engine implements the Replay Pipeline (P4.1).
type Engine struct {
	State State
}

// SnapshotImpl implements replayv1.Snapshot.
type SnapshotImpl struct {
	stateHash    string
	replayOffset uint64
	payload      []byte
}

func (s *SnapshotImpl) StateHash() string    { return s.stateHash }
func (s *SnapshotImpl) ReplayOffset() uint64 { return s.replayOffset }
func (s *SnapshotImpl) Payload() []byte      { return s.payload }

func NewEngine() *Engine {
	return &Engine{
		State: State{
			Values: make(map[string]string),
		},
	}
}

// Replay processes a stream of events to reconstruct state (P4.1).
func (e *Engine) Replay(ctx context.Context, envelopes []eventsv1.EventEnvelope) error {
	// 1. Ordering (P4.1)
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i].GetReplaySequence() < envelopes[j].GetReplaySequence()
	})

	for _, env := range envelopes {
		// 2. Validation (P4.1)
		if env.GetReplaySequence() < e.State.LogicalTime {
			return fmt.Errorf("non-monotonic logical time: event %d < state %d", env.GetReplaySequence(), e.State.LogicalTime)
		}

		// 3. Application (P4.1)
		if err := e.Apply(env); err != nil {
			return err
		}
	}

	return nil
}

// Apply updates the engine's internal state based on an event.
func (e *Engine) Apply(env eventsv1.EventEnvelope) error {
	var payload []byte
	
	// Attempt to retrieve payload via interface checks
	if pEv, ok := env.(interface{ GetPayload() []byte }); ok {
		payload = pEv.GetPayload()
	} else if pEv, ok := env.(interface{ GetPayload() json.RawMessage }); ok {
		payload = []byte(pEv.GetPayload())
	} else if pEv, ok := env.(interface{ GetPayload() string }); ok {
		payload = []byte(pEv.GetPayload())
	}

	var values map[string]string
	if err := json.Unmarshal(payload, &values); err != nil {
		// Fallback for non-map payloads
		e.State.Values["last_payload"] = string(payload)
	} else {
		for k, v := range values {
			e.State.Values[k] = v
		}
	}
	e.State.LogicalTime = env.GetReplaySequence()
	return nil
}

// GetCurrentState returns the current snapshot of the Replay Engine.
func (e *Engine) GetCurrentState() (replayv1.Snapshot, error) {
	stateHash := e.CalculateStateHash()
	payload, _ := json.Marshal(e.State.Values)

	return &SnapshotImpl{
		stateHash:    stateHash,
		replayOffset: e.State.LogicalTime,
		payload:      payload,
	}, nil
}

// SetLogicalTime sets the logical time offset directly.
func (e *Engine) SetLogicalTime(time uint64) {
	e.State.LogicalTime = time
}

// CalculateStateHash produces a canonical SHA-256 hash of the state using the WorldStateHasher (P4.2).
func (e *Engine) CalculateStateHash() string {
	// 1. Define Subsystem Manifest (this would typically be provided by Genesis/Determinism Contract)
	manifest := &hash.SubsystemManifest{Subsystems: []string{"State"}}
	hasher := &hash.WorldStateHasher{Manifest: manifest}

	// 2. Hash internal state deterministically
	b, _ := serialization.StableMarshal(e.State)
	stateHash := fmt.Sprintf("%x", sha256.Sum256(b))
	subsystemHashes := map[string]string{"State": stateHash}
	
	return hasher.ComputeHash(subsystemHashes)
}

// DivergenceDetector detects mismatch between expected and actual state (P4.3).
type DivergenceDetector struct{}

func (d *DivergenceDetector) Check(expectedHash string, actualHash string) error {
	if expectedHash != actualHash {
		return fmt.Errorf("REPLAY_DIVERGENCE: expected %s, got %s", expectedHash, actualHash)
	}
	return nil
}
