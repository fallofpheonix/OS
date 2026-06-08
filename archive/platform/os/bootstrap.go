/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/fallofpheonix/phoenix/foundation/ledger/src"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

// PURPOSE: Orchestrates the secure boot sequence of the PhoenixOS substrate.
// CONTRACT: Components must be initialized in strict dependency order. 
//           The Ledger must exist before the Bus; the Applier must exist 
//           before any eBPF probes are loaded.
// FAILURE: Any initialization failure result in an immediate os.Exit(1).
// CONNECTS: This is the system entry point.

func main() {
	fmt.Println("🦅 PhoenixOS Bootstrap: Initiating Sovereign Sequence")

	// 1. Initialize Substrate (The Memory/Durable Layer)
	// WHY: Without a bounded memory allocator, the system is vulnerable to 
	// OOM-based denial of service via ledger bloat.
	allocator := resource.NewBoundedAllocator(1024*1024*512, 1024) // 512MB limit
	substrateLedger := ledger.NewLedger(allocator)
	fmt.Println("[BOOT] Substrate Ledger initialized (Memory-Bounded)")

	// 2. Initialize Distribution (The Event Bus)
	// WHY: The Bus is the nervous system through which all telemetry flows.
	eventBus := bus.NewBus()
	fmt.Println("[BOOT] Event Bus active")

	// 3. Initialize Serialization (The Applier)
	// WHY: The Applier is the single-threaded gatekeeper that ensures 
	// deterministic order for the ledger.
	applier := bus.NewApplier(bus.ApplierConfig{
		BufferSize: 4096,
		Ledger:     substrateLedger,
		Bus:        eventBus,
		Topics:     []string{"SYSCALL", "ACTUATION", "SECURITY"},
	})

	// 4. Start the Ingestion Loop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := applier.Start(ctx); err != nil {
		fmt.Printf("[BOOT FATAL] Failed to start Event Applier: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[BOOT] Event Applier loop started")

	// 5. Handle System Signals
	// WHY: Graceful shutdown is required to drain the Applier's channel 
	// and ensure forensic integrity of the last events.
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🦅 PHOENIX SYSTEM READY (Operational Mode)")

	<-sigChan
	fmt.Println("\n🦅 PhoenixOS Shutdown Initiated...")
	cancel()
	
	// Wait for drain (managed by Applier.drain via context cancellation)
	fmt.Println("🦅 Shutdown Complete.")
}
