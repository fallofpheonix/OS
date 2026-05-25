package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Event struct {
	Timestamp time.Time      `json:"timestamp"`
	EventID   string         `json:"event_id"`
	Category  string         `json:"category"`
	EventType string         `json:"event_type"`
	HostID    string         `json:"host_id"`
	PID       uint32         `json:"pid"`
	PPID      uint32         `json:"ppid"`
	Comm      string         `json:"comm"`
	Payload   map[string]any `json:"payload"`
}

func main() {
	// 1. Hospital Scenario (Ransomware Rename Storm)
	generateHospital()
}

func generateHospital() {
	f, _ := os.Create("hospital.jsonl")
	defer f.Close()

	now := time.Now()
	pid := uint32(5000)

	// Benign activity
	for i := 0; i < 100; i++ {
		e := Event{
			Timestamp: now.Add(time.Duration(i) * time.Millisecond),
			EventID:   fmt.Sprintf("hosp-benign-%d", i),
			Category:  "filesystem",
			EventType: "write",
			PID:       100,
			Comm:      "syslog",
			Payload:   map[string]any{"entropy_score": 3.2},
		}
		data, _ := json.Marshal(e)
		f.Write(append(data, '\n'))
	}

	// Attack starts
	for i := 0; i < 50; i++ {
		e := Event{
			Timestamp: now.Add(time.Duration(100+i) * time.Millisecond),
			EventID:   fmt.Sprintf("hosp-attack-%d", i),
			Category:  "filesystem",
			EventType: "rename",
			PID:       pid,
			Comm:      "wanadecrypt",
			Payload:   map[string]any{"entropy_score": 8.9},
		}
		data, _ := json.Marshal(e)
		f.Write(append(data, '\n'))
	}
}
