package monitor

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

type DriftScore struct {
	EventType string
	Severity  float64
	PID       int
	UID       int
	ZScore    float64
	EventID   int64
}

type MonitorService struct {
	Bus *bus.Bus
}

func NewMonitorService(i interface{}, b *bus.Bus) *MonitorService {
	return &MonitorService{Bus: b}
}

func (m *MonitorService) Process(e bus.TelemetryEvent) DriftScore {
	return DriftScore{
		EventType: e.EventType, 
		Severity: e.Severity, 
		EventID: e.SeqID,
		ZScore: e.Severity * 2.0, // Simulation: multiply severity by 2.0
	}
}
