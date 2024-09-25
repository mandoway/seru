package transform

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"github.com/mandoway/seru/languages/cue/language"
)

// ApplyTransformationToEveryApplicableStatement modifies a statement in-place
//
// # Casts a node to the given type and only calls transformOrNil if the type of node is equal to the type parameter T
func ApplyTransformationToEveryApplicableStatement[T ast.Node](input []byte, buildTransformation func(node T) Transformation) []*ast.File {
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

		// Actual Transformation
		// No need to continue when we only ever modify one statement per run
		nodeTransformation := buildTransformation(filteredStatement)
		modifiedStatementInCurrentRun = nodeTransformation.ProcessNode(cursor)

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
