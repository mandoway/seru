package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/astext"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type ConstantPropagationReduction struct {
}

func (l ConstantPropagationReduction) Apply(input []byte) []*ast.File {
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Ident](input, func(node *ast.Ident, scope string) transform.Transformation {
		resolvedValueNode := astext.ResolveIdentifierInExpression(node)
		if resolvedValueNode == nil {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(resolvedValueNode)
	})
}
