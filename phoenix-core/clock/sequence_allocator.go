package clock

import (
	"sync/atomic"
)

// SequenceAllocator provides monotonic sequence IDs for events.
type SequenceAllocator struct {
	nextID uint64
}

func NewSequenceAllocator() *SequenceAllocator {
	return &SequenceAllocator{nextID: 0}
}

// Allocate returns the next monotonic sequence ID.
func (a *SequenceAllocator) Allocate() uint64 {
	return atomic.AddUint64(&a.nextID, 1)
}
