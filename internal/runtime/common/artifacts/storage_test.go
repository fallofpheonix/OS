/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package artifacts

import (
	"os"
	"path/filepath"
	"testing"
)

func TestSaveRun(t *testing.T) {
	tmpDir, _ := os.MkdirTemp("", "phoenix_test")
	defer os.RemoveAll(tmpDir)

	data := RunData{
		Summary:    map[string]string{"status": "ok"},
		Metrics:    map[string]int{"cpu": 5},
		Events:     []string{"evt1", "evt2"},
		Validation: true,
	}

	runID := "run-001"
	err := SaveRun(tmpDir, runID, data)
	if err != nil {
		t.Fatalf("SaveRun failed: %v", err)
	}

	// Verify files
	runDir := filepath.Join(tmpDir, runID)
	expectedFiles := []string{"summary.json", "metrics.json", "events.json", "validation.json", "hashes.sha256"}
	for _, f := range expectedFiles {
		if _, err := os.Stat(filepath.Join(runDir, f)); os.IsNotExist(err) {
			t.Errorf("Expected file %s does not exist", f)
		}
	}
}
