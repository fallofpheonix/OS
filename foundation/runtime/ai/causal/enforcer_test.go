/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package causal

import (
	"testing"
	"time"
)

func TestEnforcer_AddNode(t *testing.T) {
	e := NewEnforcer()

	t1 := time.Now().Add(-1 * time.Hour)
	t2 := time.Now()

	n1 := &Node{ID: "1", Timestamp: t1}
	if err := e.AddNode(n1); err != nil {
		t.Fatalf("Failed to add first node: %v", err)
	}

	n2 := &Node{ID: "2", Timestamp: t2, Parents: []string{"1"}}
	if err := e.AddNode(n2); err != nil {
		t.Fatalf("Failed to add second node: %v", err)
	}

	// Test unique ID
	if err := e.AddNode(n1); err == nil {
		t.Error("Expected error for duplicate node ID, got nil")
	}

	// Test missing parent
	n3 := &Node{ID: "3", Timestamp: t2, Parents: []string{"missing"}}
	if err := e.AddNode(n3); err == nil {
		t.Error("Expected error for missing parent, got nil")
	}

	// Test temporal regression
	t_early := t1.Add(-10 * time.Minute)
	n4 := &Node{ID: "4", Timestamp: t_early, Parents: []string{"1"}}
	if err := e.AddNode(n4); err == nil {
		t.Error("Expected error for temporal regression, got nil")
	}

	// Test Merkle Hash Mismatch
	n5 := &Node{ID: "5", Timestamp: t2, Parents: []string{"1"}, Hash: "invalid-hash"}
	if err := e.AddNode(n5); err == nil {
		t.Error("Expected error for hash mismatch, got nil")
	}

	// Test Valid Hash Generation
	n6 := &Node{ID: "6", Timestamp: t2, Parents: []string{"1"}}
	if err := e.AddNode(n6); err != nil {
		t.Fatalf("Failed to add node with auto-generated hash: %v", err)
	}
	if n6.Hash == "" {
		t.Error("Expected hash to be auto-generated, but it's empty")
	}
}
