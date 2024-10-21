package context

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/language"
	"github.com/mandoway/seru/languages/cue/strategy"
	"github.com/mandoway/seru/reduction/semantic"
)

var Strategies = []semantic.Strategy[ast.File]{
	strategy.LetReduction{},
	strategy.EmptyDeclarationReduction{},
	strategy.PackageReduction{},
	strategy.RedundantNestingReduction{},
	strategy.ListReduction{},
	strategy.TrivialIfReduction{},
	strategy.IfReduction{},
	strategy.EllipsisReduction{},
	strategy.ConstantPropagationReduction{},
	strategy.StringInterpolationReduction{},
	strategy.LoopUnrollingReduction{},
	strategy.UnificationReduction{},
	strategy.UnionReduction{},
	strategy.ImportReduction{},
}

var Context = semantic.Context[ast.File]{
	Parser:     language.Parser{},
	Strategies: Strategies,
	Serializer: language.Serializer{},
}
