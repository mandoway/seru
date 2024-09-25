package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type NullReduction struct{}

func (r NullReduction) Apply(input []byte) []*ast.File {
	removeIfDeclIsNull := func(node *ast.EmbedDecl) transform.Transformation {
		basicLit, ok := node.Expr.(*ast.BasicLit)

		if !ok || basicLit.Kind != token.NULL {
			return transform.NewNoopTransformation()
		}

		return transform.NewDeleteTransformation()
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.EmbedDecl](input, removeIfDeclIsNull)
}
