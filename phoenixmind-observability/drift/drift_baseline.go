package drift

import "sync"

type Baseline struct {
	mu      sync.RWMutex
	Modules map[string]float64
}

func (b *Baseline) Set(module string, val float64) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.Modules[module] = val
}

func (b *Baseline) Get(module string) float64 {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.Modules[module]
}
