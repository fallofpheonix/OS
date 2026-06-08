/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — RESOURCE QUOTA MANAGEMENT
//
// The QuotaManager enforces deterministic resource limits for I/O and workers.
// It prevents any single component from consuming more than its fair share
// of system resources.
//
// WORKFLOW:
//   RequestIO(size) → check ioUsed + size <= ioLimit → record or reject
//   AcquireWorker() → check workerActive < workerLimit → reserve or reject
//   ReleaseWorker() → decrement workerActive
//
// THREAD SAFETY: Uses sync/atomic for lock-free updates.
// The CAS (CompareAndSwap) pattern ensures correct concurrent access.
//
// RESOURCE BUDGETS:
//   - ioLimit: maximum bytes that can be consumed by I/O operations
//   - workerLimit: maximum concurrent workers
//   - Both are configurable at creation time
//
// USAGE: DeterministicWorkerPool uses QuotaManager to enforce limits.
// =========================================================================
package resource

import (
	"fmt"
	"sync/atomic"
)

// QuotaManager implements deterministic resource limits for I/O and Workers.
// This addresses [RSC-004], [RSC-005], and [RSC-007].
type QuotaManager struct {
	ioLimit      uint64
	ioUsed       uint64
	workerLimit  uint32
	workerActive uint32
}

func NewQuotaManager(ioLimit uint64, workerLimit uint32) *QuotaManager {
	return &QuotaManager{
		ioLimit:     ioLimit,
		workerLimit: workerLimit,
	}
}

// RequestIO checks if an I/O operation of 'size' bytes is allowed.
func (q *QuotaManager) RequestIO(size uint64) error {
	newUsed := atomic.AddUint64(&q.ioUsed, size)
	if q.ioLimit > 0 && newUsed > q.ioLimit {
		atomic.AddUint64(&q.ioUsed, ^uint64(size-1))
		return fmt.Errorf("I/O quota exceeded: %d > %d", newUsed, q.ioLimit)
	}
	return nil
}

// AcquireWorker attempts to reserve a worker slot.
func (q *QuotaManager) AcquireWorker() error {
	active := atomic.LoadUint32(&q.workerActive)
	if q.workerLimit > 0 && active >= q.workerLimit {
		return fmt.Errorf("worker quota exceeded: %d >= %d", active, q.workerLimit)
	}

	if atomic.CompareAndSwapUint32(&q.workerActive, active, active+1) {
		return nil
	}
	return q.AcquireWorker() // Retry on race (deterministic in logical tick context)
}

// ReleaseWorker frees a worker slot.
func (q *QuotaManager) ReleaseWorker() {
	atomic.AddUint32(&q.workerActive, ^uint32(0))
}

func (q *QuotaManager) Status() (used uint64, active uint32) {
	return atomic.LoadUint64(&q.ioUsed), atomic.LoadUint32(&q.workerActive)
}
