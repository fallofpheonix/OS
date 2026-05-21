package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Phoenix Trace: 3-Tier Storage Engine starting...")
	ts := NewTraceStorage()

	// 1. Create active processes (HOT)
	ts.AddProcess(1, 0, "init")
	ts.AddProcess(101, 1, "sh")
	ts.AddProcess(102, 101, "ls")
	ts.Stats()

	// 2. Terminate short-lived process (Move to WARM)
	fmt.Println("\nProcess 'ls' terminating...")
	ts.TerminateProcess(102)
	ts.Stats()

	// 3. Trigger Lifecycle Pruning (Move to COLD)
	fmt.Println("\nTriggering Lifecycle Pruning (WARM -> COLD)...")
	ts.PruneToCold()
	ts.Stats()

	// 4. Verify Critical Node Retention
	fmt.Println("\nVerifying 'init' remains active (HOT)...")
	ts.Stats()
	
	time.Sleep(100 * time.Millisecond)
}
