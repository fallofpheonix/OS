/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: warden_compat_test.go
 *
 * Purpose:
 * Tests for the compatibility layer added to the root `warden` package.
 * Covers RegisterInvariant, logViolation, and the GraphProvider /
 * CertificateValidator interface contracts.
 *
 * Subsystem:
 * PhoenixGuard (Tactical & Enforcement Layer) — Root Warden compat shim
 *
 * Workflow:
 *   [CYCLE 8] Actuate() → Invariant.Verify() → on violation → logViolation()
 *   → Bus.Publish("warden.violation", ...) → downstream subscribers react
 *
 * Created For:
 * P0 buildability: gap closure for missing types/methods referenced by
 * active runtime callers and existing tests.
 */

package warden

import (
	"bytes"
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"testing"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
)

// ---------------------------------------------------------------------------
// RegisterInvariant
// ---------------------------------------------------------------------------

func TestRegisterInvariant_AppendsToSlice(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

	if got := len(w.Invariants); got != 0 {
		t.Fatalf("expected empty invariants on fresh Warden, got %d", got)
	}

	inv1 := &EvidenceWeightInvariant{StateThresholds: map[SystemState]float64{StateCritical: 0.8}}
	inv2 := &ContextualInvariant{}
	inv3 := &CertificateInvariant{}

	w.RegisterInvariant(inv1)
	w.RegisterInvariant(inv2)
	w.RegisterInvariant(inv3)

	if got := len(w.Invariants); got != 3 {
		t.Fatalf("expected 3 invariants after Register, got %d", got)
	}
	if w.Invariants[0] != inv1 {
		t.Errorf("Invariants[0]: expected inv1")
	}
	if w.Invariants[1] != inv2 {
		t.Errorf("Invariants[1]: expected inv2")
	}
	if w.Invariants[2] != inv3 {
		t.Errorf("Invariants[2]: expected inv3")
	}
}

func TestRegisterInvariant_ThreadSafe(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)

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

	if got := len(w.Invariants); got != goroutines {
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
	w := NewWarden(b)

	if got := len(w.Actuators); got != 0 {
		t.Fatalf("expected empty actuators on fresh Warden, got %d", got)
	}

	a1 := &stubActuator{name: "process"}
	a2 := &stubActuator{name: "ebpf"}

	w.RegisterActuator(a1)
	w.RegisterActuator(a2)

	if got := len(w.Actuators); got != 2 {
		t.Fatalf("expected 2 actuators after Register, got %d", got)
	}
	if name := w.Actuators[0].(interface{ Name() string }).Name(); name != "process" {
		t.Errorf("Actuators[0]: expected name=process, got %q", name)
	}
	if name := w.Actuators[1].(interface{ Name() string }).Name(); name != "ebpf" {
		t.Errorf("Actuators[1]: expected name=ebpf, got %q", name)
	}
}

// ---------------------------------------------------------------------------
// EnableDiagnostics
// ---------------------------------------------------------------------------

func TestEnableDiagnostics_EmptyPathNoOp(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	if err := w.EnableDiagnostics(""); err != nil {
		t.Errorf("expected no error for empty path, got %v", err)
	}
	if w.DiagnosticLogger != nil {
		t.Errorf("expected DiagnosticLogger to remain nil for empty path")
	}
}

func TestEnableDiagnostics_WritesToFile(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	tmp := filepath.Join(t.TempDir(), "warden_diag.log")

	if err := w.EnableDiagnostics(tmp); err != nil {
		t.Fatalf("EnableDiagnostics failed: %v", err)
	}
	if w.DiagnosticLogger == nil {
		t.Fatalf("expected DiagnosticLogger to be set after EnableDiagnostics")
	}

	w.DiagnosticLogger.Printf("hello diagnostic")
	w.logViolation("test violation %d", 7)

	// Flush by re-opening
	data, err := os.ReadFile(tmp)
	if err != nil {
		t.Fatalf("read diag log: %v", err)
	}
	content := string(data)
	if !strings.Contains(content, "hello diagnostic") {
		t.Errorf("expected 'hello diagnostic' in %q", content)
	}
	if !strings.Contains(content, "[VIOLATION]") {
		t.Errorf("expected '[VIOLATION]' tag in %q", content)
	}
	if !strings.Contains(content, "test violation 7") {
		t.Errorf("expected formatted violation message in %q", content)
	}
}

func TestEnableDiagnostics_BadPathReturnsError(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	// /nonexistent-root/warden.log: parent dir does not exist
	err := w.EnableDiagnostics("/nonexistent-root-12345/warden.log")
	if err == nil {
		t.Errorf("expected error for unopenable path, got nil")
	}
	if w.DiagnosticLogger != nil {
		t.Errorf("expected DiagnosticLogger to remain nil on error")
	}
}

// ---------------------------------------------------------------------------
// logViolation
// ---------------------------------------------------------------------------

func TestLogViolation_UsesDiagnosticLoggerWhenSet(t *testing.T) {
	var buf bytes.Buffer
	b := bus.NewBus()
	w := NewWarden(b)
	w.DiagnosticLogger = log.New(&buf, "", 0)

	w.logViolation("test %s with %d", "message", 42)

	out := buf.String()
	if !strings.Contains(out, "[VIOLATION]") {
		t.Errorf("expected log to contain [VIOLATION] tag, got %q", out)
	}
	if !strings.Contains(out, "test message with 42") {
		t.Errorf("expected formatted message in log, got %q", out)
	}
}

func TestLogViolation_DoesNotPanicWhenDiagnosticLoggerNil(t *testing.T) {
	b := bus.NewBus()
	w := NewWarden(b)
	// DiagnosticLogger is nil — must fall back to standard log without panic.
	w.logViolation("fallback case %s", "ok")
}

func TestLogViolation_PublishesTelemetryEventOnBus(t *testing.T) {
	b := bus.NewBus()
	var sinkBuf bytes.Buffer
	w := NewWarden(b)
	w.DiagnosticLogger = log.New(&sinkBuf, "", 0)

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
	if !v.VerifyCertificate("evt-1", 0.5, []byte{0x01}) {
		t.Errorf("expected stub validator to return true for well-formed cert")
	}
}

type stubCertValidator struct{}

func (s *stubCertValidator) VerifyCertificate(eventID string, weight float64, cert []byte) bool {
	return eventID != "" && weight >= 0 && len(cert) > 0
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
