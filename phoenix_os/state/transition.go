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
