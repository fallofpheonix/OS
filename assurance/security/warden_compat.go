/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: warden_compat.go
 *
 * Purpose:
 * Compatibility shim for the root `warden` package. Adds the types,
 * interfaces, and methods that are referenced by invariant.go, warden.go,
 * and downstream callers (active PhoenixOS runtime) but were not yet
 * declared in this package. The newer, parallel `engine/warden.go`
 * implementation remains the long-term canonical interface; this file
 * only closes the build gap so the workspace compiles.
 *
 * Subsystem:
 * PhoenixGuard (Tactical & Enforcement Layer)
 *
 * Workflow:
 *   ai/orchestrator → Warden.Actuate()
 *     → Invariant.Verify() (Evidence / Certificate / Contextual)
 *       → on failure → Warden.logViolation() → Bus.Publish("warden.violation")
 *
 * Created For:
 * P0 buildability. Required to unblock the active Phoenix.Terminus/PhoenixOS
 * module, which imports github.com/fallofpheonix/phoenix/assurance/security (root pkg).
 *
 * Known Limitations:
 * - logViolation emits only to DiagnosticLogger and the bus; no persistent
 *   audit sink. A future hardening pass should write to the Audit package
 *   with cryptographic chaining.
 *
 * Security Considerations:
 * - Every violation MUST be observable. If both DiagnosticLogger and Bus
 *   are nil, the violation is still logged to the standard logger, never
 *   silently dropped.
 */

package warden

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

/**
 * GraphProof
 *
 * Causal proof object passed by the Oracle to back an
 * AuthorityEscalationRequest. Captures the reasoning path the Oracle
 * followed and the namespace the target was expected to be in.
 *
 * Lifecycle:
 * Constructed by the Oracle; consumed by ContextualInvariant.Verify().
 */
type GraphProof struct {
	Path            []string
	ExpectedNsproxy uint32
}

/**
 * GraphProvider
 *
 * Verifies causal paths extracted from GraphProof.
 *
 * Implemented by:
 * - The trace DAG component (production)
 * - MockGraphProvider in tests
 */
type GraphProvider interface {
	VerifyPath(path []string) (bool, error)
}

/**
 * CertificateValidator
 *
 * Verifies the cryptographic certificate attached to an
 * AuthorityEscalationRequest against the originating event.
 *
 * Implemented by:
 * - Phoenix.Nucleus/PhoenixCore/ledger (production)
 * - Test stubs
 */
type CertificateValidator interface {
	VerifyCertificate(eventID string, weight float64, cert []byte) bool
}

/**
 * RegisterInvariant
 *
 * Appends an Invariant to the Warden's enforcement chain.
 * Invariants are evaluated in registration order during Actuate().
 *
 * Thread Safety:
 * Acquires the Warden's write lock; safe for concurrent registration.
 */
func (w *Warden) RegisterInvariant(inv Invariant) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Invariants = append(w.Invariants, inv)
}

/**
 * RegisterActuator
 *
 * Appends an Actuator to the Warden's physical response chain.
 * Actuators are invoked in registration order during Actuate() when
 * TargetPID > 0 and ActuationClass is Isolate or Throttle.
 *
 * Thread Safety:
 * Acquires the Warden's write lock; safe for concurrent registration.
 */
func (w *Warden) RegisterActuator(act Actuator) {
	w.mu.Lock()
	defer w.mu.Unlock()
	w.Actuators = append(w.Actuators, act)
}

/**
 * EnableDiagnostics
 *
 * Routes Warden diagnostic output (state transitions, violations,
 * actuator events) to a file at the given path.
 *
 * Behavior:
 * - If path is empty, the file logger is left nil and logViolation
 *   falls back to standard log.
 * - If the file cannot be opened, the file logger remains nil and an
 *   error is returned; caller's existing log sink is preserved.
 * - Idempotent: calling EnableDiagnostics twice replaces the previous
 *   file handle. The first handle is NOT closed (the original caller
 *   is responsible for its lifetime).
 *
 * Thread Safety:
 * Acquires the Warden's write lock.
 */
func (w *Warden) EnableDiagnostics(path string) error {
	if path == "" {
		return nil
	}
	f, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0o644)
	if err != nil {
		return fmt.Errorf("enable diagnostics: open %s: %w", path, err)
	}
	w.mu.Lock()
	w.DiagnosticLogger = log.New(f, "[Warden] ", log.LstdFlags|log.Lmicroseconds)
	w.mu.Unlock()
	return nil
}

/**
 * logViolation
 *
 * Emits a Warden violation to all available sinks:
 *   1. DiagnosticLogger (if configured)
 *   2. standard log (fallback)
 *   3. Bus topic "warden.violation" (if Bus is configured)
 *
 * Format: Printf-style, tagged with [VIOLATION].
 *
 * Failure mode:
 * Never panics. If all sinks are unavailable, the standard logger still
 * receives the message.
 */
func (w *Warden) logViolation(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)

	if w.DiagnosticLogger != nil {
		w.DiagnosticLogger.Printf("[VIOLATION] %s", msg)
	} else {
		log.Printf("[VIOLATION] %s", msg)
	}

	if w.Bus != nil {
		payload, err := json.Marshal(map[string]string{"message": msg})
		if err != nil {
			log.Printf("[Warden] failed to marshal violation payload: %v", err)
			return
		}
		w.Bus.Publish("warden.violation", bus.TelemetryEvent{
			EventType: "warden.violation",
			Severity:  0.85,
			Source:    "PhoenixGuard/warden",
			Payload:   payload,
		})
	}
}
