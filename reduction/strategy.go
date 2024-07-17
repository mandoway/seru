package reduction

type Strategy[AST any] interface {
	Apply(input *AST) []*AST
}
