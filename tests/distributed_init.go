package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"
	"github.com/fallofpheonix/PheonixDistributed/discovery"
	"github.com/fallofpheonix/PheonixDistributed/ledger"
)

func main() {
	fmt.Println("--- PhoenixOS Stage D: Networking Instantiation Test ---")

	// 1. Setup Orchestrator
	o := ai.NewAIOrchestrator()

	// 2. Mock Distributed Components
	disc := discovery.NewBeaconTransport(9999, "test-node-0")
	cons := &ledger.DistributedLedger{} // Simplified mock

	// 3. Start Networking
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	fmt.Println("[Test] Starting Peer Discovery and Consensus...")
	err := o.StartNetworking(ctx, disc, cons)
	if err != nil {
		fmt.Printf("[Test] FAILURE: %v\n", err)
		return
	}

	fmt.Println("[Test] Networking started successfully.")
	time.Sleep(1 * time.Second)
	fmt.Println("[Test] SUCCESS: Distributed primitives active.")
}
