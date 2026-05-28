package main

import (
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/ai"
	"github.com/fallofpheonix/PheonixTruth/src"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/telemetry/process_graphs"
)
func main() {
	fmt.Println("Starting DAG Re-genesis...")
	
	// This would normally be connected to the running runtime.
	// For this script, we'll simulate the graph and ledger connection.
	
	// 1. Initialize
	l := ledger.NewLedger(nil)
	gf := &ai.GraphFeature{Graph: process_graphs.NewGraph()}
	
	// 2. Suspend
	gf.SetPaused(true)
	
	// 3. Rebuild
	gf.RebuildFromLedger(l)
	fmt.Println("DAG Rebuilt successfully.")
	
	// 4. Resume
	gf.SetPaused(false)
	fmt.Println("DAG Re-genesis complete.")
}
