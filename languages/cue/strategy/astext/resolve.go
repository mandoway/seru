package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/util/collection"
)

// ResolveIdentifierValueInExpression NOT COMPLETE
// Check source code for supported expressions
func ResolveIdentifierValueInExpression(expr ast.Node) ast.Node {
	switch typedExpr := expr.(type) {
	case *ast.Ident:
		resolved := resolveIdentValue(typedExpr)
		if resolved == nil {
			return nil
		}
		return resolved.(ast.Expr)
	case *ast.Interpolation:
		return NewInterpolation(resolveExpressions(typedExpr.Elts))
	case *ast.BinaryExpr:
		return ast.NewBinExpr(typedExpr.Op, ResolveIdentifierValueInExpression(typedExpr.X).(ast.Expr), ResolveIdentifierValueInExpression(typedExpr.Y).(ast.Expr))
	case *ast.UnaryExpr:
		return CopyUnaryExpression(typedExpr, ResolveIdentifierValueInExpression(typedExpr.X).(ast.Expr))
	case *ast.CallExpr:
		return ast.NewCall(typedExpr.Fun, resolveExpressions(typedExpr.Args)...)
	case *ast.IndexExpr:
		return CopyIndexExpression(typedExpr, ResolveIdentifierValueInExpression(typedExpr.Index).(ast.Expr))
	case *ast.SliceExpr:
		return CopySliceExpression(typedExpr, ResolveIdentifierValueInExpression(typedExpr.Low).(ast.Expr), ResolveIdentifierValueInExpression(typedExpr.High).(ast.Expr))
	case *ast.ListLit:
		return ast.NewList(resolveExpressions(typedExpr.Elts)...)
	case *ast.ParenExpr:
		return CopyParenExpression(typedExpr, ResolveIdentifierValueInExpression(typedExpr.X).(ast.Expr))
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
		return ResolveIdentifierValueInExpression(expr.Expr)
	default:
		ast.SetRelPos(expr, token.NoRelPos)
		return expr
	}
}

func resolveExpressions(items []ast.Expr) []ast.Expr {
	return collection.MapSlice[ast.Expr, ast.Expr](items, func(it ast.Expr) (ast.Expr, error) {
		if ident, ok := it.(*ast.Ident); ok {
			value := resolveIdentValue(ident)
			if value == nil {
				return it, nil
			}
			return value.(ast.Expr), nil
		}

		return it, nil
	})
}
