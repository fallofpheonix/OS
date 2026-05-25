package validation

type TruthState string

const (
	StateValidated TruthState = "VALIDATED"
	StateWarning   TruthState = "WARNING"
	StateEscalated TruthState = "ESCALATED"
	StateBlocked   TruthState = "BLOCKED"
)

type ValidationResult struct {
	Entity        string
	State         TruthState
	Valid         bool
	Confidence    float64
	Risk          float64
	Drift         float64
	EvidenceCount int
	Errors        []string
	Warnings      []string
}
