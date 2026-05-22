package main

import (
	"crypto/sha256"
	"fmt"
	"testing"

	"phoenix/arbiter"
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
		
		// Initial state: Add a genesis entry
		l.AddEntry("GENESIS", "SYSTEM", []byte("boot"))
		
		// Hash the initial ledger state
		h := sha256.New()
		cp, _ := l.Checkpoint()
		h.Write(cp)
		return fmt.Sprintf("%x", h.Sum(nil))
	}
	
	hash1 := initSystem()
	hash2 := initSystem()
	
	if hash1 != hash2 {
		t.Errorf("Boot non-determinism detected!\nRun 1: %s\nRun 2: %s", hash1, hash2)
	}
}
