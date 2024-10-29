package strategy

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"errors"
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
	evaluated, err := evaluate(expr)
	if err != nil {
		return nil, err
	}

	fields, err := evaluated.Fields()
	if err != nil {
		return nil, err
	}

	var decls []ast.Decl
	for fields.Next() {
		valueSource := fields.Value().Source()

		switch value := valueSource.(type) {
		case *ast.Field:
			decls = append(decls, value)
		case ast.Expr:
			label := fields.Label()
			decls = append(decls, &ast.Field{
				Label: ast.NewIdent(label),
				Value: value,
			})
		default:
			return nil, errors.New("unknown declaration type was empty")
		}
	}

	return decls, err
}
