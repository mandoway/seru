package reduction

type Context[AST any] struct {
	Parser     Parser[AST]
	Strategies []Strategy[AST]
	Serializer[AST]
}
