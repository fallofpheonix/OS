package artifacts

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

type RunData struct {
	Summary    interface{}
	Metrics    interface{}
	Events     interface{}
	Validation interface{}
}

func SaveRun(baseDir string, runID string, data RunData) error {
	runDir := filepath.Join(baseDir, runID)
	if err := os.MkdirAll(runDir, 0755); err != nil {
		return err
	}

	files := map[string]interface{}{
		"summary.json":    data.Summary,
		"metrics.json":    data.Metrics,
		"events.json":     data.Events,
		"validation.json": data.Validation,
	}

	var hashes string
	for name, content := range files {
		path := filepath.Join(runDir, name)
		jsonData, err := json.MarshalIndent(content, "", "  ")
		if err != nil {
			return err
		}
		if err := os.WriteFile(path, jsonData, 0644); err != nil {
			return err
		}
		
		h := sha256.Sum256(jsonData)
		hashes += fmt.Sprintf("%x  %s\n", h, name)
	}

	return os.WriteFile(filepath.Join(runDir, "hashes.sha256"), []byte(hashes), 0644)
}
