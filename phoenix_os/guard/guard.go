package guard

import (
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

// TODO: Implement Fast-Path (<100ms) enforcement logic for high-entropy signals.
// FIXME: Direct SDI mapping is prohibited; must integrate with Warden FSM state.

const (
	ModeSaturation = "saturation"
)

type GuardAdapter struct {
	Bus *bus.Bus
}

func NewGuardAdapter(b *bus.Bus, file string, mode string, scale float64, seed int) *GuardAdapter {
	return &GuardAdapter{Bus: b}
}

func (g *GuardAdapter) FetchEvents() ([]bus.TelemetryEvent, error) {
	return []bus.TelemetryEvent{}, nil
}

func (g *GuardAdapter) GetSequenceHash(events []bus.TelemetryEvent) string {
	return "mock-hash"
}
