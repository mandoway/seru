package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/reduction"
)

type LetReduction struct {
	reduction.Strategy[ast.File]
}

func (r LetReduction) Apply(input *ast.File) []*ast.File {
	modified := []*ast.File{
		input,
		input,
	}
	return modified
}
