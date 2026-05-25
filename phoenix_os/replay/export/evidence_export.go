package export

import (
	"encoding/json"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/replay/identity"
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
