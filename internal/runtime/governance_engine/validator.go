/*
 * PHOENIX MATRIX SOVEREIGN ARCHITECTURE
 * [STATUS]: 18-Repository Substrate Consolidated
 * [FUTURE ENHANCEMENT]: Needs continuous formal verification scaling and HDF5 vector optimizations.
 * [POTENTIAL LOOPHOLE]: Ensure strict hardware isolation when deploying to bare-metal. Watch for timing side-channels.
 * [ERROR PRONE AREA]: Concurrency bottlenecks in event bus and race conditions in cross-domain memory mappings.
 */
// =========================================================================
// WORKFLOW POSITION: UTILITY — GOVERNANCE VALIDATION
//
// The governance engine validates that the repository structure and
// mandatory axioms are present and correct.
//
// WORKFLOW:
//   ValidateRoadmapIntegrity(root) → check required directories exist
//   CheckMandatoryAxioms(root) → verify 6 mandatory axioms in GEMINI.md
//
// MANDATORY AXIOMS:
//   1. "Determinism is sacred"
//   2. "Replay is authoritative"
//   3. "AI is advisory"
//   4. "Control must remain bounded"
//   5. "Telemetry correctness"
//   6. "Never scale instability"
//
// PURPOSE: Ensures the codebase adheres to its foundational principles.
// If any axiom is missing, the system should not be deployed.
//
// LIMITATION: The axiom check is a text search — easily bypassed by
// adding the axiom strings as comments anywhere in the file.
// =========================================================================
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
