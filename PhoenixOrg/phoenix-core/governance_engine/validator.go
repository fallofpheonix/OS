package gov

import (
	"fmt"
	"os"
	"strings"
)

// ValidateRoadmapIntegrity checks if the core subsystems mentioned in documentation are present.
func ValidateRoadmapIntegrity(root string) ([]string, error) {
	requiredDirs := map[string]string{
		"ai":      "phoenix_os/ai",
		"bus":     "phoenix_os/bus",
		"common":  "phoenix_os/common",
		"ledger":  "phoenix_os/ledger",
		"arbiter": "03_repositories/integrated/phoenix-control/arbiter",
		"warden":  "03_repositories/integrated/phoenix-control/warden",
		"guard":   "03_repositories/integrated/phoenix-runtime/guard",
		"monitor": "03_repositories/integrated/phoenix-logic/monitor",
		"tcs":     "03_repositories/integrated/phoenix-logic/tcs",
		"trace":   "03_repositories/integrated/phoenix-logic/trace",
	}

	var missing []string
	for name, relPath := range requiredDirs {
		path := fmt.Sprintf("%s/%s", root, relPath)
		if _, err := os.Stat(path); os.IsNotExist(err) {
			missing = append(missing, name)
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
