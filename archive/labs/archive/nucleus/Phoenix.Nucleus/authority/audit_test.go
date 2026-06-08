/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package authority

import (
	"testing"
	"github.com/fallofpheonix/Phoenix.Nucleus/ledger"
)

func TestAuthority_Conservation(t *testing.T) {
	reg := NewRegistry()
	rootID := "ROOT"
	var total ledger.AuthorityAtom = 100
	
	reg.IssueRootAuthority(rootID, total)
	
	children := map[string]ledger.AuthorityAtom{
		"CHILD_A": 40,
		"CHILD_B": 30,
	}
	
	if err := reg.DelegateAuthority(rootID, children); err != nil {
		t.Fatalf("delegation failed: %v", err)
	}
	
	if err := reg.AuditConservation(rootID, total); err != nil {
		t.Errorf("audit failed: %v", err)
	}
	
	// Try to over-delegate
	over := map[string]ledger.AuthorityAtom{
		"CHILD_C": 40,
	}
	if err := reg.DelegateAuthority(rootID, over); err == nil {
		t.Errorf("failed to detect over-delegation")
	}
}
