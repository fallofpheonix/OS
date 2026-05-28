package trace

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

type TraceStorage struct{}

func NewTraceStorage(path string, i interface{}) (*TraceStorage, error) {
	return &TraceStorage{}, nil
}

func (t *TraceStorage) Write(e bus.TelemetryEvent) error { return nil }
func (t *TraceStorage) Close() error { return nil }
