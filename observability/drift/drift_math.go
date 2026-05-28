package drift

import (
	"math"
)

// CalculateStatisticalDrift computes Z-score based drift analysis
func CalculateStatisticalDrift(data []float64) float64 {
	if len(data) == 0 {
		return 0.0
	}
	var sum float64
	for _, v := range data {
		sum += v
	}
	mean := sum / float64(len(data))
	
	var sqDiffSum float64
	for _, v := range data {
		sqDiffSum += math.Pow(v-mean, 2)
	}
	stdDev := math.Sqrt(sqDiffSum / float64(len(data)))
	
	// Return normalized drift coefficient
	if stdDev == 0 {
		return 0.0
	}
	return (data[len(data)-1] - mean) / stdDev
}
