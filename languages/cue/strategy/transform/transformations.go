package transform

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"github.com/mandoway/seru/languages/cue/language"
)

// ModifyApplicableStatements modifies a statement in-place
//
// # Casts a node to the given type and only calls transformOrNil if the type of node is equal to the type parameter T
//
// input full AST of current program
// transformOrNil returns modified node or nil if no modification should be applied to current node
func ModifyApplicableStatements[T ast.Node](input []byte, transformOrNil func(node T) ast.Node) []*ast.File {
	return applyTransformationToEveryApplicableStatement(input, func(node T, cursor astutil.Cursor) bool {
		transformed := transformOrNil(node)

		if transformed == nil {
			return false
		}

		cursor.Replace(transformed)
		return true
	})
}

func RemoveApplicableStatements[T ast.Node](input []byte, isApplicable func(node T) bool) []*ast.File {
	return applyTransformationToEveryApplicableStatement(input, func(node T, cursor astutil.Cursor) bool {
		if isApplicable(node) {
			cursor.Delete()
			return true
		}
		return false
	})
}

func applyTransformationToEveryApplicableStatement[T ast.Node](input []byte, action func(node T, cursor astutil.Cursor) bool) []*ast.File {
	var (
		transformedFiles                 []*ast.File
		applicableStatementsInCurrentRun int
		lastSeenApplicableStatement      int
		modifiedStatementInCurrentRun    bool
	)

	modifyApplicableStatement := func(cursor astutil.Cursor) bool {
		// Filter applicable nodes
		filteredStatement, ok := cursor.Node().(T)
		if !ok {
			return true
		}

		// Filter already applied nodes
		statementWasCheckedInPreviousRun := applicableStatementsInCurrentRun < lastSeenApplicableStatement
		applicableStatementsInCurrentRun++
		if modifiedStatementInCurrentRun {
			return false
		}
		if statementWasCheckedInPreviousRun {
			return true
		}
		lastSeenApplicableStatement = max(lastSeenApplicableStatement, applicableStatementsInCurrentRun)

		// Actual transformation
		// No need to continue when we only ever modify one statement per run
		modifiedStatementInCurrentRun = action(filteredStatement, cursor)

		return !modifiedStatementInCurrentRun
	}

	// Main reduction loop
	parser := language.Parser{}
	for {
		applicableStatementsInCurrentRun = 0
		modifiedStatementInCurrentRun = false

		workingCopy, _ := parser.Parse(input)
		transformedCode := astutil.Apply(workingCopy, modifyApplicableStatement, nil).(*ast.File)

		if !modifiedStatementInCurrentRun {
			break
		}

		transformedFiles = append(transformedFiles, transformedCode)
	}

	return transformedFiles
}
