package semantic

type Strategy[AST any] interface {
	Apply(input []byte) []*AST
}
