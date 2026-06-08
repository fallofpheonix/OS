/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package unit

import (
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/runtime/telemetry/detector"
	"github.com/fallofpheonix/phoenix/foundation/runtime/telemetry/events"
	lineage "github.com/fallofpheonix/phoenix/foundation/observability/engine/process_lineage"
)

func TestDetector(t *testing.T) {
	d := detector.NewDetector()
	g := lineage.NewLineageGraph()
	now := time.Now()

	// Case 1: Benign event
	evt1 := events.Event{
		PID:       100,
		Comm:      "ls",
		EventType: "execve",
		Payload:   map[string]any{"entropy_score": 3.0},
	}
	res1 := d.Analyze(evt1, g)
	if res1.IsThreat {
		t.Error("ls should not be a threat")
	}

	// Case 2: High entropy write from 'gpg'
	g.AddProcess(200, 1, "gpg", "/usr/bin/gpg", now)
	evt2 := events.Event{
		PID:       200,
		Comm:      "gpg",
		EventType: "write",
		Payload:   map[string]any{"entropy_score": 8.5},
	}
	res2 := d.Analyze(evt2, g)
	if !res2.IsThreat {
		t.Errorf("gpg high entropy write should be a threat. Score: %f", res2.ImportanceScore)
	}
	if res2.ImportanceScore < 0.8 {
		t.Errorf("Expected score >= 0.8, got %f", res2.ImportanceScore)
	}

	// Case 3: FastPath
	if !d.FastPath(evt2) {
		t.Error("FastPath should trigger for entropy 8.5")
	}
}
