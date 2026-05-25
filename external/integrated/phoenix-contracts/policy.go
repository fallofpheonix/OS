package contracts

// Policy defines the interface for all PhoenixOS decision policies.
type Policy interface {
	Name() string
	Version() Version
	Evaluate(ctx PolicyContext) (PolicyResult, error)
}

// PolicyResult contains the outcome of a policy evaluation.
type PolicyResult struct {
	Score        float64
	Class        ActuationClass
	Confidence   float64
	Rationale    string
	ContractHash string
}
