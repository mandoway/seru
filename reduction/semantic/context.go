package semantic

type Context[AST any] struct {
	Parser     Parser[AST]
	Strategies []Strategy[AST]
	Serializer Serializer[AST]
}
