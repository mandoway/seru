package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/util/collection"
)

// ResolveIdentifierInExpression NOT COMPLETE
// Check source code for supported expressions
func ResolveIdentifierInExpression(expr ast.Expr) ast.Expr {
	switch typedExpr := expr.(type) {
	case *ast.Ident:
		return resolveIdentValue(typedExpr).(ast.Expr)
	case *ast.Interpolation:
		return NewInterpolation(resolveExpressions(typedExpr.Elts))
	case *ast.BinaryExpr:
		return ast.NewBinExpr(typedExpr.Op, ResolveIdentifierInExpression(typedExpr.X), ResolveIdentifierInExpression(typedExpr.Y))
	case *ast.UnaryExpr:
		return CopyUnaryExpression(typedExpr, ResolveIdentifierInExpression(typedExpr.X))
	case *ast.CallExpr:
		return ast.NewCall(typedExpr.Fun, resolveExpressions(typedExpr.Args)...)
	case *ast.IndexExpr:
		return CopyIndexExpression(typedExpr, ResolveIdentifierInExpression(typedExpr.Index))
	case *ast.SliceExpr:
		return CopySliceExpression(typedExpr, ResolveIdentifierInExpression(typedExpr.Low), ResolveIdentifierInExpression(typedExpr.High))
	case *ast.ListLit:
		return ast.NewList(resolveExpressions(typedExpr.Elts)...)
	case *ast.ParenExpr:
		return CopyParenExpression(typedExpr, ResolveIdentifierInExpression(typedExpr.X))
	default:
		return expr
	}
}

func resolveIdentValue(node *ast.Ident) ast.Node {
	resolvedValueNode := node.Node
	if resolvedValueNode == nil {
		return nil
	}

	switch expr := resolvedValueNode.(type) {
	case *ast.Ident:
		return resolveIdentValue(expr)
	case *ast.LetClause:
		ast.SetRelPos(expr.Expr, token.NoRelPos)
		return expr.Expr
	default:
		ast.SetRelPos(expr, token.NoRelPos)
		return expr
	}
}

func resolveExpressions(items []ast.Expr) []ast.Expr {
	return collection.MapSlice[ast.Expr, ast.Expr](items, func(it ast.Expr) (ast.Expr, error) {
		if ident, ok := it.(*ast.Ident); ok {
			return resolveIdentValue(ident).(ast.Expr), nil
		}

		return it, nil
	})
}
