package kernel

import (
	"sync/atomic"
)

// DropTracker defines the interface for tracking dropped events.
type DropTracker interface {
	// Record drops a specified number of events.
	Record(count uint64)
	// TotalDrops returns the cumulative count of all dropped events.
	TotalDrops() uint64
	// Reset clears the total count of dropped events.
	Reset()
}

// NewDropTracker creates a new instance of DropTracker.
func NewDropTracker() DropTracker {
	return &simpleDropTracker{
		totalDrops: atomic.Uint64{},
	}
}

type simpleDropTracker struct {
	totalDrops atomic.Uint64
}

// Record drops a specified number of events.
func (dt *simpleDropTracker) Record(count uint64) {
	dt.totalDrops.Add(count)
}

// TotalDrops returns the cumulative count of all dropped events.
func (dt *simpleDropTracker) TotalDrops() uint64 {
	return dt.totalDrops.Load()
}

// Reset clears the total count of dropped events.
func (dt *simpleDropTracker) Reset() {
	dt.totalDrops.Store(0)
}
