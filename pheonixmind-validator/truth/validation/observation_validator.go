package validation

type ObservationValidator struct{}

func (v *ObservationValidator) Validate(driftStr string) ValidationResult {
	// Simple drift parsing for demonstration
	drift := 0.0
	// Assume some parsing logic here
	if drift < 0.2 {
		return ValidationResult{
			Entity:     "Observation",
			State:      StateValidated,
			Valid:      true,
			Drift:      drift,
			Confidence: 0.9,
		}
	}
	return ValidationResult{
		Entity:   "Observation",
		State:    StateWarning,
		Valid:    false,
		Drift:    drift,
		Warnings: []string{"High drift detected"},
	}
}
