package findIdorUrl

import (
	"go/ast"
	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

var Analyzer = &analysis.Analyzer{
	Name: "findIdORUrl",
	Doc: "findIdORUrl finds identifiers which name are Id or Url",
	Run: run,
	Requires: []*analysis.Analyzer{},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)
	nodeFilter := []ast.Node{ (*ast.Ident)(nil) }
	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
			case *ast.Ident:
				if n.Name == "Id" || n.Name == "Url" {
					pass.Reportf(n.Pos(), "identifier %s is not allowed", n.Name)
				}
		}
	})

	return nil, nil
}