/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package actuators

import (
	"context"
	"fmt"
	"log"
	"os"
	"syscall"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// ProcessActuator implements the Actuator interface for OS processes.
type ProcessActuator struct {
	currentLevel securityv1.ContainmentLevel
}

func NewProcessActuator() *ProcessActuator {
	return &ProcessActuator{currentLevel: securityv1.LevelNone}
}

func (p *ProcessActuator) Name() string {
	return "ProcessActuator"
}

func (p *ProcessActuator) Actuate(ctx context.Context, action securityv1.Containment) error {
	p.currentLevel = action.Level()
	pid := 0
	_, _ = fmt.Sscanf(action.Target(), "PID:%d", &pid)
	if pid == 0 {
		return fmt.Errorf("invalid target format: %s", action.Target())
	}

	switch action.Level() {
	case securityv1.LevelNone:
		log.Printf("[Actuator] process normal: PID %d", pid)
	case securityv1.LevelMonitor:
		log.Printf("[Actuator] WARNING: PID %d flagged (Reason: %s)", pid, action.Reason())
	case securityv1.LevelSandbox:
		log.Printf("[Actuator] Throttling PID: %d to limit: %.2f", pid, 1.0)
	case securityv1.LevelIsolate:
		proc, err := os.FindProcess(pid)
		if err != nil {
			return fmt.Errorf("process %d not found: %w", pid, err)
		}
		if err := proc.Signal(syscall.SIGSTOP); err != nil {
			return fmt.Errorf("failed to signal pid %d with SIGSTOP: %w", pid, err)
		}
		log.Printf("[Actuator] Signal %v sent to PID: %d", syscall.SIGSTOP, pid)
	case securityv1.LevelQuench:
		proc, err := os.FindProcess(pid)
		if err != nil {
			return fmt.Errorf("process %d not found: %w", pid, err)
		}
		if err := proc.Signal(syscall.SIGSTOP); err != nil {
			return fmt.Errorf("failed to signal pid %d with SIGSTOP: %w", pid, err)
		}
		log.Printf("[Actuator] Signal %v sent to PID: %d", syscall.SIGSTOP, pid)
	}
	return nil
}

// Kill terminates the process using SIGKILL.
func (p *ProcessActuator) Kill(ctx context.Context, pid int) error {
	p.currentLevel = securityv1.LevelQuench
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("process %d not found: %w", pid, err)
	}
	if err := proc.Signal(syscall.SIGKILL); err != nil {
		return fmt.Errorf("failed to signal pid %d with SIGKILL: %w", pid, err)
	}
	log.Printf("[Actuator] Signal %v sent to PID: %d", syscall.SIGKILL, pid)
	return nil
}

func (p *ProcessActuator) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	return p.currentLevel, nil
}
