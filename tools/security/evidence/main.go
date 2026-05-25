package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/fallofpheonix/phoenix-os/phoenix_os/truth_ledger/src"
)

func main() {
	l := ledger.NewLedger(nil)

	// Mocking workflow using V2 ledger
	payload1, _ := json.Marshal(map[string]interface{}{
		"trace_hash":    "trace-001",
		"sdi":           0.95,
		"policy_id":     "POL-001",
		"action":        "FREEZE",
		"result":        "SUCCESS",
		"confidence":    0.98,
		"replay_id":     "replay-d992fd",
		"experiment_id": "R022",
	})

	err := l.AddEntryV2("event-001", "cause-001", payload1, "trace-001", "NORMAL", "FREEZE", "POL-001")
	if err != nil {
		fmt.Printf("Failed to add entry 1: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Committed Evidence 1\n")

	payload2, _ := json.Marshal(map[string]interface{}{
		"trace_hash":    "trace-002",
		"sdi":           0.45,
		"policy_id":     "POL-002",
		"action":        "OBSERVE",
		"result":        "SUCCESS",
		"confidence":    0.90,
		"replay_id":     "replay-d992fd",
		"experiment_id": "R022",
	})

	err = l.AddEntryV2("event-002", "cause-002", payload2, "trace-002", "FREEZE", "OBSERVE", "POL-002")
	if err != nil {
		fmt.Printf("Failed to add entry 2: %v\n", err)
		os.Exit(1)
	}
	fmt.Printf("Committed Evidence 2\n")

	if err := l.Verify(); err == nil {
		fmt.Println("[PASS] Ledger integrity verified.")
	} else {
		fmt.Printf("[FAIL] Ledger integrity compromised: %v\n", err)
		os.Exit(1)
	}

	// Output ledger for inspection
	data, _ := json.MarshalIndent(l.SortedEntries(), "", "  ")
	fmt.Println(string(data))
}
