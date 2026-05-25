package observations

import (
	"bufio"
	"encoding/json"
	"os"
)

// Observation represents the standardized structure for an observability cycle.
type Observation struct {
	Cycle            string  `json:"cycle"`
	RuntimeHealth    float64 `json:"runtime_health"`
	AverageDrift     float64 `json:"average_drift"`
	CriticalFailures int     `json:"critical_failures"`
	Coverage         float64 `json:"coverage"`
	Status           string  `json:"status"`
}

type RuntimeEntry struct {
	Timestamp string `json:"timestamp"`
	Module    string `json:"module"`
	Status    string `json:"status"`
	Error     string `json:"error"`
}

type ModuleStatus struct {
	Arbiter  string `json:"arbiter"`
	LastScan string `json:"last_scan"`
}

// Ingest reads input data and generates an Observation record.
func Ingest(cycle string) (*Observation, error) {
	// 1. Load data
	auditEntries, _ := LoadRuntimeAudit("phoenix_os/logs/runtime_audit.jsonl")
	moduleStatus, _ := LoadModuleStatus("PhoenixMind-Org/phoenixmind-core/reality/MODULE_STATUS.json")

	// 2. Compute metrics
	health := ComputeHealth(auditEntries, moduleStatus)
	drift := 0.11 // TODO: Need drift.ComputeDrift(history, baseline)
	coverage := ComputeCoverage(auditEntries)
	failures := ComputeCriticalFailures(auditEntries)
	
	return &Observation{
		Cycle:            cycle,
		RuntimeHealth:    health,
		AverageDrift:     drift,
		CriticalFailures: failures,
		Coverage:         coverage,
		Status:           "STABLE",
	}, nil
}

func LoadRuntimeAudit(path string) ([]RuntimeEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []RuntimeEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry RuntimeEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err == nil {
			entries = append(entries, entry)
		}
	}
	return entries, nil
}

func LoadModuleStatus(path string) (*ModuleStatus, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}
	var status ModuleStatus
	if err := json.Unmarshal(data, &status); err != nil {
		return nil, err
	}
	return &status, nil
}

func ComputeHealth(audit []RuntimeEntry, status *ModuleStatus) float64 {
	if status == nil {
		return 0.0
	}
	if status.Arbiter == "RUNNING" {
		return 1.0
	}
	return 0.5
}

// ComputeCoverage calculates observed vs expected modules.
func ComputeCoverage(audit []RuntimeEntry) float64 {
	uniqueModules := make(map[string]bool)
	for _, entry := range audit {
		uniqueModules[entry.Module] = true
	}
	
	observed := float64(len(uniqueModules))
	expected := 8.0 // Based on RUNTIME_HEALTH.json
	
	if expected == 0 {
		return 0.0
	}
	return observed / expected
}

func ComputeCriticalFailures(audit []RuntimeEntry) int {
	count := 0
	for _, entry := range audit {
		if entry.Status == "FAILED" || entry.Error != "None" || entry.Status == "MISSING" || entry.Status == "UNKNOWN" || entry.Status == "ESCALATED" {
			count++
		}
	}
	return count
}
