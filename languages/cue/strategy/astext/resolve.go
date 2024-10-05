package astext

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/util/collection"
	"strings"
)

// ResolveIdentifierValueInExpression replaces all identifiers with their respective values
// Returns node with all identifiers replaced, if possible and whether a (sub-) expression was changed
func ResolveIdentifierValueInExpression(expr ast.Node) (ast.Node, bool) {
	return resolveIdentifierValueInExpressionWithBanlist(expr, NewIdentifierSet())
}

func ResolveIdentifierValueInExpressionInScope(expr ast.Node, scope string) (ast.Node, bool) {
	parts := strings.Split(scope, ".")
	bannedIdentifiers := NewIdentifierSet()
	for _, part := range parts {
		bannedIdentifiers.put(part)
	}
	return resolveIdentifierValueInExpressionWithBanlist(expr, bannedIdentifiers)
}

func resolveIdentifierValueInExpressionWithBanlist(expr ast.Node, bannedIdentifiers *IdentifierSet) (ast.Node, bool) {
	switch typedExpr := expr.(type) {
	case *ast.Ident:
		resolved, changed := resolveIdentValue(typedExpr, bannedIdentifiers)
		if !changed {
			return expr, false
		}
		return resolved.(ast.Expr), true
	case *ast.Interpolation:
		expressions, changed := resolveExpressions(typedExpr.Elts, bannedIdentifiers)
		return NewInterpolation(expressions), changed
	case *ast.BinaryExpr:
		x, changedX := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, bannedIdentifiers)
		y, changedY := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Y, bannedIdentifiers)
		return ast.NewBinExpr(typedExpr.Op, x.(ast.Expr), y.(ast.Expr)), changedX || changedY
	case *ast.UnaryExpr:
		x, changedX := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, bannedIdentifiers)
		return CopyUnaryExpression(typedExpr, x.(ast.Expr)), changedX
	case *ast.CallExpr:
		expressions, anyChanged := resolveExpressions(typedExpr.Args, bannedIdentifiers)
		return ast.NewCall(typedExpr.Fun, expressions...), anyChanged
	case *ast.IndexExpr:
		expression, changed := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Index, bannedIdentifiers)
		return CopyIndexExpression(typedExpr, expression.(ast.Expr)), changed
	case *ast.SliceExpr:
		low, changedLow := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Low, bannedIdentifiers)
		high, changedHigh := resolveIdentifierValueInExpressionWithBanlist(typedExpr.High, bannedIdentifiers)
		return CopySliceExpression(typedExpr, low.(ast.Expr), high.(ast.Expr)), changedLow || changedHigh
	case *ast.ListLit:
		expressions, anyChanged := resolveExpressions(typedExpr.Elts, bannedIdentifiers)
		return ast.NewList(expressions...), anyChanged
	case *ast.ParenExpr:
		expression, changed := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, bannedIdentifiers)
		return CopyParenExpression(typedExpr, expression.(ast.Expr)), changed
	default:
		return expr, false
	}
}

// resolveIdentValue returns the backing value of an identifier and whether a new value was found or not
func resolveIdentValue(node *ast.Ident, bannedIdentifiers *IdentifierSet) (ast.Node, bool) {
	if bannedIdentifiers.has(node.Name) {
		bannedIdentifiers.markFound(node.Name)
		return node, false
	}
	resolvedValueNode := node.Node
	if resolvedValueNode == nil {
		return node, false
	}
	isRecursive := resolvedValueNode.Pos().Offset() <= node.Pos().Offset() && resolvedValueNode.End().Offset() >= node.End().Offset()
	if isRecursive {
		return node, false
	}

	expr := resolvedValueNode
	if let, ok := expr.(*ast.LetClause); ok {
		expr = let.Expr
	}
	value := expr

	bannedIdentifiers.put(node.Name)
	switch expr := resolvedValueNode.(type) {
	case *ast.Ident:
		value, _ = resolveIdentValue(expr, bannedIdentifiers)
	case *ast.LetClause:
		value, _ = resolveIdentifierValueInExpressionWithBanlist(expr.Expr, bannedIdentifiers)
	}

	if bannedIdentifiers.wasFound(node.Name) {
		bannedIdentifiers.delete(node.Name)
		return node, false
	}
	bannedIdentifiers.delete(node.Name)

	ast.SetRelPos(value, token.NoRelPos)
	return value, true
}

func resolveExpressions(items []ast.Expr, bannedIdentifiers *IdentifierSet) ([]ast.Expr, bool) {
	anyChanged := false
	return collection.MapSlice[ast.Expr, ast.Expr](items, func(it ast.Expr) (ast.Expr, error) {
		if ident, ok := it.(*ast.Ident); ok {
			value, changed := resolveIdentValue(ident, bannedIdentifiers)
			anyChanged = anyChanged || changed
			if !changed {
				return it, nil
			}
			return value.(ast.Expr), nil
		}

		return it, nil
	}), anyChanged
}

type IdentifierSet struct {
	foundByName map[string]bool
}

func NewIdentifierSet() *IdentifierSet {
	return &IdentifierSet{foundByName: map[string]bool{}}
}

func (set *IdentifierSet) put(key string) {
	set.foundByName[key] = false
}

func (set *IdentifierSet) markFound(key string) {
	set.foundByName[key] = true
}

func (set *IdentifierSet) wasFound(key string) bool {
	val, found := set.foundByName[key]
	return val && found
}

func (set *IdentifierSet) has(key string) bool {
	_, found := set.foundByName[key]
	return found
}

func (set *IdentifierSet) delete(key string) {
	delete(set.foundByName, key)
}
