package semantic

type Serializer[AST any] interface {
	Serialize(abstractTree *AST) ([]byte, error)
}
