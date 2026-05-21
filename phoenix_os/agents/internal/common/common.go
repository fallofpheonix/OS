package common

import (
	"encoding/json"
	"fmt"
	"time"
)

type InternalCommand struct {
	Source    string      `json:"source"`
	Command   string      `json:"command"`
	Params    interface{} `json:"params"`
	Timestamp time.Time   `json:"timestamp"`
}

func DispatchCommand(cmd InternalCommand) {
	data, _ := json.Marshal(cmd)
	fmt.Printf("[PHOENIX_CMD] %s\n", string(data))
}
