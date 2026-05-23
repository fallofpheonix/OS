package telemetry

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"phoenix/agents/internal/types"
)

type TelemetryAgent interface {
	Start() error
	Stop() error
	RecordEvent(ev types.TelemetryEvent)
	GetLineage(pid uint32) ([]types.TelemetryEvent, error)
	GenerateMockEvent() types.TelemetryEvent
}

type Agent struct {
	mu       sync.Mutex
	running  bool
	history  []types.TelemetryEvent
	historyLimit int
}

func NewTelemetryAgent(limit int) *Agent {
	if limit <= 0 {
		limit = 1000
	}
	return &Agent{
		history:      make([]types.TelemetryEvent, 0),
		historyLimit: limit,
	}
}

func (a *Agent) Start() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.running = true
	return nil
}

func (a *Agent) Stop() error {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.running = false
	return nil
}

func (a *Agent) GetLineage(pid uint32) ([]types.TelemetryEvent, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	var lineage []types.TelemetryEvent
	currentPID := pid
	visited := make(map[uint32]bool)

	// Keep searching until we hit pid 0 or a circular ref or max depth
	for currentPID > 0 && len(lineage) < 50 && !visited[currentPID] {
		visited[currentPID] = true
		found := false
		for i := len(a.history) - 1; i >= 0; i-- {
			ev := a.history[i]
			if ev.PID == currentPID {
				lineage = append(lineage, ev)
				currentPID = ev.PPID
				found = true
				break
			}
		}
		if !found {
			break
		}
	}
	return lineage, nil
}

func (a *Agent) RecordEvent(ev types.TelemetryEvent) {
	a.mu.Lock()
	defer a.mu.Unlock()

	if len(a.history) >= a.historyLimit {
		// Evict oldest (ring-buffer logic)
		a.history = a.history[1:]
	}
	a.history = append(a.history, ev)
}

func (a *Agent) GenerateMockEvent() types.TelemetryEvent {
	categories := []string{"process", "syscall", "filesystem", "network"}
	cat := categories[rand.Intn(len(categories))]

	pid := uint32(rand.Intn(10000) + 1000)
	ppid := uint32(rand.Intn(1000) + 1)
	comm := []string{"bash", "python3", "nginx", "curl", "malware"}[rand.Intn(5)]

	ev := types.TelemetryEvent{
		Timestamp: time.Now(),
		EventID:   fmt.Sprintf("evt-%d", rand.Int63()),
		Category:  cat,
		EventType: "mock." + cat,
		HostID:    "phoenix-dev-node",
		PID:       pid,
		PPID:      ppid,
		UID:       1000,
		GID:       1000,
		Comm:      comm,
		ExePath:   "/usr/bin/" + comm,
	}

	switch cat {
	case "process":
		ev.Process = &types.ProcessPayload{
			Args:    []string{"-v", "run"},
			EnvVars: []string{"PATH=/usr/bin"},
		}
	case "syscall":
		ev.Syscall = &types.SyscallPayload{
			SyscallNr: uint64(rand.Intn(300)),
			Args:      []uint64{0x0, 0x1, 0x2},
			RetVal:    0,
		}
	case "filesystem":
		ev.Filesystem = &types.FilesystemPayload{
			FilePath:         "/etc/passwd",
			Flags:            1,
			Mode:             0644,
			BytesRequested:   256,
			BytesTransferred: 256,
		}
	case "network":
		ev.Network = &types.NetworkPayload{
			SAddr:       "127.0.0.1",
			DAddr:       "8.8.8.8",
			SPort:       uint16(rand.Intn(60000) + 1024),
			DPort:       443,
			Protocol:    "TCP",
			StateChange: "ESTABLISHED",
		}
	}

	a.RecordEvent(ev)
	return ev
}
