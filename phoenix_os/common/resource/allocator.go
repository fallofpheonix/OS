package resource

import (
	"fmt"
	"sync/atomic"
)

// BoundedAllocator tracks memory and resource usage deterministically.
type BoundedAllocator struct {
	memoryLimit uint64
	memoryUsed  uint64
	itemCount   uint64
	itemLimit   uint64
}

func NewBoundedAllocator(memLimit uint64, itemLimit uint64) *BoundedAllocator {
	return &BoundedAllocator{
		memoryLimit: memLimit,
		itemLimit:   itemLimit,
	}
}

// Allocate records an allocation and returns an error if limits are exceeded.
func (a *BoundedAllocator) Allocate(size uint64) error {
	newUsed := atomic.AddUint64(&a.memoryUsed, size)
	if a.memoryLimit > 0 && newUsed > a.memoryLimit {
		atomic.AddUint64(&a.memoryUsed, ^uint64(size-1)) // Rollback
		return fmt.Errorf("memory quota exceeded: %d > %d", newUsed, a.memoryLimit)
	}

	newItems := atomic.AddUint64(&a.itemCount, 1)
	if a.itemLimit > 0 && newItems > a.itemLimit {
		atomic.AddUint64(&a.itemCount, ^uint64(0)) // Rollback
		atomic.AddUint64(&a.memoryUsed, ^uint64(size-1)) // Rollback
		return fmt.Errorf("item quota exceeded: %d > %d", newItems, a.itemLimit)
	}

	return nil
}

// Deallocate records a deallocation.
func (a *BoundedAllocator) Deallocate(size uint64) {
	atomic.AddUint64(&a.memoryUsed, ^uint64(size-1))
	atomic.AddUint64(&a.itemCount, ^uint64(0))
}

func (a *BoundedAllocator) Usage() (uint64, uint64) {
	return atomic.LoadUint64(&a.memoryUsed), atomic.LoadUint64(&a.itemCount)
}
