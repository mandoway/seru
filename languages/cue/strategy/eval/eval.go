package eval

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
	"github.com/mandoway/seru/languages/cue/strategy/astext"
)

func BuildBooleanEvaluator(bytes []byte) func(expr ast.Expr) (bool, error) {
	eval := BuildEvaluator(bytes)

	return func(expr ast.Expr) (bool, error) {
		result, err := eval(expr)
		if err != nil {
			return false, err
		}
		return result.Bool()
	}
}

func BuildStringEvaluator(bytes []byte) func(expr ast.Expr) (string, error) {
	eval := BuildEvaluator(bytes)
	return func(expr ast.Expr) (string, error) {
		result, err := eval(expr)
		if err != nil {
			return "", err
		}
		return result.String()
	}
}

func BuildEvaluator(bytes []byte) func(expr ast.Expr) (cue.Value, error) {
	ctx := cuecontext.New()
	scope := cue.Scope(ctx.CompileBytes(bytes))

	return func(expr ast.Expr) (cue.Value, error) {
		resolved, _ := astext.ResolveIdentifierValueInExpression(expr)
		formattedExpr, err := format.Node(resolved)
		if err != nil {
			return cue.Value{}, err
		}
		return ctx.CompileBytes(formattedExpr, scope), nil
	}
}
