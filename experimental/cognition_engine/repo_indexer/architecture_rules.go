package repo_indexer

// ArchitectureRule defines a constraint for the repository
type ArchitectureRule struct {
	Name        string
	Description string
	Severity    string
}

// EvaluateRule checks if a symbol graph violates a rule
func EvaluateRule(rule ArchitectureRule, graph *SymbolGraph) bool {
	// Implementation to follow based on graph analysis
	return true
}
