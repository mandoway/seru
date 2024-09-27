package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type RedundantNestingReduction struct {
}

func (r RedundantNestingReduction) Apply(input []byte) []*ast.File {
	removeOneNestedLayer := func(node *ast.StructLit, _ string) transform.Transformation {
		if len(node.Elts) != 1 {
			return transform.NewNoopTransformation()
		}

		elementAsEmbedDecl, ok := node.Elts[0].(*ast.EmbedDecl)
		if !ok {
			return transform.NewNoopTransformation()
		}

		elementExprAsStructLit, ok := elementAsEmbedDecl.Expr.(*ast.StructLit)
		if !ok {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(&ast.StructLit{
			Lbrace: node.Lbrace,
			Elts:   elementExprAsStructLit.Elts,
		})
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.StructLit](input, removeOneNestedLayer)
}
