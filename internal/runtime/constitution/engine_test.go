package constitution

import (
	"github.com/fallofpheonix/phoenix/foundation/events"
	"testing"
)

func TestBootValidator_Failures(t *testing.T) {
	bv := NewBootValidator(NewEngine())

	// Test 1: Missing Constitution
	err := bv.ValidateBoot(event.Checkpoint{StateHash: "0x123"}, []string{"key1"}, "")
	if err == nil {
		t.Error("expected error for missing constitution, got nil")
	}

	// Test 2: Invalid Ledger
	err = bv.ValidateBoot(event.Checkpoint{}, []string{"key1"}, "hash123")
	if err == nil {
		t.Error("expected error for invalid ledger, got nil")
	}

	// Test 3: No Keys
	err = bv.ValidateBoot(event.Checkpoint{StateHash: "0x123"}, []string{}, "hash123")
	if err == nil {
		t.Error("expected error for no keys, got nil")
	}
}

func TestValidateTransition_CausalIntegrity(t *testing.T) {
	engine := NewEngine()

	current := event.Event{EventID: "E1"}

	// Valid child
	nextValid := event.Event{EventID: "E2", ParentID: "E1", Signature: "SIG", AuthorityID: "AUTH", IdentityID: "ID", Payload: []byte("{}")}
	err := engine.ValidateTransition(current, nextValid)
	if err != nil {
		t.Errorf("expected nil for valid transition, got %v", err)
	}

	// Invalid child (ParentID mismatch)
	nextInvalid := event.Event{EventID: "E3", ParentID: "WRONG"}
	err = engine.ValidateTransition(current, nextInvalid)
	if err == nil {
		t.Error("expected causal mismatch error, got nil")
	}
}
