package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net"
	"os"
	"sync"
	"syscall"
	"time"

	"sentinel/telemetry/bus/normalizer"
)

// --- RFC-006 Process Graph Implementation ---

type ProcessNode struct {
	PID          uint32     `json:"pid"`
	PPID         uint32     `json:"ppid"`
	OriginalPPID uint32     `json:"original_ppid"`
	Comm         string     `json:"comm"`
	ExePath      string     `json:"exe_path"`
	Args         []string   `json:"args"`
	UID          uint32     `json:"uid"`
	GID          uint32     `json:"gid"`
	StartTime    time.Time  `json:"start_time"`
	ExitTime     *time.Time `json:"exit_time,omitempty"`
	ExitCode     int32      `json:"exit_code,omitempty"`
	IsActive     bool       `json:"is_active"`
	Children     []uint32   `json:"children"`
}

type ProcessGraph struct {
	mu    sync.RWMutex
	Nodes map[uint32]*ProcessNode
}

func NewProcessGraph() *ProcessGraph {
	return &ProcessGraph{
		Nodes: make(map[uint32]*ProcessNode),
	}
}

func (g *ProcessGraph) Update(evt *normalizer.Event) {
	g.mu.Lock()
	defer g.mu.Unlock()

	switch evt.EventType {
	case "fork":
		childPID := evt.PID
		ppid := evt.PPID

		node := &ProcessNode{
			PID:          childPID,
			PPID:         ppid,
			OriginalPPID: ppid,
			Comm:         evt.Comm,
			ExePath:      evt.ExePath,
			UID:          evt.UID,
			GID:          evt.GID,
			StartTime:    evt.Timestamp,
			IsActive:     true,
			Children:     []uint32{},
		}
		g.Nodes[childPID] = node

		if parent, exists := g.Nodes[ppid]; exists {
			parent.Children = append(parent.Children, childPID)
		}

	case "execve":
		pid := evt.PID
		node, exists := g.Nodes[pid]
		if !exists {
			node = &ProcessNode{
				PID:          pid,
				PPID:         evt.PPID,
				OriginalPPID: evt.PPID,
				StartTime:    evt.Timestamp,
				IsActive:     true,
				Children:     []uint32{},
			}
			g.Nodes[pid] = node
		}

		node.Comm = evt.Comm
		node.ExePath = evt.ExePath
		node.UID = evt.UID
		node.GID = evt.GID
		if argsVal, ok := evt.Payload["args"]; ok {
			if argsSlice, ok := argsVal.([]interface{}); ok {
				node.Args = make([]string, len(argsSlice))
				for i, v := range argsSlice {
					if s, ok := v.(string); ok {
						node.Args[i] = s
					}
				}
			}
		}

	case "exit":
		pid := evt.PID
		if node, exists := g.Nodes[pid]; exists {
			node.IsActive = false
			now := evt.Timestamp
			node.ExitTime = &now
			if codeVal, ok := evt.Payload["exit_code"]; ok {
				if f, ok := codeVal.(float64); ok {
					node.ExitCode = int32(f)
				}
			}
		}
	}
}

// --- RFC-008 Simulation Telemetry Agent (Generator) ---

type MockGenerator struct {
	hostID string
	rng    *rand.Rand
}

