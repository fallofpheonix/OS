package contracts

// Evidence defines the contract for forensic validation.
type Evidence interface {
	Score() float64
	Confidence() float64
	GetPID() int
	Verify() bool
	GetDivergence() int
	GetHashIntegrity() bool
}
