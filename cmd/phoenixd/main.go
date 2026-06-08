package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"

	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
	"github.com/fallofpheonix/phoenix/foundation/runtime/kernel"
)

// PURPOSE: Orchestrates the secure boot sequence of the PhoenixOS substrate.
// CONTRACT: Components must be initialized in strict dependency order.
//           The Persistor must exist before the Ledger; the Ledger must exist
//           before the Bus; the Applier must exist before any probes are loaded.
// FAILURE: Any initialization failure result in an immediate os.Exit(1).
// CONNECTS: This is the system entry point.

// Adapter to bridge kernel.EventPublisher to bus.Bus
type kernelBusAdapter struct {
	b *bus.Bus
}

func (kba *kernelBusAdapter) Publish(topic string, event kernel.TelemetryEvent) {
	// Map kernel.TelemetryEvent to bus.TelemetryEvent
	busEvent := bus.TelemetryEvent{
		EventID:     event.EventID,
		CausalID:    "KERNEL",
		LogicalTick: uint64(event.SeqID), // Note: Map SeqID to LogicalTick for now
		Payload:     event.Payload,
	}
	kba.b.Publish(topic, busEvent)
}

func main() {
	fmt.Println("🦅 PhoenixOS Bootstrap: Initiating Sovereign Sequence")


	// T1.1: Wire Persistor
	dbPath := filepath.Join(os.Getenv("HOME"), ".phoenix", "state.log")
	if err := os.MkdirAll(filepath.Dir(dbPath), 0700); err != nil {
		fmt.Printf("[BOOT FATAL] Failed to create data dir: %v\n", err)
		os.Exit(1)
	}

	persistor, err := ledger.NewPersistor(dbPath)
	if err != nil {
		fmt.Printf("[BOOT FATAL] Failed to create persistor: %v\n", err)
		os.Exit(1)
	}

	// Initialize new ledger or recover? For this slice, just write header if missing,
	// but production would Replay() first. Since we just need to prove wiring, 
	// we initialize and append.
	if _, err := os.Stat(dbPath); os.IsNotExist(err) {
		err = persistor.WriteHeader(ledger.LedgerFileHeader{
			Version:           "1.0",
			GenesisID:         "PHOENIX-GENESIS",
			Timestamp:         0,
			Algorithm:         "SHA256",
			FixedPointDivisor: 1000000,
		})
		if err != nil {
			fmt.Printf("[BOOT FATAL] Failed to write header: %v\n", err)
			os.Exit(1)
		}
	}

	// 1. Initialize Substrate (The Memory/Durable Layer)
	allocator := resource.NewBoundedAllocator(1024*1024*512, 1024) // 512MB limit
	substrateLedger := ledger.NewLedger(allocator).WithPersistor(persistor)
	fmt.Println("[BOOT] Substrate Ledger initialized (Disk-Backed)")

	// 2. Initialize Distribution (The Event Bus)
	eventBus := bus.NewBus()
	fmt.Println("[BOOT] Event Bus active")

	// 3. Initialize Serialization (The Applier)
	applier := bus.NewApplier(bus.ApplierConfig{
		BufferSize: 4096,
		Ledger:     substrateLedger,
		Bus:        eventBus,
		Topics:     []string{"SYSCALL", "ACTUATION", "SECURITY", "MANUAL_TEST"},
	})

	// 4. Start the Ingestion Loop
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	if err := applier.Start(ctx); err != nil {
		fmt.Printf("[BOOT FATAL] Failed to start Event Applier: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("[BOOT] Event Applier loop started")

	// T1.2: Wire eBPF loader to the Bus
	// Note: We use a mocked path for the eBPF object for now since we aren't compiling the C code here.
	bpfLoader := kernel.NewLoader(&kernelBusAdapter{b: eventBus})
	_ = bpfLoader

	// Manual Event Publish (T1.1 Verification)
	event := bus.TelemetryEvent{
		EventID:     "T1.1-BOOT-TEST",
		CausalID:    "GENESIS",
		LogicalTick: 1,
		Payload:     []byte(`{"msg": "System Boot Successful"}`),
	}
	eventBus.Publish("MANUAL_TEST", event)

	// Manual Syscall Publish (T1.2 Verification Mock)
	syscallEvent := bus.TelemetryEvent{
		EventID:     "T1.2-SYSCALL-MOCK",
		CausalID:    "T1.1-BOOT-TEST",
		LogicalTick: 2,
		Payload:     []byte(`{"syscall": "execve", "pid": 1234}`),
	}
	eventBus.Publish("SYSCALL", syscallEvent)

	// 5. Handle System Signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	fmt.Println("🦅 PHOENIX SYSTEM READY (Operational Mode)")
	fmt.Println("   Press Ctrl+C to exit and flush state to disk.")

	<-sigChan
	fmt.Println("\n🦅 PhoenixOS Shutdown Initiated...")
	cancel()

	// Wait for drain (managed by Applier.drain via context cancellation)
	// In a real shutdown we would WaitGroup here, but Applier ctx cancellation works.
	fmt.Println("🦅 Shutdown Complete.")
}
