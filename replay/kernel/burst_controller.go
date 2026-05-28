package kernel

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// BurstController defines the interface for managing event bursts.
type BurstController interface {
	// Allow attempts to allow a certain number of events.
	// Returns true if events are allowed, false otherwise (burst limit exceeded).
	Allow(events uint64) bool
	// Capacity returns the maximum number of events allowed in a burst.
	Capacity() uint64
	// Available returns the current number of available slots in the burst.
	Available() uint64
	// Reset resets the burst controller's state.
	Reset()
}

// NewBurstController creates a new instance of BurstController.
// capacity: maximum number of events allowed in a burst.
// refillRate: number of events refilled per second.
func NewBurstController(capacity uint64, refillRate uint64) (BurstController, error) {
	if capacity == 0 {
		return nil, fmt.Errorf("burst capacity cannot be zero")
	}
	if refillRate == 0 {
		return nil, fmt.Errorf("refill rate cannot be zero")
	}

	bc := &tokenBucketBurstController{
		capacity:        capacity,
		refillRate:      refillRate,
		currentTokens:   atomic.Uint64{},
		lastRefillTime:  time.Now(),
	}
	bc.currentTokens.Store(capacity) // Start with full capacity
	return bc, nil
}

type tokenBucketBurstController struct {
	capacity        uint64
	refillRate      uint64 // events per second
	currentTokens   atomic.Uint64
	mu              sync.Mutex // Protects lastRefillTime for refills
	lastRefillTime  time.Time
}

// Allow attempts to allow a certain number of events.
func (tbc *tokenBucketBurstController) Allow(events uint64) bool {
	if events == 0 {
		return true
	}

	tbc.mu.Lock()
	defer tbc.mu.Unlock()

	tbc.refillTokens()

	current := tbc.currentTokens.Load()
	if current >= events {
		tbc.currentTokens.Store(current - events) // Directly subtract
		return true
	}
	return false
}

// Capacity returns the maximum number of events allowed in a burst.
func (tbc *tokenBucketBurstController) Capacity() uint64 {
	return tbc.capacity
}

// Available returns the current number of available slots in the burst.
func (tbc *tokenBucketBurstController) Available() uint64 {
	tbc.mu.Lock()
	defer tbc.mu.Unlock()
	tbc.refillTokens()
	return tbc.currentTokens.Load()
}

// Reset resets the burst controller's state to full capacity.
func (tbc *tokenBucketBurstController) Reset() {
	tbc.mu.Lock()
	defer tbc.mu.Unlock()
	tbc.currentTokens.Store(tbc.capacity)
	tbc.lastRefillTime = time.Now()
}

// refillTokens calculates and adds tokens based on elapsed time and refill rate.
// This method must be called under a lock.
func (tbc *tokenBucketBurstController) refillTokens() {
	now := time.Now()
	elapsed := now.Sub(tbc.lastRefillTime)
	if elapsed > 0 {
		tokensToAdd := uint64(float64(tbc.refillRate) * elapsed.Seconds())
		if tokensToAdd > 0 {
			current := tbc.currentTokens.Load()
			newTokens := current + tokensToAdd
			if newTokens > tbc.capacity {
				newTokens = tbc.capacity
			}
			tbc.currentTokens.Store(newTokens)
		}
		tbc.lastRefillTime = now
	}
}
