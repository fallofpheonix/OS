package concurrency

import (
	"phoenix/common/resource"
	"sync"
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
