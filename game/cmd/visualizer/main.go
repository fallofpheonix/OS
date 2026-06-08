package main

import (
	"log"
	"time"

	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/game/engine"
	"github.com/fallofpheonix/phoenix/game/multiplayer"
	"github.com/fallofpheonix/phoenix/game/pscript"
)

func main() {
	// 1. Start WebSocket Server
	wsServer := multiplayer.NewWebSocketServer()
	go wsServer.Start(":8080")

	// 2. Prepare Script
	script := `
		let target = 100
		move("agent_01", target)
		verify("agent_01")
	`

	// 3. Parse Script to Instructions
	l := pscript.NewLexer(script)
	p := pscript.NewParser(l)
	instructions, err := p.Parse()
	if err != nil {
		log.Fatalf("Parse error: %v", err)
	}

	// 4. Initialize VM
	v := engine.NewVM(instructions)
	v.State.GetEntities()["agent_01"] = &engine.Entity{
		ID:     "agent_01",
		Pos:    phxmath.NewFixedPoint(0),
		Status: "IDLE",
	}

	log.Println("Simulation started. Connect Godot to ws://localhost:8080/ws")

	// 5. Simulation Loop
	for v.PC < len(instructions) {
		err := v.Step()
		if err != nil {
			log.Fatalf("VM error: %v", err)
		}

		// Broadcast state
		wsServer.BroadcastWorldState(v.State)

		log.Printf("Tick %d: Agent pos %v, status %s", v.State.Tick, v.State.GetEntities()["agent_01"].Pos.Float64(), v.State.GetEntities()["agent_01"].Status)

		time.Sleep(1 * time.Second) // Slow down for visualization
	}

	log.Println("Simulation finished.")
	// Keep server alive for a bit so client can see final state
	time.Sleep(10 * time.Second)
}
