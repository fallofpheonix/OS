package kernel

import (
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

func TestNewBurstController(t *testing.T) {
	// Valid creation
	bc, err := NewBurstController(10, 1)
	if err != nil {
		t.Fatalf("NewBurstController failed: %v", err)
	}
	if bc.Capacity() != 10 {
		t.Errorf("Expected capacity 10, got %d", bc.Capacity())
	}
	if bc.Available() != 10 { // Starts full
		t.Errorf("Expected available 10, got %d", bc.Available())
	}

	// Zero capacity (should fail)
	_, err = NewBurstController(0, 1)
	if err == nil {
		t.Error("Expected error for zero capacity, got nil")
	} else if err.Error() != "burst capacity cannot be zero" {
		t.Errorf("Expected 'burst capacity cannot be zero' error, got: %v", err)
	}

	// Zero refill rate (should fail)
	_, err = NewBurstController(10, 0)
	if err == nil {
		t.Error("Expected error for zero refill rate, got nil")
	} else if err.Error() != "refill rate cannot be zero" {
		t.Errorf("Expected 'refill rate cannot be zero' error, got: %v", err)
	}
}

func TestBurstController_Allow(t *testing.T) {
	bc, _ := NewBurstController(10, 1) // Capacity 10, refill 1/sec

	// Allow within capacity
	if !bc.Allow(5) {
		t.Error("Expected to allow 5 events, but denied")
	}
	if bc.Available() != 5 {
		t.Errorf("Expected available 5, got %d", bc.Available())
	}

	// Allow exact remaining capacity
	if !bc.Allow(5) {
		t.Error("Expected to allow 5 events, but denied")
	}
	if bc.Available() != 0 {
		t.Errorf("Expected available 0, got %d", bc.Available())
	}

	// Deny when not enough capacity
	if bc.Allow(1) {
		t.Error("Expected to deny 1 event, but allowed")
	}
	if bc.Available() != 0 {
		t.Errorf("Expected available 0, got %d", bc.Available())
	}

	// Allow zero events
	if !bc.Allow(0) {
		t.Error("Expected to allow 0 events, but denied")
	}
	if bc.Available() != 0 {
		t.Errorf("Expected available 0, got %d", bc.Available())
	}
}

func TestBurstController_Refill(t *testing.T) {
	bc, _ := NewBurstController(10, 1) // Capacity 10, refill 1/sec
	_ = bc.Allow(10)                    // Empty the bucket

	if bc.Available() != 0 {
		t.Fatalf("Expected 0 available after emptying, got %d", bc.Available())
	}

	// Wait for refill (1 second should add 1 token)
	time.Sleep(1 * time.Second)
	if bc.Available() != 1 {
		t.Errorf("Expected 1 available after 1 sec refill, got %d", bc.Available())
	}

	// Wait for more refill (5 seconds should add 5 tokens)
	time.Sleep(5 * time.Second)
	if bc.Available() != 6 { // 1 (from previous) + 5
		t.Errorf("Expected 6 available after 5 sec refill, got %d", bc.Available())
	}

	// Wait to refill beyond capacity (should cap at capacity)
	time.Sleep(10 * time.Second) // Should refill 10 tokens, but cap at 10 (total 6+10=16, capped to 10)
	if bc.Available() != 10 {
		t.Errorf("Expected available to cap at 10, got %d", bc.Available())
	}
}

func TestBurstController_Reset(t *testing.T) {
	bc, _ := NewBurstController(10, 1)
	_ = bc.Allow(5) // Consume some tokens
	if bc.Available() != 5 {
		t.Fatalf("Expected available 5, got %d", bc.Available())
	}

	bc.Reset()
	if bc.Available() != 10 { // Should be full after reset
		t.Errorf("Expected available 10 after reset, got %d", bc.Available())
	}

	// Ensure refill time is also reset, so next Allow doesn't immediately refill
	_ = bc.Allow(10)
	time.Sleep(1 * time.Second)
	if bc.Available() != 1 { // Should be 1 due to refillRate of 1/sec
		t.Errorf("Expected available 1 after 1s, got %d", bc.Available())
	}
}

func TestBurstController_Concurrency(t *testing.T) {
	capacity := uint64(100)
	refillRate := uint64(10) // 10 tokens/sec
	bc, _ := NewBurstController(capacity, refillRate)

	numGoroutines := 10
	allowAttemptsPerGoroutine := 20

	var wg sync.WaitGroup
	wg.Add(numGoroutines)

	allowedCount := uint64(0)

	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < allowAttemptsPerGoroutine; j++ {
				// Each attempt tries to allow 1 event
				if bc.Allow(1) {
					atomic.AddUint64(&allowedCount, 1)
				}
				time.Sleep(10 * time.Millisecond) // Simulate some work
			}
		}()
	}
	wg.Wait()

	// Since refillRate is 10/sec, and goroutines run for ~200ms (20 * 10ms),
	// it's tricky to get an exact expected count without a precise simulation.
	// We primarily want to ensure no panics/race conditions and that `allowedCount` is reasonable.

	t.Logf("Total events allowed: %d", allowedCount)
	t.Logf("Remaining available tokens: %d", bc.Available())

	// A rough check: allowedCount should not exceed the sum of initial capacity + refills during test
	// The test runs for ~200ms. So roughly 2 refills * 10 tokens/sec = 2 tokens refilled.
	// But since time.Sleep is not precise and goroutines schedule varies, we can't be exact.
	// For simplicity, let's just ensure allowedCount <= initial capacity + total possible refills
	// Total test time approx: 200ms. Max possible refills in 200ms = 10 tokens/sec * 0.2 sec = 2 tokens.
	maxPossibleAllowed := capacity + (refillRate * 2) // Roughly 2 seconds passes during the test
	if allowedCount > maxPossibleAllowed {
		t.Errorf("Allowed count %d exceeded max possible %d", allowedCount, maxPossibleAllowed)
	}
	if allowedCount == 0 {
		t.Error("Expected some events to be allowed, but allowedCount is 0")
	}

	// Test a large single Allow request concurrently
	bc.Reset()
	allowedCount = 0
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			if bc.Allow(capacity / uint64(numGoroutines)) {
				atomic.AddUint64(&allowedCount, capacity/uint64(numGoroutines))
			}
		}()
	}
	wg.Wait()
	t.Logf("Total events allowed in burst concurrency: %d", allowedCount)
	if allowedCount > capacity {
		t.Errorf("Allowed count %d exceeded capacity %d", allowedCount, capacity)
	}
	if allowedCount == 0 {
		t.Error("Expected some events to be allowed, but allowedCount is 0")
	}
}
