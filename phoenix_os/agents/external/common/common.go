package common

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type AgentTask struct {
	AgentID   string    `json:"agent_id"`
	Action    string    `json:"action"`
	Target    string    `json:"target"`
	Timestamp time.Time `json:"timestamp"`
}

func LogAction(task AgentTask, logPath string) error {
	f, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer f.Close()

	data, _ := json.Marshal(task)
	fmt.Fprintln(f, string(data))
	return nil
}
