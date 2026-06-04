/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"encoding/json"
	"testing"
)

func TestMetricsExport(t *testing.T) {
	GlobalMetrics.Reset()
	r := NewStateRegistry(StateSafe)

	// Perform operations
	r.Transition(StateWatch, "b1", "e1", "d1")
	r.Transition(StateAlert, "b2", "e2", "d2")
	r.Rollback("r1", "e3", "d3")
	r.CreateSnapshot()
	GlobalMetrics.IncReplayMismatch()

	data, err := GlobalMetrics.ExportMetrics()
	if err != nil {
		t.Fatalf("export failed: %v", err)
	}

	var m map[string]interface{}
	if err := json.Unmarshal(data, &m); err != nil {
		t.Fatalf("unmarshal failed: %v", err)
	}

	// Verify all metric categories
	expectations := map[string]float64{
		"TransitionCount": 2,
		"Rollbacks":       1,
		"SnapshotCreates": 1,
		"ReplayMismatch":  1,
		"WatchEntries":    2,
		"AlertEntries":    1,
	}

	for k, v := range expectations {
		if val, ok := m[k].(float64); !ok || val != v {
			t.Errorf("expected %s to be %f, got %v", k, v, m[k])
		}
	}
}
