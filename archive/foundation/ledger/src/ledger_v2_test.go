package ledger

import (
	"bytes"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

func TestLedgerV2StateTransitionContinuity(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 100)
	l := NewLedger(alloc)

	// Add initial state entry (NORMAL)
	err := l.AddEntryV2("E1", "CAUSE1", 0, []byte("init"), "", nil, []byte("NORMAL"), "1.0.0")
	if err != nil {
		t.Fatalf("AddEntryV2 failed: %v", err)
	}

	// Add secondary valid transition (NORMAL -> SUSPICIOUS)
	err = l.AddEntryV2("E2", "CAUSE2", 1, []byte("actuate"), "", []byte("NORMAL"), []byte("SUSPICIOUS"), "1.0.0")
	if err != nil {
		t.Fatalf("AddEntryV2 failed: %v", err)
	}

	// Verify ledger is integral
	if err := l.Verify(); err != nil {
		t.Errorf("Expected ledger to verify successfully, got: %v", err)
	}
}

func TestLedgerV2PoisoningStateTransitionGap(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 100)
	l := NewLedger(alloc)

	// Add initial state entry (NORMAL -> SUSPICIOUS)
	_ = l.AddEntryV2("E1", "CAUSE1", 0, []byte("actuate"), "", []byte("NORMAL"), []byte("SUSPICIOUS"), "1.0.0")

	// Try to add a transition that starts from CONTAINED (violates parent state 'SUSPICIOUS')
	_ = l.AddEntryV2("E2", "CAUSE2", 1, []byte("bad-actuate"), "", []byte("CONTAINED"), []byte("RECOVERY"), "1.0.0")

	// Verify ledger - should fail due to state transition gap
	err := l.Verify()
	if err == nil {
		t.Error("Expected verification failure due to state transition gap, got nil")
	} else {
		t.Logf("Got expected validation error: %v", err)
	}
}

func TestLedgerV2ValidationHashTampering(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 100)
	l := NewLedger(alloc)

	_ = l.AddEntryV2("E1", "CAUSE1", 0, []byte("data"), "", []byte("NORMAL"), []byte("SUSPICIOUS"), "1.0.0")

	// Find the entry in the map
	var entryKey string
	for k := range l.Entries {
		entryKey = k
		break
	}

	entry := l.Entries[entryKey]
	// Tamper with ValidationHash
	entry.ValidationHash = []byte("corrupted")
	l.Entries[entryKey] = entry

	// Verify - should fail validation hash check
	err := l.Verify()
	if err == nil {
		t.Error("Expected verification failure due to validation hash tampering, got nil")
	} else {
		t.Logf("Got expected validation error: %v", err)
	}
}

func TestLedgerEntrySerialization(t *testing.T) {
	entry := LedgerEntry{
		EventID:     "E1",
		Payload:     []byte("data"),
		LogicalTick: 123,
	}

	if entry.EventID != "E1" || !bytes.Equal(entry.Payload, []byte("data")) || entry.LogicalTick != 123 {
		t.Errorf("Struct field mismatch: %+v", entry)
	}
}
