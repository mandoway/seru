package strategy

import (
	"cuelang.org/go/cue/ast"
	"github.com/mandoway/seru/languages/cue/strategy/transform"
)

type PackageReduction struct {
}

func (p PackageReduction) Apply(input []byte) []*ast.File {
	alwaysRemove := func(node *ast.Package) transform.Transformation {
		return transform.NewDeleteTransformation()
	}
	return transform.ApplyTransformationToEveryApplicableStatement[*ast.Package](input, alwaysRemove)
}
