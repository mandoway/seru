package strategy

import (
	"cuelang.org/go/cue/ast"
	"fmt"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type ConstantPropagationReduction struct {
}

func (l ConstantPropagationReduction) Apply(input []byte) []*ast.File {
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Ident](input, func(node *ast.Ident, scope string) transform.Transformation {
		fmt.Println(scope)
		resolvedValueNode := resolveIdentValue(node)
		if resolvedValueNode == nil {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(resolvedValueNode)
	})
}

func resolveIdentValue(node *ast.Ident) ast.Node {
	resolvedValueNode := node.Node
	if resolvedValueNode == nil {
		return nil
	}

	switch expr := resolvedValueNode.(type) {
	case *ast.LetClause:
		return expr.Expr
	case *ast.Ident:
		return resolveIdentValue(expr)
	default:
		return expr
	}
}
