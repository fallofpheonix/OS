package validators

import (
	"fmt"
	"math"
	"github.com/fallofpheonix/PheonixTruth/src"
)

// EntropyValidator implements Dynamic Information Flow Control (DIFC) by flagging high-entropy data flows.
type EntropyValidator struct {
	Threshold float64
}

// NewEntropyValidator creates a new EntropyValidator with a default threshold.
func NewEntropyValidator(threshold float64) *EntropyValidator {
	if threshold == 0 {
		threshold = 7.5 // Default threshold for high entropy (e.g. encrypted or compressed data)
	}
	return &EntropyValidator{
		Threshold: threshold,
	}
}

// Name returns the validator name.
func (v *EntropyValidator) Name() string {
	return "EntropyValidator"
}

// Validate calculates the Shannon entropy of the payload and checks if it violates DIFC rules.
func (v *EntropyValidator) Validate(entry *ledger.LedgerEntry) ValidationResult {
	entropy := calculateShannonEntropy(entry.Payload)

	// DIFC Rule: High entropy data from untrusted sources should not flow to restricted sinks.
	// For this first implementation, we flag any high-entropy payload that might indicate
	// unauthorized encryption or data exfiltration.
	
	if entropy > v.Threshold {
		// In a full DIFC implementation, we would check entry.Source and the destination in the payload.
		// For now, we flag high entropy as a potential anomaly.
		return ValidationResult{
			Valid:  false,
			Reason: fmt.Sprintf("High entropy detected: %.2f > %.2f (Possible encrypted exfiltration or ransomware)", entropy, v.Threshold),
		}
	}

	return ValidationResult{Valid: true}
}

// Reset is a no-op for the stateless EntropyValidator.
func (v *EntropyValidator) Reset() {}

// calculateShannonEntropy computes the Shannon entropy of a byte slice.
func calculateShannonEntropy(data []byte) float64 {
	if len(data) == 0 {
		return 0
	}

	frequencies := make(map[byte]int)
	for _, b := range data {
		frequencies[b]++
	}

	var entropy float64
	for _, count := range frequencies {
		p := float64(count) / float64(len(data))
		entropy -= p * math.Log2(p)
	}

	return entropy
}
