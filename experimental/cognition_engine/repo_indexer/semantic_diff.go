package repo_indexer

type DiffResult struct {
	ChangedSymbols []string
	RiskLevel      string
}

func Compare(oldGraph, newGraph *SymbolGraph) DiffResult {
	return DiffResult{
		ChangedSymbols: []string{},
		RiskLevel:      "LOW",
	}
}
