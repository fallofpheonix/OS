package logs

import (
	"bufio"
	"encoding/json"
	"os"
)

type LogEntry struct {
	Timestamp int64  `json:"timestamp"`
	Module    string `json:"module"`
	Level     string `json:"level"`
	Message   string `json:"message"`
	Drift     float64 `json:"drift,omitempty"`
}

func IngestJSONL(path string) ([]LogEntry, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var entries []LogEntry
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var entry LogEntry
		if err := json.Unmarshal(scanner.Bytes(), &entry); err == nil {
			entries = append(entries, entry)
		}
	}
	return entries, scanner.Err()
}
