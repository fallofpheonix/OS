/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package causal

import (
	"fmt"
	"testing"
	"time"
)

func TestCausalLineageSimulator(t *testing.T) {
	e := NewEnforcer()
	startTime := time.Now().Add(-24 * time.Hour)

	fmt.Println("--- Starting Causal Lineage Simulation ---")

	// Create a chain of events
	prevID := ""
	for i := 0; i < 5; i++ {
		id := fmt.Sprintf("event-%d", i)
		ts := startTime.Add(time.Duration(i) * time.Hour)
		parents := []string{}
		if prevID != "" {
			parents = append(parents, prevID)
		}

		node := &Node{
			ID:        id,
			Timestamp: ts,
			Data:      fmt.Sprintf("Matrix event %d data", i),
			Parents:   parents,
		}

		err := e.AddNode(node)
		if err != nil {
			t.Fatalf("Simulation failed at step %d: %v", i, err)
		}
		fmt.Printf("Sealed Node: %s | Time: %v | Hash: %s\n", node.ID, node.Timestamp.Format(time.RFC3339), node.Hash)
		prevID = id
	}

	// Attempt a temporal anomaly (backwards in time)
	anomaly := &Node{
		ID:        "anomaly-1",
		Timestamp: startTime, // Way before its parent's timestamp
		Parents:   []string{"event-4"},
	}
	err := e.AddNode(anomaly)
	if err == nil {
		t.Error("SIMULATION FAILURE: Matrix failed to reject temporal anomaly")
	} else {
		fmt.Printf("Matrix Protected: Successfully rejected temporal anomaly: %v\n", err)
	}

	// Attempt a cyclic anomaly
	// Since event-0 is a parent of event-1, etc.,
	// we try to make event-0 a child of event-4.
	// This will be caught by temporal validation first.
	cyclic := &Node{
		ID:        "event-0", // Existing ID
		Timestamp: time.Now(),
		Parents:   []string{"event-4"},
	}
	err = e.AddNode(cyclic)
	if err == nil {
		t.Error("SIMULATION FAILURE: Matrix failed to reject duplicate/cyclic node ID")
	} else {
		fmt.Printf("Matrix Protected: Successfully rejected duplicate/cyclic ID: %v\n", err)
	}

	fmt.Println("--- Causal Lineage Simulation Complete ---")
}
