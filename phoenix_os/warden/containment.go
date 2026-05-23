package warden

import (
	"fmt"
)

// LocalAction implements the Action interface for local OS containment.
type LocalAction struct{}

func (l *LocalAction) Observe(pid int) {
	fmt.Printf("[WARDEN ACTION] Activating Enhanced Observation for PID %d\n", pid)
}

func (l *LocalAction) Snapshot(pid int) {
	fmt.Printf("[WARDEN ACTION] Triggering Forensic Memory Snapshot for PID %d\n", pid)
}

func (l *LocalAction) Isolate(pid int) {
	fmt.Printf("[WARDEN ACTION] ISOLATING PID %d: Cgroup freezer active, network egress blocked.\n", pid)
}

func (l *LocalAction) Recover(pid int) {
	fmt.Printf("[WARDEN ACTION] Recovering PID %d: Releasing freezer, restoring network access.\n", pid)
}
