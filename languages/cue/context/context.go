package context

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy"
	"github.com/mandoway/seru/reduction"
)

func BuildContext() reduction.Context[ast.File] {
	return reduction.Context[ast.File]{
		Parser: Parser{},
		Strategies: []reduction.Strategy[ast.File]{
			strategy.LetReduction{},
		},
		Serializer: Serializer{},
	}
}
