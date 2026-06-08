/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package ordering

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// Event represents a single event with an associated sequence number.
type Event struct {
	SequenceID uint64
	Payload    []byte
	// Additional metadata could go here, e.g., Timestamp, SourceID
}

// EventOrderer defines the interface for managing and verifying event order.
type EventOrderer interface {
	// RecordEvent assigns a sequence ID to an event and stores it.
	// Returns the assigned sequence ID.
	RecordEvent(payload []byte) (uint64, error)
	// GetEvent retrieves an event by its sequence ID.
	GetEvent(sequenceID uint64) (*Event, error)
	// LastSequenceID returns the highest sequence ID recorded.
	LastSequenceID() uint64
}

// NewEventOrderer creates a new instance of EventOrderer.
func NewEventOrderer() EventOrderer {
	return &simpleEventOrderer{
		events: make(map[uint64]*Event),
	}
}

type simpleEventOrderer struct {
	nextSequenceID atomic.Uint64
	events         map[uint64]*Event
	mu             sync.RWMutex // Protects events map
}

// RecordEvent assigns a sequence ID to an event and stores it.
func (seo *simpleEventOrderer) RecordEvent(payload []byte) (uint64, error) {
	if payload == nil {
		return 0, fmt.Errorf("event payload cannot be nil")
	}

	sequenceID := seo.nextSequenceID.Add(1)
	event := &Event{
		SequenceID: sequenceID,
		Payload:    payload,
	}

	seo.mu.Lock()
	defer seo.mu.Unlock()
	seo.events[sequenceID] = event
	return sequenceID, nil
}

// GetEvent retrieves an event by its sequence ID.
func (seo *simpleEventOrderer) GetEvent(sequenceID uint64) (*Event, error) {
	seo.mu.RLock()
	defer seo.mu.RUnlock()
	event, ok := seo.events[sequenceID]
	if !ok {
		return nil, fmt.Errorf("event with sequence ID %d not found", sequenceID)
	}
	return event, nil
}

// LastSequenceID returns the highest sequence ID recorded.
func (seo *simpleEventOrderer) LastSequenceID() uint64 {
	return seo.nextSequenceID.Load()
}
