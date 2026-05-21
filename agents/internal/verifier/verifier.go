package verifier

import (
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
)

// Violation represents a boundary rule check failure
type Violation struct {
	FilePath string
	Import   string
	Rule     string
}

// VerifyImports scans all Go files under rootPath and checks them against PhoenixOS architectural boundaries.
func VerifyImports(rootPath string) ([]Violation, error) {
	var violations []Violation

	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip hidden dirs, build dirs, node_modules, etc.
		if info.IsDir() {
			name := info.Name()
			if strings.HasPrefix(name, ".") || name == "node_modules" || name == "vendor" || name == "artifacts" {
				return filepath.SkipDir
			}
			return nil
		}

		// Only inspect Go files
		if !strings.HasSuffix(info.Name(), ".go") {
			return nil
		}

		// Parse the file's AST (imports only)
		fset := token.NewFileSet()
		fileAST, err := parser.ParseFile(fset, path, nil, parser.ImportsOnly)
		if err != nil {
			// If we fail to parse, skip or log it. Don't fail the whole walk for one bad file.
			return nil
		}

		// Check the imports in this file
		for _, spec := range fileAST.Imports {
			// Clean import path (remove quotes)
			impPath := strings.Trim(spec.Path.Value, `"`)

			// Normalize paths to determine layers
			relPath, _ := filepath.Rel(rootPath, path)
			relPathLower := strings.ToLower(relPath)

			// Rule 1: No Research / Experiments -> Control / Kernel / Game loop paths
			// (barred from importing control, kernel, or game directly to trigger actions)
			if strings.Contains(relPathLower, "research") || strings.Contains(relPathLower, "experiments") {
				// Allow verifier or integration test files which might reside in agents/tests or agents/internal/verifier
				if !strings.Contains(relPathLower, "verifier") && !strings.Contains(relPathLower, "agents/tests") {
					if impPath == "phoenix/agents/internal/control" {
						violations = append(violations, Violation{
							FilePath: relPath,
							Import:   impPath,
							Rule:     "Rule 1 Violation: Research/Experiment code cannot import Control Agent directly.",
						})
					}
					if impPath == "phoenix/agents/internal/kernel" {
						violations = append(violations, Violation{
							FilePath: relPath,
							Import:   impPath,
							Rule:     "Rule 1 Violation: Research/Experiment code cannot import Kernel Agent directly.",
						})
					}
				}
			}

			// Rule 2: Strict path for Kernel Agent imports
			// Only phoenix/sentinel (Runtime orchestrator) and agents internal/kernel or tests/verifier can import Kernel Agent
			if impPath == "phoenix/agents/internal/kernel" {
				isAllowedImport := false
				if strings.Contains(relPathLower, "internal/kernel") ||
					strings.Contains(relPathLower, "internal/verifier") ||
					strings.Contains(relPathLower, "agents/tests") ||
					strings.Contains(relPathLower, "sentinel") ||
					strings.Contains(relPathLower, "agents/src") {
					isAllowedImport = true
				}
				if !isAllowedImport {
					violations = append(violations, Violation{
						FilePath: relPath,
						Import:   impPath,
						Rule:     "Rule 2 Violation: Kernel Agent can only be imported by internal/kernel, verifier, agents tests, or sentinel orchestrator.",
					})
				}
			}

			// Rule 3: Telemetry Agent cannot import Control or Kernel Agents
			if strings.Contains(relPathLower, "internal/telemetry") {
				if impPath == "phoenix/agents/internal/control" || impPath == "phoenix/agents/internal/kernel" {
					violations = append(violations, Violation{
						FilePath: relPath,
						Import:   impPath,
						Rule:     "Rule 3 Violation: Telemetry Agent cannot import Control or Kernel Agents directly.",
					})
				}
			}

			// Rule 4: No direct External -> Kernel path.
			// External layer packages (like any surface/API/gateway client) cannot import kernel agent.
			// We define "external" as anything outside the agents module or phoenix_os core (sentinel, bus, monitor, warden, arbiter).
			if impPath == "phoenix/agents/internal/kernel" {
				// If the file path is not within agents/ or phoenix_os/
				isInCore := strings.HasPrefix(relPath, "agents/") || strings.HasPrefix(relPath, "phoenix_os/")
				if !isInCore {
					violations = append(violations, Violation{
						FilePath: relPath,
						Import:   impPath,
						Rule:     "Rule 4 Violation: External components cannot import Kernel Agent directly.",
					})
				}
			}
		}

		return nil
	})

	return violations, err
}
