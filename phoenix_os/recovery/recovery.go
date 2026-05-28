package recovery

import (
	"log"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
)

type Orchestrator interface{}

type RecoveryLoop struct {
	Bus  *bus.Bus
	Orch Orchestrator
}

func NewRecoveryLoop(b *bus.Bus, orch Orchestrator) *RecoveryLoop {
	return &RecoveryLoop{Bus: b, Orch: orch}
}

func (rl *RecoveryLoop) Start() {
	log.Printf("[Recovery] Loop active")
}
