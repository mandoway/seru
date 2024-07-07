package reduction

type Serializer[AST any] interface {
	Serialize(abstractTree *AST) ([]byte, error)
}
