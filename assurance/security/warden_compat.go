/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
/**
 * FILE: warden_compat.go
 */

package security

import (
	"encoding/json"
	"fmt"
	"log"

	securityv1 "github.com/fallofpheonix/phoenix/foundation/contracts/security/v1"
	phxmath "github.com/fallofpheonix/phoenix/foundation/math"
	"github.com/fallofpheonix/phoenix/foundation/runtime/bus"
)

type Actuator = securityv1.Actuator

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
	VerifyCertificate(eventID string, weight phxmath.FixedPoint, cert []byte) bool
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
	w.invariants = append(w.invariants, inv)
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
	w.actuators = append(w.actuators, act)
}

/**
 * logViolation
 *
 * Emits a Warden violation to all available sinks:
 *   1. standard log (fallback)
 *   2. Bus topic "warden.violation" (if Bus is configured)
 *
 * Format: Printf-style, tagged with [VIOLATION].
 *
 * Failure mode:
 * Never panics. If all sinks are unavailable, the standard logger still
 * receives the message.
 */
func (w *Warden) logViolation(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	log.Printf("[VIOLATION] %s", msg)

	if w.bus != nil {
		payload, err := json.Marshal(map[string]string{"message": msg})
		if err != nil {
			log.Printf("[Warden] failed to marshal violation payload: %v", err)
			return
		}
		w.bus.Publish("warden.violation", bus.TelemetryEvent{
			EventType: "warden.violation",
			Severity:  phxmath.NewFixedPointRaw(850000), // 0.85
			Source:    "PhoenixGuard/warden",
			Payload:   payload,
		})
	}
}
