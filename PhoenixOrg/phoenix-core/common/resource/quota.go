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

func (q *QuotaManager) Status() (uint64, uint32) {
	return atomic.LoadUint64(&q.ioUsed), atomic.LoadUint32(&q.workerActive)
}
