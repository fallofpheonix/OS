/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package truth_test

import (
	"testing"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestEvidenceRegistry(t *testing.T) {
	registry := evidence.NewEvidenceRegistry()
	
	e := &evidence.Evidence{
		EntityID: "arbiter",
		State:    evidence.VALIDATED,
	}

	registry.Add(e)

	got, ok := registry.Get("arbiter")
	if !ok {
		t.Fatal("expected evidence to be found")
	}

	if got.State != evidence.VALIDATED {
		t.Errorf("expected state %s, got %s", evidence.VALIDATED, got.State)
	}
}
