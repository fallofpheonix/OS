package contracts

import "fmt"

// ContractValidator ensures runtime compliance with contract definitions.
type ContractValidator struct {
	APILevel int
	Hash     string
}

// ValidatePolicy ensures a policy implementation matches the required contract.
func (v *ContractValidator) ValidatePolicy(p Policy) error {
	if p.Name() == "" {
		return fmt.Errorf("policy must have a name")
	}
	// Future: Check ContractHash matching
	return nil
}
