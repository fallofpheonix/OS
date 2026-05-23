package gov

import (
	"fmt"
	"os"
	"strings"
)

// ValidateRoadmapIntegrity checks if the core subsystems mentioned in documentation are present.
func ValidateRoadmapIntegrity(root string) ([]string, error) {
	requiredDirs := []string{
		"ai", "arbiter", "bus", "common", "guard", "ledger", "monitor", "tcs", "trace", "warden",
	}

	var missing []string
	for _, dir := range requiredDirs {
		path := fmt.Sprintf("%s/phoenix_os/%s", root, dir)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			missing = append(missing, dir)
		}
	}

	return missing, nil
}

// CheckMandatoryAxioms verifies if GEMINI.md exists and contains the 6 axioms.
func CheckMandatoryAxioms(root string) error {
	path := fmt.Sprintf("%s/GEMINI.md", root)
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	axioms := []string{
		"Determinism is sacred",
		"Replay is authoritative",
		"AI is advisory",
		"Control must remain bounded",
		"Telemetry correctness",
		"Never scale instability",
	}

	for _, axiom := range axioms {
		if !strings.Contains(string(content), axiom) {
			return fmt.Errorf("missing mandatory axiom: %s", axiom)
		}
	}

	return nil
}
