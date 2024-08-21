package main

import (
	"cuelang.org/go/cue/ast"
	"cuelang.org/go/cue/ast/astutil"
	"cuelang.org/go/cue/scanner"
	"cuelang.org/go/cue/token"
	"fmt"
	"github.com/mandoway/seru/languages/cue/language"
	"testing"
)

func TestParser(t *testing.T) {
	file := `
	#B: {
	n: string
	v: {
		r: {
		}
	}
	}
	
	#F: {
		t: [...string]
		...
	}
	
	#N: {
		n: string
		p: #F
		v: #V
	}
	
	#C: {
		p: #F
		v: #V
	}
	
	#C2: {
		c: #C
		t: #N
		ns: {
			[NS=string]: #N & {n: NS}
		}
	
		g: {
			t: {...} | *{}
			...
		}
	}
	
	#V: {
		let l = {[string]: _}
		t: [string]: {}
		ns: [string]: {}
		r: l
	}
	
	fs: f1: #C2 & {
		ns: m: {
			p: {}
			v: {
				r: {
					for s in ["e"] {
						(L & {n: "m", sc: s}).j
					}
				}
			}
	
		}
	
		t: {
			NS=n: string
	
			v: {
				r: {
					"\(NS)/b": {
						metadata: {
							n: NS
						}
					}
					(L & {n: NS, sc: "o"}).j
				}
			}
		}
	}
	
	let L = {
		NS=n: string
		sc:   string
	
		let l = #B & {
			n: NS
		}
	
		j: {
			l.v.r
		}
	}
	`
	source := []byte(file)

	tree, err := language.Parser{}.Parse(source)
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
	// yields wrong results
	count := 0
	astutil.Apply(tree, func(cursor astutil.Cursor) bool {
		count++
		return true
	}, nil)

	return count
}

func countTokensUsingApplyAfter(tree *ast.File) int {
	// yields wrong results
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
	if s.ErrorCount > 0 {
		t.Errorf("error count is %d", s.ErrorCount)
	}

	return count
}
