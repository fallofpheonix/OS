package main

import (
	"fmt"
	"time"
	"phoenix/agents/internal/common"
)

func main() {
	fmt.Println("Phoenix Self-Healer Online")
	
	// Simulate Rollback
	cmd := common.InternalCommand{
		Source:    "self-healer",
		Command:   "RESTORE_SNAPSHOT",
		Params:    map[string]string{"target": "container-01", "version": "v1.0.2"},
		Timestamp: time.Now(),
	}
	common.DispatchCommand(cmd)
}
