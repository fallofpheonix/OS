package main

import (
	"fmt"
	"time"
	"phoenix/agents/external/common"
)

func main() {
	fmt.Println("Phoenix Research Agent Online")
	task := common.AgentTask{
		AgentID:   "research-01",
		Action:    "INGEST_PAPER",
		Target:    "01_research/foundations/information_theory.md",
		Timestamp: time.Now(),
	}
	common.LogAction(task, "research_log.json")
	fmt.Printf("Researching: %s\n", task.Target)
}
