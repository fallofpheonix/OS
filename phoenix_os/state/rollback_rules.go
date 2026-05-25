package state

func isRollbackTransition(current, next SystemState) bool {
	switch current {
	case StateWatch:
		return next == StateSafe
	case StateAlert:
		return next == StateWatch
	case StateContain:
		return next == StateAlert
	case StateRecovery:
		return next == StateContain
	default:
		return false
	}
}
