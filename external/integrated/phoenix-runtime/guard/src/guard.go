package main

import (
	"fmt"
	"time"
)

type GuardAction uint32

const (
	ActionAllow   GuardAction = 0
	ActionSuspend GuardAction = 1
	ActionKill    GuardAction = 2
)

type GuardService struct {
	PolicyMap map[uint32]GuardAction
}

func NewGuardService() *GuardService {
	return &GuardService{
		PolicyMap: make(map[uint32]GuardAction),
	}
}

// FastPathDetect simulates high-speed heuristic detection
func (g *GuardService) FastPathDetect(pid uint32, entropy float64, renameRate int) {
	start := time.Now()

	// HEU-01: High Entropy
	if entropy > 7.9 {
		g.UpdateKernelPolicy(pid, ActionSuspend)
		fmt.Printf("[FAST_PATH] HEU-01 Triggered for PID %d (Latency: %v)\n", pid, time.Since(start))
		return
	}

	// HEU-02: Rename Burst
	if renameRate > 50 {
		g.UpdateKernelPolicy(pid, ActionKill)
		fmt.Printf("[FAST_PATH] HEU-02 Triggered for PID %d (Latency: %v)\n", pid, time.Since(start))
		return
	}
}

func (g *GuardService) UpdateKernelPolicy(pid uint32, action GuardAction) {
	// In a real env, this would update the BPF Hash Map
	g.PolicyMap[pid] = action
	fmt.Printf("[GUARD] BPF Map Updated: PID %d -> Action %d\n", pid, action)
}

func main() {
	fmt.Println("Phoenix Guard Runtime starting...")
	guard := NewGuardService()

	// Simulate high-entropy ransomware write
	fmt.Println("Simulating Ransomware Signal...")
	guard.FastPathDetect(5001, 7.95, 0)

	// Simulate bulk rename burst
	fmt.Println("Simulating Rename Burst Signal...")
	guard.FastPathDetect(5002, 3.5, 100)
}
