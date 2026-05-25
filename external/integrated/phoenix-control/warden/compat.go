package warden

// Legacy aliases for backward compatibility with earlier Phase 3 tracks.
// These should be phased out in favor of the new FSM states.
const (
	StateNormal     SystemState = StateSafe
	StateSuspicious SystemState = StateWatch
	StateContained  SystemState = StateContain
)
