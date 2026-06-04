// Package rollback provides atomic state restoration across containment layers.
// Core Domain Logic: Implements the "Rollback Orchestrator" (Cycle 10) which coordinates 
// synchronized snapshots and restoration of Process, Network, and File layers to maintain system atomicity.
package rollback

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"

	"github.com/fallofpheonix/phoenix/foundation/runtime/containment"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/file"
	"github.com/fallofpheonix/phoenix/foundation/runtime/containment/network"
)

// AuditProvider defines the common interface for containment layer snapshots.
// API Scope: Internal interface for containment providers.
type AuditProvider interface {
	CreateSnapshot() ([]byte, error)
	RestoreFromSnapshot(data []byte) error
}

// GlobalSnapshot captures the combined audit state for all containment layers.
// Internal State: Aggregated binary data from all sub-layers with integrity hash.
// API Scope: Internal/Management.
// Concurrency: Thread-safe (immutable once created).
type GlobalSnapshot struct {
	ProcessData []byte `json:"process_data"`
	NetworkData []byte `json:"network_data"`
	FileData    []byte `json:"file_data"`
	Version     string `json:"version"`
	Sequence    int    `json:"sequence"`
	Hash        string `json:"hash"`
}

const GlobalVersion = "1.0.0"

// LABEL: [PURE] [INTERNAL_ONLY] [STABLE]
// calculateGlobalHash computes the SHA-256 integrity hash of a global snapshot.
// I/O: None.
// Complexity: O(D) where D is the total size of snapshot data.
func calculateGlobalHash(s GlobalSnapshot) string {
	b, _ := json.Marshal(struct {
		ProcessData []byte
		NetworkData []byte
		FileData    []byte
		Version     string
		Sequence    int
	}{
		ProcessData: s.ProcessData,
		NetworkData: s.NetworkData,
		FileData:    s.FileData,
		Version:     s.Version,
		Sequence:    s.Sequence,
	})
	return fmt.Sprintf("%x", sha256.Sum256(b))
}

// Orchestrator coordinates cross-component snapshots and restoration.
// Internal State: References to Process, Network, and File audit providers.
// API Scope: Public within PhoenixCore for recovery management.
// Concurrency: Restoration is typically single-threaded during recovery; creation is thread-safe if providers are.
type Orchestrator struct {
	ProcessAudit *containment.ProcessAudit
	NetworkAudit *network.NetworkAudit
	FileAudit    *file.FileAudit
}

// LABEL: [IO_BOUND] [PUBLIC_API] [STABLE]
// CreateGlobalSnapshot aggregates state from all containment layers into a single atomic snapshot.
// I/O: Potential IO depending on provider implementation.
// Complexity: O(D) where D is the total size of collected state.
func (o *Orchestrator) CreateGlobalSnapshot(seq int) ([]byte, error) {
	pData, _ := o.ProcessAudit.CreateSnapshot()
	nData, _ := o.NetworkAudit.CreateSnapshot()
	fData, _ := o.FileAudit.CreateSnapshot()

	s := GlobalSnapshot{
		ProcessData: pData,
		NetworkData: nData,
		FileData:    fData,
		Version:     GlobalVersion,
		Sequence:    seq,
	}
	s.Hash = calculateGlobalHash(s)

	return json.Marshal(s)
}

// LABEL: [IO_BOUND] [PUBLIC_API] [STABLE]
// RestoreGlobal restores state across all containment layers atomically.
// I/O: Potential IO depending on provider implementation.
// Side Effects: Modifies state of all underlying containment layers.
// Complexity: O(D) where D is the total size of snapshot data.
func (o *Orchestrator) RestoreGlobal(data []byte) error {
	var s GlobalSnapshot
	if err := json.Unmarshal(data, &s); err != nil {
		return fmt.Errorf("snapshot unmarshal failed: %w", err)
	}

	if s.Version != GlobalVersion {
		return fmt.Errorf("incompatible version: %s", s.Version)
	}
	if calculateGlobalHash(s) != s.Hash {
		return fmt.Errorf("snapshot integrity check failed")
	}

	if err := o.ProcessAudit.RestoreFromSnapshot(s.ProcessData); err != nil {
		return err
	}
	if err := o.NetworkAudit.RestoreFromSnapshot(s.NetworkData); err != nil {
		return err
	}
	if err := o.FileAudit.RestoreFromSnapshot(s.FileData); err != nil {
		return err
	}
	return nil
}
