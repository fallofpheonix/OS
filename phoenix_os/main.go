package main

import (
	"crypto/sha256"
	"fmt"
	"log"
	"os"
	"time"

	"phoenix/arbiter"
	"phoenix/bus"
	"phoenix/common/logical_clock"
	"phoenix/common/resource"
	"phoenix/common/serialization"
	"phoenix/guard"
	"phoenix/ledger"
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

	mon := monitor.NewMonitorService(nil, b)
	arb := arbiter.NewArbiter(b)
	w := warden.NewWarden(b)

	// ── 2. Guard Replay Adapter (Seeded for Determinism) ─────────
	eventsFile := "/Users/fallofpheonix/os/test_events.jsonl"
	const replaySeed = 42
	fmt.Printf("[BOOT] Starting Guard Adapter (Seed: %d)\n", replaySeed)
	guardAdapter := guard.NewGuardAdapter(b, eventsFile, guard.ModeSaturation, 1.0, replaySeed)

	events, err := guardAdapter.FetchEvents()
	if err != nil {
		log.Fatalf("[FATAL] Failed to fetch events: %v", err)
	}
	fmt.Printf("[REPLAY] Loaded %d events for deterministic execution\n", len(events))

	// ── 3. Deterministic Main Loop (Logical Tick) ────────────────
	for _, event := range events {
		// Logical Tick Increment
		event.LogicalTick = logicalClock.Tick()

		// Tick Step 1: Ingest into Trace Storage
		if err := traceStore.Write(event); err != nil {
			log.Printf("[TRACE ERROR] %v", err)
		}

		// Tick Step 2: Process through Monitor (L3)
		score := mon.Process(event)

		// Tick Step 3: Update TCS (L3 Confidence)
		telemetryWindow.AddEvent(tcs.TelemetryEvent{
			Timestamp:  time.Unix(event.WallTimeUnix, 0),
			SequenceID: uint64(event.SeqID),
			Payload:    event.Payload,
			JitterNS:   0,
		})
		tcsScore := telemetryWindow.Evaluate()
		
		// Tick Step 4: Evaluate Degradation (Circuit Breaker)
		oldDegraded := degMon.IsDegraded()
		degMon.Evaluate(tcsScore)
		if degMon.IsDegraded() != oldDegraded {
			action := "ENTER_NORMAL_MODE"
			if degMon.IsDegraded() {
				action = "ENTER_DEGRADED_MODE"
			}
			payload, _ := serialization.CanonicalJSON(map[string]interface{}{
				"action": action,
				"score":  tcsScore,
			})
			evLedger.AddEntry("STATE-TRANSITION", "TCS-VIOLATION", payload)
		}

		// Tick Step 5: Strategic Policy Evaluation (L5.5 Arbiter)
		targetState, authorized := arb.Evaluate(score, tcsScore)
		
		// Tick Step 6: Tactical Actuation (L5 Warden)
		if authorized && !degMon.IsDegraded() {
			transitioned := w.Actuate(targetState, event.SeqID, event.WallTimeUnix)
			if transitioned {
				payload, _ := serialization.CanonicalJSON(map[string]interface{}{
					"state": string(w.State),
				})
				evLedger.AddEntry(fmt.Sprintf("WARDEN-ACTION-%d", event.SeqID), "POLICY-ACTUATION", payload)
			}
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
}
