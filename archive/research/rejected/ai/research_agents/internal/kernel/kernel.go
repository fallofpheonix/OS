package kernel

import (
	"errors"
	"fmt"
	"sync"
)

type KernelAgent interface {
	IsLocked() bool
	RegisterLSMHook(hookName string) error
	ApplySchedulerOverride(pid uint32, priority int) error
	ActuateContainmentPolicy(pid uint32, action string) error
}

type Agent struct {
	mu     sync.RWMutex
	locked bool
}

func NewKernelAgent() *Agent {
	// Security Locked by default until Telemetry, Graph, and Validation exist and are verified.
	return &Agent{
		locked: true,
	}
}

func (a *Agent) IsLocked() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.locked
}

// Unlock can be called only when the validation pipeline has successfully run and verified the userspace state.
func (a *Agent) Unlock() {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.locked = false
}

func (a *Agent) RegisterLSMHook(hookName string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.locked {
		return errors.New("security constraint violation: Kernel Agent is locked (LSM hook registry disabled)")
	}

	// Simulated register
	fmt.Printf("[KERNEL] Registered LSM Hook: %s\n", hookName)
	return nil
}

func (a *Agent) ApplySchedulerOverride(pid uint32, priority int) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.locked {
		return errors.New("security constraint violation: Kernel Agent is locked (Scheduler override disabled)")
	}

	// Simulated override
	fmt.Printf("[KERNEL] Applied scheduler priority %d to PID %d\n", priority, pid)
	return nil
}

func (a *Agent) ActuateContainmentPolicy(pid uint32, action string) error {
	a.mu.RLock()
	defer a.mu.RUnlock()

	if a.locked {
		return errors.New("security constraint violation: Kernel Agent is locked (Containment policy actuation disabled)")
	}

	// Simulated actuation
	fmt.Printf("[KERNEL] Actuated containment policy '%s' on PID %d\n", action, pid)
	return nil
}
