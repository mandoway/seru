package reduction

type Parser[AST any] interface {
	Parse(fileContent []byte) (*AST, error)
}
