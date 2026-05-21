package sandbox

import (
	"errors"
	"fmt"
)

// KernelSimulator mocks eBPF and kernel runtime constraints.
type KernelSimulator struct {
	MaxMapEntries      int
	CurrentEntries     int
	StackDepth         int
	MaxStackDepth      int
	GlobalEnergyBudget float64
	ConsumedEnergy     float64
}

// NewKernelSimulator initializes a mock kernel environment.
func NewKernelSimulator() *KernelSimulator {
	return &KernelSimulator{
		MaxMapEntries:      1024,
		MaxStackDepth:      512, // eBPF limit
		GlobalEnergyBudget: 1000.0, // Default energy budget
		ConsumedEnergy:     0.0,
	}
}

// RequestEnergy attempts to consume energy from the global budget.
func (k *KernelSimulator) RequestEnergy(amount float64) error {
	if k.ConsumedEnergy+amount > k.GlobalEnergyBudget {
		return fmt.Errorf("energy budget exceeded: requested %.2f, available %.2f", amount, k.GlobalEnergyBudget-k.ConsumedEnergy)
	}
	k.ConsumedEnergy += amount
	return nil
}

// UpdateMap simulates writing to an eBPF map.
func (k *KernelSimulator) UpdateMap(key string, value interface{}) error {
	if k.CurrentEntries >= k.MaxMapEntries {
		return errors.New("eBPF map limit reached (memory exhaustion)")
	}
	k.CurrentEntries++
	return nil
}

// CheckStackDepth simulates eBPF verifier stack depth checks.
func (k *KernelSimulator) CheckStackDepth(depth int) error {
	if depth > k.MaxStackDepth {
		return fmt.Errorf("eBPF verifier error: stack depth %d exceeds limit %d", depth, k.MaxStackDepth)
	}
	k.StackDepth = depth
	return nil
}

// Panic simulates a kernel panic.
func (k *KernelSimulator) Panic(reason string) {
	fmt.Printf("!!! KERNEL PANIC: %s !!!\n", reason)
}
