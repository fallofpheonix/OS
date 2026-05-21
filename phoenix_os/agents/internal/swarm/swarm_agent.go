package main

import (
	"fmt"
	"time"
	"phoenix/agents/internal/common"
)

func main() {
	fmt.Println("Phoenix Swarm Coordinator Online")
	
	// Simulate consensus check
	cmd := common.InternalCommand{
		Source:    "swarm-coord",
		Command:   "SYNC_NODES",
		Params:    map[string]int{"nodes": 5},
		Timestamp: time.Now(),
	}
	common.DispatchCommand(cmd)
}
