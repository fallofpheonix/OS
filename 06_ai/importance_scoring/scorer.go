package importance_scoring

// Scorer calculates the Importance Score based on weighted metrics.
type Scorer struct {
	Weights [3]float64
}

// NewScorer creates a new scorer with default weights.
func NewScorer() *Scorer {
	return &Scorer{
		Weights: [3]float64{0.5, 0.3, 0.2}, // Entropy, Network, Resource
	}
}

// Calculate returns a normalized score [0, 1].
func (s *Scorer) Calculate(entropy, network, resource float64) float64 {
	return (entropy * s.Weights[0]) + (network * s.Weights[1]) + (resource * s.Weights[2])
}
