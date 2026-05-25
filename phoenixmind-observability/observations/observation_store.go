package observations

import (
	"encoding/json"
	"fmt"
	"os"
)

// Store saves the Observation to a file.
func Store(obs *Observation) error {
	filename := fmt.Sprintf("%s.json", obs.Cycle)
	data, err := json.MarshalIndent(obs, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}
