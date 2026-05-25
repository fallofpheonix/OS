package validation

type State string

const (
	AllowState    State = "ALLOW"
	BlockState    State = "BLOCK"
	EscalateState State = "ESCALATE"
	RejectState   State = "REJECT"
)

type SecurityValidator struct{}
type TrustValidator struct{}
type PhaseValidator struct{}
