package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type NullReduction struct{}

func (r NullReduction) Apply(input []byte) []*ast.File {
	isDeclarationJustNull := func(node *ast.EmbedDecl) bool {
		basicLit, ok := node.Expr.(*ast.BasicLit)
		return ok && basicLit.Kind == token.NULL
	}
	return transform.RemoveApplicableStatements[*ast.EmbedDecl](input, isDeclarationJustNull)
}
