package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"github.com/mandoway/seru/reduction"
)

// LetReduction iterates over let-clauses and transforms them to regular fields
type LetReduction struct {
	reduction.Strategy[ast.File]
}

// Apply returns variants of the input where exactly one let statement was transformed
func (r LetReduction) Apply(input *ast.File) []*ast.File {

	var transformed []*ast.File
	var seenClausesInThisRun, modifiedClauses int

	modifyNextLetClause := func(cursor astutil.Cursor) bool {
		// Filter applicable nodes
		letClause, ok := cursor.Node().(*ast.LetClause)
		if !ok {
			return true
		}
		// Filter already applied nodes
		seenClausesInThisRun++

		alreadyModifiedStatementInThisRun := modifiedClauses > len(transformed)
		statementWasModifiedInPreviousRun := seenClausesInThisRun-1 < modifiedClauses
		if alreadyModifiedStatementInThisRun || statementWasModifiedInPreviousRun {
			return false
		}

		fieldFromClause := &ast.Field{
			Label: ast.NewIdent(letClause.Ident.Name),
			Value: letClause.Expr,
		}

		cursor.Replace(fieldFromClause)

		modifiedClauses++
		return false
	}

	modifiedClauses = 0
	previous := -1
	for {
		seenClausesInThisRun = 0
		previous = modifiedClauses

		workingCopy := *input
		transformedCode := astutil.Apply(&workingCopy, modifyNextLetClause, nil).(*ast.File)

		if modifiedClauses == previous {
			break
		}

		transformed = append(transformed, transformedCode)
	}

	return transformed
}
