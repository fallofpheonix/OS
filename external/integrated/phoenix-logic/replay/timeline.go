package replay

import (
	"encoding/json"
	"fmt"
	"github.com/fallofpheonix/phoenix-os/phoenix_os/bus"
	"sort"
)

// ProcessNode represents a single process in the causal timeline.
type ProcessNode struct {
	PID       int
	Comm      string
	Args      []string
	ParentPID int
	Events    []bus.TelemetryEvent
	Children  []*ProcessNode
}

// Timeline represents the reconstructed causal graph of all system activity.
type Timeline struct {
	RootProcesses []*ProcessNode
	ProcessMap    map[int]*ProcessNode
}

// PayloadProcessExec represents the expected payload for process.exec events.
type PayloadProcessExec struct {
	PPID int      `json:"ppid"`
	Comm string   `json:"comm"`
	Args []string `json:"args"`
}

// ReconstructTimeline builds a causal graph from a slice of events.
func ReconstructTimeline(events []bus.TelemetryEvent) (*Timeline, error) {
	tl := &Timeline{
		ProcessMap: make(map[int]*ProcessNode),
	}

	// Sort events by monotonic time to ensure we process them in order
	sort.Slice(events, func(i, j int) bool {
		return events[i].MonotonicNs < events[j].MonotonicNs
	})

	for _, ev := range events {
		node, ok := tl.ProcessMap[ev.PID]
		if !ok {
			node = &ProcessNode{
				PID:    ev.PID,
				Events: []bus.TelemetryEvent{},
			}
			tl.ProcessMap[ev.PID] = node
		}

		node.Events = append(node.Events, ev)

		// Handle process start/exec to establish parentage
		if ev.EventType == "process.exec" || ev.EventType == "process.fork" {
			var payload PayloadProcessExec
			if err := json.Unmarshal(ev.Payload, &payload); err == nil {
				node.Comm = payload.Comm
				node.Args = payload.Args
				node.ParentPID = payload.PPID

				// Link to parent if it exists
				if parent, ok := tl.ProcessMap[payload.PPID]; ok {
					parent.Children = append(parent.Children, node)
				}
			}
		}
	}

	// Identify root processes (those whose parents we haven't seen or PPID=1/0)
	for _, node := range tl.ProcessMap {
		isRoot := true
		if node.ParentPID > 1 {
			if _, ok := tl.ProcessMap[node.ParentPID]; ok {
				isRoot = false
			}
		}
		if isRoot {
			tl.RootProcesses = append(tl.RootProcesses, node)
		}
	}

	return tl, nil
}

// PrintTimeline prints a text representation of the causal graph.
func (tl *Timeline) PrintTimeline() {
	for _, root := range tl.RootProcesses {
		tl.printNode(root, 0)
	}
}

func (tl *Timeline) printNode(node *ProcessNode, indent int) {
	prefix := ""
	for i := 0; i < indent; i++ {
		prefix += "  "
	}
	fmt.Printf("%s[PID %d] %s (Events: %d)\n", prefix, node.PID, node.Comm, len(node.Events))
	for _, ev := range node.Events {
		fmt.Printf("%s  - %s: %s\n", prefix, ev.EventType, string(ev.Payload))
	}
	for _, child := range node.Children {
		tl.printNode(child, indent+1)
	}
}
