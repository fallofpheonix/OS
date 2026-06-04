/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — ARTIFACT STORAGE WITH INTEGRITY
//
// SaveRun persists run data (summary, metrics, events, validation) to disk
// with SHA-256 integrity hashes. This provides tamper-evident storage for
// forensic analysis and replay verification.
//
// WORKFLOW:
//   SaveRun(baseDir, runID, data)
//     → Create run directory
//     → Write summary.json, metrics.json, events.json, validation.json
//     → Compute SHA-256 hash of each file
//     → Write hashes.sha256 manifest
//
// INTEGRITY: The hashes.sha256 file contains SHA-256 hashes of all artifacts.
// Any modification to an artifact will be detected by re-hashing and comparing.
//
// USAGE: Called by the Replay Engine and Validation system to persist
// deterministic replay results for later verification.
// =========================================================================
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

func SaveRun(baseDir, runID string, data RunData) error {
	runDir := filepath.Join(baseDir, runID)
	if err := os.MkdirAll(runDir, 0o755); err != nil {
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
		if err := os.WriteFile(path, jsonData, 0o644); err != nil {
			return err
		}

		h := sha256.Sum256(jsonData)
		hashes += fmt.Sprintf("%x  %s\n", h, name)
	}

	return os.WriteFile(filepath.Join(runDir, "hashes.sha256"), []byte(hashes), 0o644)
}
