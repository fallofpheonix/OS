package main

import (
	"encoding/json"
	"fmt"
	"os"

	"phoenix/ledger/src"
)

func main() {
	l := ledger.NewLedger()

	// Mocking Stage 5 workflow: event -> trace -> decision -> action
	e1 := ledger.Evidence{
		TraceHash:    "trace-001",
		SDI:          0.95,
		PolicyID:     "POL-001",
		Action:       "FREEZE",
		Result:       "SUCCESS",
		Confidence:   0.98,
		ReplayID:     "replay-d992fd",
		ExperimentID: "R022",
	}

	hash1 := l.Commit(e1)
	fmt.Printf("Committed Evidence 1: %s\n", hash1)

	e2 := ledger.Evidence{
		TraceHash:    "trace-002",
		SDI:          0.45,
		PolicyID:     "POL-002",
		Action:       "OBSERVE",
		Result:       "SUCCESS",
		Confidence:   0.90,
		ReplayID:     "replay-d992fd",
		ExperimentID: "R022",
	}

	hash2 := l.Commit(e2)
	fmt.Printf("Committed Evidence 2: %s\n", hash2)

	if l.Verify() {
		fmt.Println("[PASS] Ledger integrity verified.")
	} else {
		fmt.Println("[FAIL] Ledger integrity compromised!")
		os.Exit(1)
	}

	// Output ledger for inspection
	data, _ := json.MarshalIndent(l.Entries, "", "  ")
	fmt.Println(string(data))
}
