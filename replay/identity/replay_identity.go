package identity

import (
	"fmt"
	"time"
)

// DivergenceType classifies the nature of the detected state deviation.
type DivergenceType string

const (
	DivNone        DivergenceType = "NONE"
	DivHash        DivergenceType = "HASH_MISMATCH"
	DivSequence    DivergenceType = "SEQUENCE_GAP"
	DivTransition  DivergenceType = "INVALID_TRANSITION"
	DivByzantine   DivergenceType = "BYZANTINE_FAULT"
)

// ReplayReport represents the outcome of a replay identity check with forensic detail.
type ReplayReport struct {
	ReplayID       string         `json:"replay_id"`
	HashBefore     string         `json:"hash_before"`
	HashAfter      string         `json:"hash_after"`
	Divergence     bool           `json:"divergence"`
	DivergenceTick uint64         `json:"divergence_tick,omitempty"`
	DivType        DivergenceType `json:"divergence_type,omitempty"`
	Status         string         `json:"status"` // e.g., "VERIFIED", "DIVERGED", "ERROR"
	Timestamp      time.Time      `json:"timestamp"`
	Message        string         `json:"message,omitempty"`
}

// ReplaySession represents an active replay process.
type ReplaySession struct {
	ID          string
	StartTime   time.Time
	InitialHash string
	CurrentTick uint64
}

// ReplayIdentity defines the concrete struct for verifying the identity and determinism of a replay.
type ReplayIdentity struct {
	Hash      string
	SessionID string
	Active    bool

	// Metrics for active session
	TicksProcessed uint64
	LastVerifiedTick uint64

	// Backward Compatibility Fields (for evidence_export.go)
	InputHash  string
	OutputHash string
	Divergence bool
	ReplayID   string
	Timestamp  int64
}

// NewReplayIdentity creates a new instance of ReplayIdentity.
func NewReplayIdentity() *ReplayIdentity {
	return &ReplayIdentity{
		Hash: "initial_placeholder_hash",
	}
}

// StartSession initiates a new replay session.
func (r *ReplayIdentity) StartSession(replayID string) (*ReplaySession, error) {
	if replayID == "" {
		return nil, fmt.Errorf("replay ID cannot be empty")
	}
	r.Active = true
	r.SessionID = replayID
	r.ReplayID = replayID
	session := &ReplaySession{
		ID:          replayID,
		StartTime:   time.Now(),
		InitialHash: r.Hash,
	}
	return session, nil
}

// EndSession concludes a replay session.
func (r *ReplayIdentity) EndSession(session *ReplaySession, finalHash string) (*ReplayReport, error) {
	if session == nil {
		return nil, fmt.Errorf("replay session cannot be nil")
	}
	if finalHash == "" {
		return nil, fmt.Errorf("final state hash cannot be empty")
	}
	
	r.Active = false
	divergence := session.InitialHash != finalHash
	r.Divergence = divergence
	r.OutputHash = finalHash
	r.Timestamp = time.Now().Unix()

	status := "VERIFIED"
	if divergence {
		status = "DIVERGED"
	}

	return &ReplayReport{
		ReplayID:   session.ID,
		HashBefore: session.InitialHash,
		HashAfter:  finalHash,
		Divergence: divergence,
		Status:     status,
		Timestamp:  time.Now(),
	}, nil
}

// CheckDivergence compares two replay states to detect divergence.
func (r *ReplayIdentity) CheckDivergence(hash1, hash2 string) bool {
    return hash1 != hash2
}

// ValidateHashChain verifies the integrity of a sequence of hashes.
// It ensures that no hash is empty and (if implemented) validates Merkle links.
func (r *ReplayIdentity) ValidateHashChain(hashes [][]byte) error {
	if len(hashes) == 0 {
		return nil
	}

	for i, h := range hashes {
		if len(h) == 0 {
			return fmt.Errorf("empty hash detected at index %d", i)
		}
	}

	// Future: Implement recursive Merkle-DAG link verification here.
	return nil
}

// RecordDivergence creates a report for an early termination due to state divergence.
func (r *ReplayIdentity) RecordDivergence(session *ReplaySession, tick uint64, divType DivergenceType, msg string) *ReplayReport {
	r.Active = false
	r.Divergence = true
	return &ReplayReport{
		ReplayID:       session.ID,
		HashBefore:     session.InitialHash,
		Divergence:     true,
		DivergenceTick: tick,
		DivType:        divType,
		Status:         "DIVERGED",
		Timestamp:      time.Now(),
		Message:        msg,
	}
}
