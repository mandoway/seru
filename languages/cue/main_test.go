package main

import (
	"cuelang.org/go/cue/ast/astutil"
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

	ast, err := context.Parser{}.Parse([]byte(file))
	if err != nil {
		return
	}

	count := 0
	astutil.Apply(ast, func(cursor astutil.Cursor) bool {
		count++
		return true
	}, nil)

	fmt.Println(count)
}
