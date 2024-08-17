package semantic

const (
	ReduceFunction = "Reduce"
	CountFunction  = "CountTokens"
)

type ReduceFunctionType func(fileContent []byte) ([][]byte, error)

// TokenCountFunctionType TODO move out of semantic package
type TokenCountFunctionType func(fileContent []byte) int

type Functions struct {
	ReduceFunction ReduceFunctionType
	CountFunction  TokenCountFunctionType
}
