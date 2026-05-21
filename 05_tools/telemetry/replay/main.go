package main

import (
	"bufio"
	"crypto/sha256"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"phoenix/telemetry/events"
	"phoenix/telemetry/process_lineage"
)

type ReplayResult struct {
	TotalEvents int    `json:"total_events"`
	GraphSize   int    `json:"graph_size"`
	Duration    string `json:"duration"`
	GraphHash   string `json:"graph_hash"`
}

func main() {
	inputFile := flag.String("input", "../../../09_telemetry/replay_events_large.jsonl", "Path to JSONL events file")
	outputFile := flag.String("output", "replay_result.json", "Path to output result file")
	flag.Parse()

	log.Printf("Starting replay of %s", *inputFile)

	file, err := os.Open(*inputFile)
	if err != nil {
		log.Fatalf("Failed to open input: %v", err)
	}
	defer file.Close()

	graph := lineage.NewLineageGraph()
	reader := bufio.NewReader(file)
	count := 0
	start := time.Now()

	for {
		line, err := reader.ReadBytes('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("Error reading line: %v", err)
		}

		var evt events.Event
		if err := json.Unmarshal(line, &evt); err != nil {
			log.Printf("Skip invalid line: %v", err)
			continue
		}

		// Process event types relevant to lineage
		switch evt.EventType {
		case "execve", "fork":
			graph.AddProcess(evt.PID, evt.PPID, evt.Comm, evt.ExePath, evt.Timestamp)
		case "exit":
			graph.ExitProcess(evt.PID, evt.Timestamp)
		}

		count++
		if count%10000 == 0 {
			fmt.Printf("\rProcessed %d events...", count)
		}
	}
	fmt.Println()

	elapsed := time.Since(start)
	
	// Calculate a deterministic hash of the final graph state for verification
	hash := calculateGraphHash(graph)

	result := ReplayResult{
		TotalEvents: count,
		GraphSize:   graph.Size(),
		Duration:    elapsed.String(),
		GraphHash:   hash,
	}

	resData, _ := json.MarshalIndent(result, "", "  ")
	if err := os.WriteFile(*outputFile, resData, 0644); err != nil {
		log.Fatalf("Failed to write result: %v", err)
	}

	fmt.Printf("Replay complete. Processed %d events in %v. Graph Size: %d\n", count, elapsed, graph.Size())
	fmt.Printf("Graph Hash: %s\n", hash)
}

func calculateGraphHash(g *lineage.LineageGraph) string {
	// For determinism, we'd need to sort nodes by PID and serialize.
	// This is a simplified version.
	data, _ := json.Marshal(g.Nodes)
	h := sha256.New()
	h.Write(data)
	return fmt.Sprintf("%x", h.Sum(nil))
}
