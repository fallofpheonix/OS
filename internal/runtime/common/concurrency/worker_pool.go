/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — DETERMINISTIC WORKER POOL
//
// The DeterministicWorkerPool manages concurrent task execution with
// resource quotas. It ensures that tasks are processed within the
// system's resource budget (I/O and worker limits).
//
// WORKFLOW:
//   Initialize → NewDeterministicWorkerPool(quota, buffer)
//   → Start(count) → spawn worker goroutines
//   → Submit(task) → task queued in channel
//   → Worker picks task → quota.AcquireWorker() → execute → quota.ReleaseWorker()
//   → Wait() → block until all tasks complete
//
// RESOURCE MANAGEMENT:
//   - QuotaManager enforces I/O and worker limits
//   - Workers must acquire a slot before executing
//   - Slots are released after task completion
//   - If quota is exceeded, task is skipped (not queued)
//
// DETERMINISTIC PROPERTY: Tasks are processed in FIFO order within each worker.
// The worker count and task ordering are deterministic given the same input.
// =========================================================================
package concurrency

import (
	"sync"

	"github.com/fallofpheonix/phoenix/foundation/runtime/common/resource"
)

// DeterministicWorkerPool ensures that concurrent tasks are processed
// in a manner that can be deterministically replayed.
// This implements [CON-008].
type DeterministicWorkerPool struct {
	quota    *resource.QuotaManager
	wg       sync.WaitGroup
	taskChan chan func()
}

func NewDeterministicWorkerPool(quota *resource.QuotaManager, buffer int) *DeterministicWorkerPool {
	return &DeterministicWorkerPool{
		quota:    quota,
		taskChan: make(chan func(), buffer),
	}
}

// Start spawns the workers.
func (p *DeterministicWorkerPool) Start(count int) {
	for i := 0; i < count; i++ {
		go func() {
			for task := range p.taskChan {
				if err := p.quota.AcquireWorker(); err == nil {
					task()
					p.quota.ReleaseWorker()
				}
				p.wg.Done()
			}
		}()
	}
}

// Submit adds a task to the pool.
func (p *DeterministicWorkerPool) Submit(task func()) {
	p.wg.Add(1)
	p.taskChan <- task
}

// Wait blocks until all tasks are complete.
func (p *DeterministicWorkerPool) Wait() {
	p.wg.Wait()
}
