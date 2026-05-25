package governance

// ReadinessInput holds the metrics needed to calculate a readiness score.
type ReadinessInput struct {
    ObservationCycles int
    RuntimeHealth     float64
    DriftValue        float64
    CriticalAlerts    int
}

// CalculateReadinessScore evaluates the input metrics and returns a readiness status string.
func CalculateReadinessScore(input ReadinessInput) string {
    if input.CriticalAlerts > 0 {
        return "BLOCKED"
    }

    score := 0.0
    
    // Weight RuntimeHealth (e.g., 40%)
    score += input.RuntimeHealth * 40.0

    // Weight DriftValue (e.g., 40%) - lower drift is better.
    // We invert the drift value for scoring (1.0 is no drift, 0.0 is high drift).
    score += (1.0 - input.DriftValue) * 40.0

    // Weight ObservationCycles (e.g., 20%) - more cycles are better, cap at 10.
    cycleScore := float64(input.ObservationCycles)
    if cycleScore > 10 {
        cycleScore = 10
    }
    score += (cycleScore / 10.0) * 20.0

    if score >= 85.0 {
        return "READY"
    }
    if score >= 60.0 {
        return "REVIEW"
    }
    
    return "PENDING"
}
