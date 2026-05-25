package validation

type SecurityValidator struct{}

func (v *SecurityValidator) Validate(threat string) ValidationResult {
	if threat == "NONE" {
		return ValidationResult{
			Entity:     "Security",
			State:      StateValidated,
			Valid:      true,
			Confidence: 1.0,
		}
	}
	return ValidationResult{
		Entity: "Security",
		State:  StateBlocked,
		Valid:  false,
		Errors: []string{"Security violation: " + threat},
	}
}
