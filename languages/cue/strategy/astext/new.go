package astext

import "cuelang.org/go/cue/ast"

func NewInterpolation(items []ast.Expr) *ast.Interpolation {
	return &ast.Interpolation{
		Elts: items,
	}
}
