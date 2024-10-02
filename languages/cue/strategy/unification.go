package strategy

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/eval"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
	"github.com/mandoway/seru/util/collection"
)

type UnificationReduction struct {
}

func (u UnificationReduction) Apply(input []byte) []*ast.File {
	evaluate := eval.BuildEvaluator(input)

	return transform.ApplyTransformationToEveryApplicableStatement[*ast.BinaryExpr](input, func(node *ast.BinaryExpr, _ string) transform.Transformation {
		isAnd := node.Op == token.AND
		if !isAnd {
			return transform.NewNoopTransformation()
		}

		leftDecls, err := evaluateAsStructLit(node.X, evaluate)
		if err != nil {
			return transform.NewNoopTransformation()
		}

		rightDecls, err := evaluateAsStructLit(node.Y, evaluate)
		if err != nil {
			return transform.NewNoopTransformation()
		}

		combinedDecls := append(collection.MapToInterface(leftDecls), collection.MapToInterface(rightDecls)...)
		unified := ast.NewStruct(combinedDecls...)
		return transform.NewReplacementTransformation(unified)
	})
}

func evaluateAsStructLit(expr ast.Expr, evaluate func(expr ast.Expr) (cue.Value, error)) ([]ast.Decl, error) {
	value, err := evaluate(expr)
	if err != nil {
		return nil, err
	}

	fields, err := value.Fields()
	if err != nil {
		return nil, err
	}

	var decls []ast.Decl
	for fields.Next() {
		decls = append(decls, fields.Value().Source().(ast.Decl))
	}

	return decls, err
}
