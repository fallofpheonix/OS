package main

import (
	"fmt"
	"time"
	"phoenix/agents/external/common"
)

func main() {
	fmt.Println("Phoenix Benchmark Agent Online")
	task := common.AgentTask{
		AgentID:   "bench-01",
		Action:    "RUN_REPLAY",
		Target:    "14_experiments/R001/results.json",
		Timestamp: time.Now(),
	}
	common.LogAction(task, "bench_log.json")
	fmt.Printf("Benchmarking: %s\n", task.Target)
}
