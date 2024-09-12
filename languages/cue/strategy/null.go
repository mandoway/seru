package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
)

type NullReduction struct{}

func (r NullReduction) Apply(input *ast.File) []*ast.File {
	return removeApplicableStatements[*ast.EmbedDecl](input, func(node *ast.EmbedDecl) bool {
		basicLit, ok := node.Expr.(*ast.BasicLit)
		return ok && basicLit.Kind == token.NULL
	})
}
