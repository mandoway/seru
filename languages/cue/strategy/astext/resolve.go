package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/util/collection"
)

// ResolveIdentifierValueInExpression replaces all identifiers with their respective values
// Returns node with all identifiers replaced, if possible and whether a (sub-) expression was changed
func ResolveIdentifierValueInExpression(expr ast.Node) (ast.Node, bool) {
	switch typedExpr := expr.(type) {
	case *ast.Ident:
		resolved, changed := resolveIdentValue(typedExpr)
		if !changed {
			return expr, false
		}
		return resolved.(ast.Expr), true
	case *ast.Interpolation:
		expressions, changed := resolveExpressions(typedExpr.Elts)
		return NewInterpolation(expressions), changed
	case *ast.BinaryExpr:
		x, changedX := ResolveIdentifierValueInExpression(typedExpr.X)
		y, changedY := ResolveIdentifierValueInExpression(typedExpr.Y)
		return ast.NewBinExpr(typedExpr.Op, x.(ast.Expr), y.(ast.Expr)), changedX || changedY
	case *ast.UnaryExpr:
		x, changedX := ResolveIdentifierValueInExpression(typedExpr.X)
		return CopyUnaryExpression(typedExpr, x.(ast.Expr)), changedX
	case *ast.CallExpr:
		expressions, anyChanged := resolveExpressions(typedExpr.Args)
		return ast.NewCall(typedExpr.Fun, expressions...), anyChanged
	case *ast.IndexExpr:
		expression, changed := ResolveIdentifierValueInExpression(typedExpr.Index)
		return CopyIndexExpression(typedExpr, expression.(ast.Expr)), changed
	case *ast.SliceExpr:
		low, changedLow := ResolveIdentifierValueInExpression(typedExpr.Low)
		high, changedHigh := ResolveIdentifierValueInExpression(typedExpr.High)
		return CopySliceExpression(typedExpr, low.(ast.Expr), high.(ast.Expr)), changedLow || changedHigh
	case *ast.ListLit:
		expressions, anyChanged := resolveExpressions(typedExpr.Elts)
		return ast.NewList(expressions...), anyChanged
	case *ast.ParenExpr:
		expression, changed := ResolveIdentifierValueInExpression(typedExpr.X)
		return CopyParenExpression(typedExpr, expression.(ast.Expr)), changed
	default:
		return expr, false
	}
}

// resolveIdentValue returns the backing value of an identifier and whether a new value was found or not
func resolveIdentValue(node *ast.Ident) (ast.Node, bool) {
	resolvedValueNode := node.Node
	if resolvedValueNode == nil {
		return node, false
	}

	switch expr := resolvedValueNode.(type) {
	case *ast.Ident:
		value, _ := resolveIdentValue(expr)
		ast.SetRelPos(value, token.NoRelPos)
		return value, true
	case *ast.LetClause:
		ast.SetRelPos(expr.Expr, token.NoRelPos)
		value, _ := ResolveIdentifierValueInExpression(expr.Expr)
		return value, true
	default:
		ast.SetRelPos(expr, token.NoRelPos)
		return expr, true
	}
}

func resolveExpressions(items []ast.Expr) ([]ast.Expr, bool) {
	anyChanged := false
	return collection.MapSlice[ast.Expr, ast.Expr](items, func(it ast.Expr) (ast.Expr, error) {
		if ident, ok := it.(*ast.Ident); ok {
			value, changed := resolveIdentValue(ident)
			anyChanged = anyChanged || changed
			if !changed {
				return it, nil
			}
			return value.(ast.Expr), nil
		}

		return it, nil
	}), anyChanged
}
