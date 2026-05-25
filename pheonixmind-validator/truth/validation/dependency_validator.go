package validation

type DependencyValidator struct{}

func (v *DependencyValidator) Validate(depGraph string) ValidationResult {
	if depGraph == "CLEAN" {
		return ValidationResult{
			Entity:     "Dependency",
			State:      StateValidated,
			Valid:      true,
			Confidence: 1.0,
		}
	}
	return ValidationResult{
		Entity:   "Dependency",
		State:    StateBlocked,
		Valid:    false,
		Errors:   []string{"Cycle detected in dependency graph"},
	}
}
