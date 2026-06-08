package security

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

// mockActuator tracks calls for E2E verification
type mockActuator struct {
	mu        sync.Mutex
	killCount int
}

func (m *mockActuator) Kill(ctx context.Context, pid int) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.killCount++
	return nil
}
func (m *mockActuator) Actuate(ctx context.Context, a securityv1.Containment) error { return nil }
func (m *mockActuator) GetCurrentLevel() (securityv1.ContainmentLevel, error)       { return 0, nil }

func (m *mockActuator) getKillCount() int {
	m.mu.Lock()
	defer m.mu.Unlock()
	return m.killCount
}

func TestE2E_CausalClosure(t *testing.T) {
	// [SETUP]: Boot Warden with real ledger but NO Start() (Simulates immediate crash)
	l := ledger.NewLedger(nil)
	w := NewWarden(bus.NewBus(), l)
	w.shadowMode.Store(false)

	policyHash := computePolicyHash(w.policies.Load())

	// [ACT]: Fire 5 actuations
	for i := 0; i < 5; i++ {
		req := AuthorityEscalationRequest{
			EventID:        fmt.Sprintf("EVT-CAS-%d", i),
			TargetPID:      uint32(100 + i),
			TargetState:    StateWatch,
			ActuationClass: ClassKill,
			PolicyHash:     policyHash,
			EvidenceWeight: phxmath.NewFixedPointRaw(300000), // 0.30 > 0.20
		}
		_, err := w.ActuateRequest(req, 0, uint64(i))
		if err != nil {
			t.Fatalf("ActuateRequest failed: %v", err)
		}
	}

	// [ASSERT]: Scan ledger for orphaned PENDING entries
	entries := l.SortedEntries()
	unpaired := make(map[string]bool)
	for _, e := range entries {
		// RECTIFIED: Q175 - Distinguish PENDING vs COMPLETION
		if e.CauseID == "ROOT" {
			unpaired[e.EventID] = true
		} else {
			delete(unpaired, e.CauseID)
		}
	}

	if len(unpaired) != 5 {
		t.Errorf("Orphan detection failed: expected 5 orphans, got %d", len(unpaired))
	}
}

func TestE2E_NormalClosure(t *testing.T) {
	// [SETUP]: Boot Warden and Start() async drainer
	l := ledger.NewLedger(nil)
	w := NewWarden(bus.NewBus(), l)
	w.shadowMode.Store(false) // EXPLICIT OPT-OUT for physical actuation test
	act := &mockActuator{}
	w.actuators = append(w.actuators, act)
	w.Start()
	defer w.Stop()

	policyHash := computePolicyHash(w.policies.Load())

	// [ACT]: Fire 5 actuations and wait for completion
	for i := 0; i < 5; i++ {
		req := AuthorityEscalationRequest{
			EventID:        fmt.Sprintf("EVT-NOR-%d", i),
			ActuationClass: ClassKill,
			TargetState:    StateWatch,
			PolicyHash:     policyHash,
			EvidenceWeight: phxmath.NewFixedPointRaw(300000),
		}
		_, _ = w.ActuateRequest(req, 0, uint64(i))
	}

	// Busy wait for async completion (max 2s)
	deadline := time.Now().Add(2 * time.Second)
	for time.Now().Before(deadline) {
		if act.getKillCount() == 5 {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}

	if act.getKillCount() != 5 {
		t.Fatalf("Actuator only fired %d times, want 5", act.getKillCount())
	}

	// [ASSERT]: Verify 0 orphans
	entries := l.SortedEntries()
	unpaired := make(map[string]bool)
	for _, e := range entries {
		if e.CauseID == "ROOT" {
			unpaired[e.EventID] = true
		} else {
			delete(unpaired, e.CauseID)
		}
	}

	if len(unpaired) != 0 {
		for id := range unpaired {
			t.Errorf("Orphan remaining: %s", id)
		}
		t.Errorf("Causal loop mismatch: %d orphans remain", len(unpaired))
	}
}

func TestE2E_ShadowModeNoActuation(t *testing.T) {
	l := ledger.NewLedger(nil)
	w := NewWarden(bus.NewBus(), l)
	w.shadowMode.Store(true) // Explicitly shadow
	act := &mockActuator{}
	w.actuators = append(w.actuators, act)
	w.Start()
	defer w.Stop()

	policyHash := computePolicyHash(w.policies.Load())
	req := AuthorityEscalationRequest{
		EventID: "SHADOW-1", ActuationClass: ClassKill, TargetState: StateWatch,
		PolicyHash: policyHash, EvidenceWeight: phxmath.NewFixedPointRaw(300000),
	}

	_, err := w.ActuateRequest(req, 0, 100)
	if err != nil {
		t.Fatalf("ActuateRequest failed: %v", err)
	}

	time.Sleep(200 * time.Millisecond)

	if act.getKillCount() != 0 {
		t.Error("Physical actuation occurred in Shadow Mode!")
	}
}

func TestE2E_PolicyMismatchRejects(t *testing.T) {
	l := ledger.NewLedger(nil)
	w := NewWarden(bus.NewBus(), l)

	req := AuthorityEscalationRequest{
		EventID: "BAD-POLICY", TargetState: StateWatch,
		PolicyHash:     [32]byte{0xDE, 0xAD}, // Wrong hash
		EvidenceWeight: phxmath.NewFixedPointRaw(300000),
	}

	_, err := w.ActuateRequest(req, 0, 100)
	if err == nil || err.Error() != "ERR_POLICY_MISMATCH" {
		t.Errorf("Expected ERR_POLICY_MISMATCH, got %v", err)
	}

	if len(l.SortedEntries()) != 0 {
		t.Error("Ledger entry written for rejected request!")
	}
}