func NewMockGenerator(hostID string) *MockGenerator {
	return &MockGenerator{
		hostID: hostID,
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (m *MockGenerator) GenerateEventSequence(seqID string, count int) [][]byte {
	events := make([][]byte, 0, count)
	pids := make([]uint32, count)
	for i := 0; i < count; i++ {
		pids[i] = uint32(1000 + m.rng.Intn(90000))
	}

	for i := 0; i < count; i++ {
		evtType := "execve"
		category := "process"
		
		roll := m.rng.Float32()
		var payload map[string]interface{}
		pid := pids[i]
		ppid := pids[m.rng.Intn(count)]
		if ppid == pid {
			ppid = 500
		}
		comm := "sh"
		exePath := "/bin/sh"

		if roll < 0.60 {
			subRoll := m.rng.Float32()
			if subRoll < 0.33 {
				evtType = "fork"
				payload = map[string]interface{}{
					"args":     []string{},
					"env_vars": []string{"PATH=/usr/bin"},
				}
			} else if subRoll < 0.66 {
				evtType = "execve"
				comm = "gpg"
				exePath = "/usr/bin/gpg"
				payload = map[string]interface{}{
					"args":     []string{"--encrypt", "doc.txt"},
					"env_vars": []string{"PATH=/usr/bin"},
				}
			} else {
				evtType = "exit"
				exitCode := int32(0)
				payload = map[string]interface{}{
					"exit_code": exitCode,
				}
			}
		} else {
			category = "filesystem"
			evtType = "write"
			comm = "gpg"
			exePath = "/usr/bin/gpg"
			
			entropy := 4.2
			if m.rng.Float32() > 0.5 {
				entropy = 7.85
			}
			payload = map[string]interface{}{
				"file_path":         fmt.Sprintf("/home/user/document_%d.txt", m.rng.Intn(1000)),
				"flags":             578,
				"mode":              438,
				"bytes_requested":   4096,
				"bytes_transferred": 4096,
				"entropy_score":     entropy,
			}
		}

		rawEvt := map[string]interface{}{
			"timestamp":  time.Now().Format(time.RFC3339Nano),
			"event_id":   fmt.Sprintf("e0000000-0000-0000-0000-%012d", m.rng.Intn(1000000000)),
			"category":   category,
			"event_type": evtType,
			"host_id":    m.hostID,
			"pid":        pid,
			"ppid":       ppid,
			"uid":        501,
			"gid":        20,
			"comm":       comm,
			"exe_path":   exePath,
			"payload":    payload,
		}

		data, _ := json.Marshal(rawEvt)
		events = append(events, append(data, '\n'))
	}

	return events
}

// --- Benchmark & Server Harness ---

func runServer(socketPath string, graph *ProcessGraph, norm normalizer.Normalizer, wg *sync.WaitGroup, totalEvents int, processedChan chan<- int, droppedChan chan<- int) {
	defer wg.Done()

	_ = os.Remove(socketPath)

	listener, err := net.Listen("unix", socketPath)
	if err != nil {
		fmt.Printf("Server failed to listen on UDS: %v\n", err)
		return
	}
	defer listener.Close()
	defer os.Remove(socketPath)

	for {
		conn, err := listener.Accept()
		if err != nil {
			return
		}

		go func(c net.Conn) {
			defer c.Close()
			reader := bufio.NewReaderSize(c, 65536)
			localProcessed := 0
			localDropped := 0

			for {
				line, err := reader.ReadBytes('\n')
				if err != nil {
					if err != io.EOF {
						fmt.Printf("Error reading from UDS: %v\n", err)
					}
					break
				}

				if len(line) == 0 {
					continue
				}

				evt, err := norm.Normalize(line)
				if err != nil {
					localDropped++
					continue
				}

				if err := norm.Enrich(evt); err != nil {
					localDropped++
					continue
				}

				graph.Update(evt)

				localProcessed++
				if localProcessed+localDropped >= totalEvents {
					break
				}
			}

			processedChan <- localProcessed
			droppedChan <- localDropped
		}(conn)
	}
}

func getCPUTime() (time.Duration, error) {
	var rusage syscall.Rusage
	err := syscall.Getrusage(syscall.RUSAGE_SELF, &rusage)
	if err != nil {
		return 0, err
	}
	userTime := time.Duration(rusage.Utime.Sec)*time.Second + time.Duration(rusage.Utime.Usec)*time.Microsecond
	sysTime := time.Duration(rusage.Stime.Sec)*time.Second + time.Duration(rusage.Stime.Usec)*time.Microsecond
	return userTime + sysTime, nil
}

func main() {
	socketPath := "./sentinel_sim.sock"

	fmt.Println("==================================================")
	fmt.Println("   Pheonix Telemetry Simulation Benchmark       ")
	fmt.Println("==================================================")

	// Mode 1: Peak Throughput & Integrity Check (Unthrottled, 100k events)
	{
		totalEvents := 100000
		fmt.Printf("--- Mode 1: Peak Throughput & Integrity (Unthrottled %d events) ---\n", totalEvents)
		
		graph := NewProcessGraph()
		norm := normalizer.NewEventNormalizer()
		generator := NewMockGenerator("macos-dev-host")

		mockPayloads := generator.GenerateEventSequence("seq-1", totalEvents)

		processedChan := make(chan int, 1)
		droppedChan := make(chan int, 1)

		var wg sync.WaitGroup
		wg.Add(1)
		go runServer(socketPath, graph, norm, &wg, totalEvents, processedChan, droppedChan)

		time.Sleep(100 * time.Millisecond)

		conn, err := net.Dial("unix", socketPath)
		if err != nil {
			fmt.Printf("Client failed to dial UDS: %v\n", err)
			return
		}

		startWallTime := time.Now()
		startCPUTime, _ := getCPUTime()

		writer := bufio.NewWriterSize(conn, 65536)
		for _, payload := range mockPayloads {
			_, _ = writer.Write(payload)
		}
		_ = writer.Flush()
		_ = conn.Close()

		processed := <-processedChan
		dropped := <-droppedChan
		
		endWallTime := time.Now()
		endCPUTime, _ := getCPUTime()

		elapsedWall := endWallTime.Sub(startWallTime)
		elapsedCPU := endCPUTime - startCPUTime
		dropRate := (float64(dropped) / float64(totalEvents)) * 100.0
		throughput := float64(processed) / elapsedWall.Seconds()

		fmt.Printf("  Wall Time:     %v\n", elapsedWall)
		fmt.Printf("  CPU Time:      %v\n", elapsedCPU)
		fmt.Printf("  Throughput:    %.2f events/sec\n", throughput)
		fmt.Printf("  Dropped:       %d (%.2f%% drop rate)\n", dropped, dropRate)
		if dropRate < 2.0 {
			fmt.Println("  [PASS] Event drop rate within budget (< 2.0%)")
		} else {
			fmt.Printf("  [FAIL] Event drop rate exceeded budget (%.2f%%)\n", dropRate)
		}
		fmt.Println()
	}

	// Mode 2: Resource Overhead Verification (Throttled 1,000 events/sec target)
	// We run 3 reproducible iterations as requested.
	{
		totalEvents := 1000
		fmt.Printf("--- Mode 2: CPU Budget Verification (Throttled target 1,000 events/sec) ---\n")

		for run := 1; run <= 3; run++ {
			fmt.Printf("  Run %d of 3:\n", run)
			graph := NewProcessGraph()
			norm := normalizer.NewEventNormalizer()
			generator := NewMockGenerator("macos-dev-host")

			mockPayloads := generator.GenerateEventSequence("seq-2", totalEvents)

			processedChan := make(chan int, 1)
			droppedChan := make(chan int, 1)

			var wg sync.WaitGroup
			wg.Add(1)
			go runServer(socketPath, graph, norm, &wg, totalEvents, processedChan, droppedChan)

			time.Sleep(100 * time.Millisecond)

			conn, err := net.Dial("unix", socketPath)
			if err != nil {
				fmt.Printf("Client failed to dial UDS: %v\n", err)
				return
			}

			startWallTime := time.Now()
			startCPUTime, _ := getCPUTime()

			writer := bufio.NewWriterSize(conn, 65536)
			
			// We stream 10 batches of 100 events, sleeping 100ms between each, total 1.0 second duration
			batchSize := 100
			for i := 0; i < 10; i++ {
				for j := 0; j < batchSize; j++ {
					idx := i*batchSize + j
					_, _ = writer.Write(mockPayloads[idx])
				}
				_ = writer.Flush()
				time.Sleep(100 * time.Millisecond)
			}
			_ = conn.Close()

			processed := <-processedChan
			dropped := <-droppedChan

			endWallTime := time.Now()
			endCPUTime, _ := getCPUTime()

			elapsedWall := endWallTime.Sub(startWallTime)
			elapsedCPU := endCPUTime - startCPUTime
			cpuPercent := (float64(elapsedCPU) / float64(elapsedWall)) * 100.0
			dropRate := (float64(dropped) / float64(totalEvents)) * 100.0

			fmt.Printf("    Wall Time:   %v\n", elapsedWall)
			fmt.Printf("    CPU Time:    %v\n", elapsedCPU)
			fmt.Printf("    Avg CPU %%:   %.2f%%\n", cpuPercent)
			fmt.Printf("    Processed:   %d\n", processed)
			fmt.Printf("    Dropped:     %d (%.2f%% drop rate)\n", dropped, dropRate)

			if cpuPercent <= 5.0 {
				fmt.Println("    [PASS] CPU overhead within budget (< 5.0%)")
			} else {
				fmt.Printf("    [FAIL] CPU overhead exceeded budget (%.2f%% > 5.0%%)\n", cpuPercent)
			}
			if dropRate <= 2.0 {
				fmt.Println("    [PASS] Event drop rate within budget (< 2.0%)")
			} else {
				fmt.Printf("    [FAIL] Event drop rate exceeded budget (%.2f%%)\n", dropRate)
			}
			fmt.Println()
		}
	}
}
