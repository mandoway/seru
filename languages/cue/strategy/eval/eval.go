package eval

import (
	"cuelang.org/go/cue"
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/cuecontext"
	"cuelang.org/go/cue/format"
)

func BuildBooleanEvaluator(bytes []byte) func(expr ast.Expr) (bool, error) {
	ctx := cuecontext.New()
	scope := cue.Scope(ctx.CompileBytes(bytes))

	return func(expr ast.Expr) (bool, error) {
		formattedExpr, err := format.Node(expr)
		if err != nil {
			return false, nil
		}
		return ctx.CompileBytes(formattedExpr, scope).Bool()
	}
}
