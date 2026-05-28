package kernel

import (
	"fmt"
	"sync/atomic"
)

// OverflowDetector defines the interface for detecting overflow conditions.
type OverflowDetector interface {
	// Configure sets the overflow threshold.
	Configure(threshold uint64) error
	// Record increments an internal counter and checks for overflow.
	// Returns true if overflow is detected (i.e., current count exceeds threshold).
	Record(count uint64) (bool, error)
	// CurrentCount returns the current recorded count.
	CurrentCount() uint64
	// Reset clears the current count.
	Reset()
	// Threshold returns the configured threshold.
	Threshold() uint64
}

// NewOverflowDetector creates a new instance of OverflowDetector.
func NewOverflowDetector() OverflowDetector {
	return &simpleOverflowDetector{
		threshold: atomic.Uint64{},
		current:   atomic.Uint64{},
	}
}

type simpleOverflowDetector struct {
	threshold atomic.Uint64
	current   atomic.Uint64
}

// Configure sets the overflow threshold.
func (od *simpleOverflowDetector) Configure(threshold uint64) error {
	if threshold == 0 {
		return fmt.Errorf("overflow threshold cannot be zero")
	}
	od.threshold.Store(threshold)
	return nil
}

// Record increments an internal counter and checks for overflow.
func (od *simpleOverflowDetector) Record(count uint64) (bool, error) {
	if od.threshold.Load() == 0 {
		return false, fmt.Errorf("overflow detector not configured: threshold is zero")
	}
	newCount := od.current.Add(count)
	return newCount > od.threshold.Load(), nil
}

// CurrentCount returns the current recorded count.
func (od *simpleOverflowDetector) CurrentCount() uint64 {
	return od.current.Load()
}

// Reset clears the current count.
func (od *simpleOverflowDetector) Reset() {
	od.current.Store(0)
}

// Threshold returns the configured threshold.
func (od *simpleOverflowDetector) Threshold() uint64 {
	return od.threshold.Load()
}
