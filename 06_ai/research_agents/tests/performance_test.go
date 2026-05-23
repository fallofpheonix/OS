package tests

import (
	"os"
	"testing"
	"time"

	"phoenix/agents/internal/control"
	"phoenix/agents/internal/game"
	"phoenix/agents/internal/graph"
	"phoenix/agents/internal/kernel"
	"phoenix/agents/internal/physics"
	"phoenix/agents/internal/types"
)

func BenchmarkEndToEndLoop(b *testing.B) {
	// Setup
	forensicsDir := "./bench_forensics"
	os.MkdirAll(forensicsDir, 0755)
	defer os.RemoveAll(forensicsDir)

	graphAgent := graph.NewGraphAgent()
	physicsAgent := physics.NewPhysicsAgent()
	gameAgent := game.NewGameAgent()
	controlAgent := control.NewControlAgent(1.5, 0.2, 0.1, 2.0)
	kernelAgent := kernel.NewKernelAgent()
	kernelAgent.Unlock()

	now := time.Now()
	event := types.TelemetryEvent{
		Timestamp: now,
		EventID:   "bench-evt",
		Category:  "filesystem",
		EventType: "file.write",
		PID:       8888,
		PPID:      1,
		Comm:      "bench",
		ExePath:   "/usr/bin/bench",
		Filesystem: &types.FilesystemPayload{
			FilePath:       "/tmp/bench",
			BytesRequested: 1024,
		},
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Single loop iteration
		_ = graphAgent.UpdateGraph(event)
		dag := graphAgent.GetAttackDAG()
		state, _ := physicsAgent.GetSecurityState(dag, now)
		gameAgent.UpdateBeliefs(state, "high_entropy_write")
		strategy, _ := gameAgent.SolveBestStrategy(state, dag, now)
		_ = controlAgent.EnforceStrategy(strategy, state.ThreatTemperature, now)
		_ = kernelAgent.ActuateContainmentPolicy(8888, "KILL")
	}
}
