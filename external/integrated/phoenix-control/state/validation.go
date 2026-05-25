package state

import "fmt"

// RecoveryValidator ensures that a transition from RECOVERY to SAFE is valid.
type RecoveryValidator struct {
	Checklist []string
}

func NewRecoveryValidator() *RecoveryValidator {
	return &RecoveryValidator{
		Checklist: []string{
			"kernel_integrity_verified",
			"process_quarantine_cleared",
			"telemetry_stream_stable",
		},
	}
}

// ValidateRecovery checks if the provided evidence satisfies recovery requirements.
func (v *RecoveryValidator) ValidateRecovery(evidence map[string]bool) error {
	for _, item := range v.Checklist {
		if !evidence[item] {
			return fmt.Errorf("recovery validation failed: missing evidence for %s", item)
		}
	}
	return nil
}
