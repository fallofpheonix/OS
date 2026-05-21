package common

// CalculateImportanceScore calculates the Importance Score based on weighted metrics.
// Experts recommend weighting criticality and entropy contribution highest.
// Criticality (0.6) is the primary driver to prevent mimicry.
func CalculateImportanceScore(rank, criticality, entropy, spread, depth float64) float64 {
	return 0.1*rank + 0.6*criticality + 0.1*entropy + 0.1*spread + 0.1*depth
}
