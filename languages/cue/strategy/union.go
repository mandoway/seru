package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/token"
	"github.com/mandoway/seru/languages/cue/strategy/astext"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type UnionReduction struct {
}

func (t UnionReduction) Apply(input []byte) []*ast.File {

	var (
		resolvedOptionsByNodePos = map[int]*ExpressionUnion{}
		maxOptions               int
		transformations          []*ast.File
	)

	for i := 0; i == 0 || i < maxOptions; i++ {
		current := transform.ApplyTransformationToEveryApplicableStatement[*ast.BinaryExpr](input, func(node *ast.BinaryExpr, _ string) transform.Transformation {
			isOr := node.Op == token.OR
			if !isOr {
				return transform.NewNoopTransformation()
			}

			key := node.Pos().Offset()
			if i == 0 {
				if isSubExpression(node, resolvedOptionsByNodePos) {
					return transform.NewNoopTransformation()
				}
				options, err := NewExpressionUnion(node)
				if err != nil {
					return transform.NewNoopTransformation()
				}
				resolvedOptionsByNodePos[key] = options
				maxOptions = max(maxOptions, len(options.Expressions))
			}

			options, found := resolvedOptionsByNodePos[key]
			isDifferentExpr := options.EndOffset != node.End().Offset()
			if !found || isDifferentExpr || i >= len(options.Expressions) {
				return transform.NewNoopTransformation()
			}

			return transform.NewReplacementTransformation(options.Expressions[i])
		})

		transformations = append(transformations, current...)
	}

	return transformations
}

func isSubExpression(node *ast.BinaryExpr, resolvedOptionsByNodePos map[int]*ExpressionUnion) bool {
	start := node.Pos().Offset()
	end := node.End().Offset()

	for pos, exprUnion := range resolvedOptionsByNodePos {
		curStart := pos
		curEnd := exprUnion.EndOffset

		if curStart <= start && end <= curEnd {
			return true
		}
	}

	return false
}

type ExpressionUnion struct {
	Expressions []ast.Expr
	EndOffset   int
}

func NewExpressionUnion(expr ast.Expr) (*ExpressionUnion, error) {
	options, err := extractInlinedOptions(expr)
	if err != nil {
		return nil, err
	}
	return &ExpressionUnion{
		Expressions: options,
		EndOffset:   expr.End().Offset(),
	}, nil
}

func extractInlinedOptions(node ast.Expr) ([]ast.Expr, error) {
	if bin, ok := node.(*ast.BinaryExpr); ok && bin.Op == token.OR {
		leftOptions, err := extractInlinedOptions(bin.X)
		if err != nil {
			return nil, err
		}
		rightOptions, err := extractInlinedOptions(bin.Y)
		if err != nil {
			return nil, err
		}
		return append(leftOptions, rightOptions...), nil
	}

	returnedExpr := node
	if un, ok := node.(*ast.UnaryExpr); ok && un.Op == token.MUL {
		returnedExpr = un.X
	}

	expr, _ := astext.ResolveIdentifierValueInExpression(returnedExpr)

	return []ast.Expr{expr.(ast.Expr)}, nil
}
