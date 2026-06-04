/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
package state

func isForwardTransition(current, next SystemState) bool {
	switch current {
	case StateSafe:
		return next == StateWatch
	case StateWatch:
		return next == StateAlert
	case StateAlert:
		return next == StateContain
	case StateContain:
		return next == StateRecovery
	case StateRecovery:
		return next == StateSafe
	default:
		return false
	}
}

func isValidTransition(current, next SystemState, isRollback bool) bool {
	if isRollback {
		return isRollbackTransition(current, next)
	}
	return isForwardTransition(current, next)
}
