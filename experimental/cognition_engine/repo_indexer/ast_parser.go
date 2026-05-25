package repo_indexer

import (
	"go/ast"
	"go/parser"
	"go/token"
)

type Symbol struct {
	Name string
	Kind string
}

func ParseRepo(path string) ([]Symbol, error) {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, path, nil, parser.ParseComments)
	if err != nil {
		return nil, err
	}

	var symbols []Symbol
	for _, pkg := range pkgs {
		for _, file := range pkg.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				switch x := n.(type) {
				case *ast.FuncDecl:
					symbols = append(symbols, Symbol{Name: x.Name.Name, Kind: "function"})
				case *ast.TypeSpec:
					symbols = append(symbols, Symbol{Name: x.Name.Name, Kind: "type"})
				}
				return true
			})
		}
	}
	return symbols, nil
}
