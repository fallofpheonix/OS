package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/fallofpheonix/phoenix/foundation/ledger"
)

func main() {
	path := flag.String("path", "ledger.jsonl", "Path to the ledger file")
	replay := flag.Bool("replay", false, "Replay the ledger and verify state")
	flag.Parse()

	if *replay {
		fmt.Printf("[RESTART] Opening ledger at %s for reconstruction...\n", *path)
		p, err := ledger.NewPersistor(*path)
		if err != nil {
			log.Fatalf("Failed to open persistor: %v", err)
		}
		defer p.Close()

		replayer := ledger.NewReplayer(p)
		hash, err := replayer.Reconstruct()
		if err != nil {
			log.Fatalf("Reconstruction failed: %v", err)
		}

		fmt.Printf("[SUCCESS] Head Hash Verified: %s\n", hash)
		fmt.Println("[SUCCESS] State survives restart.")
	} else {
		// Clean up old ledger for clean demo
		os.Remove(*path)

		fmt.Printf("[START] Initializing new ledger at %s...\n", *path)
		p, err := ledger.NewPersistor(*path)
		if err != nil {
			log.Fatalf("Failed to create persistor: %v", err)
		}
		defer p.Close()

		// Genesis
		e0 := ledger.NewEvent(0, ledger.EventGenesis, []byte("{}"), "", "SYSTEM")
		if err := p.Append(e0); err != nil {
			log.Fatalf("Failed to append genesis: %v", err)
		}

		// Fact
		e1 := ledger.NewEvent(1, ledger.EventFact, []byte(`{"fact_id":"SENSE-001","value":"HIGH_HEAT"}`), e0.Hash, "SENSOR_GATEWAY")
		if err := p.Append(e1); err != nil {
			log.Fatalf("Failed to append e1: %v", err)
		}

		fmt.Printf("[PERSISTED] Head Hash: %s\n", e1.Hash)
		fmt.Println("[EXIT] Process terminating.")
	}
}
