package replay

import (
	"database/sql"
	"github.com/fallofpheonix/phoenix-os/telemetry/bus"
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
	return &TraceStorage{db: db, busCh: busCh}, nil
}

func (t *TraceStorage) Write(event bus.TelemetryEvent) error {
	return nil
}

func (t *TraceStorage) Close() error {
	return t.db.Close()
}
