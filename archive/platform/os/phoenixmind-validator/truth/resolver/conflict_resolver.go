/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package resolver

import "github.com/fallofpheonix/phoenixmind-validator/truth/evidence"

// Priority map for truth states. Higher value is higher priority.
var statePriority = map[evidence.TruthState]int{
	evidence.REJECTED:  7,
	evidence.BLOCKED:   6,
	evidence.ESCALATED: 5,
	evidence.WARNING:   4,
	evidence.VALIDATED: 3,
	evidence.OBSERVED:  2,
	evidence.UNKNOWN:   1,
}

// ResolveConflict takes two truth states and returns the one with higher priority.
func ResolveConflict(stateA, stateB evidence.TruthState) evidence.TruthState {
	if statePriority[stateA] > statePriority[stateB] {
		return stateA
	}
	return stateB
}

// MergeTruth combines multiple pieces of evidence for an entity into a single, resolved TruthState.
func MergeTruth(evidenceSet []evidence.Evidence) evidence.TruthState {
	if len(evidenceSet) == 0 {
		return evidence.UNKNOWN
	}

	finalState := evidence.UNKNOWN
	for _, ev := range evidenceSet {
		finalState = ResolveConflict(finalState, ev.State)
	}

	return finalState
}
