/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package watcher

import (
	"strings"
	"github.com/fallofpheonix/PhoenixSimulation/schema"
)

// GuardianWatcher implements the ArchitecturalConstraint interface
// and acts as the gatekeeper for system integrity.
type GuardianWatcher struct {
	// ForbiddenPaths defines the root paths that the Builder is strictly 
	// prohibited from modifying (e.g., Warden logic, telemetry subsystem).
	ForbiddenPaths []string
}

// NewGuardianWatcher initializes the gatekeeper with hard-coded invariants.
func NewGuardianWatcher() *GuardianWatcher {
	return &GuardianWatcher{
		ForbiddenPaths: []string{
			"phoenix_os/warden/",   // Warden invariants are immutable
			"phoenix_os/telemetry/", // Telemetry plumbing is immutable
			"truth_ledger/",         // Ledger history is immutable
		},
	}
}

// IsViolating checks if the proposed patch attempts to modify restricted code paths.
func (g *GuardianWatcher) IsViolating(proposal schema.PatchProposal) bool {
	for _, path := range g.ForbiddenPaths {
		if strings.HasPrefix(proposal.TargetFile, path) {
			return true // Violation: Attempting to modify core security infrastructure
		}
	}
	return false
}
