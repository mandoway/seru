package strategy

import (
	"cuelang.org/go/cue/ast"
)

// LetReduction iterates over let-clauses and transforms them to regular fields
type LetReduction struct{}

// Apply returns variants of the input where exactly one let statement was transformed
func (r LetReduction) Apply(input *ast.File) []*ast.File {
	return applyTransformationToEveryApplicableStatement[*ast.LetClause](input, func(node *ast.LetClause) ast.Node {
		return &ast.Field{
			Label: ast.NewIdent(node.Ident.Name),
			Value: node.Expr,
		}
	})
}
