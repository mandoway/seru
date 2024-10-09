package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type TrivialIfReduction struct{}

func (i TrivialIfReduction) Apply(input []byte) []*ast.File {
	alwaysFalse := func(node *ast.IfClause, _ string) transform.Transformation {
		if isBooleanLit(node) {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(&ast.IfClause{
			If:        node.If,
			Condition: ast.NewBool(false),
		})
	}
	alwaysTrue := func(node *ast.IfClause, _ string) transform.Transformation {
		if isBooleanLit(node) {
			return transform.NewNoopTransformation()
		}
		return transform.NewReplacementTransformation(&ast.IfClause{
			If:        node.If,
			Condition: ast.NewBool(true),
		})
	}

	var transformations []*ast.File
	transformations = append(transformations, transform.ApplyTransformationToEveryApplicableStatement[*ast.IfClause](input, alwaysFalse)...)
	transformations = append(transformations, transform.ApplyTransformationToEveryApplicableStatement[*ast.IfClause](input, alwaysTrue)...)
	return transformations
}

func isBooleanLit(node *ast.IfClause) bool {
	lit, ok := node.Condition.(*ast.BasicLit)
	if !ok {
		return false
	}

	return lit.Kind == token.TRUE || lit.Kind == token.FALSE
}
