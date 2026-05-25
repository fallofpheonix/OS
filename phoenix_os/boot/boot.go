package boot

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/truth_ledger/src"
)

// SubsystemInfo holds metadata about a subsystem's initial state.
type SubsystemInfo struct {
	Name    string                 `json:"name"`
	Version string                 `json:"version"`
	Config  map[string]interface{} `json:"config"`
	Status  string                 `json:"status"`
}

// BootTelemetry represents the full state of the system at boot.
type BootTelemetry struct {
	Timestamp  int64           `json:"timestamp"`
	Subsystems []SubsystemInfo `json:"subsystems"`
	Checksum   string          `json:"checksum"`
}

// CaptureBootTelemetry aggregates subsystem info and records it to the ledger.
func CaptureBootTelemetry(evLedger *ledger.Ledger, subsystems []SubsystemInfo) (*BootTelemetry, error) {
	bt := &BootTelemetry{
		Timestamp:  time.Now().Unix(),
		Subsystems: subsystems,
	}

	// Calculate Deterministic Checksum
	checksum, err := bt.CalculateChecksum()
	if err != nil {
		return nil, err
	}
	bt.Checksum = checksum

	// Record GENESIS event in Ledger
	payload, _ := json.Marshal(bt)
	evLedger.AddEntry("GENESIS", "BOOT-SEQUENCE", payload)

	return bt, nil
}

// CalculateChecksum generates a SHA-256 hash of the boot telemetry (excluding the checksum itself).
func (bt *BootTelemetry) CalculateChecksum() (string, error) {
	// To ensure determinism, we marshal the subsystems which are ordered.
	data, err := json.Marshal(bt.Subsystems)
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(data)
	return fmt.Sprintf("%x", hash), nil
}

// VerifyBoot verifies if the current boot matches a target checksum.
func VerifyBoot(current, expected string) error {
	if current != expected {
		return fmt.Errorf("boot checksum mismatch! expected %s, got %s", expected, current)
	}
	return nil
}

// NewSubsystemInfo is a helper to create info for a subsystem.
func NewSubsystemInfo(name, version string, config map[string]interface{}) SubsystemInfo {
	return SubsystemInfo{
		Name:    name,
		Version: version,
		Config:  config,
		Status:  "INITIALIZED",
	}
}
