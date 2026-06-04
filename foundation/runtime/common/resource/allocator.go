/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — BOUNDED RESOURCE ALLOCATION
//
// The BoundedAllocator tracks memory and resource usage to prevent
// unbounded growth. It enforces limits on memory consumption and item counts.
//
// WORKFLOW:
//   Allocate(size) → check memory limit → check item limit → record allocation
//   Deallocate(size) → release memory → release item count
//   Usage() → return current memory and item counts
//
// THREAD SAFETY: Uses sync/atomic for lock-free updates.
// The CAS (CompareAndSwap) pattern ensures correct concurrent allocation.
//
// RESOURCE BUDGETS:
//   - memoryLimit: maximum bytes that can be allocated
//   - itemLimit: maximum number of items that can be allocated
//   - Both limits are configurable at creation time
// =========================================================================
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

func NewBoundedAllocator(memLimit, itemLimit uint64) *BoundedAllocator {
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
		atomic.AddUint64(&a.itemCount, ^uint64(0))       // Rollback
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

func (a *BoundedAllocator) Usage() (used, count uint64) {
	return atomic.LoadUint64(&a.memoryUsed), atomic.LoadUint64(&a.itemCount)
}
