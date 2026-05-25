package unit

import (
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestLineageGraph(t *testing.T) {
	g := NewLineageGraph()
	now := time.Now()

	// Build a simple tree
	// 1 -> 10 -> 100
	// 1 -> 11
	g.AddProcess(1, 0, "systemd", "/sbin/init", now)
	g.AddProcess(10, 1, "bash", "/bin/bash", now.Add(time.Second))
	g.AddProcess(100, 10, "ls", "/bin/ls", now.Add(2*time.Second))
	g.AddProcess(11, 1, "sshd", "/usr/sbin/sshd", now.Add(time.Second))

	if g.Size() != 4 {
		t.Errorf("Expected size 4, got %d", g.Size())
	}

	ancestors := g.GetAncestors(100)
	if len(ancestors) != 2 || ancestors[0] != 10 || ancestors[1] != 1 {
		t.Errorf("Unexpected ancestors for 100: %v", ancestors)
	}

	descendants := g.GetDescendants(1)
	if len(descendants) != 3 {
		t.Errorf("Expected 3 descendants for PID 1, got %d", len(descendants))
	}

	g.ExitProcess(100, now.Add(3*time.Second))
	if g.Nodes[100].IsActive {
		t.Error("Process 100 should be inactive")
	}
}

// Memory Stress Test (Lineage Memory Tests)
func TestLineageMemoryStress(t *testing.T) {
	g := NewLineageGraph()
	now := time.Now()
	
	const numProcesses = 100000
	
	var m1, m2 runtime.MemStats
	runtime.GC()
	runtime.ReadMemStats(&m1)

	for i := uint32(1); i <= numProcesses; i++ {
		ppid := uint32(0)
		if i > 1 {
			ppid = i / 2 // Create a binary tree structure
		}
		g.AddProcess(i, ppid, "test-proc", "/usr/bin/test", now)
	}

	runtime.GC()
	runtime.ReadMemStats(&m2)

	memUsedMB := float64(m2.Alloc-m1.Alloc) / 1024 / 1024
	fmt.Printf("Memory used for %d processes: %.2f MB (%.2f bytes/node)\n", 
		numProcesses, memUsedMB, float64(m2.Alloc-m1.Alloc)/float64(numProcesses))

	if memUsedMB > 100 { // Arbitrary limit for 100k nodes
		t.Errorf("Memory usage too high: %.2f MB", memUsedMB)
	}

	// Test traversal performance on large graph
	start := time.Now()
	descendants := g.GetDescendants(1)
	elapsed := time.Since(start)
	
	fmt.Printf("BFS traversal of %d nodes took %v\n", len(descendants), elapsed)
	if elapsed > 500*time.Millisecond {
		t.Errorf("Traversal too slow: %v", elapsed)
	}
}

func BenchmarkAddProcess(b *testing.B) {
	g := NewLineageGraph()
	now := time.Now()
	for i := 0; i < b.N; i++ {
		g.AddProcess(uint32(i+1), uint32(i/2), "bench", "/bin/bench", now)
	}
}
