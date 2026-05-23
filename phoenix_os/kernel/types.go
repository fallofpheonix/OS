package kernel

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

// KernelEvent matches the deliverable schema in 02_docs/05_kernel/EVENT_SCHEMA.md
type KernelEvent struct {
	EventID   string `json:"event_id"`
	PID       int    `json:"pid"`
	PPID      int    `json:"ppid"`
	Process   string `json:"process"`
	Syscall   string `json:"syscall"`
	Timestamp int64  `json:"timestamp"`
	CPU       int    `json:"cpu"`
}

func (e KernelEvent) ToJSON() []byte {
	b, _ := json.Marshal(e)
	return b
}

// MockGenerator simulates kernel telemetry for non-Linux environments (macOS)
type MockGenerator struct {
	running bool
}

func NewMockGenerator() *MockGenerator {
	return &MockGenerator{}
}

func (m *MockGenerator) Generate() KernelEvent {
	syscalls := []string{"execve", "fork", "open", "read", "write", "connect", "exit"}
	processes := []string{"bash", "curl", "python3", "nginx", "sudo", "apt"}
	
	pid := rand.Intn(30000) + 1000
	
	return KernelEvent{
		EventID:   fmt.Sprintf("mock-%d", time.Now().UnixNano()),
		PID:       pid,
		PPID:      1,
		Process:   processes[rand.Intn(len(processes))],
		Syscall:   syscalls[rand.Intn(len(syscalls))],
		Timestamp: time.Now().UnixNano(),
		CPU:       rand.Intn(8),
	}
}
