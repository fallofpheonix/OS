package forensics

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/agents/internal/types"
)

type ForensicsAgent interface {
	TriggerSnapshot(pid uint32) (types.SnapshotResult, error)
	VerifyHash(path string) (string, error)
	GetSnapshots() []types.SnapshotResult
}

type Agent struct {
	mu           sync.Mutex
	snapshots    []types.SnapshotResult
	snapshotDir  string
}

func NewForensicsAgent(snapshotDir string) (*Agent, error) {
	if snapshotDir == "" {
		snapshotDir = "/Users/fallofpheonix/os/agents/artifacts/forensics"
	}
	// Ensure snapshot directory exists
	if err := os.MkdirAll(snapshotDir, 0755); err != nil {
		return nil, fmt.Errorf("failed to create snapshot dir: %w", err)
	}

	return &Agent{
		snapshots:   make([]types.SnapshotResult, 0),
		snapshotDir: snapshotDir,
	}, nil
}

// TriggerSnapshot captures memory/process state simulation metadata and writes it to an immutable file
func (a *Agent) TriggerSnapshot(pid uint32) (types.SnapshotResult, error) {
	a.mu.Lock()
	defer a.mu.Unlock()

	snapshotID := fmt.Sprintf("snap-%d-%d", pid, time.Now().UnixNano())
	filename := fmt.Sprintf("snapshot_%s.json", snapshotID)
	snapshotPath := filepath.Join(a.snapshotDir, filename)

	// Simulated snapshot metadata content
	content := fmt.Sprintf(`{
	"snapshot_id": "%s",
	"pid": %d,
	"timestamp": "%s",
	"system": "PhoenixOS",
	"memory_map": [
		{"start": "0x00400000", "end": "0x00450000", "perm": "r-xp", "name": "app_bin"},
		{"start": "0x00650000", "end": "0x00660000", "perm": "rw-p", "name": "app_data"},
		{"start": "0x7ffeb000", "end": "0x7ffed000", "perm": "rwxp", "name": "stack"}
	],
	"open_files": [
		"/etc/passwd",
		"/var/log/syslog"
	]
}`, snapshotID, pid, time.Now().Format(time.RFC3339))

	// Write simulated snapshot to file
	if err := os.WriteFile(snapshotPath, []byte(content), 0644); err != nil {
		return types.SnapshotResult{}, fmt.Errorf("failed to write snapshot: %w", err)
	}

	// Calculate SHA-256 hash of the written snapshot
	hashVal, err := a.calculateHash(snapshotPath)
	if err != nil {
		return types.SnapshotResult{}, fmt.Errorf("failed to calculate snapshot hash: %w", err)
	}

	res := types.SnapshotResult{
		SnapshotID: snapshotID,
		PID:        pid,
		Hash:       hashVal,
		Path:       snapshotPath,
		Timestamp:  time.Now(),
	}

	a.snapshots = append(a.snapshots, res)
	return res, nil
}

// VerifyHash computes the SHA-256 hash of the file at path and returns it
func (a *Agent) VerifyHash(path string) (string, error) {
	return a.calculateHash(path)
}

func (a *Agent) GetSnapshots() []types.SnapshotResult {
	a.mu.Lock()
	defer a.mu.Unlock()
	res := make([]types.SnapshotResult, len(a.snapshots))
	copy(res, a.snapshots)
	return res
}

func (a *Agent) calculateHash(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	h := sha256.New()
	if _, err := io.Copy(h, f); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}
