package strategy

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
)

func transformApplicableStatements[T ast.Node](input *ast.File, transform func(node T) ast.Node) []*ast.File {
	return applyTransformationToEveryApplicableStatement(input, func(node T, cursor astutil.Cursor) {
		transformed := transform(node)
		cursor.Replace(transformed)
	})
}

func removeApplicableStatements[T ast.Node](input *ast.File, predicate func(node T) bool) []*ast.File {
	return applyTransformationToEveryApplicableStatement(input, func(node T, cursor astutil.Cursor) {
		if predicate(node) {
			cursor.Delete()
		}
	})
}

func applyTransformationToEveryApplicableStatement[T ast.Node](input *ast.File, action func(node T, cursor astutil.Cursor)) []*ast.File {
	var (
		transformedFiles                 []*ast.File
		applicableStatementsInCurrentRun int
		modifiedStatementInCurrentRun    bool
	)

	modifyApplicableClause := func(cursor astutil.Cursor) bool {
		// Filter applicable nodes
		filteredStatement, ok := cursor.Node().(T)
		if !ok {
			return true
		}

		// Filter already applied nodes
		applicableStatementsInCurrentRun++
		statementWasModifiedInPreviousRun := applicableStatementsInCurrentRun-1 < len(transformedFiles)
		if modifiedStatementInCurrentRun || statementWasModifiedInPreviousRun {
			return false
		}

		// Actual transformation
		action(filteredStatement, cursor)

		// No need to continue when we only ever modify one statement per run
		modifiedStatementInCurrentRun = true
		return false
	}

	// Main reduction loop
	for {
		applicableStatementsInCurrentRun = 0
		modifiedStatementInCurrentRun = false

		workingCopy := *input
		transformedCode := astutil.Apply(&workingCopy, modifyApplicableClause, nil).(*ast.File)

		if !modifiedStatementInCurrentRun {
			break
		}

		transformedFiles = append(transformedFiles, transformedCode)
	}

	return transformedFiles
}
