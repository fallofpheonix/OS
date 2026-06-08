/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package export

import (
	"encoding/json"
	"testing"
	"github.com/fallofpheonix/replay/identity"
)

func TestEvidenceExport(t *testing.T) {
	replay := identity.ReplayIdentity{
		ReplayID:   "replay-123",
		InputHash:  "abc",
		OutputHash: "def",
		Divergence: true,
		Timestamp:  123456789,
	}

	evidence := CreateEvidenceFromReplay("arbiter", replay)
	
	if evidence.Entity != "arbiter" {
		t.Errorf("Expected entity 'arbiter', got '%s'", evidence.Entity)
	}
	if evidence.InputHash != "abc" {
		t.Errorf("Expected input hash 'abc', got '%s'", evidence.InputHash)
	}
	if !evidence.Divergence {
		t.Errorf("Expected divergence to be true, but it was false")
	}
	if evidence.TruthState != "OBSERVED" {
		t.Errorf("Expected truth state 'OBSERVED', got '%s'", evidence.TruthState)
	}

	jsonData, err := evidence.ToJSON()
	if err != nil {
		t.Fatalf("Failed to serialize evidence to JSON: %v", err)
	}

	var decoded Evidence
	err = json.Unmarshal(jsonData, &decoded)
	if err != nil {
		t.Fatalf("Failed to deserialize JSON to evidence: %v", err)
	}

	if decoded.OutputHash != "def" {
		t.Errorf("Expected output hash 'def' after JSON roundtrip, got '%s'", decoded.OutputHash)
	}
}
