package transform

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"github.com/mandoway/seru/languages/cue/language"
)

// ApplyTransformationToEveryApplicableStatement modifies a statement in-place
//
// # Casts a node to the given type and only calls transformOrNil if the type of node is equal to the type parameter T
func ApplyTransformationToEveryApplicableStatement[T ast.Node](input []byte, buildTransformation func(node T, parentSelector string) Transformation) []*ast.File {
	var (
		transformedFiles                 []*ast.File
		applicableStatementsInCurrentRun int
		lastSeenApplicableStatement      int
		modifiedStatementInCurrentRun    bool
		scopeStack                       ScopeStack
	)

	modifyApplicableStatement := func(cursor astutil.Cursor) bool {
		adjustScope(cursor, &scopeStack)

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
		parentSelectorOfCurrent := scopeStack.ScopeOfCurrentPosition(cursor)
		nodeTransformation := buildTransformation(filteredStatement, parentSelectorOfCurrent)
		modifiedStatementInCurrentRun = nodeTransformation.ProcessNode(cursor)

		return !modifiedStatementInCurrentRun
	}

	// Main reduction loop
	parser := language.Parser{}
	for {
		applicableStatementsInCurrentRun = 0
		modifiedStatementInCurrentRun = false
		scopeStack.Clear()

		workingCopy, _ := parser.Parse(input)
		transformedCode := astutil.Apply(workingCopy, modifyApplicableStatement, nil).(*ast.File)

		if !modifiedStatementInCurrentRun {
			break
		}

		transformedFiles = append(transformedFiles, transformedCode)
	}

	return transformedFiles
}

func adjustScope(cursor astutil.Cursor, scopeStack *ScopeStack) {
	if field, nodeIsField := cursor.Node().(*ast.Field); nodeIsField {
		ident, labelIsIdent := field.Label.(*ast.Ident)
		structLit, valueIsStruct := field.Value.(*ast.StructLit)

		if labelIsIdent && valueIsStruct {
			scopeStack.Push(ScopeLayer{
				Name:  ident.Name,
				Start: structLit.Lbrace,
				End:   structLit.Rbrace,
			})
		}
	}

	if top, err := scopeStack.Top(); err == nil && top.End.Before(cursor.Node().Pos()) {
		scopeStack.Pop()
	}
}
