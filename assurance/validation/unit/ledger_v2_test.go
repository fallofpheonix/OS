/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

func TestLedgerV2StateTransitionContinuity(t *testing.T) {
	alloc := resource.NewBoundedAllocator(1024*1024, 100)
	l := NewLedger(alloc)

	// Add initial state entry (NORMAL)
	err := l.AddEntryV2("E1", "CAUSE1", []byte("init"), "", "", "NORMAL", "1.0.0")
	if err != nil {
		t.Fatalf("AddEntryV2 failed: %v", err)
	}

	// Add secondary valid transition (NORMAL -> SUSPICIOUS)
	err = l.AddEntryV2("E2", "CAUSE2", []byte("actuate"), "", "NORMAL", "SUSPICIOUS", "1.0.0")
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
	_ = l.AddEntryV2("E1", "CAUSE1", []byte("actuate"), "", "NORMAL", "SUSPICIOUS", "1.0.0")

	// Try to add a transition that starts from CONTAINED (violates parent state 'SUSPICIOUS')
	_ = l.AddEntryV2("E2", "CAUSE2", []byte("bad-actuate"), "", "CONTAINED", "RECOVERY", "1.0.0")

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

	_ = l.AddEntryV2("E1", "CAUSE1", []byte("data"), "", "NORMAL", "SUSPICIOUS", "1.0.0")

	// Find the entry in the map
	var entryKey string
	for k := range l.Entries {
		entryKey = k
		break
	}

	entry := l.Entries[entryKey]
	// Tamper with validation hash
	entry.ValidationHash = []byte("tampered-validation-hash")
	l.Entries[entryKey] = entry

	// Verify should fail due to invalid validation hash
	err := l.Verify()
	if err == nil {
		t.Error("Expected verification failure due to tampered validation hash, got nil")
	} else {
		t.Logf("Got expected validation error: %v", err)
	}
}
