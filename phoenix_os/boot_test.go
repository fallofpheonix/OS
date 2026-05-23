package main

import (
	"testing"

	"phoenix/arbiter"
	"phoenix/boot"
	"phoenix/bus"
	"phoenix/common/resource"
	"phoenix/ledger/src"
	"phoenix/monitor"
	"phoenix/warden"
)

func TestBootReproducibility(t *testing.T) {
	// Axiom: System initialization must result in identical state hashes.
	
	initSystem := func() string {
		b := bus.NewBus()
		alloc := resource.NewBoundedAllocator(1024, 10)
		l := ledger.NewLedger(alloc)
		_ = monitor.NewMonitorService(nil, b)
		_ = arbiter.NewArbiter(b)
		_ = warden.NewWarden(b)
		
		// Capture boot telemetry
		bootInfo := []boot.SubsystemInfo{
			boot.NewSubsystemInfo("Bus", "1.0.0", nil),
			boot.NewSubsystemInfo("Ledger", "2.0.0", nil),
		}
		bt, _ := boot.CaptureBootTelemetry(l, bootInfo)
		return bt.Checksum
	}
	
	hash1 := initSystem()
	hash2 := initSystem()
	
	if hash1 != hash2 {
		t.Errorf("Boot non-determinism detected!\nRun 1: %s\nRun 2: %s", hash1, hash2)
	}
}
