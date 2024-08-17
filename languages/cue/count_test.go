package main

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/scanner"
	"cuelang.org/go/cue/token"
	"fmt"
	"github.com/mandoway/seru/languages/cue/context"
	"testing"
)

func TestParser(t *testing.T) {
	file := `
		data: forms: [{
		fooID: "00-0000001"
	}]
	
	form1040: (#compute & {in: data}).out
	
	#K1: {
		#_base: common: 3
		#FormFoo: {
			#_base
			fooID: string
		}
		#FormBar: {
			#_base
			barID: string
		}
		#Form: {
			#FormFoo | #FormBar
		}
	}
	
	#Input: {
		forms: [...#K1.#Form]
	}
	
	#summarizeReturn: {
		in:  #Input
		out: [ for k in in.forms { k.common } ]
	}
	
	#compute: {
		in:  #Input
		out: (#summarizeReturn & {"in": in}).out
	}

	`
	source := []byte(file)

	tree, err := context.Parser{}.Parse(source)
	if err != nil {
		return
	}

	applyCount := countTokensUsingApply(tree)
	applyCountAfter := countTokensUsingApplyAfter(tree)
	scannerCount := countTokensUsingScanner(t, source)

	fmt.Println(applyCount)
	fmt.Println(applyCountAfter)
	fmt.Println(scannerCount) // equal to Perses result
}

func countTokensUsingApply(tree *ast.File) int {
	count := 0
	astutil.Apply(tree, func(cursor astutil.Cursor) bool {
		count++
		return true
	}, nil)

	return count
}

func countTokensUsingApplyAfter(tree *ast.File) int {
	count := 0
	astutil.Apply(tree, nil, func(cursor astutil.Cursor) bool {
		count++
		return true
	})

	return count
}

func countTokensUsingScanner(t *testing.T, source []byte) int {
	eh := func(_ token.Pos, msg string, args []interface{}) {
		t.Errorf("error handler called (msg = %s)", fmt.Sprintf(msg, args...))
	}

	// verify scan
	var s scanner.Scanner
	s.Init(token.NewFile("", -1, len(source)), source, eh, scanner.ScanComments|scanner.DontInsertCommas)

	count := 0

	for {
		_, tok, _ := s.Scan()
		if tok == token.EOF {
			break
		}

		count++
	}

	return count
}
