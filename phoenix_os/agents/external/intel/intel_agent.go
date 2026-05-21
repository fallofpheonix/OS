package main

import (
	"fmt"
	"time"
	"phoenix/agents/external/common"
)

func main() {
	fmt.Println("Phoenix Threat Intel Agent Online")
	task := common.AgentTask{
		AgentID:   "intel-01",
		Action:    "FETCH_FEED",
		Target:    "07_security/feeds/mitre_vfs_attack.json",
		Timestamp: time.Now(),
	}
	common.LogAction(task, "intel_log.json")
	fmt.Printf("Updating Feed: %s\n", task.Target)
}
