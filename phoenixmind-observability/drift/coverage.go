package drift

import (
	"encoding/json"
	"os"
)

type ModuleStatus struct {
	Active bool `json:"active"`
}

func GetCoverageRatio(auditPath string) (float64, error) {
	file, err := os.Open(auditPath)
	if err != nil {
		return 0.0, err
	}
	defer file.Close()

	var status map[string]ModuleStatus
	if err := json.NewDecoder(file).Decode(&status); err != nil {
		return 0.0, err
	}

	active := 0
	for _, s := range status {
		if s.Active {
			active++
		}
	}
	
	if len(status) == 0 {
		return 0.0, nil
	}
	return float64(active) / float64(len(status)), nil
}
