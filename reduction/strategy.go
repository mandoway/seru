package reduction

import "cuelang.org/go/cue/ast"

type Action[AST any] interface {
	Apply(input *AST) []*ast.File
}

type Strategy[AST any] struct{}
