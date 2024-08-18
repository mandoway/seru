package context

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/languages/cue/strategy"
	"github.com/mandoway/seru/reduction/semantic"
)

var Strategies = []semantic.Strategy[ast.File]{
	strategy.LetReduction{},
}

var Context = semantic.Context[ast.File]{
	Parser:     language.Parser{},
	Strategies: Strategies,
	Serializer: language.Serializer{},
}
