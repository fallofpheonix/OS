package rollback

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/file"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/network"
)

// AuditProvider defines the common interface for containment layer snapshots.
type AuditProvider interface {
	CreateSnapshot() ([]byte, error)
	RestoreFromSnapshot(data []byte) error
}

// GlobalSnapshot captures the combined audit state for all containment layers.
type GlobalSnapshot struct {
	ProcessData []byte `json:"process_data"`
	NetworkData []byte `json:"network_data"`
	FileData    []byte `json:"file_data"`
	Version     string `json:"version"`
	Sequence    int    `json:"sequence"`
	Hash        string `json:"hash"`
}

const GlobalVersion = "1.0.0"

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

// Orchestrator coordinates cross-component snapshots.
type Orchestrator struct {
	ProcessAudit *containment.ProcessAudit
	NetworkAudit *network.NetworkAudit
	FileAudit    *file.FileAudit
}

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

// RestoreGlobal restores state across all containment layers.
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
