package main

import (
	"context"
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/arbiter"
	"github.com/fallofpheonix/PheonixGuard"
	"github.com/fallofpheonix/PheonixGuard/actuators"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/monitor"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/tcs"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/trace"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/guard"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/boot"
	kernel "github.com/fallofpheonix/PheonixKernel"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"github.com/fallofpheonix/PheonixDistributed/discovery"
	"github.com/fallofpheonix/PheonixDistributed/ledger"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/common/resource"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/containment/rollback"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/game"
	localLedger "github.com/fallofpheonix/PheonixTruth/src"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/recovery"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/telemetry/process_graphs"
)

type KernelBusAdapter struct {
	Bus *bus.Bus
}

func (a *KernelBusAdapter) Publish(topic string, event kernel.TelemetryEvent) {
	bEvent := bus.TelemetryEvent{
		EventID:   event.EventID,
		SeqID:     event.SeqID,
		Source:    event.Source,
		EventType: event.EventType,
		Severity:  event.Severity,
		Payload:   event.Payload,
	}
	a.Bus.Publish(topic, bEvent)
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("✦ PHOENIX CYBERNETIC SECURITY RUNTIME (STAGE 1)")
	fmt.Println("==================================================")

	// ── 1. Initialize Subsystems (Synchronous Mode) ─────────────
	b := bus.NewBus()

	// Deterministic Resource Bounding (1GB Limit, 1M Entries)
	alloc := resource.NewBoundedAllocator(1024*1024*1024, 1000000)
	evLedger := localLedger.NewLedger(alloc)

	telemetryWindow := tcs.NewSlidingWindow(60 * time.Second)
	degMon := tcs.NewDegradationMonitor(telemetryWindow, nil)

	dbPath := "/tmp/phoenix_trace.db"
	_ = os.Remove(dbPath)
	traceStore, err := trace.NewTraceStorage(dbPath, nil)
	if err != nil {
		log.Fatalf("[FATAL] Trace storage init failed: %v", err)
	}
	defer traceStore.Close()

	// F1: Causal Graph & Recovery
	graph := process_graphs.NewGraph()
	orch := &rollback.Orchestrator{} // Simplified for F1
	recoveryLoop := recovery.NewRecoveryLoop(b, orch)
	recoveryLoop.Start()

	// Register the Bus Overflow callback to log snapshots
	b.OnOverflow = func(topic string, pressure float64, event bus.TelemetryEvent) {
		log.Printf("[BUS OVERFLOW] Triggered snapshot on %s. Pressure: %.2f", topic, pressure)
		snapshotPayload := []byte(fmt.Sprintf(`{"topic":"%s","pressure":%f,"trigger_event_id":"%s"}`, topic, pressure, event.EventID))

		overflowEvent := bus.TelemetryEvent{
			SeqID:        -event.SeqID, // negative sequence ID to designate overflow anomaly
			SourceEpoch:  event.SourceEpoch,
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
	// Instantiate the AI Orchestrator as the primary driver
	aiOrch := ai.NewAIOrchestrator()

	// ── 1.2 Initialize Distributed Coordination (Stage D) ─────
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	nodeIdentity := fmt.Sprintf("node-%x", sha256.Sum256([]byte(os.Getenv("HOSTNAME"))))
	beacon := discovery.NewBeaconTransport(9999, nodeIdentity)
	cons := ledger.NewStubConsensusLedger()

	if err := aiOrch.StartNetworking(ctx, beacon, cons); err != nil {
		log.Printf("[WARNING] Distributed networking failed: %v", err)
	}
	defer aiOrch.StopNetworking()

	w := warden.NewWarden(b)
	w.RegisterActuator(actuators.NewProcessActuator())

	// ── 1.1 Initialize eBPF Layer (Reflexive Actuation) ─────────
	adapter := &KernelBusAdapter{Bus: b}
	ebpfLoader := kernel.NewLoader(adapter)
	ebpfPath := "PheonixKernel/src/phoenix_exec.o"
	if err := ebpfLoader.Load(ebpfPath); err != nil {
		log.Printf("[WARNING] eBPF Loader failed (Reflexive Actuation disabled): %v", err)
	} else {
		w.RegisterActuator(actuators.NewEBPFActuator(ebpfLoader))
		defer ebpfLoader.Close()
	}

	w.RegisterInvariant(&warden.EvidenceWeightInvariant{
		StateThresholds: map[warden.SystemState]float64{
			warden.StateWatch:       0.5,
			warden.StateSuspicious:  0.7,
			warden.StateCritical:    0.9,
			warden.StateCompromised: 1.0,
		},
	})
	w.RegisterInvariant(&warden.CertificateInvariant{Validator: evLedger})

	aiOrch.RegisterFeature(&ai.LedgerFeature{Ledger: evLedger})
	aiOrch.RegisterFeature(&ai.TraceFeature{Store: traceStore})
	aiOrch.RegisterFeature(&ai.GraphFeature{Graph: graph})
	aiOrch.RegisterFeature(&ai.MonitorFeature{Service: mon})
	aiOrch.RegisterFeature(&ai.TCSFeature{Window: telemetryWindow, DegMon: degMon})
	aiOrch.RegisterFeature(&ai.ArbiterFeature{Arb: arb})
	aiOrch.RegisterFeature(&ai.WardenFeature{Warden: w})
	aiOrch.RegisterFeature(&ai.RealityFeature{AuditPath: "02_docs/02_validation/RUNTIME_REALITY_AUDIT.MD"})

	// Find the graph feature explicitly to use as the GraphProvider
	var graphProvider warden.GraphProvider
	if gf, ok := aiOrch.GetFeature("graph").(*ai.GraphFeature); ok {
		graphProvider = gf
	}
	w.RegisterInvariant(&warden.ContextualInvariant{Provider: graphProvider})

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
	eventsFile := "research/accepted/datasets/logs/test_events.jsonl"

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
		// Orchestrate the tick using the AI orchestrator (coordinates trace, monitor, tcs, arbiter, warden)
		aiOrch.OrchestrateTick(event, event.LamportClock)

		// Tick Step 7: Automatic Checkpointing (Every 1000 ticks)
		if event.LamportClock%1000 == 0 {
			checkpoint, _ := evLedger.Checkpoint()
			log.Printf("[REPLAY] Checkpoint at tick %d: %x", event.LamportClock, checkpoint)
		}
	}

	// ── 4. Verification ───────────────────────────────────────────
	fmt.Println("\n[REPLAY] Execution Complete. Verifying Invariants...")

	// Wait for background AI advisors to finish processing
	aiOrch.Wg.Wait()

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
