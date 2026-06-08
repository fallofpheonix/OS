/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 */

package security

import (
	"context"
	"strings"
	"sync"
	"testing"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	ledger "github.com/fallofpheonix/phoenix/foundation/ledger/src"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

// ---------------------------------------------------------------------------
// RegisterInvariant
// ---------------------------------------------------------------------------

func TestRegisterInvariant_AppendsToSlice(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b, &ledger.Ledger{})

	if got := len(w.getInvariants()); got != 0 {
		t.Fatalf("expected empty invariants on fresh Warden, got %d", got)
	}

	inv1 := &EvidenceWeightInvariant{StateThresholds: map[SystemState]phxmath.FixedPoint{StateCritical: phxmath.NewFixedPointRaw(800000)}}
	inv2 := &ContextualInvariant{}
	inv3 := &CertificateInvariant{}

	w.RegisterInvariant(inv1)
	w.RegisterInvariant(inv2)
	w.RegisterInvariant(inv3)

	invs := w.getInvariants()
	if got := len(invs); got != 3 {
		t.Fatalf("expected 3 invariants after Register, got %d", got)
	}
	if invs[0] != inv1 {
		t.Errorf("Invariants[0]: expected inv1")
	}
	if invs[1] != inv2 {
		t.Errorf("Invariants[1]: expected inv2")
	}
	if invs[2] != inv3 {
		t.Errorf("Invariants[2]: expected inv3")
	}
}

func TestRegisterInvariant_ThreadSafe(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b, &ledger.Ledger{})

	const goroutines = 32
	var wg sync.WaitGroup
	wg.Add(goroutines)
	for i := 0; i < goroutines; i++ {
		go func() {
			defer wg.Done()
			w.RegisterInvariant(&EvidenceWeightInvariant{})
		}()
	}
	wg.Wait()

	if got := len(w.getInvariants()); got != goroutines {
		t.Fatalf("expected %d invariants after concurrent register, got %d", goroutines, got)
	}
}

// ---------------------------------------------------------------------------
// RegisterActuator
// ---------------------------------------------------------------------------

type stubActuator struct {
	name         string
	currentLevel securityv1.ContainmentLevel
}

func (s *stubActuator) Actuate(ctx context.Context, action securityv1.Containment) error {
	s.currentLevel = action.Level()
	return nil
}

func (s *stubActuator) Kill(ctx context.Context, pid int) error {
	s.currentLevel = securityv1.LevelQuench
	return nil
}

func (s *stubActuator) GetCurrentLevel() (securityv1.ContainmentLevel, error) {
	return s.currentLevel, nil
}

func (s *stubActuator) Name() string {
	return s.name
}

func TestRegisterActuator_AppendsToSlice(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b, &ledger.Ledger{})

	if got := len(w.getActuators()); got != 0 {
		t.Fatalf("expected empty actuators on fresh Warden, got %d", got)
	}

	a1 := &stubActuator{name: "process"}
	a2 := &stubActuator{name: "ebpf"}

	w.RegisterActuator(a1)
	w.RegisterActuator(a2)

	acts := w.getActuators()
	if got := len(acts); got != 2 {
		t.Fatalf("expected 2 actuators after Register, got %d", got)
	}
	if name := acts[0].(interface{ Name() string }).Name(); name != "process" {
		t.Errorf("Actuators[0]: expected name=process, got %q", name)
	}
	if name := acts[1].(interface{ Name() string }).Name(); name != "ebpf" {
		t.Errorf("Actuators[1]: expected name=ebpf, got %q", name)
	}
}

// ---------------------------------------------------------------------------
// logViolation
// ---------------------------------------------------------------------------

func TestLogViolation_PublishesTelemetryEventOnBus(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b, &ledger.Ledger{})

	ch := b.Subscribe("warden.violation")

	var received bus.TelemetryEvent
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		received = <-ch
	}()

	w.logViolation("audit-trail event")
	wg.Wait()

	if received.EventType != "warden.violation" {
		t.Errorf("expected EventType warden.violation, got %q", received.EventType)
	}
	if !strings.Contains(string(received.Payload), "audit-trail event") {
		t.Errorf("expected payload to contain 'audit-trail event', got %q", string(received.Payload))
	}
}

// ---------------------------------------------------------------------------
// Interface contracts
// ---------------------------------------------------------------------------

func TestGraphProvider_InterfaceSatisfiedByMock(t *testing.T) {
	var _ GraphProvider = (*MockGraphProvider)(nil)
}

func TestCertificateValidator_InterfaceSatisfiedByStub(t *testing.T) {
	stub := &stubCertValidator{}
	var v CertificateValidator = stub
	if !v.VerifyCertificate("evt-1", phxmath.NewFixedPointRaw(500000), []byte{0x01}) {
		t.Errorf("expected stub validator to return true for well-formed cert")
	}
}

type stubCertValidator struct{}

func (s *stubCertValidator) VerifyCertificate(eventID string, weight phxmath.FixedPoint, cert []byte) bool {
	return eventID != "" && weight.V >= 0 && len(cert) > 0
}

// ---------------------------------------------------------------------------
// GraphProof value semantics
// ---------------------------------------------------------------------------

func TestGraphProof_FieldsPreserved(t *testing.T) {
	p := &GraphProof{
		Path:            []string{"a", "b", "c"},
		ExpectedNsproxy: 42,
	}
	if len(p.Path) != 3 || p.Path[0] != "a" || p.Path[2] != "c" {
		t.Errorf("Path not preserved: %v", p.Path)
	}
	if p.ExpectedNsproxy != 42 {
		t.Errorf("ExpectedNsproxy not preserved: %d", p.ExpectedNsproxy)
	}
}
