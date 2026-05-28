package repo_indexer

type SymbolGraph struct {
	Symbols map[string]string
	Edges   map[string][]string
}

func BuildGraph(symbols []Symbol) *SymbolGraph {
	g := &SymbolGraph{
		Symbols: make(map[string]string),
		Edges:   make(map[string][]string),
	}
	for _, s := range symbols {
		g.Symbols[s.Name] = s.Kind
	}
	return g
}
