package reduction

type Action[AST any] interface {
	Apply(input *AST) []*AST
}

type Strategy[AST any] struct{}
