package main

import (
	"fmt"
	"sync"
	"time"
)

type Node struct {
	PID       uint32
	PPID      uint32
	Name      string
	StartTime time.Time
	EndTime   time.Time
	Active    bool
}

type TraceStorage struct {
	mu     sync.RWMutex
	Hot    map[uint32]*Node
	Warm   []Node // Simplified: using slices for demo
	Cold   []Node // Skeleton nodes
	Policy EvictionPolicy
}

type EvictionPolicy struct {
	WarmThreshold time.Duration
	ColdThreshold time.Duration
}

func NewTraceStorage() *TraceStorage {
	return &TraceStorage{
		Hot:  make(map[uint32]*Node),
		Warm: make([]Node, 0),
		Cold: make([]Node, 0),
		Policy: EvictionPolicy{
			WarmThreshold: 1 * time.Hour,
			ColdThreshold: 24 * time.Hour,
		},
	}
}

func (s *TraceStorage) AddProcess(pid, ppid uint32, name string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	s.Hot[pid] = &Node{
		PID:       pid,
		PPID:      ppid,
		Name:      name,
		StartTime: time.Now(),
		Active:    true,
	}
}

func (s *TraceStorage) TerminateProcess(pid uint32) {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	node, exists := s.Hot[pid]
	if !exists { return }
	
	node.Active = false
	node.EndTime = time.Now()
	
	// Move to WARM
	s.Warm = append(s.Warm, *node)
	delete(s.Hot, pid)
	fmt.Printf("[TRACE] Moved PID %d to WARM tier\n", pid)
}

func (s *TraceStorage) PruneToCold() {
	s.mu.Lock()
	defer s.mu.Unlock()
	
	// In a real implementation, we would check timestamps.
	// For the demo, we move everything from WARM to COLD except critical nodes.
	for i := 0; i < len(s.Warm); i++ {
		node := s.Warm[i]
		if node.Name == "init" || node.Name == "systemd" || node.Name == "kernel" {
			continue // Keep critical nodes in warm/accessible tiers if needed
		}
		
		// Compress to COLD (Skeleton)
		skeleton := Node{
			PID:  node.PID,
			PPID: node.PPID,
			Name: node.Name,
		}
		s.Cold = append(s.Cold, skeleton)
		fmt.Printf("[TRACE] Compressed PID %d to COLD tier (Skeleton)\n", node.PID)
	}
	s.Warm = s.Warm[:0] // Clear warm tier after pruning
}

func (s *TraceStorage) Stats() {
	s.mu.RLock()
	defer s.mu.RUnlock()
	fmt.Printf("Trace Stats: HOT=%d, WARM=%d, COLD=%d\n", len(s.Hot), len(s.Warm), len(s.Cold))
}
