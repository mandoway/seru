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
	parentIdentifiers := NewIdentifierSet()
	for _, part := range parts {
		parentIdentifiers.put(part)
	}
	return resolveIdentifierValueInExpressionWithBanlist(expr, parentIdentifiers)
}

func resolveIdentifierValueInExpressionWithBanlist(expr ast.Node, parentIdentifiers *IdentifierSet) (ast.Node, bool) {
	switch typedExpr := expr.(type) {
	case *ast.Ident:
		resolved, changed := resolveIdentValue(typedExpr, parentIdentifiers)
		if !changed {
			return expr, false
		}
		return resolved.(ast.Expr), true
	case *ast.Interpolation:
		expressions, changed := resolveExpressions(typedExpr.Elts, parentIdentifiers)
		return NewInterpolation(expressions), changed
	case *ast.BinaryExpr:
		x, changedX := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, parentIdentifiers)
		y, changedY := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Y, parentIdentifiers)
		return ast.NewBinExpr(typedExpr.Op, x.(ast.Expr), y.(ast.Expr)), changedX || changedY
	case *ast.UnaryExpr:
		x, changedX := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, parentIdentifiers)
		return CopyUnaryExpression(typedExpr, x.(ast.Expr)), changedX
	case *ast.CallExpr:
		expressions, anyChanged := resolveExpressions(typedExpr.Args, parentIdentifiers)
		return ast.NewCall(typedExpr.Fun, expressions...), anyChanged
	case *ast.IndexExpr:
		expression, changed := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Index, parentIdentifiers)
		return CopyIndexExpression(typedExpr, expression.(ast.Expr)), changed
	case *ast.SliceExpr:
		low, changedLow := resolveIdentifierValueInExpressionWithBanlist(typedExpr.Low, parentIdentifiers)
		high, changedHigh := resolveIdentifierValueInExpressionWithBanlist(typedExpr.High, parentIdentifiers)
		return CopySliceExpression(typedExpr, low.(ast.Expr), high.(ast.Expr)), changedLow || changedHigh
	case *ast.ListLit:
		expressions, anyChanged := resolveExpressions(typedExpr.Elts, parentIdentifiers)
		return ast.NewList(expressions...), anyChanged
	case *ast.ParenExpr:
		expression, changed := resolveIdentifierValueInExpressionWithBanlist(typedExpr.X, parentIdentifiers)
		return CopyParenExpression(typedExpr, expression.(ast.Expr)), changed
	default:
		return expr, false
	}
}

// resolveIdentValue returns the backing value of an identifier and whether a new value was found or not
func resolveIdentValue(node *ast.Ident, parentIdentifiers *IdentifierSet) (ast.Node, bool) {
	if parentIdentifiers.has(node.Name) {
		parentIdentifiers.markFound(node.Name)
		return node, false
	}

	resolvedValueNode := node.Node
	if resolvedValueNode == nil {
		return node, false
	}
	hasBannedIdentifier := containsParentIdentifier(resolvedValueNode, parentIdentifiers)
	isRecursive := isIdentContainedIn(node, resolvedValueNode)
	_, isImportSpec := resolvedValueNode.(*ast.ImportSpec)
	if hasBannedIdentifier || isRecursive || isImportSpec {
		return node, false
	}

	returnValue := resolvedValueNode

	parentIdentifiers.put(node.Name)
	switch expr := resolvedValueNode.(type) {
	case *ast.Ident:
		returnValue, _ = resolveIdentValue(expr, parentIdentifiers)
	case *ast.LetClause:
		returnValue, _ = resolveIdentifierValueInExpressionWithBanlist(expr.Expr, parentIdentifiers)
	case *ast.Field:
		if alias, ok := expr.Label.(*ast.Alias); ok {
			returnValue, _ = resolveIdentifierValueInExpressionWithBanlist(alias.Expr, parentIdentifiers)
		}
	}

	if parentIdentifiers.wasFound(node.Name) {
		parentIdentifiers.delete(node.Name)
		return node, false
	}
	parentIdentifiers.delete(node.Name)

	ast.SetRelPos(returnValue, token.NoRelPos)
	return returnValue, true
}

func resolveExpressions(items []ast.Expr, parentIdentifiers *IdentifierSet) ([]ast.Expr, bool) {
	anyChanged := false
	return collection.MapSlice[ast.Expr, ast.Expr](items, func(it ast.Expr) (ast.Expr, error) {
		if ident, ok := it.(*ast.Ident); ok {
			value, changed := resolveIdentValue(ident, parentIdentifiers)
			anyChanged = anyChanged || changed
			if !changed {
				return it, nil
			}
			return value.(ast.Expr), nil
		}

		return it, nil
	}), anyChanged
}

func isIdentContainedIn(ident *ast.Ident, node ast.Node) bool {
	return node.Pos().Offset() <= ident.Pos().Offset() && node.End().Offset() >= ident.End().Offset()
}

func containsParentIdentifier(node ast.Node, parentIdentifiers *IdentifierSet) bool {
	contained := false
	ast.Walk(node, func(node ast.Node) bool {
		if ident, ok := node.(*ast.Ident); ok && parentIdentifiers.has(ident.Name) {
			contained = true
			return false
		}
		return true
	}, nil)

	return contained
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
