package trace

import (
	"database/sql"
	"fmt"

	"github.com/fallofpheonix/phoenix-os/phoenix-core/telemetry/bus"
	_ "github.com/mattn/go-sqlite3"
)

type TraceStorage struct {
	db    *sql.DB
	busCh chan bus.TelemetryEvent
}

func NewTraceStorage(dbPath string, busCh chan bus.TelemetryEvent) (*TraceStorage, error) {
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}

	pragmas := []string{
		"PRAGMA journal_mode=WAL;",
		"PRAGMA synchronous=NORMAL;",
		"PRAGMA temp_store=MEMORY;",
		"PRAGMA mmap_size=268435456;",
		"PRAGMA cache_size=-200000;",
	}
	for _, p := range pragmas {
		if _, err := db.Exec(p); err != nil {
			return nil, fmt.Errorf("failed to set %s: %v", p, err)
		}
	}

	schema := `CREATE TABLE IF NOT EXISTS events (
		seq_id INTEGER PRIMARY KEY,
		monotonic_ns INTEGER NOT NULL,
		wall_timestamp INTEGER NOT NULL,
		source TEXT NOT NULL,
		pid INTEGER,
		event_type TEXT NOT NULL,
		payload BLOB NOT NULL,
		prev_hash TEXT NOT NULL,
		hash TEXT NOT NULL
	);`
	if _, err := db.Exec(schema); err != nil {
		return nil, fmt.Errorf("failed to create schema: %v", err)
	}

	return &TraceStorage{db: db, busCh: busCh}, nil
}

func (t *TraceStorage) StartWriter() {
	go func() {
		for event := range t.busCh {
			t.Write(event)
		}
	}()
}

func (t *TraceStorage) Write(event bus.TelemetryEvent) error {
	stmt, err := t.db.Prepare(`INSERT INTO events (seq_id, monotonic_ns, wall_timestamp, source, pid, event_type, payload, prev_hash, hash)
	                           VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		event.SeqID, event.MonotonicNs, event.WallTimeUnix,
		event.Source, event.PID, event.EventType,
		event.Payload, event.PrevHash, event.Hash,
	)
	return err
}

func (t *TraceStorage) Close() error {
	return t.db.Close()
}
