package transform

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"errors"
	"fmt"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/languages/cue/logging"
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
		pushIfFieldWithStruct(cursor, &scopeStack)

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

	adjustScopeAfter := func(cursor astutil.Cursor) bool {
		popIfFieldWithStruct(cursor, &scopeStack)
		return true
	}

	// Main reduction loop
	parser := language.Parser{}
	for {
		applicableStatementsInCurrentRun = 0
		modifiedStatementInCurrentRun = false
		scopeStack.Clear()

		workingCopy, _ := parser.Parse(input)
		transformedCode, err := applyOrRecover(workingCopy, modifyApplicableStatement, adjustScopeAfter)
		if err != nil {
			logging.CueError.Printf("Skipping candidate due to error during transformation: %s", err.Error())
			continue
		}

		if !modifiedStatementInCurrentRun {
			break
		}

		transformedFiles = append(transformedFiles, transformedCode)
	}

	return transformedFiles
}

func applyOrRecover(workingCopy *ast.File, beforeCB, afterCB func(astutil.Cursor) bool) (modified *ast.File, err error) {
	defer func() {
		if r := recover(); r != nil {
			modified = nil
			err = errors.New(fmt.Sprint(r))
		}
	}()

	return astutil.Apply(workingCopy, beforeCB, afterCB).(*ast.File), nil
}

func pushIfFieldWithStruct(cursor astutil.Cursor, stack *ScopeStack) {
	if scopeLayer, isField := getScopeLayerIfFieldWithStruct(cursor); isField {
		stack.Push(*scopeLayer)
	}
}

func popIfFieldWithStruct(cursor astutil.Cursor, stack *ScopeStack) {
	if _, isField := getScopeLayerIfFieldWithStruct(cursor); isField {
		stack.Pop()
	}
}

func getScopeLayerIfFieldWithStruct(cursor astutil.Cursor) (*ScopeLayer, bool) {
	if field, nodeIsField := cursor.Node().(*ast.Field); nodeIsField {
		ident, labelIsIdent := field.Label.(*ast.Ident)
		structLit, valueIsStruct := field.Value.(*ast.StructLit)

		if labelIsIdent && valueIsStruct {
			return &ScopeLayer{
				Name:  ident.Name,
				Start: structLit.Lbrace,
				End:   structLit.Rbrace,
			}, true
		}
	}

	return nil, false
}
