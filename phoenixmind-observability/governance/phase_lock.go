package governance

// PhaseStatus represents the phase status.
type PhaseStatus string

const (
    CLOSED PhaseStatus = "CLOSED"
    ACTIVE PhaseStatus = "ACTIVE"
    LOCKED PhaseStatus = "LOCKED"
)

// PhaseLock manages phase transitions.
type PhaseLock struct {
    CurrentPhase string
    Status       PhaseStatus
}
