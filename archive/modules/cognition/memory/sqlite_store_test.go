package memory

import (
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
)

func TestVectorStore_CRUD(t *testing.T) {
	store, err := NewVectorStore(":memory:")
	if err != nil {
		t.Fatalf("failed to create vector store: %v", err)
	}
	defer store.Close()

	fact := &Fact{
		ID:         "F1",
		Version:    1,
		State:      StateActive,
		Confidence: ledger.ConfidenceScore{V: 900000},
		Timestamp:  12345,
		Data:       []byte("raw-data"),
	}

	// Test Persist
	err = store.PersistFact(TierSemantic, fact, []float32{0.1, 0.2, 0.3})
	if err != nil {
		t.Fatalf("failed to persist fact: %v", err)
	}

	// Test Search
	retrieved, err := store.SearchSemantic(nil)
	if err != nil {
		t.Fatalf("failed to search semantic memory: %v", err)
	}

	if retrieved == nil || retrieved.ID != "F1" {
		t.Errorf("expected F1, got %v", retrieved)
	}

	// Test Update via Persist (ON CONFLICT)
	fact.Confidence = ledger.ConfidenceScore{V: 950000}
	err = store.PersistFact(TierSemantic, fact, nil)
	if err != nil {
		t.Fatalf("failed to update fact: %v", err)
	}

	retrieved, _ = store.SearchSemantic(nil)
	if phxmath.FixedPoint(retrieved.Confidence).Float64() != 0.95 {
		t.Errorf("expected updated confidence 0.95, got %v", retrieved.Confidence)
	}
}

func TestVectorStore_PersistOtherTiers(t *testing.T) {
	store, _ := NewVectorStore(":memory:")
	defer store.Close()

	f := &Fact{ID: "F2", Version: 1, Data: []byte("test")}
	
	if err := store.PersistFact(TierEpisodic, f, nil); err != nil {
		t.Errorf("failed to persist episodic: %v", err)
	}
	
	if err := store.PersistFact(TierProcedural, f, nil); err != nil {
		t.Errorf("failed to persist procedural: %v", err)
	}
}

func TestVectorStore_InitErrors(t *testing.T) {
	// Attempt to open a database in a non-existent, uncreatable directory to trigger sql.Open or Exec error
	_, err := NewVectorStore("/root/invalid_dir/db.sqlite")
	if err == nil {
		t.Error("expected error when initializing store in bad path")
	}
}

func TestVectorStore_SearchEmpty(t *testing.T) {
	store, _ := NewVectorStore(":memory:")
	defer store.Close()

	// Search empty db
	retrieved, err := store.SearchSemantic(nil)
	if err != nil {
		t.Errorf("expected no error for empty search, got %v", err)
	}
	if retrieved != nil {
		t.Errorf("expected nil result for empty search, got %v", retrieved)
	}
}

func TestVectorStore_UnsupportedTier(t *testing.T) {
	store, _ := NewVectorStore(":memory:")
	defer store.Close()

	err := store.PersistFact("INVALID_TIER", &Fact{ID: "F1"}, nil)
	if err == nil {
		t.Error("expected error for unsupported tier, got nil")
	}
}
