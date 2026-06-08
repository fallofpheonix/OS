package security

import (
	"bytes"
	"encoding/binary"
	"testing"
)

func TestPayload_BinaryIntegrity(t *testing.T) {
	// Case 1: Known input -> Exact bytes
	req := AuthorityEscalationRequest{
		ActuationClass: ClassIsolate,
		TargetPID:      1234,
		TargetState:    StateCritical,
		PolicyHash:     [32]byte{0xAA, 0xBB, 0xCC},
	}

	got, err := serializePayload(req, true, StateWatch)
	if err != nil {
		t.Fatalf("Serialize failed: %v", err)
	}

	if len(got) != 43 {
		t.Fatalf("Size mismatch: got %d", len(got))
	}
	if got[0] != 0x01 {
		t.Errorf("Version mismatch: got %d", got[0])
	}
	if got[1] != 0x04 {
		t.Errorf("Class mismatch: got %d", got[1])
	}

	pid := binary.BigEndian.Uint32(got[2:6])
	if pid != 1234 {
		t.Errorf("PID mismatch: got %d", pid)
	}

	if got[6] != 0xAA || got[7] != 0xBB || got[8] != 0xCC {
		t.Errorf("PolicyHash mismatch")
	}

	if got[38] != 0x04 {
		t.Errorf("TargetState mismatch: got %d", got[38])
	}
	if got[39] != 0x01 {
		t.Errorf("Shadow bit mismatch: got %d", got[39])
	}
	if got[40] != 0x02 {
		t.Errorf("ShadowCurr mismatch: got %d", got[40])
	}

	// Verify deserialization
	reqD, shadowD, currD, err := deserializePayload(got)
	if err != nil {
		t.Fatalf("Deserialize failed: %v", err)
	}

	if reqD.ActuationClass != req.ActuationClass {
		t.Error("Deserialized class mismatch")
	}
	if reqD.TargetPID != req.TargetPID {
		t.Error("Deserialized PID mismatch")
	}
	if !bytes.Equal(reqD.PolicyHash[:], req.PolicyHash[:]) {
		t.Error("Deserialized hash mismatch")
	}
	if reqD.TargetState != req.TargetState {
		t.Error("Deserialized target state mismatch")
	}
	if shadowD != true {
		t.Error("Deserialized shadow flag mismatch")
	}
	if currD != StateWatch {
		t.Error("Deserialized current state mismatch")
	}
}

func TestPayload_ZeroValueTrap(t *testing.T) {
	// A zero-initialized buffer must fail deserialization
	data := make([]byte, 43)
	data[0] = 0x01 // Set version to bypass initial check

	_, _, _, err := deserializePayload(data)
	if err == nil {
		t.Error("Expected error for zero-initialized payload (ClassInvalid/StateInvalid)")
	}
}

func TestPayload_Capping(t *testing.T) {
	// PID cap at 0x7FFFFFFF
	req := AuthorityEscalationRequest{
		TargetPID:      0xFFFFFFFF,
		ActuationClass: ClassNone,
		TargetState:    StateSafe,
	}

	got, _ := serializePayload(req, false, StateSafe)
	pid := binary.BigEndian.Uint32(got[2:6])
	if pid != 0x7FFFFFFF {
		t.Errorf("PID cap failed: got 0x%x", pid)
	}
}

func TestPayload_Validation(t *testing.T) {
	req := AuthorityEscalationRequest{
		ActuationClass: ClassLog,
		TargetState:    StateSafe,
	}
	got, _ := serializePayload(req, false, StateSafe)

	// Test invalid shadow byte
	badShadow := make([]byte, 43)
	copy(badShadow, got)
	badShadow[39] = 0x02
	_, _, _, err := deserializePayload(badShadow)
	if err == nil {
		t.Error("Expected error for invalid isShadow byte")
	}

	// Test invalid target state
	badState := make([]byte, 43)
	copy(badState, got)
	badState[38] = 0x06
	_, _, _, err = deserializePayload(badState)
	if err == nil {
		t.Error("Expected error for invalid target state")
	}
}
