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
	"github.com/fallofpheonix/replay/identity"
)

// Evidence represents the structured data to be exported.
type Evidence struct {
	Entity      string `json:"entity"`
	InputHash   string `json:"input_hash"`
	OutputHash  string `json:"output_hash"`
	Divergence  bool   `json:"divergence"`
	TruthState  string `json:"truth"`
}

// ToJSON serializes the evidence struct to a JSON byte slice.
func (e *Evidence) ToJSON() ([]byte, error) {
	return json.Marshal(e)
}

// CreateEvidenceFromReplay converts a ReplayIdentity to an Evidence struct.
func CreateEvidenceFromReplay(entity string, replayID identity.ReplayIdentity) Evidence {
	return Evidence{
		Entity:      entity,
		InputHash:   replayID.InputHash,
		OutputHash:  replayID.OutputHash,
		Divergence:  replayID.Divergence,
		TruthState:  "OBSERVED", // As per requirement, hardcode to OBSERVED for now
	}
}
