package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	"phoenix/ai"
	"phoenix/arbiter"
	"phoenix/boot"
	"phoenix/bus"
	"phoenix/common/logical_clock"
	"phoenix/common/resource"
	"phoenix/guard"
	"phoenix/ledger/src"
	"phoenix/game"
	"phoenix/monitor"
	"phoenix/tcs"
	"phoenix/trace"
	"phoenix/warden"
)

func main() {
	fmt.Println("==================================================")
	fmt.Println("✦ PHOENIX CYBERNETIC SECURITY RUNTIME (STAGE 1)")
	fmt.Println("==================================================")

	// ── 1. Initialize Subsystems (Synchronous Mode) ─────────────
	b := bus.NewBus()
	
	// Deterministic Resource Bounding (1GB Limit, 1M Entries)
	alloc := resource.NewBoundedAllocator(1024*1024*1024, 1000000)
	evLedger := ledger.NewLedger(alloc)
	
	logicalClock := logical_clock.NewClock()
	telemetryWindow := tcs.NewSlidingWindow(60 * time.Second)
	degMon := tcs.NewDegradationMonitor(telemetryWindow, nil)

	dbPath := "/tmp/phoenix_trace.db"
	_ = os.Remove(dbPath)
	traceStore, err := trace.NewTraceStorage(dbPath, nil)
	if err != nil {
		log.Fatalf("[FATAL] Trace storage init failed: %v", err)
	}
	defer traceStore.Close()

	// Register the Bus Overflow callback to log snapshots
	b.OnOverflow = func(topic string, pressure float64, event bus.TelemetryEvent) {
		log.Printf("[BUS OVERFLOW] Triggered snapshot on %s. Pressure: %.2f", topic, pressure)
		snapshotPayload := []byte(fmt.Sprintf(`{"topic":"%s","pressure":%f,"trigger_event_id":"%s"}`, topic, pressure, event.EventID))
		
		overflowEvent := bus.TelemetryEvent{
			SeqID:        -event.SeqID, // negative sequence ID to designate overflow anomaly
			MonotonicNs:  event.MonotonicNs,
			WallTimeUnix: event.WallTimeUnix,
			Source:       "phoenix.bus",
			EventType:    "system.overflow_snapshot",
			Severity:     1.0,
			Payload:      snapshotPayload,
		}
		_ = traceStore.Write(overflowEvent)
		_ = evLedger.AddEntry("OVERFLOW-SNAPSHOT", event.EventID, snapshotPayload)
	}

	mon := monitor.NewMonitorService(nil, b)
	arb := arbiter.NewArbiter(b)
	w := warden.NewWarden(b)

	// Instantiate the AI Orchestrator as the primary driver
	aiOrch := ai.NewAIOrchestrator()
	aiOrch.RegisterFeature(&ai.LedgerFeature{Ledger: evLedger})
	aiOrch.RegisterFeature(&ai.TraceFeature{Store: traceStore})
	aiOrch.RegisterFeature(&ai.MonitorFeature{Service: mon})
	aiOrch.RegisterFeature(&ai.TCSFeature{Window: telemetryWindow, DegMon: degMon})
	aiOrch.RegisterFeature(&ai.ArbiterFeature{Arb: arb})
	aiOrch.RegisterFeature(&ai.WardenFeature{Warden: w})

	// ── 1.5 Capture Boot Telemetry (Genesis) ────────────────────
	bootInfo := []boot.SubsystemInfo{
		boot.NewSubsystemInfo("Bus", "1.0.0", map[string]interface{}{"capacity": bus.QueueCapacity}),
		boot.NewSubsystemInfo("Ledger", "2.0.0", map[string]interface{}{"alloc_limit": 1024 * 1024 * 1024}),
		boot.NewSubsystemInfo("Warden", "1.0.0", nil),
		boot.NewSubsystemInfo("Arbiter", "1.0.0", nil),
		boot.NewSubsystemInfo("AIOrchestrator", "1.0.0", nil),
	}
	bt, err := boot.CaptureBootTelemetry(evLedger, bootInfo)
	if err != nil {
		log.Printf("[BOOT] Telemetry capture failed: %v", err)
	} else {
		fmt.Printf("[BOOT] Genesis Checksum: %s\n", bt.Checksum)
		
		// ── 1.6 Verify Boot Integrity (Optional) ──────────────────
		expectedChecksum := os.Getenv("PHOENIX_BOOT_EXPECTED")
		if expectedChecksum != "" {
			if err := boot.VerifyBoot(bt.Checksum, expectedChecksum); err != nil {
				log.Fatalf("[FATAL] Boot Integrity Violation: %v", err)
			}
			fmt.Println("[BOOT] Integrity Verified")
		}
	}

	// ── 2. Guard Replay Adapter (Seeded for Determinism) ─────────
	eventsFile := "/Users/fallofpheonix/os/04_datasets/logs/test_events.jsonl"
	
	// Initialize and start the Game Server for SOC Simulation
	scoreState := game.NewScoreState()
	gameServer := game.NewGameServer(scoreState, w, evLedger, eventsFile)
	gameServer.Start(":8080")

	const replaySeed = 42
	fmt.Printf("[BOOT] Starting Guard Adapter (Seed: %d)\n", replaySeed)
	guardAdapter := guard.NewGuardAdapter(b, eventsFile, guard.ModeSaturation, 1.0, replaySeed)

	events, err := guardAdapter.FetchEvents()
	if err != nil {
		log.Fatalf("[FATAL] Failed to fetch events: %v", err)
	}
	fmt.Printf("[REPLAY] Loaded %d events for deterministic execution\n", len(events))
	
	seqHash := guardAdapter.GetSequenceHash(events)
	fmt.Printf("[REPLAY] Sequence Proof Hash: %s\n", seqHash)
	evLedger.AddEntry("REPLAY-SEQUENCE-PROOF", "GUARD", []byte(seqHash))

	// ── 3. Deterministic Main Loop (Logical Tick) ────────────────
	for _, event := range events {
		// Logical Tick Increment
		event.LogicalTick = logicalClock.Tick()

		// Orchestrate the tick using the AI orchestrator (coordinates trace, monitor, tcs, arbiter, warden)
		aiOrch.OrchestrateTick(event, event.LogicalTick)

		// Tick Step 7: Automatic Checkpointing (Every 1000 ticks)
		if event.LogicalTick % 1000 == 0 {
			checkpoint, _ := evLedger.Checkpoint()
			log.Printf("[REPLAY] Checkpoint at tick %d: %x", event.LogicalTick, checkpoint)
		}
	}

	// ── 4. Verification ───────────────────────────────────────────
	fmt.Println("\n[REPLAY] Execution Complete. Verifying Invariants...")

	if verr := evLedger.Verify(); verr != nil {
		fmt.Printf("[LEDGER] INTEGRITY VIOLATION: %v\n", verr)
	} else {
		fmt.Printf("[LEDGER] Merkle-DAG verified (%d nodes)\n", len(evLedger.Entries))
	}

	finalTCS := telemetryWindow.Evaluate()
	fmt.Printf("[TCS] Final Telemetry Confidence Score: %.4f\n", finalTCS)

	// Authority Hash (Root Hash of the DAG heads)
	outputHash := sha256.New()
	for _, head := range evLedger.Heads {
		outputHash.Write(head)
	}
	fmt.Printf("[REPLAY] Authoritative Output Hash: %x\n", outputHash.Sum(nil))

	fmt.Println("==================================================")
	fmt.Println("✦ RUNTIME HALTED")
	fmt.Println("==================================================")

	// Keep game server alive if running in game mode
	if os.Getenv("GAME_MODE") == "true" {
		fmt.Println("\n[GAME] Server running at http://localhost:8080. Press Ctrl+C to terminate.")
		select {}
	}
}
