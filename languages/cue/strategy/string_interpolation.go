package strategy

import (
	"cuelang.org/go/cue/ast"
	"errors"
	"github.com/mandoway/seru/languages/cue/strategy/astext"
	"github.com/mandoway/seru/languages/cue/strategy/eval"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type StringInterpolationReduction struct {
}

func (s StringInterpolationReduction) Apply(input []byte) []*ast.File {
	evaluator := eval.BuildStringEvaluator(input)

	var transformations []*ast.File
	transformations = append(transformations, evaluateFieldLabelInterpolations(input, evaluator)...)
	transformations = append(transformations, evaluateFieldValueInterpolations(input, evaluator)...)
	transformations = append(transformations, evaluateLetValueInterpolations(input, evaluator)...)
	return transformations
}

func evaluateFieldLabelInterpolations(input []byte, evaluator func(expr ast.Expr) (string, error)) []*ast.File {
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Field](input, func(node *ast.Field, _ string) transform.Transformation {
		result, err := evaluateInterpolation(node.Label, evaluator)
		if err != nil {
			return transform.NewNoopTransformation()
		}

		newField := astext.CopyFieldWithLabel(node, ast.NewIdent(result))
		return transform.NewReplacementTransformation(newField)
	})
}

func evaluateFieldValueInterpolations(input []byte, evaluator func(expr ast.Expr) (string, error)) []*ast.File {
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Field](input, func(node *ast.Field, _ string) transform.Transformation {
		result, err := evaluateInterpolation(node.Value, evaluator)
		if err != nil {
			return transform.NewNoopTransformation()
		}

		newField := astext.CopyFieldWithValue(node, ast.NewString(result))
		return transform.NewReplacementTransformation(newField)
	})
}

func evaluateLetValueInterpolations(input []byte, evaluator func(expr ast.Expr) (string, error)) []*ast.File {
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.LetClause](input, func(node *ast.LetClause, _ string) transform.Transformation {
		result, err := evaluateInterpolation(node.Expr, evaluator)
		if err != nil {
			return transform.NewNoopTransformation()
		}

		newLet := astext.CopyLetClauseWithValue(node, ast.NewString(result))
		return transform.NewReplacementTransformation(newLet)
	})
}

func evaluateInterpolation(node ast.Node, evaluator func(expr ast.Expr) (string, error)) (string, error) {
	valueInterpolation, ok := node.(*ast.Interpolation)
	if !ok {
		return "", errors.New("not interpolation")
	}

	result, err := evaluator(valueInterpolation)
	if err != nil {
		return "", err
	}

	return result, nil
}
