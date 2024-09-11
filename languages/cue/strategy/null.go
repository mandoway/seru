package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
)

type NullReduction struct{}

func (r NullReduction) Apply(input *ast.File) []*ast.File {
	return removeApplicableStatements[*ast.EmbedDecl](input, func(node *ast.EmbedDecl) bool {
		return node.Expr.(*ast.BasicLit).Kind == token.NULL
	})
}
