package evidence

// TruthState defines the current validity state of a system observation.
type TruthState string

const (
	UNKNOWN    TruthState = "UNKNOWN"
	OBSERVED   TruthState = "OBSERVED"
	VALIDATED  TruthState = "VALIDATED"
	WARNING    TruthState = "WARNING"
	ESCALATED  TruthState = "ESCALATED"
	BLOCKED    TruthState = "BLOCKED"
	REJECTED   TruthState = "REJECTED"
)

// Evidence represents a single unit of system observation.
type Evidence struct {
	Entity        string     `json:"entity"`
	Source        string     `json:"source"`
	EvidenceCount int        `json:"evidence_count"`
	Confidence    float64    `json:"confidence"`
	Drift         float64    `json:"drift"`
	Risk          float64    `json:"risk"`
	Status        TruthState `json:"status"`
}
