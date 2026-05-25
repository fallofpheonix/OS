package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/control"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/forensics"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/game"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/graph"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/kernel"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/physics"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/telemetry"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/types"
)

func main() {
	replayPath := flag.String("replay", "", "Path to telemetry replay file")
	flag.Parse()

	if *replayPath != "" {
		runReplay(*replayPath)
		return
	}

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
			now := time.Now()
			// 1. Mock event
			ev := telemetryAgent.GenerateMockEvent()
			
			// 2. Update Graph
			graphAgent.UpdateGraph(ev)
			
			// 3. Get Security State
			dag := graphAgent.GetAttackDAG()
			state, _ := physicsAgent.GetSecurityState(dag, now)
			
			// 4. Update Game Beliefs
			evidence := ""
			if ev.Category == "filesystem" && ev.Filesystem != nil && ev.Filesystem.BytesRequested > 100 {
				evidence = "high_entropy_write"
			}
			gameAgent.UpdateBeliefs(state, evidence)
			
			// 5. Solve Strategy
			strategy, _ := gameAgent.SolveBestStrategy(state, dag, now)
			
			// 6. Enforce Control
			controlAgent.EnforceStrategy(strategy, state.ThreatTemperature, now)

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

func runReplay(path string) {
	fmt.Fprintf(os.Stderr, "Running deterministic replay from %s\n", path)
	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatalf("Failed to read replay file: %v", err)
	}

	var events []types.TelemetryEvent
	if err := json.Unmarshal(data, &events); err != nil {
		log.Fatalf("Failed to unmarshal replay events: %v", err)
	}

	// Initialize Agents
	graphAgent := graph.NewGraphAgent()
	physicsAgent := physics.NewPhysicsAgent()
	gameAgent := game.NewGameAgent()
	controlAgent := control.NewControlAgent(1.5, 0.2, 0.1, 2.0)

	type ReplayOutput struct {
		EventID       string               `json:"event_id"`
		State         types.SecurityState  `json:"state"`
		Strategy      types.Strategy       `json:"strategy"`
		PIDMetrics    types.PIDMetrics     `json:"pid_metrics"`
		ActionHistory []string             `json:"action_history"`
	}

	var results []ReplayOutput

	for _, ev := range events {
		graphAgent.UpdateGraph(ev)
		dag := graphAgent.GetAttackDAG()
		state, _ := physicsAgent.GetSecurityState(dag, ev.Timestamp)
		
		evidence := ""
		if ev.Category == "filesystem" && ev.Filesystem != nil && ev.Filesystem.BytesRequested > 100 {
			evidence = "high_entropy_write"
		}
		gameAgent.UpdateBeliefs(state, evidence)
		strategy, _ := gameAgent.SolveBestStrategy(state, dag, ev.Timestamp)
		controlAgent.EnforceStrategy(strategy, state.ThreatTemperature, ev.Timestamp)

		results = append(results, ReplayOutput{
			EventID:       ev.EventID,
			State:         state,
			Strategy:      strategy,
			PIDMetrics:    controlAgent.GetPIDMetrics(),
			ActionHistory: controlAgent.GetActionHistory(),
		})
	}

	output, _ := json.MarshalIndent(results, "", "  ")
	fmt.Println(string(output))
}
