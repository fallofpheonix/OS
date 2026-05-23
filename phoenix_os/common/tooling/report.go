package tooling

import (
	"encoding/json"
	"fmt"
	"os"
	"phoenix/ledger/src"
	"phoenix/replay"
)

// ReproducibilityReport holds the results of a replay validation run.
// This implements [TLG-010].
type ReproducibilityReport struct {
	AuthoritativeHash string  `json:"authoritative_hash"`
	EventCount        int     `json:"event_count"`
	DivergenceIndex   int     `json:"divergence_index"`
	Status            string  `json:"status"`
	DurationMs        int64   `json:"duration_ms"`
}

// GenerateReport creates a JSON report comparing two replay runs.
func GenerateReport(idxA, idxB *replay.ReplayIndex, ledger *ledger.Ledger, duration int64) error {
	div, err := replay.Diff(idxA, idxB)
	
	status := "PASS"
	if err != nil {
		status = "FAIL"
	}

	report := ReproducibilityReport{
		AuthoritativeHash: fmt.Sprintf("%x", ledger.Heads[0]),
		EventCount:        len(idxA.Events),
		DivergenceIndex:   div,
		Status:            status,
		DurationMs:        duration,
	}

	data, _ := json.MarshalIndent(report, "", "  ")
	fmt.Println("\n--- REPRODUCIBILITY REPORT ---")
	fmt.Println(string(data))
	
	return os.WriteFile("artifacts/reproducibility_report.json", data, 0644)
}
