/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package trace

import (
	"testing"
	"time"
	"github.com/fallofpheonix/phoenixmind-validator/truth/evidence"
)

func TestLineageGraph(t *testing.T) {
	lg := &LineageGraph{}
	lg.AddLink("A", "B", "DERIVED_FROM")
	lg.AddLink("B", "C", "TRANSFORMED_BY")

	links := lg.GetLineageForID("B")
	if len(links) != 2 {
		t.Errorf("Expected 2 links for B, got %d", len(links))
	}
}

func TestCausalGraph(t *testing.T) {
	cg := &CausalGraph{}
	cg.AddLink("Event1", "Event2", "CAUSED_BY", 0.8)
	cg.AddLink("Event2", "Event3", "PRECEDES", 0.5)

	effects := cg.GetEffectsOf("Event1")
	if len(effects) != 1 {
		t.Errorf("Expected 1 effect for Event1, got %d", len(effects))
	}
}

func TestTimeline(t *testing.T) {
	tl := &Timeline{}
	tl.AddEvent(TimelineEvent{ID: "TE1", EntityID: "E1", Timestamp: time.Now()})
	time.Sleep(1 * time.Millisecond) // Ensure distinct timestamps
	tl.AddEvent(TimelineEvent{ID: "TE2", EntityID: "E1", Timestamp: time.Now()})

	events := tl.GetEventsForEntity("E1")
	if len(events) != 2 {
		t.Errorf("Expected 2 events for E1, got %d", len(events))
	}
}

func TestHistory(t *testing.T) {
	h := &History{}
	h.AddEntry("EntityX", evidence.VALIDATED, nil)
	h.AddEntry("EntityY", evidence.WARNING, nil)

	entries := h.GetHistoryForEntity("EntityX")
	if len(entries) != 1 {
		t.Errorf("Expected 1 history entry for EntityX, got %d", len(entries))
	}
}

func TestCheckpointManager(t *testing.T) {
	cm := &CheckpointManager{}
	cp := cm.CreateCheckpoint("root-hash-1", map[string]evidence.TruthState{"E1": evidence.VALIDATED})

	retrievedCp, ok := cm.GetCheckpoint(cp.ID)
	if !ok || retrievedCp.ID != cp.ID {
		t.Errorf("Failed to retrieve checkpoint %s", cp.ID)
	}
}
