package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"phoenix/agents/internal/control"
	"phoenix/agents/internal/forensics"
	"phoenix/agents/internal/game"
	"phoenix/agents/internal/graph"
	"phoenix/agents/internal/kernel"
	"phoenix/agents/internal/physics"
	"phoenix/agents/internal/telemetry"
)

func main() {
	fmt.Println("Starting PhoenixOS Internal Agents...")

	// Initialize Agents
	telemetryAgent := telemetry.NewTelemetryAgent(1000)
	graphAgent := graph.NewGraphAgent()
	physicsAgent := physics.NewPhysicsAgent()
	gameAgent := game.NewGameAgent()
	controlAgent := control.NewControlAgent(1.5, 0.2, 0.1, 2.0)
	
	forensicsDir := "./artifacts/forensics"
	forensicsAgent, err := forensics.NewForensicsAgent(forensicsDir)
	if err != nil {
		log.Fatalf("Failed to initialize forensics agent: %v", err)
	}
	_ = forensicsAgent // Silence unused warning if not used in loop yet
	fmt.Printf("Forensics Agent initialized (Storage: %s)\n", forensicsDir)
	
	kernelAgent := kernel.NewKernelAgent()

	// Start Telemetry
	if err := telemetryAgent.Start(); err != nil {
		log.Fatalf("Failed to start telemetry agent: %v", err)
	}

	fmt.Println("Agents initialized and running.")
	fmt.Printf("Kernel Agent Status: Locked=%v\n", kernelAgent.IsLocked())

	// Loop simulation (mocking the event loop)
	go func() {
		for {
			// 1. Mock event
			ev := telemetryAgent.GenerateMockEvent()
			
			// 2. Update Graph
			graphAgent.UpdateGraph(ev)
			
			// 3. Get Security State
			dag := graphAgent.GetAttackDAG()
			state, _ := physicsAgent.GetSecurityState(dag)
			
			// 4. Update Game Beliefs
			if ev.Category == "filesystem" && ev.Filesystem != nil && ev.Filesystem.BytesRequested > 100 {
				gameAgent.UpdateBeliefs(state, "high_entropy_write")
			}
			
			// 5. Solve Strategy
			strategy, _ := gameAgent.SolveBestStrategy(state, dag)
			
			// 6. Enforce Control
			controlAgent.EnforceStrategy(strategy, state.ThreatTemperature)

			time.Sleep(1 * time.Second)
		}
	}()

	// Wait for termination
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT, syscall.SIGTERM)
	<-stop

	fmt.Println("Shutting down agents...")
	telemetryAgent.Stop()
}
