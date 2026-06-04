/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: CYCLE 10 — EMERGENCY HALT (Layer 3.5)
//
// The KillSwitch is the ULTIMATE FAILSAFE of the Phoenix Matrix.
// When engaged, it freezes the Warden FSM permanently, preventing ANY
// further state transitions. This is the "pull the plug" mechanism.
//
// WORKFLOW:
//   Threat detected → KillSwitch.Engage()
//     → Warden.Lock() → FSM frozen in current state
//       → System enters "preservation mode"
//         → RecoveryLoop analyzes logs and attempts root cause identification
//           → If safe: manual operator intervention required
//             → Process restart to unlock FSM
//
// DESIGN DECISION: The KillSwitch is IRREVERSIBLE within a process lifetime.
// This prevents an attacker from disengaging the failsafe after triggering it.
// Recovery requires a full process restart, which forces re-initialization of
// all kernel hooks, Bus connections, and FSM state from known-good defaults.
//
// MISSING IMPLEMENTATIONS:
//   - Kernel-level read-only remounts (planned, not implemented)
//   - Non-critical process freezing (planned, not implemented)
//   - Audit log flushing before halt (planned, not implemented)
// =========================================================================
package emergency

import (
	"fmt"
	"sync"

	"github.com/fallofpheonix/phoenix/assurance/security/engine"
)

// KillSwitch manages the global system halt mechanism.
type KillSwitch struct {
	mu      sync.Mutex
	engaged bool
	warden  *engine.Warden
}

// NewKillSwitch creates the emergency halt mechanism.
// Called once during system startup, after Warden initialization.
//
// WORKFLOW: startup → NewWarden() → NewKillSwitch(warden) → ready for emergency use
func NewKillSwitch(w *engine.Warden) *KillSwitch {
	return &KillSwitch{
		warden: w,
	}
}

// Engage triggers the thermodynamic failsafe — the POINT OF NO RETURN.
// Once engaged, the Warden FSM is permanently locked and no further
// state transitions are possible until process restart.
//
// WORKFLOW: Called when:
//  1. Manual operator intervention (button press in dashboard)
//  2. Automatic trigger if confidence score drops below critical threshold
//  3. Recovery loop detects irrecoverable state corruption
//
// SEQUENCE:
//
//	Engage() → set engaged=true → Warden.Lock() → FSM frozen
//	→ All future Transition() calls return error "FSM LOCKOUT"
//	→ System is now in "preservation mode"
//
// IDEMPOTENT: Calling Engage() multiple times is safe — second call is a no-op.
func (k *KillSwitch) Engage() {
	k.mu.Lock()
	defer k.mu.Unlock()

	if k.engaged {
		return
	}

	fmt.Println("[KILL SWITCH] ENGAGED. HALTING ALL TRANSITIONS. ENTERING PRESERVATION MODE.")
	k.engaged = true

	// Lock the FSM
	k.warden.Lock()

	// In a real implementation, this would trigger kernel-level read-only remounts
	// and freeze all non-critical process execution.
}

// IsEngaged returns whether the KillSwitch has been triggered.
// Used by the Arbiter to check if the system is in preservation mode.
// If engaged, the Arbiter should NOT attempt any state transitions.
func (k *KillSwitch) IsEngaged() bool {
	k.mu.Lock()
	defer k.mu.Unlock()
	return k.engaged
}

// Status returns a human-readable status string for the KillSwitch.
// Used by the Game Server dashboard to display system status to operators.
func (k *KillSwitch) Status() string {
	k.mu.Lock()
	defer k.mu.Unlock()
	if k.engaged {
		return "SYSTEM HALTED: READ-ONLY PRESERVATION MODE ACTIVE"
	}
	return "KILL SWITCH ARMED: NOMINAL"
}
