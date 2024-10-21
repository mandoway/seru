package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

// LetReduction iterates over let-clauses and transforms them to regular fields
type LetReduction struct{}

// Apply returns variants of the input where exactly one let statement was transformed
func (r LetReduction) Apply(input []byte) []*ast.File {
	letClausePositionsInsideComprehension := collectLetClausesInsideComprehensions(input)

	createFieldFromLetClause := func(node *ast.LetClause, _ string) transform.Transformation {
		if _, found := letClausePositionsInsideComprehension[position(node)]; found {
			return transform.NewNoopTransformation()
		}

		return transform.NewReplacementTransformation(&ast.Field{
			Label: ast.NewIdent(node.Ident.Name),
			Value: node.Expr,
		})
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.LetClause](input, createFieldFromLetClause)
}

func collectLetClausesInsideComprehensions(input []byte) map[int]struct{} {
	startPositions := map[int]struct{}{}

	transform.ApplyTransformationToEveryApplicableStatement[*ast.Comprehension](input, func(node *ast.Comprehension, _ string) transform.Transformation {
		for _, clause := range node.Clauses {
			if c, ok := clause.(*ast.LetClause); ok {
				startPositions[position(c)] = struct{}{}
			}
		}

		return transform.NewNoopTransformation()
	})

	return startPositions
}

func position(node ast.Node) int {
	return node.Pos().Offset()
}
