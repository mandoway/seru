package semantic

const (
	ReduceFunction = "Reduce"
	CountFunction  = "CountTokens"
)

type ReduceFunctionType func(fileContent []byte) ([][]byte, error)

type CountFunctionType func(fileContent []byte) int

type Functions struct {
	ReduceFunction ReduceFunctionType
	CountFunction  CountFunctionType
}
