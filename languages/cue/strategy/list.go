package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type ListReduction struct {
}

func (l ListReduction) Apply(input []byte) []*ast.File {
	reduceToEmpty := func(node *ast.ListLit, _ string) transform.Transformation {
		if len(node.Elts) == 0 {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(ast.NewList())
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.ListLit](input, reduceToEmpty)
}
