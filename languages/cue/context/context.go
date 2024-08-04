package context

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy"
	"github.com/mandoway/seru/reduction/semantic"
)

func BuildContext() semantic.Context[ast.File] {
	return semantic.Context[ast.File]{
		Parser: Parser{},
		Strategies: []semantic.Strategy[ast.File]{
			strategy.LetReduction{},
		},
		Serializer: Serializer{},
	}
}
