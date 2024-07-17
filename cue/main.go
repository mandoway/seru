package cue

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/cue/strategy"
	"github.com/mandoway/seru/reduction"
)

func Main(fileContent []byte) ([][]byte, error) {
	ctx := reduction.Context[ast.File]{
		Parser: Parser{},
		Strategies: []reduction.Strategy[ast.File]{
			strategy.LetReduction{},
		},
		Serializer: Serializer{},
	}

	return reduction.Reduce(fileContent, ctx)
}
