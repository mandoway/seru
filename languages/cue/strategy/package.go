package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type PackageReduction struct {
}

func (p PackageReduction) Apply(input *ast.File) []*ast.File {
	alwaysTrue := func(node *ast.Package) bool {
		return true
	}
	return transform.RemoveApplicableStatements[*ast.Package](input, alwaysTrue)
}
