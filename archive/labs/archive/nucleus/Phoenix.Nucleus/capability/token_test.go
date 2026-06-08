/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package capability

import (
	"testing"
)

func TestToken_Attenuate(t *testing.T) {
	rootToken := &Token{
		ID:          "ROOT_TOKEN",
		AuthorityID: "AUTH_0",
		Atoms:       100,
	}

	child, err := rootToken.Attenuate("CHILD_TOKEN", 40, []string{"read"})
	if err != nil {
		t.Fatalf("attenuation failed: %v", err)
	}

	if child.Atoms != 40 {
		t.Errorf("child atoms mismatch: expected 40, got %v", child.Atoms)
	}

	if rootToken.Atoms != 60 {
		t.Errorf("parent atoms mismatch: expected 60, got %v", rootToken.Atoms)
	}

	// Try to over-attenuate
	if _, err := rootToken.Attenuate("FAIL", 70, nil); err == nil {
		t.Errorf("failed to detect over-attenuation")
	}
}
