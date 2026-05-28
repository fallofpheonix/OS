package kernel

import (
	"fmt"
	"sync/atomic"
)

// RingMonitorState represents the current state of a monitored ring buffer.
type RingMonitorState struct {
	Capacity      uint64
	CurrentSize   uint64
	OverflowCount uint64
	DropCount     uint64
}

// RingMonitor defines the interface for monitoring a kernel-level ring buffer.
type RingMonitor interface {
	// StartMonitoring initializes and begins monitoring a ring buffer.
	StartMonitoring(capacity uint64) error
	// StopMonitoring halts the monitoring process.
	StopMonitoring() error
	// ReportState returns the current state of the monitored ring buffer.
	ReportState() (RingMonitorState, error)
	// SimulateEvent adds an event to the ring buffer, potentially causing overflow or drops.
	SimulateEvent(numEvents uint64) error
}

// NewRingMonitor creates a new instance of RingMonitor.
func NewRingMonitor() RingMonitor {
	return &ringMonitorImpl{
		isRunning: atomic.Bool{},
	}
}

type ringMonitorImpl struct {
	capacity      atomic.Uint64
	currentSize   atomic.Uint64
	overflowCount atomic.Uint64
	dropCount     atomic.Uint64
	isRunning     atomic.Bool
}

// StartMonitoring initializes and begins monitoring a ring buffer.
func (rm *ringMonitorImpl) StartMonitoring(capacity uint64) error {
	if rm.isRunning.Load() {
		return fmt.Errorf("ring monitor is already running")
	}
	if capacity == 0 {
		return fmt.Errorf("capacity cannot be zero")
	}

	rm.capacity.Store(capacity)
	rm.currentSize.Store(0)
	rm.overflowCount.Store(0)
	rm.dropCount.Store(0)
	rm.isRunning.Store(true)
	return nil
}

// StopMonitoring halts the monitoring process.
func (rm *ringMonitorImpl) StopMonitoring() error {
	if !rm.isRunning.Load() {
		return fmt.Errorf("ring monitor is not running")
	}
	rm.isRunning.Store(false)
	return nil
}

// ReportState returns the current state of the monitored ring buffer.
func (rm *ringMonitorImpl) ReportState() (RingMonitorState, error) {
	if !rm.isRunning.Load() {
		return RingMonitorState{}, fmt.Errorf("ring monitor is not running")
	}
	return RingMonitorState{
		Capacity:      rm.capacity.Load(),
		CurrentSize:   rm.currentSize.Load(),
		OverflowCount: rm.overflowCount.Load(),
		DropCount:     rm.dropCount.Load(),
	}, nil
}

// SimulateEvent adds an event to the ring buffer, potentially causing overflow or drops.
func (rm *ringMonitorImpl) SimulateEvent(numEvents uint64) error {
	if !rm.isRunning.Load() {
		return fmt.Errorf("ring monitor is not running")
	}

	capacity := rm.capacity.Load()
	currentSize := rm.currentSize.Load()
	newSize := currentSize + numEvents

	if newSize <= capacity {
		rm.currentSize.Store(newSize)
	} else {
		overflow := newSize - capacity
		rm.overflowCount.Add(overflow)
		rm.dropCount.Add(overflow) // Assuming overflow implies drops
		rm.currentSize.Store(capacity) // Ring buffer remains at capacity
	}
	return nil
}
