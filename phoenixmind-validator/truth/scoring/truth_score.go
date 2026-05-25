package scoring

import "github.com/fallofpheonix/phoenix-os/phoenixmind-validator/truth/evidence"

// Score represents the calculated belief in a piece of evidence.
type Score struct {
	Confidence float64 // 0.0 to 1.0
	Risk       float64 // 0.0 to 1.0
	Drift      float64 //
}

// CalculateTruthScore evaluates a piece of evidence and assigns a score.
func CalculateTruthScore(ev evidence.Evidence) Score {
	confidence := 0.5 // Default confidence
	risk := 0.5       // Default risk

	switch ev.Source {
	case "replay":
		confidence = 0.95
		risk = 0.1
	case "runtime_audit":
		confidence = 0.8
		risk = 0.3
	case "security_scan":
		confidence = 0.9
		risk = 0.7 // High risk if vulnerabilities are found
	}

	return Score{
		Confidence: confidence,
		Risk:       risk,
		Drift:      0.0, // To be calculated separately
	}
}
