package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

// LetReduction iterates over let-clauses and transforms them to regular fields
type LetReduction struct{}

// Apply returns variants of the input where exactly one let statement was transformed
func (r LetReduction) Apply(input []byte) []*ast.File {
	createFieldFromLetClause := func(node *ast.LetClause) ast.Node {
		return &ast.Field{
			Label: ast.NewIdent(node.Ident.Name),
			Value: node.Expr,
		}
	}
	return transform.ModifyApplicableStatements[*ast.LetClause](input, createFieldFromLetClause)
}
