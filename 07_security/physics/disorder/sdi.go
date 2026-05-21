package disorder

import (
	"math"
)

// CalculateSDI computes the Security Disorder Index using Shannon Entropy.
// In this context, probabilities p_s are derived from event category distributions.
func CalculateSDI(distribution map[string]float64) float64 {
	var sdi float64
	var total float64

	for _, count := range distribution {
		total += count
	}

	if total == 0 {
		return 0.0
	}

	for _, count := range distribution {
		p := count / total
		if p > 0 {
			sdi -= p * math.Log(p)
		}
	}

	return sdi
}
