package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fallofpheonix/phoenix/assurance/security/engine"
	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

// WardenStatePayload represents the state of the Warden persisted in the ledger.
type WardenStatePayload struct {
	State engine.SystemState `json:"state"`
}

func main() {
	path := flag.String("path", "warden_slice.jsonl", "Path to the ledger file")
	replay := flag.Bool("replay", false, "Replay the ledger and reconstruct Warden state")
	flag.Parse()

	if *replay {
		fmt.Printf("[RESTART] Opening ledger at %s for Warden reconstruction...\n", *path)
		p, err := ledger.NewPersistor(*path)
		if err != nil {
			log.Fatalf("Failed to open persistor: %v", err)
		}
		defer p.Close()

		// 1. Reconstruct the Ledger Chain
		chain := ledger.NewChain()
		if _, err := p.Load(chain); err != nil {
			log.Fatalf("Failed to load ledger: %v", err)
		}

		// 2. Initialize a fresh Warden
		w := engine.NewWarden()
		fmt.Printf("[REPLAY] Initial Warden state: %s\n", w.GetState())

		// 3. Replay Events into the Warden
		// In a real system, the Replayer or Applier would handle this.
		// For this slice, we iterate the chain and apply ENFORCEMENT actions.
		for i := uint64(0); ; i++ {
			ev, ok := chain.GetBySequence(i)
			if !ok {
				break
			}
			if ev.Type == ledger.EventEnforce {
				var payload WardenStatePayload
				if err := json.Unmarshal(ev.Payload, &payload); err == nil {
					fmt.Printf("[REPLAY] Applying transition to %s...\n", payload.State)
					// We use Transition directly because Reconstruct doesn't know
					// the history, but the Warden FSM enforces the ladder.
					if err := w.Transition(payload.State); err != nil {
						log.Fatalf("[FATAL] Replay failed FSM ladder: %v", err)
					}
				}
			}
		}

		fmt.Printf("[SUCCESS] Warden State Reconstructed: %s\n", w.GetState())
		fmt.Printf("[SUCCESS] Ledger Head Hash: %s\n", chain.GetHead().Hash)
	} else {
		os.Remove(*path)
		fmt.Printf("[START] Initializing new Warden vertical slice at %s...\n", *path)

		p, _ := ledger.NewPersistor(*path)
		defer p.Close()

		w := engine.NewWarden()
		fmt.Printf("[WARDEN] Current state: %s\n", w.GetState())

		// 1. Persist Genesis
		e0 := ledger.NewEvent(0, ledger.EventGenesis, []byte("{}"), "", "SYSTEM")
		p.Append(e0)

		// 2. Trigger Transition: SAFE -> WATCH
		target := engine.StateWatch
		fmt.Printf("[WARDEN] Transitioning: %s -> %s\n", w.GetState(), target)
		if err := w.Transition(target); err != nil {
			log.Fatalf("Transition failed: %v", err)
		}

		// 3. Persist Transition
		payload, _ := json.Marshal(WardenStatePayload{State: w.GetState()})
		e1 := ledger.NewEvent(1, ledger.EventEnforce, payload, e0.Hash, "WARDEN")
		if err := p.Append(e1); err != nil {
			log.Fatalf("Failed to persist transition: %v", err)
		}

		fmt.Printf("[PERSISTED] Warden state %s at hash %s\n", w.GetState(), e1.Hash)
		fmt.Println("[EXIT] Process terminating.")
	}
}
