/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package validators

import (
	"testing"
	"github.com/fallofpheonix/phoenix/governance/truth/engine"
	"github.com/fallofpheonix/phoenix/assurance/security"
)

func TestTransitionValidator(t *testing.T) {
	v := NewTransitionValidator()

	tests := []struct {
		before string
		after  string
		valid  bool
	}{
		{string(warden.StateSafe), string(warden.StateWatch), true},
		{string(warden.StateSafe), string(warden.StateSuspicious), false},
		{string(warden.StateWatch), string(warden.StateSafe), true},
		{string(warden.StateWatch), string(warden.StateSuspicious), true},
		{string(warden.StateSuspicious), string(warden.StateSafe), false},
		{"", "", true},
		{string(warden.StateSafe), string(warden.StateSafe), true},
	}

	for _, tt := range tests {
		entry := &ledger.LedgerEntry{
			StateBefore: tt.before,
			StateAfter:  tt.after,
		}
		res := v.Validate(entry)
		if res.Valid != tt.valid {
			t.Errorf("Validate(%s -> %s) = %v; want %v (Reason: %s)", tt.before, tt.after, res.Valid, tt.valid, res.Reason)
		}
	}
}

func TestSequenceValidator(t *testing.T) {
	v := NewSequenceValidator()

	entries := []*ledger.LedgerEntry{
		{LogicalTick: 10},
		{LogicalTick: 11},
		{LogicalTick: 10}, // Regression
		{LogicalTick: 12},
	}

	results := []bool{true, true, false, true}

	for i, entry := range entries {
		res := v.Validate(entry)
		if res.Valid != results[i] {
			t.Errorf("Step %d: Validate(tick=%d) = %v; want %v (Reason: %s)", i, entry.LogicalTick, res.Valid, results[i], res.Reason)
		}
	}
}

func TestEntropyValidator(t *testing.T) {
	v := NewEntropyValidator(4.0) // Low threshold for testing

	tests := []struct {
		payload []byte
		valid   bool
	}{
		{[]byte("aaaaa"), true},                                     // Zero entropy
		{[]byte("abcdefghijklmnopqrstuvwxyz1234567890"), false},    // Higher entropy
	}

	for _, tt := range tests {
		entry := &ledger.LedgerEntry{Payload: tt.payload}
		res := v.Validate(entry)
		if res.Valid != tt.valid {
			t.Errorf("Validate(payload=%s) = %v; want %v (Reason: %s)", string(tt.payload), res.Valid, tt.valid, res.Reason)
		}
	}
}

func TestValidatorRegistry(t *testing.T) {
	registry := NewValidatorRegistry()
	registry.Register(NewSequenceValidator())
	registry.Register(NewTransitionValidator())

	// Valid sequence and transition
	entry1 := &ledger.LedgerEntry{
		LogicalTick: 10,
		StateBefore: string(warden.StateSafe),
		StateAfter:  string(warden.StateWatch),
	}
	failures := registry.ValidateAll(entry1)
	if len(failures) > 0 {
		t.Errorf("Expected 0 failures, got %d: %v", len(failures), failures)
	}

	// Invalid sequence, valid transition
	entry2 := &ledger.LedgerEntry{
		LogicalTick: 9,
		StateBefore: string(warden.StateWatch),
		StateAfter:  string(warden.StateSuspicious),
	}
	failures = registry.ValidateAll(entry2)
	if len(failures) != 1 {
		t.Errorf("Expected 1 failure (sequence), got %d", len(failures))
	} else if failures[0].Reason == "" {
		t.Errorf("Failure reason should not be empty")
	}

	// Valid sequence, invalid transition
	entry3 := &ledger.LedgerEntry{
		LogicalTick: 11,
		StateBefore: string(warden.StateSuspicious),
		StateAfter:  string(warden.StateSafe),
	}
	failures = registry.ValidateAll(entry3)
	if len(failures) != 1 {
		t.Errorf("Expected 1 failure (transition), got %d", len(failures))
	}
}
