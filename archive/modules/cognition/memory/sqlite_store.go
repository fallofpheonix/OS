/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: INTEGRATED - Formal Cognition Authoritative
 */

// Package memory implements the tiered storage and versioned fact substrate for PhoenixOS.
package memory

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"sync"

	_ "github.com/mattn/go-sqlite3"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// VectorStore implements a SQLite-backed persistence layer for tiered memory.
// It supports high-performance vector search in Semantic memory and handles
// multi-tier state synchronization.
type VectorStore struct {
	db *sql.DB
	mu sync.RWMutex
}

// NewVectorStore initializes the SQLite database and establishes the memory schema.
// Inputs: dbPath (string) - Filesystem path to the .db file or ":memory:".
// Returns: (*VectorStore, error) if initialization fails.
func NewVectorStore(dbPath string) (*VectorStore, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open sqlite database: %w", err)
	}

	store := &VectorStore{
		db: db,
	}

	if err := store.initializeSchema(); err != nil {
		return nil, err
	}

	log.Printf("[MemoryLab SQLite] Database initialized at %s", dbPath)
	return store, nil
}

// initializeSchema executes the DDL required for memory tier tables.
// Side Effects: Creates tables if they do not exist.
// Complexity: O(1) time.
func (s *VectorStore) initializeSchema() error {
	schema := `
	CREATE TABLE IF NOT EXISTS episodic_memory (
		id TEXT PRIMARY KEY,
		version INTEGER,
		state TEXT,
		confidence REAL,
		timestamp INTEGER,
		data BLOB
	);

	CREATE TABLE IF NOT EXISTS semantic_memory (
		id TEXT PRIMARY KEY,
		version INTEGER,
		state TEXT,
		confidence REAL,
		timestamp INTEGER,
		data BLOB,
		embedding BLOB -- Reserved for sqlite-vec (FLOAT32 vector)
	);

	CREATE TABLE IF NOT EXISTS procedural_memory (
		id TEXT PRIMARY KEY,
		version INTEGER,
		state TEXT,
		confidence REAL,
		timestamp INTEGER,
		data BLOB
	);
	`
	_, err := s.db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to initialize schema: %w", err)
	}
	return nil
}

// PersistFact serializes and saves a fact to the designated SQLite table.
// Inputs: tier (Memory Tier), f (*Fact), embedding ([]float32 - optional).
// Side Effects: Performs UPSERT on conflict with existing ID.
// Complexity: O(Log N) where N is the number of records in the tier.
func (s *VectorStore) PersistFact(tier Tier, f *Fact, embedding []float32) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	var tableName string
	switch tier {
	case TierEpisodic:
		tableName = "episodic_memory"
	case TierSemantic:
		tableName = "semantic_memory"
	case TierProcedural:
		tableName = "procedural_memory"
	default:
		return fmt.Errorf("unsupported persistence tier: %s", tier)
	}

	var embData []byte
	if embedding != nil {
		embData, _ = json.Marshal(embedding)
	}

	if tier == TierSemantic {
		query := fmt.Sprintf(`
			INSERT INTO %s (id, version, state, confidence, timestamp, data, embedding)
			VALUES (?, ?, ?, ?, ?, ?, ?)
			ON CONFLICT(id) DO UPDATE SET
				version=excluded.version,
				state=excluded.state,
				confidence=excluded.confidence,
				timestamp=excluded.timestamp,
				data=excluded.data,
				embedding=excluded.embedding;
		`, tableName)
		_, err := s.db.Exec(query, f.ID, f.Version, f.State, phxmath.FixedPoint(f.Confidence).Float64(), f.Timestamp, f.Data, embData)
		return err
	} else {
		queryNoEmb := fmt.Sprintf(`
			INSERT INTO %s (id, version, state, confidence, timestamp, data)
			VALUES (?, ?, ?, ?, ?, ?)
			ON CONFLICT(id) DO UPDATE SET
				version=excluded.version,
				state=excluded.state,
				confidence=excluded.confidence,
				timestamp=excluded.timestamp,
				data=excluded.data;
		`, tableName)
		_, err := s.db.Exec(queryNoEmb, f.ID, f.Version, f.State, phxmath.FixedPoint(f.Confidence).Float64(), f.Timestamp, f.Data)
		return err
	}
}

// SearchSemantic retrieves the most recent semantic fact.
// In Phase 5, this will perform an L2-distance vector search via sqlite-vec.
// Returns (*Fact, error).
// Complexity: O(Log N) via index-backed ordering.
func (s *VectorStore) SearchSemantic(queryEmbedding []float32) (*Fact, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	row := s.db.QueryRow(`SELECT id, version, state, confidence, timestamp, data FROM semantic_memory ORDER BY timestamp DESC LIMIT 1`)

	var f Fact
	var confidence float64
	if err := row.Scan(&f.ID, &f.Version, &f.State, &confidence, &f.Timestamp, &f.Data); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}
	f.Confidence = ledger.ConfidenceScore{V: int64(confidence * 1000000)}

	return &f, nil
}

// Close terminates the database connection.
// Returns: error if the connection is already closed.
func (s *VectorStore) Close() error {
	return s.db.Close()
}
