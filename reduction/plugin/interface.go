package plugin

const (
	SemanticReduceFunction   = "Reduce"
	CountFunction            = "CountTokens"
	GetStrategyCountFunction = "GetStrategyCount"
	GetStrategyNameFunction  = "GetStrategyName"
)

// SemanticReductionFunction generates candidate programs based on the content of the current program
type SemanticReductionFunction func(fileContent []byte, strategyIndex int) ([][]byte, error)

// TokenCountFunction determines the amount of tokens in a program
type TokenCountFunction func(fileContent []byte) int

// StrategyCountFunction returns the total amount of strategies in the plugin
type StrategyCountFunction func() int

// StrategyNameFunction returns the name of a strategy or empty if not found
type StrategyNameFunction func(index int) string

type Functions struct {
	SemanticReduce   SemanticReductionFunction
	CountTokens      TokenCountFunction
	GetStrategyCount StrategyCountFunction
	GetStrategyName  StrategyNameFunction
}
