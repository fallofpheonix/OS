package validation

type Validator interface {
	Validate(entity string) ValidationResult
}

type Registry struct {
	validators map[string]Validator
}

func NewRegistry() *Registry {
	return &Registry{
		validators: make(map[string]Validator),
	}
}

func (r *Registry) Register(name string, v Validator) {
	r.validators[name] = v
}

func (r *Registry) RunAll(entity string) []ValidationResult {
	results := make([]ValidationResult, 0, len(r.validators))
	for _, v := range r.validators {
		results = append(results, v.Validate(entity))
	}
	return results
}
