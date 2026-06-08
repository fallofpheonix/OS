/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */
package memory

import (
	"encoding/json"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

func TestMemory_TieredConcurrency(t *testing.T) {
	tm := NewTieredMemory()
	const iterations = 100
	const workers = 10

	var wg sync.WaitGroup
	wg.Add(workers * 3)

	for i := 0; i < workers; i++ {
		// Ingestion workers
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				tm.Ingest(&Fact{
					ID:         fmt.Sprintf("F-%d-%d", workerID, j),
					Confidence: ledger.ConfidenceScore{V: 900000},
				})
			}
		}(i)

		// Consolidation workers
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				tm.Consolidate(fmt.Sprintf("F-%d-%d", workerID, j), TierSemantic)
			}
		}(i)

		// Search workers
		go func(workerID int) {
			defer wg.Done()
			for j := 0; j < iterations; j++ {
				tm.Search(fmt.Sprintf("F-%d-%d", workerID, j))
			}
		}(i)
	}
	wg.Wait()
}

func TestMemory_TieredConsolidation(t *testing.T) {
	tm := NewTieredMemory()

	fact := &Fact{
		ID:         "FACT_CONSOLIDATE",
		Confidence: ledger.ConfidenceScore{V: 900000},
		Version:    1,
		State:      StateActive,
	}

	// 1. Ingest into Working memory
	tm.Ingest(fact)
	if _, ok := tm.Working["FACT_CONSOLIDATE"]; !ok {
		t.Errorf("fact not found in working memory")
	}

	// 2. Consolidate to Semantic
	tm.Consolidate("FACT_CONSOLIDATE", TierSemantic)
	if _, ok := tm.Working["FACT_CONSOLIDATE"]; ok {
		t.Errorf("fact still in working memory after consolidation")
	}
	if _, ok := tm.Semantic["FACT_CONSOLIDATE"]; !ok {
		t.Errorf("fact not found in semantic memory")
	}
}

func TestMemory_SQLitePersistence(t *testing.T) {
	tm := NewTieredMemory()
	store, err := NewVectorStore(":memory:") // In-memory DB for testing
	if err != nil {
		t.Fatalf("failed to create vector store: %v", err)
	}
	defer store.Close()
	tm.Store = store

	fact := &Fact{
		ID:         "FACT_PERSIST",
		Confidence: ledger.ConfidenceScore{V: 850000},
		Version:    1,
		State:      StateActive,
		Timestamp:  1000,
		Data:       []byte("test-data"),
	}

	tm.Ingest(fact)
	tm.Consolidate("FACT_PERSIST", TierSemantic)

	// Wait for async persistence
	time.Sleep(100 * time.Millisecond)

	// Verify persistence in SQLite
	persisted, err := store.SearchSemantic(nil)
	if err != nil {
		t.Fatalf("failed to search semantic memory: %v", err)
	}

	if persisted == nil {
		t.Errorf("fact was not persisted to SQLite")
	} else if persisted.ID != "FACT_PERSIST" {
		t.Errorf("persisted fact ID mismatch: expected FACT_PERSIST, got %s", persisted.ID)
	}
}

func TestMemory_Decay(t *testing.T) {
	tm := NewTieredMemory()

	tm.Ingest(&Fact{ID: "HIGH", Confidence: ledger.ConfidenceScore{V: 900000}})
	tm.Ingest(&Fact{ID: "LOW", Confidence: ledger.ConfidenceScore{V: 200000}})

	tm.Decay(0.5)

	if _, ok := tm.Working["HIGH"]; !ok {
		t.Errorf("high confidence fact incorrectly decayed")
	}
	if _, ok := tm.Working["LOW"]; ok {
		t.Errorf("low confidence fact failed to decay")
	}
}

func TestMemory_Recall(t *testing.T) {
	mem := NewMemory()
	f1 := &Fact{ID: "F1", Version: 1, State: StateActive}
	mem.Store(f1)

	f2 := &Fact{ID: "F1", Version: 2, State: StateActive}
	mem.Store(f2)

	latest, ok := mem.Recall("F1")
	if !ok || latest.Version != 2 {
		t.Errorf("expected version 2, got %d", latest.Version)
	}

	v1, ok := mem.RecallVersion("F1", 1)
	if !ok || v1.Version != 1 {
		t.Error("failed to recall version 1")
	}

	if v1.State != StateSuperseded {
		t.Errorf("expected version 1 to be superseded, got %s", v1.State)
	}
}

func TestMemory_Recall_EdgeCases(t *testing.T) {
	mem := NewMemory()

	if _, ok := mem.Recall("MISSING"); ok {
		t.Error("expected false for missing recall")
	}

	if _, ok := mem.RecallVersion("MISSING", 1); ok {
		t.Error("expected false for missing recall version")
	}

	mem.Store(&Fact{ID: "F1", Version: 1})
	if _, ok := mem.RecallVersion("F1", 2); ok {
		t.Error("expected false for wrong version")
	}
}

func TestTieredMemory_Search_EdgeCases(t *testing.T) {
	tm := NewTieredMemory()

	if _, ok := tm.Search("MISSING"); ok {
		t.Error("expected false for missing search")
	}

	// Found in Episodic
	tm.Episodic = append(tm.Episodic, &Fact{ID: "E1"})
	if _, ok := tm.Search("E1"); !ok {
		t.Error("failed to find fact in episodic")
	}

	// Consolidate missing
	tm.Consolidate("MISSING", TierSemantic) // Should just return, no panic
}

func TestMemory_Reconstruction_EdgeCases(t *testing.T) {
	mem := NewMemory()
	badEvents := []*ledger.Event{{Type: ledger.EventFact, Payload: []byte("{bad-json")}}
	if err := mem.ReconstructFromLedger(badEvents); err == nil {
		t.Error("expected error on bad json")
	}
}

func TestMemory_Reconstruction(t *testing.T) {
	mem := NewMemory()

	// Simulate a Fact event from the ledger.
	factPayload := ledger.FactPayload{
		FactID:          "FACT_001",
		ConfidenceScore: ledger.ConfidenceScore{V: 950000},
		Timestamp:       123456789,
	}
	payload, _ := json.Marshal(factPayload)

	// Simulate a FactUpdate event.
	updatePayload := ledger.FactUpdatePayload{
		FactID:          "FACT_001",
		ConfidenceScore: ledger.ConfidenceScore{V: 990000},
		Reason:          "Multi-sensor confirmation",
	}
	updatePayloadBytes, _ := json.Marshal(updatePayload)

	events := []*ledger.Event{
		{
			Type:    ledger.EventFact,
			Payload: payload,
		},
		{
			Type:    ledger.EventFactUpdate,
			Payload: updatePayloadBytes,
		},
	}

	if err := mem.ReconstructFromLedger(events); err != nil {
		t.Fatalf("reconstruction failed: %v", err)
	}

	f, ok := mem.Recall("FACT_001")
	if !ok {
		t.Errorf("fact FACT_001 not found in memory after reconstruction")
	}

	if phxmath.FixedPoint(f.Confidence).Float64() != 0.99 {
		t.Errorf("fact confidence mismatch: expected 0.99, got %v", f.Confidence)
	}

	if f.Version != 2 {
		t.Errorf("version mismatch: expected 2, got %d", f.Version)
	}

	// Verify version 1 exists and is superseded
	f1, ok := mem.RecallVersion("FACT_001", 1)
	if !ok {
		t.Errorf("version 1 not found")
	}
	if f1.State != StateSuperseded {
		t.Errorf("state mismatch for version 1: expected SUPERSEDED, got %s", f1.State)
	}
}
