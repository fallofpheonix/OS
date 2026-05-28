package boot

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"
)

// SubsystemInfo holds metadata about a booted subsystem.
type SubsystemInfo struct {
	Name      string                 `json:"name"`
	Version   string                 `json:"version"`
	Config    map[string]interface{} `json:"config,omitempty"`
	Timestamp int64                  `json:"timestamp"`
}

// NewSubsystemInfo creates a new info struct.
func NewSubsystemInfo(name, version string, config map[string]interface{}) SubsystemInfo {
	return SubsystemInfo{
		Name:      name,
		Version:   version,
		Config:    config,
		Timestamp: time.Now().UnixNano(),
	}
}

// BootTelemetry represents the state of the system at genesis.
type BootTelemetry struct {
	Subsystems []SubsystemInfo `json:"subsystems"`
	Checksum   string          `json:"checksum"`
}

// Ledger is a mock interface for the actual ledger.
type Ledger interface {
	AddEntry(eventType, eventID string, payload []byte) error
}

// CaptureBootTelemetry records the genesis state and returns the hash.
func CaptureBootTelemetry(l Ledger, info []SubsystemInfo) (BootTelemetry, error) {
	bt := BootTelemetry{Subsystems: info}
	
	data, err := json.Marshal(bt)
	if err != nil {
		return bt, err
	}
	
	hash := sha256.Sum256(data)
	bt.Checksum = fmt.Sprintf("%x", hash)
	
	if err := l.AddEntry("SYSTEM-BOOT", "GENESIS", data); err != nil {
		return bt, err
	}
	
	return bt, nil
}

// VerifyBoot compares the current boot checksum against an expected value.
func VerifyBoot(actual, expected string) error {
	if actual != expected {
		return fmt.Errorf("integrity violation: expected %s, got %s", expected, actual)
	}
	return nil
}
