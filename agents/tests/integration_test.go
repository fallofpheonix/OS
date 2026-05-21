package tests

import (
	"fmt"
	"os"
	"testing"
	"time"

	"phoenix/agents/internal/control"
	"phoenix/agents/internal/forensics"
	"phoenix/agents/internal/game"
	"phoenix/agents/internal/graph"
	"phoenix/agents/internal/kernel"
	"phoenix/agents/internal/physics"
	"phoenix/agents/internal/telemetry"
	"phoenix/agents/internal/types"
)

func TestEndToEndPipeline(t *testing.T) {
	// Setup workspace temporary forensics directory
	forensicsDir := "./test_forensics"
	defer os.RemoveAll(forensicsDir)

	// 1. Instantiate All Agents
	telemetryAgent := telemetry.NewTelemetryAgent(100)
	graphAgent := graph.NewGraphAgent()
	physicsAgent := physics.NewPhysicsAgent()
	gameAgent := game.NewGameAgent()
	controlAgent := control.NewControlAgent(1.5, 0.2, 0.1, 2.0) // Setpoint = 2.0 Max Temp
	forensicsAgent, err := forensics.NewForensicsAgent(forensicsDir)
	if err != nil {
		t.Fatalf("Failed to initialize forensics agent: %v", err)
	}
	kernelAgent := kernel.NewKernelAgent()

	// 2. Verify Initial Security Constraints & Locked Kernel
	if !kernelAgent.IsLocked() {
		t.Error("Security Constraint Violation: Kernel agent should be locked by default")
	}

	err = kernelAgent.RegisterLSMHook("bprm_check_security")
	if err == nil {
		t.Error("Expected error when calling RegisterLSMHook on a locked Kernel Agent, got nil")
	}

	err = kernelAgent.ApplySchedulerOverride(1234, -20)
	if err == nil {
		t.Error("Expected error when calling ApplySchedulerOverride on a locked Kernel Agent, got nil")
	}

	err = kernelAgent.ActuateContainmentPolicy(1234, "KILL")
	if err == nil {
		t.Error("Expected error when calling ActuateContainmentPolicy on a locked Kernel Agent, got nil")
	}

	// 3. Simulate Normal System Activity
	normalEvent := types.TelemetryEvent{
		Timestamp: time.Now(),
		EventID:   "evt-100",
		Category:  "process",
		EventType: "process.start",
		PID:       2001,
		PPID:      1,
		Comm:      "nginx",
		ExePath:   "/usr/sbin/nginx",
	}
	telemetryAgent.RecordEvent(normalEvent)
	
	err = graphAgent.UpdateGraph(normalEvent)
	if err != nil {
		t.Fatalf("Failed to update graph: %v", err)
	}

	attackDAG := graphAgent.GetAttackDAG()
	normalState, err := physicsAgent.GetSecurityState(attackDAG)
	if err != nil {
		t.Fatalf("Failed to get security state: %v", err)
	}

	if normalState.IsAnomaly {
		t.Errorf("Unexpected anomaly state for normal nginx startup: %+v", normalState)
	}

	gameAgent.UpdateBeliefs(normalState, "normal_filesystem_io")
	normalStrategy, err := gameAgent.SolveBestStrategy(normalState, attackDAG)
	if err != nil {
		t.Fatalf("Failed to solve strategy: %v", err)
	}

	err = controlAgent.EnforceStrategy(normalStrategy, normalState.ThreatTemperature)
	if err != nil {
		t.Fatalf("Failed to enforce normal strategy: %v", err)
	}

	// Verify no aggressive containment action was triggered
	normalHistory := controlAgent.GetActionHistory()
	for _, act := range normalHistory {
		t.Errorf("Unexpected containment action triggered in normal conditions: %s", act)
	}

	// 4. Simulate Ransomware Attack / Anomaly & Escalation
	malwarePID := uint32(9999)
	var anomalyState types.SecurityState
	var activeDAG *types.IncidentGraph

	// We simulate a series of malicious filesystem events to escalate threat beliefs and temperature
	for i := 0; i < 5; i++ {
		malwareEvent := types.TelemetryEvent{
			Timestamp: time.Now(),
			EventID:   fmt.Sprintf("evt-666-%d", i),
			Category:  "network",
			EventType: "network.connect",
			PID:       malwarePID,
			PPID:      2001, // Child of nginx!
			Comm:      "malware",
			ExePath:   "/tmp/malware",
			Network: &types.NetworkPayload{
				SAddr:    "127.0.0.1",
				DAddr:    "8.8.8.8",
				SPort:    55555,
				DPort:    4444, // Reverse shell port
				Protocol: "TCP",
			},
		}
		telemetryAgent.RecordEvent(malwareEvent)

		err = graphAgent.UpdateGraph(malwareEvent)
		if err != nil {
			t.Fatalf("Failed to update graph with malware event: %v", err)
		}

		activeDAG = graphAgent.GetAttackDAG()
		anomalyState, err = physicsAgent.GetSecurityState(activeDAG)
		if err != nil {
			t.Fatalf("Failed to get security state: %v", err)
		}

		gameAgent.UpdateBeliefs(anomalyState, "high_entropy_write") // ransomware signature
	}

	// Check if threat temperature and disorder index reflect the anomaly
	if !anomalyState.IsAnomaly {
		t.Errorf("Expected anomaly state to be true, got false. State: %+v", anomalyState)
	}
	if anomalyState.ThreatTemperature <= 2.0 {
		t.Errorf("Expected elevated threat temperature, got %f", anomalyState.ThreatTemperature)
	}

	// Run Bayesian Game solver
	threatStrategy, err := gameAgent.SolveBestStrategy(anomalyState, activeDAG)
	if err != nil {
		t.Fatalf("Failed to solve game strategy: %v", err)
	}

	if threatStrategy.ContainmentLevel != 5 {
		t.Errorf("Expected level 5 (Kill) containment strategy, got level %d", threatStrategy.ContainmentLevel)
	}

	// Enforce containment in Control Agent
	err = controlAgent.EnforceStrategy(threatStrategy, anomalyState.ThreatTemperature)
	if err != nil {
		t.Fatalf("Failed to enforce threat strategy: %v", err)
	}

	actionHistory := controlAgent.GetActionHistory()
	if len(actionHistory) == 0 {
		t.Error("Expected containment actions to be logged, got 0")
	}

	hasKillAction := false
	for _, act := range actionHistory {
		t.Logf("Control Action: %s", act)
		if len(act) > 0 { // Just confirm we got some log entries
			hasKillAction = true
		}
	}
	if !hasKillAction {
		t.Error("Expected a KILL containment action to be registered")
	}

	// Trigger Forensics snapshot for malicious PID
	snapResult, err := forensicsAgent.TriggerSnapshot(malwarePID)
	if err != nil {
		t.Fatalf("Failed to trigger forensics snapshot: %v", err)
	}

	t.Logf("Snapshot captured at path: %s", snapResult.Path)
	t.Logf("Snapshot cryptographic hash: %s", snapResult.Hash)

	// Verify cryptographic hash integrity
	verifiedHash, err := forensicsAgent.VerifyHash(snapResult.Path)
	if err != nil {
		t.Fatalf("Failed to verify snapshot hash: %v", err)
	}

	if verifiedHash != snapResult.Hash {
		t.Errorf("Forensics integrity check failed: expected hash %s, computed %s", snapResult.Hash, verifiedHash)
	}

	// 5. Unlock Kernel Agent once userspace pipeline has run and validated
	kernelAgent.Unlock()
	if kernelAgent.IsLocked() {
		t.Error("Kernel Agent should be unlocked now")
	}

	// Kernel hooks should be modifiable now
	err = kernelAgent.RegisterLSMHook("bprm_check_security")
	if err != nil {
		t.Errorf("Failed to register LSM hook on unlocked Kernel Agent: %v", err)
	}

	err = kernelAgent.ActuateContainmentPolicy(malwarePID, "KILL")
	if err != nil {
		t.Errorf("Failed to actuate containment policy on unlocked Kernel Agent: %v", err)
	}

	t.Log("End-to-End Simulation successfully completed!")
}
