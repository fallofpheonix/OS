package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix_os/bus"
	"github.com/fallofpheonix/phoenix_os/guard"
	"github.com/fallofpheonix/phoenix_os/ledger"
	"github.com/fallofpheonix/phoenix_os/monitor"
	"github.com/fallofpheonix/phoenix_os/tcs"
	"github.com/fallofpheonix/phoenix_os/trace"
	"github.com/fallofpheonix/phoenix_os/warden"
)

// LedgerWorker drains the evidence channel asynchronously without blocking the fast-path
func LedgerWorker(payloadChan <-chan tcs.ActuationPayload, evLedger *ledger.Ledger) {
	for p := range payloadChan {
		data, _ := json.Marshal(p)
		if err := evLedger.AddEntry(p.ActionID, p.CauseID, data); err != nil {
			log.Printf("[LEDGER ERROR] %v", err)
		}
	}
}

func main() {
	fmt.Println("==================================================")
	fmt.Println("✦ PHOENIX CYBERNETIC SECURITY RUNTIME")
	fmt.Println("==================================================")

	// ── 1. Event Bus ──────────────────────────────────────────────
	b := bus.NewBus()
	rawCh := b.Subscribe("telemetry.raw")
	scoredCh := b.Subscribe("telemetry.scored")
	wardenActionCh := b.Subscribe("warden.action")

	// ── 2. Evidence Ledger (async, channel-based) ─────────────────
	evLedger := &ledger.Ledger{}
	payloadChan := make(chan tcs.ActuationPayload, 100000)
	
	var ledgerWg sync.WaitGroup
	ledgerWg.Add(1)
	go func() {
		defer ledgerWg.Done()
		LedgerWorker(payloadChan, evLedger)
	}()

	// ── 3. TCS Sliding Window (60s window) ────────────────────────
	telemetryWindow := tcs.NewSlidingWindow(60 * time.Second)
	fmt.Println("[BOOT] TCS Sliding Window initialized (60s window)")

	// ── 4. Degradation Monitor (circuit breaker) ──────────────────
	degMon := tcs.NewDegradationMonitor(telemetryWindow, payloadChan)
	fmt.Println("[BOOT] Degradation Monitor active (threshold: 0.85)")

	// ── 5. Trace Storage (SQLite WAL) ─────────────────────────────
	dbPath := "/tmp/phoenix_trace.db"
	_ = os.Remove(dbPath)
	traceStore, err := trace.NewTraceStorage(dbPath, rawCh)
	if err != nil {
		log.Fatalf("[FATAL] Trace storage init failed: %v", err)
	}
	traceStore.StartWriter()
	fmt.Println("[BOOT] Trace WAL storage ready")

	// ── 6. Monitor (EWMA + Linear Kalman) ─────────────────────────
	monInCh := b.Subscribe("telemetry.raw")
	mon := monitor.NewMonitorService(monInCh, b)
	mon.Start()
	fmt.Println("[BOOT] Monitor engine started (EWMA + Kalman)")

	// ── 7. Warden FSM (authoritative, hysteresis-backed) ─────────
	wardenInCh := make(chan monitor.DriftScore, 100)
	
	var wardenWg sync.WaitGroup
	wardenWg.Add(1)
	w := warden.NewWarden(wardenInCh, b)
	go func() {
		defer wardenWg.Done()
		for score := range wardenInCh {
			w.Evaluate(score)
		}
	}()
	fmt.Println("[BOOT] Warden FSM active (hysteresis: 30s dwell)")

	// ── 8. Guard Replay Adapter ───────────────────────────────────
	eventsFile := "/Users/fallofpheonix/os/test_events.jsonl"
	fmt.Println("[BOOT] Starting Guard Adapter (Saturation Mode)")
	guardAdapter := guard.NewGuardAdapter(b, eventsFile, guard.ModeSaturation, 1.0)

	totalEvents, err := guardAdapter.Start()
	if err != nil {
		log.Fatalf("[FATAL] Guard failed: %v", err)
	}
	fmt.Printf("[REPLAY] Replayed %d events successfully\n", totalEvents)

	// ── 9. Drain and Process Pipeline ──────────────────────────────
	// Read and process events from scoredCh until all totalEvents are scored.
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		for ev := range scoredCh {
			var score monitor.DriftScore
			if err := json.Unmarshal(ev.Payload, &score); err == nil {
				wardenInCh <- score

				// Feed scored events into TCS as synthetic telemetry deterministically
				telemetryWindow.AddEvent(tcs.TelemetryEvent{
					Timestamp:  time.Unix(score.WallTimeUnix, 0),
					SequenceID: uint64(score.EventID),
					JitterNS:   0,
				})

				// Synchronously evaluate degradation monitor
				tcsScore := telemetryWindow.Evaluate()
				degMon.Evaluate(tcsScore)

				if score.EventID == totalEvents {
					break
				}
			}
		}
	}()

	// Wait for scored channel processing loop to finish
	wg.Wait()
	close(wardenInCh)

	// Wait for Warden FSM processing to finish
	wardenWg.Wait()

	// Drain any remaining warden actions and log them to the Ledger
	for {
		select {
		case ev := <-wardenActionCh:
			payloadChan <- tcs.ActuationPayload{
				ActionID: fmt.Sprintf("WARDEN-FSM-%d", ev.SeqID),
				CauseID:  "FSM-TRANSITION",
				TargetIP: 0,
				Action:   string(ev.Payload),
			}
		default:
			goto DRAINED
		}
	}
DRAINED:

	// Wait for LedgerWorker to finish draining
	close(payloadChan)
	ledgerWg.Wait()

	// Verify Ledger integrity
	if verr := evLedger.Verify(); verr != nil {
		fmt.Printf("[LEDGER] INTEGRITY VIOLATION: %v\n", verr)
	} else {
		fmt.Printf("[LEDGER] Hash-chain verified (%d entries)\n", len(evLedger.Entries))
	}

	// Print final TCS score
	finalTCS := telemetryWindow.Evaluate()
	fmt.Printf("[TCS] Final Telemetry Confidence Score: %.4f\n", finalTCS)

	// Compute deterministic output hash for replay reproducibility proof
	outputHash := sha256.New()
	for _, entry := range evLedger.Entries {
		outputHash.Write(entry.Hash[:])
	}
	fmt.Printf("[REPLAY] Deterministic output hash: %x\n", outputHash.Sum(nil))

	fmt.Println("==================================================")
	fmt.Println("✦ RUNTIME HALTED")
	fmt.Println("==================================================")
}
