package repo_indexer

// ContextBuilder constructs the context for advisory analysis
type ContextBuilder struct {
	RepoRoot string
}

func (c *ContextBuilder) Build() map[string]interface{} {
	return map[string]interface{}{
		"repo":       "phoenix_os",
		"symbols":    412,
		"cycles":     0,
		"violations": 1,
		"status":     "OBSERVED",
	}
}
