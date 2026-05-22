package warden

import (
	"github.com/fallofpheonix/phoenix-os/telemetry/bus"
)

type Warden struct {
	outBus *bus.Bus
}

func NewWarden(outBus *bus.Bus) *Warden {
	return &Warden{outBus: outBus}
}
